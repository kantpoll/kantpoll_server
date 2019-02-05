/**
 * Kantpoll Project
 * https://github.com/kantpoll
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/boltdb/bolt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lestrrat/go-jwx/jwk"
	"github.com/pkg/errors"
	"golang.org/x/crypto/pbkdf2"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/******************** Constants ********************/
//How many peers to connect
const HOW_MANY_ENODES = "30"

const THE_AND = "--and--"
const REGISTER_VOTER = "RegisterVoter"
const SEND_VOTE = "SendVote"
const SEND_PREVOTE = "SendPreVote"
const ENTER_GROUP = "EnterGroup"

//Standard responses to the client
const ERROR_STRING = "error"
const COMPLETE_STRING = "complete"
const TRUE_STRING = "true"
const FALSE_STRING = "false"
const OKOK_STRING = "OKOK"
const BLANK_SPACE = " "

const SECONDS_TO_INIT_GETH = 20
const SECONDS_TO_INIT_TOR = 15
const SECONDS_TO_UPDATE_INFO = 15
const MAX_SECONDS_ADD_VOTERS = 1200

//For vote cancellation
const DELETE_SUFFIX = "-delete"

/******************** Structs ********************/

type CampaignIPNS struct {
	ipns string
	id   string
}

type CampaignContract struct {
	address string
	mgz     string
}

type Candidate struct {
	Website string `json:"website"`
	Count   int    `json:"count"`
}

type Ballot struct {
	Id                string      `json:"id"`
	Closed            bool        `json:"closed"`
	HowManyCandidates int         `json:"how_many_candidates"`
	Candidates        []Candidate `json:"candidates"`
}

type ChairpersonTor struct {
	Onion     string `json:"onion"`
	PublicKey string `json:"public_key"`
}

type CampaignInfo struct {
	Chairperson                     string         `json:"chairperson"`
	Mgz                             int            `json:"mgz"`
	Rounds                          int            `json:"rounds"`
	RemainingRounds                 int            `json:"remaining_rounds"`
	CurrentBallot                   int            `json:"current_ballot"`
	CurrentMessage                  string         `json:"current_message"`
	HowManyBallots                  int            `json:"how_many_ballots"`
	HowManyGroups                   int            `json:"how_many_groups"`
	CanCancel                       bool           `json:"can_cancel"`
	StoppingAccessionToGroups       string         `json:"stopping_accession_to_groups"`
	DisableCandidateLink            bool           `json:"disable_candidate_link"`
	ChairpersonTor                  ChairpersonTor `json:"chairperson_tor"`
	Ballots                         []Ballot       `json:"ballots"`
	VotesPerBallotCandidateCategory [][]string     `json:"votes_per_ballot_candidate_category"`
	Signature                       string         `json:"signature"`
}

type PublicKeys struct {
	Prefixes []int    `json:"prefixes"`
	Keys     []string `json:"keys"`
}

type GroupBallot struct {
	Committed  []bool   `json:"committed"`
	Numbers    []string `json:"numbers"`
	Candidates []int    `json:"candidates"`
}

type GroupInfo struct {
	Index          int            `json:"index"`
	Chairperson    string         `json:"chairperson"`
	ChairpersonTor ChairpersonTor `json:"chairperson_tor"`
	Category       string         `json:"category"`
	Size           int            `json:"size"`
	Voters         []string       `json:"voters"`
	Pubkeys        PublicKeys     `json:"pubkeys"`
	Ballots        []GroupBallot  `json:"ballots"`
	Signature      string         `json:"signature"`
}

type RegisterVoterMessage struct {
	User                 *string `json:"user"`
	Address              *string `json:"address"`
	Pubkey               *string `json:"pubkey"`
	Signature            *string `json:"signature"`
	NotBefore            *string `json:"not_before"`
	CertificateSignature *string `json:"certificate_signature"`
}

type PreVoteMessage struct {
	Ballot      *int    `json:"ballot"`
	Group       *int    `json:"group"`
	Candidate   *int    `json:"candidate"`
	Address     *string `json:"address"`
	Signature   *string `json:"signature"`
	FirstNumber *string `json:"first_number"`
	Campaign    *string `json:"campaign"`
}

type VoteMessage struct {
	Ballot      *int    `json:"ballot"`
	Group       *int    `json:"group"`
	Candidate   *int    `json:"candidate"`
	Signature   *string `json:"signature"`
	FirstNumber *string `json:"first_number"`
	Campaign    *string `json:"campaign"`
}

type EnterGroupMessage struct {
	Address   *string `json:"address"`
	Category  *int    `json:"category"`
	Signature *string `json:"signature"`
}

/******************** Global variables ********************/

//Set the commands according to the os
var executables = make(map[string]map[string]string)

//In order to know if an instance of geth is already running
//And to kill these commands on exit
var gethCmd, torCmd *exec.Cmd

//For interprocess communication
var stdinGeth io.WriteCloser

//To update campaign and ipfs campaign info
var nsStarted bool

//To interact with the node.js scripts
var hashcodesPrefix string
var ecdsaPrivkey *ecdsa.PrivateKey
var rsaPrivPem string
var rsaPubJWK string
var regularExpression string
var ethPwd string

//The last chainid, which is currently active
var lastId = 0

//The captcha generator
var theCaptcha *captcha.Captcha

//A map with half of the captcha as a key and half as a value
var capMap map[string]string

//A map with pre-vote nonces to be used to send votes
var alternativeCodes = make(map[string]bool)

//The ipns of the current campaign
var currentIpns CampaignIPNS

//The id of the current campaign
var currentCampaignId string

//Data to instantiate a binding to the contract
var currentContract CampaignContract

//Info about the campaign and groups are cached in these variables
var (
	campaignInfo     string
	campaignIPFSInfo string
	groupsInfo       []string
)

//When exiting
var stopRecevingRequests bool

//Used in the methods getHiddenServiceHostname and verifyCaptcha
var myOnion = ""

//Bindings to the Ethereum blockchain
var (
	web3Conn      *ethclient.Client
	ethTransactor *bind.TransactOpts
	ethCaller     = &bind.CallOpts{Pending: true}
	mgz3          *Mgz3
	mgz27         *Mgz27
	mgz81         *Mgz81
)

//To prevent malicious software from interfering with the campaign
var adminSessionCode = ""

//To check users' pkey(address) and id
var certificateAuthority = ""

//To insert users into groups
var nextPositionAvailable = make(map[int][2]int)
var stopAccessionToGroups = false

//Databases
var votesDB *bolt.DB
var cancellationsDB *bolt.DB

//It is used to decrypt tor messages
var rsaPrivKey *rsa.PrivateKey

//The categories of the groups
var groupCategoriesArray []string

//The address of the campaign creator, which is used to figure out the nonce
var creatorAddress common.Address

/******************** Handlers *********************/

/**
 * It handles admin requests
 */
func adminHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//When exiting
		if stopRecevingRequests {
			fprint(w, ERROR_STRING)
			return
		}

		//Adding Content Security Policy and CORS
		w.Header().Add("Content-Security-Policy",
			"script-src 'self' ; child-src 'self'; object-src 'none'; "+
				"form-action 'self' ; connect-src http: https: data: wss: ws: ; worker-src 'self' ")
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:1985 http://127.0.0.1:1985")

		//Preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.Header().Add("Access-Control-Allow-Headers", "Encrypted")
			fprint(w, OKOK_STRING)
			return
		}

		path := r.RequestURI[1:]
		if r.Header.Get("Encrypted") == TRUE_STRING {
			pathBytes, _ := hexutil.Decode(path)
			path = AESDecrypt(adminSessionCode, string(pathBytes))
		}

		if strings.Index(path, "query") == 0 {
			//Verifying if the request comes from a valid administrative session
			if len(adminSessionCode) > 0 && r.Header.Get("encrypted") != TRUE_STRING {
				fprint(w, ERROR_STRING)
				return
			}

			if strings.Index(path, "queryEnode") == 0 {
				whatIsMyEnode(path, w)
			} else if strings.Index(path, "queryAddProfile") == 0 {
				addProfile(path, w)
			} else if strings.Index(path, "queryGetProfile") == 0 {
				getProfile(path, w)
			} else if strings.Index(path, "queryGetIPNS") == 0 {
				getIPNS(path, w)
			} else if strings.Index(path, "queryAddIPNSKey") == 0 {
				addIPNS(path, w)
			} else if strings.Index(path, "querySetBlockchain") == 0 {
				setBlockchain(path, w)
			} else if strings.Index(path, "queryVerifyBlockchain") == 0 {
				verifyBlockchain(path, w)
			} else if strings.Index(path, "queryRunBlockchain") == 0 {
				runBlockchain(path, w)
			} else if strings.Index(path, "querySetCampaignContract") == 0 {
				setCampaignContract(path)
				fprint(w, OKOK_STRING)
			} else if strings.Index(path, "queryInsertAccountIntoBlockchain") == 0 {
				insertAccountIntoBlockchain(path, w)
			} else if strings.Index(path, "queryCreatePwdFile") == 0 {
				createPwdFile(path, w)
			} else if strings.Index(path, "queryGetHiddenServiceHostname") == 0 {
				getHiddenServiceHostname(w)
			} else if strings.Index(path, "queryCleanVariables") == 0 {
				cleanVariables()
				fprint(w, OKOK_STRING)
			} else if strings.Index(path, "queryNewSession") == 0 {
				newSession(path, w)
			} else if strings.Index(path, "queryStartCampaign") == 0 {
				startCampaign(path)
				fprint(w, OKOK_STRING)
			} else if strings.Index(path, "querySendVotes") == 0 {
				sendVotes(w)
			} else if strings.Index(path, "queryShowCancellations") == 0 {
				showCancellations(w)
			} else if strings.Index(path, "queryPublishCampaign") == 0 {
				publishCampaign(w)
			} else if strings.Index(path, "queryAddVoterToGroup") == 0 {
				addVoterToGroupLocal(path, 0)
				fprint(w, OKOK_STRING)
			} else if strings.Index(path, "queryInitMgz") == 0 {
				initMgz(path)
				fprint(w, OKOK_STRING)
			}
			return
		}

		// Serve with the actual handler
		h.ServeHTTP(w, r)
	}
}

/**
 * It handles tor requests sent by voters
 */
func mainTorHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//When exiting
		if stopRecevingRequests {
			fprint(w, ERROR_STRING)
			return
		}

		//Adding Content Security Policy and CORS. Setting also de Transfer-Encoding so that it works with the Android app
		w.Header().Add("Content-Security-Policy",
			"script-src 'none' ; style-src 'none' ; child-src 'none'; object-src 'none'; form-action 'none' ;"+
				" connect-src 'none' ; worker-src 'none'")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

		path, passphrase := decryptPath(r)

		if strings.Index(path, "query") == 0 {
			if strings.Index(path, "queryRegisterVoterGet") == 0 {
				receiveGetMessage(path, w, REGISTER_VOTER)
			} else if strings.Index(path, "querySendPreVoteGet") == 0 {
				receiveGetMessage(path, w, SEND_PREVOTE)
			} else if strings.Index(path, "querySendVoteGet") == 0 {
				receiveGetMessage(path, w, SEND_VOTE)
			} else if strings.Index(path, "queryEnterGroupGet") == 0 {
				receiveGetMessage(path, w, ENTER_GROUP)
			} else if strings.Index(path, "queryGetCampaignInfo") == 0 {
				printCampaignInfo(w)
			} else if strings.Index(path, "queryGetCampaignIPFSInfo") == 0 {
				printCampaignIPFSInfo(w)
			} else if strings.Index(path, "queryGetGroupInfo") == 0 {
				getGroupInfo(path, w)
			} else if strings.Index(path, "queryMyGroupIndex") == 0 {
				myGroupIndex(path, w, passphrase)
			} else if strings.Index(path, "queryGenerateCaptcha") == 0 {
				generateCaptcha(w)
			} else if strings.Index(path, "queryCheckPreVote") == 0 {
				checkPreVote(path, w)
			}
		} else {
			fprint(w, BLANK_SPACE)
		}
	}
}

/**
 * The path can come encrypted
 */
func decryptPath(r *http.Request) (string, string) {
	parts := strings.Split(r.RequestURI[1:], THE_AND)
	path := parts[0]
	passphraseStr := ""

	if len(parts) > 1 {
		if rsaPrivKey == nil {
			return "", ""
		}

		passphraseEncrypted, err1 := hexutil.Decode(parts[1])
		if err1 != nil {
			return "", ""
		}

		passphrase, err2 := rsa.DecryptOAEP(sha256.New(),
			crand.Reader, rsaPrivKey, []byte(passphraseEncrypted), nil)
		if err2 != nil {
			return "", ""
		}

		pathBytes, err3 := hexutil.Decode(path)
		if err3 != nil {
			return "", ""
		}

		passphraseStr = string(passphrase)
		path = AESDecrypt(passphraseStr, string(pathBytes))
	}

	return path, passphraseStr
}

/**
 * 	An observer only receives checkPreVote requests
 */
func checkPreVoteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//When exiting
		if stopRecevingRequests {
			fprint(w, ERROR_STRING)
			return
		}

		//Adding Content Security Policy and CORS. Setting also de Transfer-Encoding so that it works with the Android app
		w.Header().Add("Content-Security-Policy",
			"script-src 'none' ; style-src 'none' ; child-src 'none'; object-src 'none'; form-action 'none' ;"+
				" connect-src 'none' ; worker-src 'none'")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

		path, _ := decryptPath(r)
		if strings.Index(path, "queryCheckPreVote") == 0 {
			checkPreVote(path, w)
		} else if strings.Index(path, "queryGetRSAPubkey") == 0 {
			fprint(w, rsaPubJWK)
		} else {
			fprint(w, BLANK_SPACE)
		}
	}
}

/******************** Functions ********************/

/**
 * Start a full server
 */
func startServer() {
	//The map with executable paths
	loadExecutables()

	//Checking if the .kantpoll subdirectories were created. If it was not created, then create it
	initDirectory()

	printline("╔-------------------------------------------------------------╗ ", true, false)
	printline("|            Kantpoll - Election Management System            |", false, false)
	printline("|-------------------------------------------------------------|", false, false)
	printline("| Open http://localhost:1985/login.html to create a vault     |", false, false)
	printline("| Open http://localhost:1985/home.html to manage a campaign   |", false, false)
	printline("| Type 'exit' to exit                                         |", false, false)
	printline("╚-------------------------------------------------------------╝ ", false, true)

	//This is used to avoid "bot voters"
	configCaptcha()

	//Handler to interact with the IPFS and Geth
	changeHeaderThenServe := adminHandler

	//Handler to interact with Tor requests
	torHandler := mainTorHandler

	initIPFS(0)

	//Initializing Tor to provide the hidden service
	initTor()

	setMyOnion(true)

	go func() {
		//Listen requests made through the Tor network
		http.ListenAndServe(":1988", torHandler())
	}()

	//Updates campaign info
	go func() {
		for !stopRecevingRequests {
			time.Sleep(SECONDS_TO_UPDATE_INFO * time.Second)

			getCampaignInfo()
			getCampaignIPFSInfo()
		}
	}()

	//Updates groups info
	go func() {
		for !stopRecevingRequests {
			time.Sleep(SECONDS_TO_UPDATE_INFO * time.Second)

			getGroupsInfo()
		}
	}()

	//It provides the web pages and services
	http.ListenAndServe("localhost:1985", changeHeaderThenServe(http.FileServer(http.Dir("./website"))))
}

/**
 * Start a minimal server (just to receive Tor requests and interact with the blockchain
 */
func startObserver() {
	//The map with executable paths
	loadExecutables()

	printline("╔-------------------------------------------------------------╗ ", true, false)
	printline("|      Kantpoll - Election Management System (Observer)       |", false, false)
	printline("|-------------------------------------------------------------|", false, false)
	printline("| Type 'exit' to exit                                         |", false, false)
	printline("╚-------------------------------------------------------------╝ ", false, true)

	//Handler to interact with Tor requests
	torHandler := checkPreVoteHandler

	//Initializing Tor to provide the hidden service
	initTor()
	printMyOnion(true)

	//To receive requests
	loadRSAFile(true)

	//Binding to the Ethereum contract used to check pre-votes
	var err1 error
	web3Conn, err1 = ethclient.Dial("http://127.0.0.1:8545")
	if err1 == nil {
		_, err1 = web3Conn.NetworkID(context.TODO())
	}

	if err1 != nil {
		printline("Failed to connect to the Ethereum node. Is it running?", false, false)
	} else {
		var err2 error
		switch currentContract.mgz {
		case "3":
			mgz3, err2 = NewMgz3(common.HexToAddress(currentContract.address), web3Conn)
			break
		case "27":
			mgz27, err2 = NewMgz27(common.HexToAddress(currentContract.address), web3Conn)
			break
		case "81":
			mgz81, err2 = NewMgz81(common.HexToAddress(currentContract.address), web3Conn)
		}

		if err2 != nil || isMgzNull() {
			printline("Failed to instantiate the contract.", false, false)
		} else {
			printline("Contract opened.", false, false)
		}
	}

	//Listen requests made through the Tor network
	http.ListenAndServe(":1988", torHandler())
}

/**
 * It generates an RSA key and prints the RSA public key on the screen
 */
func loadRSAFile(firstCall bool) {
	fileName := getHome() + "/.kantpoll/rsapriv.txt"
	file, err1 := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0700)
	defer file.Close()
	if err1 != nil {
		printline("Error creating or opening the RSA private key file", false, false)
		return
	}

	data, err2 := ioutil.ReadFile(fileName)
	if err2 != nil {
		printline("Error reading the RSA private key file", false, false)
		return
	}

	dataStr := string(data)
	if len(dataStr) > 1 {
		loadRSAPrivPem(dataStr)
		generateRSAPubJWK()
	} else if firstCall {
		reader := crand.Reader
		bitSize := 1024

		key, err3 := rsa.GenerateKey(reader, bitSize)
		if err3 != nil {
			printline("Error generating RSA private key", false, false)
			return
		}

		privKeyBytes, _ := x509.MarshalPKCS8PrivateKey(key)
		var privateKey = &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privKeyBytes,
		}

		err4 := pem.Encode(file, privateKey)
		if err4 != nil {
			printline("Error generating RSA private key file", false, false)
			return
		}

		loadRSAFile(false)
	} else {
		printline("RSA private key file generation failed", false, false)
	}
}

/**
 * It generates the RSA public key used by the observer to receive requests, and prints its hash
 */
func generateRSAPubJWK() {
	if rsaPrivKey != nil {
		jwkKey, err1 := jwk.New(&rsaPrivKey.PublicKey)
		if err1 != nil {
			printline("RSA public key error", false, false)
			return
		}
		pubJWK, err2 := json.MarshalIndent(jwkKey, "", " ")
		if err2 != nil {
			printline("RSA public key error - JSON", false, false)
			return
		}

		rsaPubJWK = string(pubJWK)
		hash := crypto.Keccak256Hash([]byte(rsaPubJWK))
		printline("RSA pubkey hash: "+hash.String(), false, false)
	}
}

/**
 * To avoid bot voters sending requests via Tor
 */
func configCaptcha() {
	theCaptcha = captcha.New()
	capMap = make(map[string]string)
	fontContenrs, err := ioutil.ReadFile("website/fonts-for-captcha/Roboto-Regular.ttf")
	if err != nil {
		printline(err.Error(), false, false)
	}
	err = theCaptcha.AddFontFromBytes(fontContenrs)
	if err != nil {
		printline(err.Error(), false, false)
	}
	theCaptcha.SetSize(75, 25)
	theCaptcha.SetDisturbance(captcha.MEDIUM)
	theCaptcha.SetFrontColor(
		color.RGBA{0, 0, 0, 255},
		color.RGBA{0, 0, 255, 255},
		color.RGBA{0, 153, 0, 255})
	theCaptcha.SetBkgColor(color.RGBA{255, 255, 255, 255})
}

/**
 * In order to start a new campaign
 */
func cleanVariables() {
	campaignInfo = ""
	campaignIPFSInfo = ""
	groupsInfo = nil
	nsStarted = false
	nextPositionAvailable = make(map[int][2]int)
	go func() {
		time.Sleep(1 * time.Second) //max: time to start the campaign on server
		mgz3 = nil
		mgz27 = nil
		mgz81 = nil
	}()
}

/**
 * Update server variables with info about some contract
 */
func setCampaignContract(path string) {
	address := getParam(path, "?address=", "&mgz=")
	mgz := getParam(path, "&mgz=", "")
	currentContract.address = address
	currentContract.mgz = mgz
}

/**
 * To generate a random adminSessionCode
 */
func randomStr(n int) string {
	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

/**
 * To prevent malicious software from interfering with the campaign
 */
func newSession(path string, w http.ResponseWriter) {
	if len(adminSessionCode) == 0 {
		adminSessionCode = randomStr(16)

		rsaPubPem := string(common.FromHex(getParam(path, "?rsa=", "")))
		if rsaPubPem == "" {
			fprint(w, ERROR_STRING)
			adminSessionCode = ""
			return
		}

		block, _ := pem.Decode([]byte(rsaPubPem))
		if block != nil && block.Type == "RSA PUBLIC KEY" {
			pub, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err == nil {
				rsaPub, ok := pub.(*rsa.PublicKey)
				if ok {
					cipherText, err2 := rsa.EncryptOAEP(sha256.New(),
						crand.Reader, rsaPub, []byte(adminSessionCode), nil)
					if err2 == nil {
						fprint(w, hexutil.Encode(cipherText))
						return
					}
				}
			}
		}
	}

	adminSessionCode = ""
	fprint(w, ERROR_STRING)
}

/*
 * These databases store votes and cancellations
 */
func initDatabases() {
	var err error
	f1 := "dbs" + string(os.PathSeparator) + "votes_" + currentCampaignId + ".db"
	votesDB, err = bolt.Open(f1, 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		printline("Cannot initialize votes db", false, false)
	}

	f2 := "dbs" + string(os.PathSeparator) + "cancellations_" + currentCampaignId + ".db"
	cancellationsDB, err = bolt.Open(f2, 0600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		printline("Cannot initialize cancellations db", false, false)
	}
}

/**
 * It returns a CAPTCHA for the voter to authenticate
 */
func generateCaptcha(w http.ResponseWriter) {
	_, capImg := genCaptcha()
	capUri := getCaptchaURI(capImg)
	fprint(w, capUri)
}

/**
 * It generates a captcha image and string, and stores this string in a map
 */
func genCaptcha() (string, image.Image) {
	img, str := theCaptcha.Create(6, captcha.NUM)
	capMap[str[0:3]] = str[3:6]

	return str, img
}

/**
 * It returns an URI representing some image
 */
func getCaptchaURI(img image.Image) string {
	out := new(bytes.Buffer)
	err := png.Encode(out, img)
	base64Img := ""

	if err != nil {
		printline("Can't encode captcha image", false, false)
		return ""
	} else {
		base64Img = base64.StdEncoding.EncodeToString(out.Bytes())
	}

	return "data:image/png;base64," + base64Img
}

/**
 * Verifying if the captcha string provided is present in the captcha map
 */
func verifyCaptcha(str string) bool {
	if len(str) != 6 {
		if val, ok := alternativeCodes[str]; ok {
			if val == true {
				delete(alternativeCodes, str)
				return true
			}
		}
		return false
	}
	if val, ok := capMap[str[0:3]]; ok {
		if val == str[3:6] {
			delete(capMap, str[0:3])
			return true
		}
	}
	return false
}

/**
 * This function receives GET messages from Tor pages
 */
func receiveGetMessage(path string, w http.ResponseWriter, messageType string) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	if !verifyParams(path, "?data=", "&captcha=") {
		fprint(w, ERROR_STRING)
		return
	}

	capStr := getParam(path, "&captcha=", "")
	if verifyCaptcha(capStr) {
		if messageType == REGISTER_VOTER {
			receiveRegisterVoterMessage(path, w)
		} else if messageType == SEND_PREVOTE {
			receivePreVoteMessage(path, w)
		} else if messageType == SEND_VOTE {
			receiveVoteMessage(path, w)
		} else if messageType == ENTER_GROUP {
			receiveEnterGroupMessage(path, w)
		}
	}
}

/**
 * It stores a vote into the db
 */
func storeVote(vote VoteMessage) error {
	currentBallot := getCurrentBallot()
	if *vote.Ballot != toInt(currentBallot) {
		return errors.New("Wrong ballot")
	}

	bucketId := []byte(strconv.Itoa(*vote.Ballot) + "-" + strconv.Itoa(*vote.Group))
	err1 := votesDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketId)
		if b != nil {
			err5 := b.ForEach(func(k, value []byte) error {
				var voteMessage VoteMessage
				err6 := json.Unmarshal(value, &voteMessage)
				if err6 == nil {
					if *voteMessage.FirstNumber == *vote.FirstNumber {
						return errors.New("Vote already inserted")
					}
				}
				return err6
			})
			return err5
		}
		return nil
	})

	if err1 != nil {
		return err1
	}

	err2 := votesDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketId)
		var err3 error
		if b == nil {
			b, err3 = tx.CreateBucketIfNotExists(bucketId)
			if err3 != nil {
				return err3
			}
		}

		id, _ := b.NextSequence()
		intId := int(id)

		// Marshal vote data into bytes
		buf, err4 := json.Marshal(vote)
		if err4 != nil {
			return err4
		}

		// Persist bytes to users bucket.
		return b.Put(itob(intId), buf)
	})

	return err2
}

/**
 * It stores a cancellation into the db
 */
func storeCancellation(canc PreVoteMessage) error {
	currentBallot := getCurrentBallot()
	if *canc.Ballot != toInt(currentBallot) {
		return errors.New("Wrong ballot")
	}
	bucketId := []byte(strconv.Itoa(*canc.Ballot) + "-" + strconv.Itoa(*canc.Group))
	err1 := cancellationsDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketId)
		if b != nil {
			err5 := b.ForEach(func(k, value []byte) error {
				var cancMessage PreVoteMessage
				err6 := json.Unmarshal(value, &cancMessage)
				if err6 == nil {
					if *cancMessage.FirstNumber == *canc.FirstNumber {
						return errors.New("Cancellation already inserted")
					}
				}
				return err6
			})
			return err5
		}
		return nil
	})

	if err1 != nil {
		return err1
	}

	err2 := cancellationsDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketId)
		var err3 error
		if b == nil {
			b, err3 = tx.CreateBucketIfNotExists(bucketId)
			if err3 != nil {
				return err3
			}
		}

		id, _ := b.NextSequence()
		intId := int(id)

		// Marshal vote data into bytes
		buf, err4 := json.Marshal(canc)
		if err4 != nil {
			return err4
		}

		// Persist bytes to users bucket.
		return b.Put(itob(intId), buf)
	})

	return err2
}

/**
 * It returns an 8-byte big endian representation of v.
 */
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

/**
 * It puts all votes of the db into the blockchain
 */
func sendVotes(w http.ResponseWriter) {
	createResultsFile()

	currentBallot := getCurrentBallot()
	ballotStr := currentBallot.String()

	hmgInt := getHowManyGroups()
	var err1 error
	for i := 0; i < hmgInt; i++ {
		err1 = votesDB.View(func(tx *bolt.Tx) error {
			bucketId := []byte(ballotStr + "-" + strconv.Itoa(i))
			b := tx.Bucket(bucketId)
			if b != nil {
				err2 := b.ForEach(func(k, value []byte) error {
					var voteMessage VoteMessage
					err3 := json.Unmarshal(value, &voteMessage)
					if err3 == nil && *voteMessage.Ballot == toInt(currentBallot) {
						index := new(big.Int)
						index = index.SetBytes(k)
						index = toBig(toInt(*index) - 1) //The index starts with 1
						err3 = sendVote(voteMessage, index)
					}
					return err3
				})
				return err2
			}
			return nil
		})

		if err1 != nil {
			printline("Error while sending the votes", true, true)
			fprint(w, ERROR_STRING)
			break
		}
	}

	if err1 == nil {
		fprint(w, COMPLETE_STRING)
	}
}

/**
 * It prints all cancellations in the current ballot
 */
func showCancellations(w http.ResponseWriter) {
	currentBallot := getCurrentBallot()
	ballotStr := currentBallot.String()
	results := "Ballot: " + ballotStr + "\r\n" + "Message: " + currentCampaignId + "-" + ballotStr + DELETE_SUFFIX +
		"\r\n" + "Groups: \r\n"

	hmgInt := getHowManyGroups()
	var err1 error
	for i := 0; i < hmgInt; i++ {
		err1 = cancellationsDB.View(func(tx *bolt.Tx) error {
			bucketId := []byte(ballotStr + "-" + strconv.Itoa(i))
			b := tx.Bucket(bucketId)
			if b != nil {
				results += "Group " + strconv.Itoa(i) + ":\r\n"
				err2 := b.ForEach(func(k, value []byte) error {
					var cancMessage PreVoteMessage
					err3 := json.Unmarshal(value, &cancMessage)
					if err3 == nil && *cancMessage.Ballot == toInt(currentBallot) {
						candidate := (*cancMessage.Candidate * -1) - 1
						results += "Address: " + *cancMessage.Address +
							"; Candidate: " + strconv.Itoa(candidate) +
							"; First number: " + *cancMessage.FirstNumber +
							"; Signature: " + *cancMessage.Signature + "\r\n"
					}
					return err3
				})
				return err2
			}
			return nil
		})

		if err1 != nil {
			fprint(w, ERROR_STRING)
			break
		}
	}
	if err1 == nil {
		fprint(w, results)
	}
}

/**
 * It creates a file with group pubkeys and vote signatures
 */
func createResultsFile() {
	currentBallot := getCurrentBallot()
	ballotStr := currentBallot.String()
	results := "Ballot: " + ballotStr + "\r\n" + "Message: " + currentCampaignId + "-" + ballotStr + "\r\n" +
		"Groups: \r\n"

	hmgInt := getHowManyGroups()
	var err1 error
	for i := 0; i < hmgInt; i++ {
		err1 = votesDB.View(func(tx *bolt.Tx) error {
			bucketId := []byte(ballotStr + "-" + strconv.Itoa(i))
			b := tx.Bucket(bucketId)
			if b != nil {
				results += "Group " + strconv.Itoa(i) + ":\r\n"
				results += "Public keys: " + getGroupPubkeys(i) + "\r\n"
				results += "Votes: \r\n"
				err2 := b.ForEach(func(k, value []byte) error {
					var voteMessage VoteMessage
					err3 := json.Unmarshal(value, &voteMessage)
					if err3 == nil && *voteMessage.Ballot == toInt(currentBallot) {
						results += "Candidate: " + strconv.Itoa(*voteMessage.Candidate) +
							"; First number: " + *voteMessage.FirstNumber +
							"; Signature: " + *voteMessage.Signature + "\r\n"
					}
					return err3
				})
				return err2
			}
			return nil
		})

		if err1 != nil {
			printline("Error while creating results file", true, true)
			break
		}
	}
	if err1 == nil {
		addFile("results"+ballotStr, results)
	}
}

/**
 * It puts a vote into the blockchain
 */
func sendVote(vote VoteMessage, position *big.Int) error {
	if isMgzNull() || !nsStarted {
		return errors.New("Blockchain not initialized")
	}

	fnumberBytes := common.FromHex(*vote.FirstNumber)
	fnumberBytes32 := new([32]byte)
	copy(fnumberBytes32[:], fnumberBytes[:32])

	ballot := vote.Ballot
	candidate := vote.Candidate
	group := vote.Group

	ethTransactor.Nonce = getEthNonce()

	var err error
	switch currentContract.mgz {
	case "3":
		_, err = mgz3.Vote(ethTransactor, toBig(*ballot), toBig(*group), position, *fnumberBytes32,
			toBig(*candidate))
		break
	case "27":
		_, err = mgz27.Vote(ethTransactor, toBig(*ballot), toBig(*group), position, *fnumberBytes32,
			toBig(*candidate))
		break
	case "81":
		_, err = mgz81.Vote(ethTransactor, toBig(*ballot), toBig(*group), position, *fnumberBytes32,
			toBig(*candidate))
	}

	if err != nil {
		printline(err.Error(), true, true)
	}

	return err
}

/**
 * Used in getHiddenServiceHostname and printHiddenServiceHostname
 */
func setMyOnion(firstCall bool) {
	myOnionBytes, err := ioutil.ReadFile(getHome() + "/.kantpoll/tor/Data/HiddenService/hostname")
	if err == nil {
		myOnion = string(myOnionBytes)
		if len(myOnion) >= 22 {
			myOnion = myOnion[0:22]
		}
	}

	if err != nil || len(myOnion) < 22 {
		if firstCall {
			go func() {
				time.Sleep(SECONDS_TO_INIT_TOR * time.Second)
				setMyOnion(false)
			}()
		} else {
			printline("Can't find Tor hostname", false, false)
		}
	}
}

/**
 * Used in getHiddenServiceHostname and printHiddenServiceHostname
 */
func printMyOnion(firstCall bool) {
	myOnionBytes, err := ioutil.ReadFile(getHome() + "/.kantpoll/tor/Data/HiddenService/hostname")
	if err == nil {
		myOnion = string(myOnionBytes)
		if len(myOnion) >= 22 {
			myOnion = myOnion[0:22]
			printline("Your hidden service hostname: "+myOnion, false, false)
		}
	}

	if err != nil || len(myOnion) < 22 {
		if firstCall {
			go func() {
				time.Sleep(SECONDS_TO_INIT_TOR * time.Second)
				printMyOnion(false)
			}()
		} else {
			printline("Can't find Tor hostname", false, false)
		}
	}
}

/**
 * It informs the .onion address for the voters to send messages
 */
func getHiddenServiceHostname(w http.ResponseWriter) {
	if len(myOnion) == 22 {
		fprint(w, string(myOnion))
	} else {
		fprint(w, ERROR_STRING)
	}
}

/**
 * It checks whether the fields of the JSON are not nil
 */
func validateEnterGroupMessage(m EnterGroupMessage) bool {
	if m.Signature == nil || m.Address == nil || m.Category == nil {
		return false
	}
	return true
}

/**
 * It checks whether the fields of the JSON are not nil
 */
func validateVoteMessage(m VoteMessage) bool {
	if m.Signature == nil || m.Ballot == nil || m.Campaign == nil || m.FirstNumber == nil || m.Candidate == nil ||
		m.Group == nil {
		return false
	}
	return true
}

/**
 * It checks whether the fields of the JSON are not nil
 */
func validatePreVoteMessage(m PreVoteMessage) bool {
	if m.Signature == nil || m.Ballot == nil || m.Campaign == nil || m.FirstNumber == nil || m.Candidate == nil ||
		m.Group == nil || m.Address == nil {
		return false
	}
	return true
}

/**
 * It checks whether the fields of the JSON are not nil
 */
func validateRegisterVoterMessage(m RegisterVoterMessage) bool {
	if m.Address == nil || m.Signature == nil || m.Pubkey == nil || m.User == nil || m.CertificateSignature == nil ||
		m.NotBefore == nil {
		return false
	}
	return true
}

/**
 * Put a message into a slice
 */
func receiveRegisterVoterMessage(path string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	data := getParam(path, "?data=", "&captcha=")

	ok := registerVoter(data)
	if !ok {
		fprint(w, ERROR_STRING)
		return
	}

	fprint(w, COMPLETE_STRING)
}

/**
 * It registers the voter in the blockchain
 */
func registerVoter(data string) bool {
	if isMgzNull() || !nsStarted || isGroupMembershipStopped() {
		return false
	}

	message, errX := hexutil.Decode(data)
	if errX != nil {
		return false
	}
	var messageJSON RegisterVoterMessage
	err := json.Unmarshal([]byte(message), &messageJSON)
	if err == nil && validateRegisterVoterMessage(messageJSON) {
		caMessage := *messageJSON.User + "," + *messageJSON.Address + "," + *messageJSON.NotBefore
		if verify(currentCampaignId, *messageJSON.Signature, *messageJSON.Address) &&
			verify(caMessage, *messageJSON.CertificateSignature, certificateAuthority) &&
			checkVoterRequirements(*messageJSON.User) {

			if len(*messageJSON.Pubkey) != 66 || !isValidHex(*messageJSON.Pubkey) || !isValidHex(*messageJSON.Address) {
				return false
			}

			sum := sha256.Sum256([]byte(hashcodesPrefix + *messageJSON.User))
			prefix := new(big.Int)
			var pubkey [32]byte
			prefixStr := *messageJSON.Pubkey
			prefix, ok := prefix.SetString(prefixStr[1:2], 10)

			if !ok {
				return false
			}

			pubkeyStr := *messageJSON.Pubkey
			copy(pubkey[:], common.Hex2Bytes(pubkeyStr[2:]))

			ethTransactor.Nonce = getEthNonce()

			var err2, err3 error
			exists := false
			switch currentContract.mgz {
			case "3":
				exists, err2 = mgz3.CheckHashcode(ethCaller, sum)
				if !exists && err2 == nil {
					_, err3 = mgz3.GiveRightToVote(ethTransactor, common.HexToAddress(*messageJSON.Address), prefix, pubkey, sum)
				}
				break
			case "27":
				exists, err2 = mgz27.CheckHashcode(ethCaller, sum)
				if !exists && err2 == nil {
					_, err3 = mgz27.GiveRightToVote(ethTransactor, common.HexToAddress(*messageJSON.Address), prefix, pubkey, sum)
				}
				break
			case "81":
				exists, err2 = mgz81.CheckHashcode(ethCaller, sum)
				if !exists && err2 == nil {
					_, err3 = mgz81.GiveRightToVote(ethTransactor, common.HexToAddress(*messageJSON.Address), prefix, pubkey, sum)
				}
			}

			if err2 == nil && err3 == nil && !exists {
				if verbose {
					printline(REGISTER_VOTER+" action completed", false, false)
				}
				return true
			}
		}
	}

	return false
}

/**
 * Put a message into a slice
 */
func receivePreVoteMessage(path string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	data := getParam(path, "?data=", "&captcha=")

	ok, isCanc := insertPreVote(data)
	if !ok {
		fprint(w, ERROR_STRING)
		return
	}

	if !isCanc {
		nonce := sha256.Sum256([]byte(data))
		nonceStr := hexutil.Encode(nonce[:])
		alternativeCodes[nonceStr[2:18]] = true //It starts at position 2 to remove the "0x"
	}

	fprint(w, COMPLETE_STRING)
}

/**
 * In order to send pre-votes and add voters to groups
 */
func isGroupMembershipStopped() bool {
	var resp [32]byte
	var err0 error
	switch currentContract.mgz {
	case "3":
		resp, err0 = mgz3.StoppingAccessionToGroups(ethCaller)
		break
	case "27":
		resp, err0 = mgz27.StoppingAccessionToGroups(ethCaller)
		break
	case "81":
		resp, err0 = mgz81.StoppingAccessionToGroups(ethCaller)
	}
	if err0 == nil && toAscii(resp) == "stopped" {
		return true
	}

	return false
}

/**
 * It inserts a pre-vote into the blockchain
 * 1st return - action completed
 * 2nd return - is cancellation
 */
func insertPreVote(data string) (bool, bool) {
	if isMgzNull() || !nsStarted || !isGroupMembershipStopped() {
		return false, false
	}

	currentBallotBig := getCurrentBallot()
	currentBallotStr := currentBallotBig.String()

	voteMessage := currentCampaignId + "-" + currentBallotStr

	message, errX := hexutil.Decode(data)
	if errX != nil {
		return false, false
	}

	var messageJSON PreVoteMessage
	err := json.Unmarshal([]byte(message), &messageJSON)
	if err == nil && validatePreVoteMessage(messageJSON) {
		if *messageJSON.Candidate < 0 {
			voteMessage += DELETE_SUFFIX
		}
		if verify(voteMessage, *messageJSON.Signature, *messageJSON.Address) {
			if *messageJSON.Candidate < 0 {
				err3 := storeCancellation(messageJSON)
				if err3 == nil {
					if verbose {
						printline(SEND_PREVOTE+" action completed (cancellation)", false, false)
					}
					return true, true
				}
			} else if candidateInt := messageJSON.Candidate; isValidHex(*messageJSON.FirstNumber) &&
				isValidHex(*messageJSON.Address) && len(*messageJSON.Address) == 42 &&
				len(*messageJSON.FirstNumber) == 66 {

				fnumberBytes := common.FromHex(*messageJSON.FirstNumber)
				fnumberBytes32 := new([32]byte)
				copy(fnumberBytes32[:], fnumberBytes[:32])

				voter := common.FromHex(*messageJSON.Address)
				ballot := toBig(*messageJSON.Ballot)
				group := toBig(*messageJSON.Group)
				candidate := toBig(*candidateInt)

				var voterBytes20 [20]byte
				copy(voterBytes20[:], voter[:20])

				ethTransactor.Nonce = getEthNonce()

				var err2 error
				switch currentContract.mgz {
				case "3":
					_, err2 = mgz3.AddPreVote(ethTransactor, ballot, group, *fnumberBytes32, voterBytes20, candidate)
					break
				case "27":
					_, err2 = mgz27.AddPreVote(ethTransactor, ballot, group, *fnumberBytes32, voterBytes20, candidate)
					break
				case "81":
					_, err2 = mgz81.AddPreVote(ethTransactor, ballot, group, *fnumberBytes32, voterBytes20, candidate)
				}
				if err2 == nil {
					if verbose {
						printline(SEND_PREVOTE+" action completed", false, false)
					}
					return true, false
				}
			}
		}
	}

	return false, false
}

/**
 * It verifies if the received string is a valid hexadecimal number
 */
func isValidHex(hexStr string) bool {
	if strings.Index(hexStr, "0x") == 0 {
		hexStr = hexStr[2:]
	}
	_, err := hex.DecodeString(hexStr)
	if err != nil {
		return false
	}
	return true
}

/**
 * Put a message into a slice
 */
func receiveVoteMessage(path string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	data := getParam(path, "?data=", "&captcha=")

	ok := insertVote(data)
	if !ok {
		fprint(w, ERROR_STRING)
		return
	}

	fprint(w, COMPLETE_STRING)
}

/**
 * The current ballot of the current campaign
 */
func getCurrentBallot() big.Int {
	currentBallotBig := new(big.Int)

	if isMgzNull() || !nsStarted {
		return *currentBallotBig
	}

	switch currentContract.mgz {
	case "3":
		currentBallotBig, _ = mgz3.CurrentBallot(ethCaller)
		break
	case "27":
		currentBallotBig, _ = mgz27.CurrentBallot(ethCaller)
		break
	case "81":
		currentBallotBig, _ = mgz81.CurrentBallot(ethCaller)
	}

	return *currentBallotBig
}

/**
 * It inserts a vote into the blockchain and also stores the vote into a file
 */
func insertVote(data string) bool {
	if !nsStarted || !isGroupMembershipStopped() {
		return false
	}

	message, errX := hexutil.Decode(data)
	if errX != nil {
		return false
	}

	var messageJSON VoteMessage
	err := json.Unmarshal([]byte(message), &messageJSON)
	if err == nil && validateVoteMessage(messageJSON) && *messageJSON.Campaign == currentCampaignId {
		currentBallotBig := getCurrentBallot()
		ballot := toInt(currentBallotBig)

		if *messageJSON.Ballot == ballot {
			voteMessage := currentCampaignId + "-" + strconv.Itoa(ballot)

			pubkeysStr := getGroupPubkeys(*messageJSON.Group)
			if pubkeysStr == "" {
				return false
			}

			if len(pubkeysStr) > 0 && verifyURS(voteMessage, pubkeysStr, *messageJSON.Signature) {
				err3 := storeVote(messageJSON)
				if err3 != nil {
					return false
				}

				if verbose {
					printline(SEND_VOTE+" action completed", false, false)
				}
				return true
			}
		}
	}
	return false
}

/**
 * It returns the public keys in JSON format
 */
func getGroupPubkeys(group int) string {
	if isMgzNull() || !nsStarted {
		return ""
	}

	var pubkeysStr string
	switch currentContract.mgz {
	case "3":
		prefixes, pubkeys, err2 := mgz3.GetGroupPubkeys(ethCaller, toBig(group))
		if err2 != nil {
			return ""
		}
		pubkeysStr = makePubkeys(prefixes[:], pubkeys[:])
		break
	case "27":
		prefixes, pubkeys, err2 := mgz27.GetGroupPubkeys(ethCaller, toBig(group))
		if err2 != nil {
			return ""
		}
		pubkeysStr = makePubkeys(prefixes[:], pubkeys[:])
	case "81":
		prefixes, pubkeys, err2 := mgz81.GetGroupPubkeys(ethCaller, toBig(group))
		if err2 != nil {
			return ""
		}
		pubkeysStr = makePubkeys(prefixes[:], pubkeys[:])
	}

	return pubkeysStr
}

/**
 * It gathers public keys in the format required to run the URS (Unique Ring Signatures) script
 */
func makePubkeys(prefixes []*big.Int, pubkeys [][32]byte) string {
	result := "{"
	for i, j := 0, 0; i < len(prefixes); i++ {
		prefix := toInt(*prefixes[i])
		key := common.Bytes2Hex(pubkeys[i][:])
		key = strings.Replace(key, "0x", "", -1)
		if prefix != 0 {
			result += "\"" + strconv.Itoa(j) + "\":\"" + "0" + strconv.Itoa(prefix) +
				key + "\","
			j++
		}
	}
	result = result[:len(result)-1] + "}"

	return result
}

/**
 * Put a message into a slice
 */
func receiveEnterGroupMessage(path string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	data := getParam(path, "?data=", "&captcha=")

	ok := addVoterToGroup(data)

	if ok {
		fprint(w, COMPLETE_STRING)
	} else {
		fprint(w, ERROR_STRING)
	}
}

/**
 * It adds a voter into a group in the blockchain
 */
func addVoterToGroup(data string) bool {
	if isMgzNull() || !nsStarted || isGroupMembershipStopped() || stopAccessionToGroups {
		return false
	}

	message, errX := hexutil.Decode(data)
	if errX != nil {
		return false
	}

	var messageJSON EnterGroupMessage
	err := json.Unmarshal([]byte(message), &messageJSON)
	if err == nil && validateEnterGroupMessage(messageJSON) &&
		verify(currentCampaignId+"_enter_group", *messageJSON.Signature, *messageJSON.Address) {
		group, position := getGroupAndPosition(*messageJSON.Category)
		if toInt(*group) == -1 || toInt(*position) == -1 || !isValidHex(*messageJSON.Address) {
			return false
		}

		ethTransactor.Nonce = getEthNonce()

		var err2 error
		switch currentContract.mgz {
		case "3":
			_, err2 = mgz3.AddVoterToGroup(ethTransactor, common.HexToAddress(*messageJSON.Address), group, position)
			break
		case "27":
			_, err2 = mgz27.AddVoterToGroup(ethTransactor, common.HexToAddress(*messageJSON.Address), group, position)
			break
		case "81":
			_, err2 = mgz81.AddVoterToGroup(ethTransactor, common.HexToAddress(*messageJSON.Address), group, position)
		}
		if err2 == nil {
			if verbose {
				printline(ENTER_GROUP+" action completed", false, false)
			}
			return true
		}
	}

	return false
}

/**
 * It adds a voter into a group in the blockchain (It is called by the campaign creator)
 */
func addVoterToGroupLocal(path string, attempt int) {
	if isMgzNull() || !nsStarted || attempt > 5 {
		return
	}

	address := getParam(path, "?address=", "&group=")
	groupStr := getParam(path, "&group=", "&position=")
	positionStr := getParam(path, "&position=", "&last=")
	last := getParam(path, "&last=", "")
	group := new(big.Int)
	position := new(big.Int)
	group, _ = group.SetString(groupStr, 10)
	position, _ = position.SetString(positionStr, 10)

	ethTransactor.Nonce = getEthNonce()

	var err2 error
	var tr *types.Transaction
	switch currentContract.mgz {
	case "3":
		tr, err2 = mgz3.AddVoterToGroup(ethTransactor, common.HexToAddress(address), group, position)
		break
	case "27":
		tr, err2 = mgz27.AddVoterToGroup(ethTransactor, common.HexToAddress(address), group, position)
		break
	case "81":
		tr, err2 = mgz81.AddVoterToGroup(ethTransactor, common.HexToAddress(address), group, position)
	}
	if err2 == nil {
		if verbose {
			printline(ENTER_GROUP+" action completed", false, false)
		}
	} else {
		gasPrice, _ := web3Conn.SuggestGasPrice(context.TODO())
		gasPrice = gasPrice.Mul(gasPrice, toBig(4+attempt))
		ethTransactor.GasPrice = gasPrice

		if verbose {
			printline(ENTER_GROUP+" - attempt "+strconv.Itoa(attempt), false, false)
		}
		addVoterToGroupLocal(path, attempt+1)
		if attempt == 0 {
			ethTransactor.GasPrice = getGasPrice()
		}
	}

	if last == TRUE_STRING && attempt == 0 {
		time.Sleep(10 * time.Second)

		go func(tr2 *types.Transaction) {
			stopAccessionToGroups = true
			timestamp := time.Now().Unix()

			for ; time.Now().Unix() < timestamp+MAX_SECONDS_ADD_VOTERS; {

				if tr2 != nil && tr2.Hash().Big() != toBig(0) {
					receipt, err := web3Conn.TransactionReceipt(context.TODO(), tr2.Hash())
					if err != nil {
						printline("Error with the transaction receipt", false, false)
					} else if receipt.Status == uint64(1) { //ReceiptStatusSuccessful
						break
					}
				}
				/*if c, _ := web3Conn.PendingTransactionCount(context.TODO()); c == 0{
					break
				}*/
				time.Sleep(10 * time.Second)
			}
			time.Sleep(1 * time.Minute)

			updateNextPositionAvailable()
			stopAccessionToGroups = false
		}(tr)
	}
}

/**
 * The nonce that should be used to send transactions
 */
func getEthNonce() *big.Int {
	howManyTransactions, _ := web3Conn.PendingNonceAt(context.TODO(), creatorAddress)
	oldNonce := ethTransactor.Nonce
	nonce := new(big.Int)
	if oldNonce != nil && howManyTransactions <= oldNonce.Uint64() {
		return nonce.SetUint64(oldNonce.Uint64() + 1)
	}
	return nonce.SetUint64(howManyTransactions)
}

/**
 * The suggested gas price is not enough
 */
func getGasPrice() *big.Int {
	gasPrice, _ := web3Conn.SuggestGasPrice(context.TODO())
	gasPrice = gasPrice.Mul(gasPrice, toBig(3))
	return gasPrice
}

/**
 * This should be enough for any transaction
 */
func getGasLimit() uint64 {
	return 150000
}

/**
 * The category of groups
 */
func getCategory(category *big.Int) string {
	c := toInt(*category)
	if c >= 0 && c < len(groupCategoriesArray) {
		return groupCategoriesArray[c]
	}
	return ""
}

/**
 * In order to correctly insert voters into groups
 */
func updateNextPositionAvailable() {
	if isMgzNull() || !nsStarted {
		if nsStarted {
			printline("Failed to update next position available", false, false)
		}
		return
	}

	hmgInt := getHowManyGroups()
	for i := 0; i < hmgInt; i++ {
		switch currentContract.mgz {
		case "3":
			group, _ := mgz3.Groups(ethCaller, toBig(i))
			category := toInt(*group.Category)
			size := toInt(*group.Size)
			if group.CPerson.String() == "" || group.CPerson.String() == "0x" {
				continue
			}
			if size > 0 {
				nextPositionAvailable[category] = [2]int{i, size}
			}
			break
		case "27":
			group, _ := mgz27.Groups(ethCaller, toBig(i))
			category := toInt(*group.Category)
			size := toInt(*group.Size)
			if group.CPerson.String() == "" || group.CPerson.String() == "0x" {
				continue
			}
			if size > 0 {
				nextPositionAvailable[category] = [2]int{i, size}
			}
			break
		case "81":
			group, _ := mgz81.Groups(ethCaller, toBig(i))
			category := toInt(*group.Category)
			size := toInt(*group.Size)
			if group.CPerson.String() == "" || group.CPerson.String() == "0x" {
				continue
			}
			if size > 0 {
				nextPositionAvailable[category] = [2]int{i, size}
			}
		}
	}
}

/**
 * It gets the last vacant position for some category and then updates the position
 */
func getGroupAndPosition(category int) (*big.Int, *big.Int) {
	if getCategory(toBig(category)) == "" {
		return toBig(-1), toBig(-1)
	}

	groupNPos := nextPositionAvailable[category]
	group := groupNPos[0]
	pos := groupNPos[1]
	mgzInt, _ := strconv.Atoi(currentContract.mgz)

	if pos >= mgzInt {
		pos = 0
		hmgInt := getHowManyGroups()
		for i := group + 1; i < hmgInt; i++ {
			var g struct {
				CPerson  common.Address
				Category *big.Int
				Size     *big.Int
			}
			var err0 error
			switch currentContract.mgz {
			case "3":
				g, err0 = mgz3.Groups(ethCaller, toBig(i))
				break
			case "27":
				g, err0 = mgz27.Groups(ethCaller, toBig(i))
				break
			case "81":
				g, err0 = mgz81.Groups(ethCaller, toBig(i))
			}

			if err0 != nil {
				return toBig(-1), toBig(-1)
			}

			c := toInt(*g.Category)
			if c == category {
				group = i
				break
			}
		}
	}

	nextPositionAvailable[category] = [2]int{group, pos + 1}
	return toBig(group), toBig(pos)
}

/**
 * It is important to execute this command before the user call "geth init", for the Firewall authorization
 */
func firstTimeGeth() {
	_, err := exec.LookPath(executables[runtime.GOOS]["geth"])

	if err == nil {
		cmd := exec.Command(executables[runtime.GOOS]["geth"], "console")
		out, _ := cmd.Output()
		outstr := string(out)
		cmd.Process.Kill()

		if verbose {
			printCommandResponse(outstr)
		}

		go func() {
			time.Sleep(25 * time.Second)
			cmd.Process.Kill()
		}()
	} else {
		printline("Geth is not installed.", false, false)
	}
}

/**
 * Writing a new genesis.json file and starting a new geth instance
 */
func setBlockchain(path string, w http.ResponseWriter) {
	chainid := getParam(path, "?chainid=", "&address=")
	address := getParam(path, "&address=", "&enodes=")
	//howManyEnodes = getParam(path, "&enodes=", "&directory=")
	directory := getParam(path, "&directory=", "&nonce=")
	nonce := getParam(path, "&nonce=", "")

	//Creating the directory where the data of this specific blockchain will be placed
	if _, err := os.Stat(getHome() + "/.kantpoll/blockchains/" + directory); os.IsNotExist(err) {
		os.Mkdir(getHome()+"/.kantpoll/blockchains/"+directory, 0700)
	}

	//Composing the genesis.json file
	genesis := "{ \"config\": {" +
		"\"chainId\": " + chainid + "," +
		"\"homesteadBlock\": 1," +
		"\"eip150Block\": 2," +
		"\"eip150Hash\": \"0x0000000000000000000000000000000000000000000000000000000000000000\"," +
		"\"eip155Block\": 3," +
		"\"eip158Block\": 3," +
		"\"byzantiumBlock\": 4" +
		"}," +
		"\"difficulty\": \"0x1\"," +
		"\"gasLimit\": \"0x8000000\"," +
		"\"nonce\": \"" + nonce + "\"," +
		"\"alloc\": {" +
		"\"" + address + "\": { \"balance\": \"10000000000000000000000\" }" + //10000 ether
		"}}"

	//Writing genesis string into the file
	data := []byte(genesis)
	err := ioutil.WriteFile(getHome()+"/.kantpoll/blockchains/"+directory+"/genesis.json", data, 0700)
	if err != nil {
		fprint(w, ERROR_STRING)
		printline("Error while creating the genesis.json file", false, false)
		return
	}

	//Initializing Geth
	initGeth(w, directory)
}

/**
 * Initialize the IPFS service to provide the users' pages
 */
func initIPFS(times int) {
	if times == 2 {
		printline("Limit of tries reached", false, false)
		return
	}
	_, err := exec.LookPath(executables[runtime.GOOS]["ipfs"])
	existDir := verifyIpfsDir()

	if err == nil && !existDir {
		cmd := exec.Command(executables[runtime.GOOS]["ipfs"], "init")
		err := cmd.Run()
		cmd.Process.Kill()

		if err == nil {
			printline("IPFS was initialized", false, false)
			initIPFS(times + 1)
		} else {
			printline("IPFS was not initialized", false, false)
		}
	} else if err != nil {
		printline("IPFS is not installed.", false, false)
	}
}

/**
 * Initialize the Tor to allow voters to send messages and nodes to comunicate
 */
func initTor() {
	_, err := exec.LookPath(executables[runtime.GOOS]["tor"])

	if err == nil {
		path := ""
		if runtime.GOOS == "windows" {
			path = "\\.kantpoll\\tor\\Data\\Tor\\torrc"
		} else if runtime.GOOS == "linux" {
			path = "/.kantpoll/tor/Data/Tor/torrc"
		}
		absolutePath := getHome() + path

		torCmd = exec.Command(executables[runtime.GOOS]["tor"], "-f", absolutePath)

		err := torCmd.Start()
		if err == nil {
			printline("Tor has started", false, false)
		} else {
			printline("Tor error", false, false)
		}
	} else {
		printline("Tor is not installed.", false, false)
	}
}

/**
 * Initialize Geth node
 */
func initGeth(w http.ResponseWriter, directory string) {
	_, err := exec.LookPath(executables[runtime.GOOS]["geth"])
	if err == nil {
		datadir := getHome() + string(os.PathSeparator) + ".kantpoll" + string(os.PathSeparator) + "blockchains" +
			string(os.PathSeparator) + directory
		genesis := getHome() + string(os.PathSeparator) + ".kantpoll" + string(os.PathSeparator) + "blockchains" +
			string(os.PathSeparator) + directory + string(os.PathSeparator) + "genesis.json"

		if _, err := os.Stat(genesis); os.IsNotExist(err) {
			printline("The datadir does not exist", false, false)
			fprint(w, ERROR_STRING)
			return
		}

		cmd := exec.Command(executables[runtime.GOOS]["geth"], "--datadir", datadir, "init", genesis)
		err := cmd.Run()
		cmd.Process.Kill()

		if err == nil {
			printline("Geth has started", false, false)
			fprint(w, COMPLETE_STRING)
		} else {
			printline("Geth has not started", false, false)
			fprint(w, ERROR_STRING)
		}
	} else {
		printline("Geth is not installed", false, false)
		fprint(w, ERROR_STRING)
	}
}

/**
 * Verifying if IPFS dir exists
 */
func verifyIpfsDir() bool {
	//Check whether the .ipfs folder was created in user's PC
	if _, err := os.Stat(getHome() + "/.ipfs/version"); !os.IsNotExist(err) {
		return true
	}
	return false
}

/**
 * Getting the HOME directory
 */
func getHome() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

/**
 * It calls geth in order to get the enode
 */
func whatIsMyEnode(path string, w http.ResponseWriter) {
	rpcPort := getParam(path, "?port=", "")

	if _, err := strconv.Atoi(rpcPort); err != nil {
		fprint(w, ERROR_STRING)
		return
	}

	cmd := exec.Command(executables[runtime.GOOS]["geth"], "attach", "http://localhost:"+rpcPort)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fprint(w, ERROR_STRING)
	}
	args := "admin.nodeInfo.enode"

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, args)
	}()

	out, err := cmd.CombinedOutput()
	outstr := string(out)

	if strings.LastIndex(outstr, "enode://") > 0 {
		begin := strings.LastIndex(outstr, "enode://")
		end := strings.LastIndex(outstr, ":30")
		outstr = outstr[begin : end+6]
	}

	fprint(w, outstr)

	go func() {
		time.Sleep(3 * time.Second)
		cmd.Process.Kill()
	}()
}

/**
 * It publishes, on the IPFS, info about the campaign, candidates, parties, group categories, and results
 */
func publishCampaign(w http.ResponseWriter) {
	//Executing IPFS
	_, err := exec.LookPath(executables[runtime.GOOS]["ipfs"])
	if err == nil {
		newIPFSKey(currentCampaignId)

		//The error used to verify whether the profile was inserted in the IPFS or not.
		var err1 *appError
		address, err1 := newIPFSPage(currentCampaignId)
		if err1 == nil {
			err1 = publishIPFS(address, currentCampaignId)
			if err1 == nil {
				//After publishing the directory, try to pin it to local storage.
				pinIPFS(address)
				fprint(w, COMPLETE_STRING)
			} else {
				fprint(w, ERROR_STRING)
			}
		}
	} else {
		fprint(w, ERROR_STRING)
		printline("IPFS not installed", false, false)
	}
}

/**
 * Insert or overwrite a new profile (person or campaign)
 */
func addProfile(path string, w http.ResponseWriter) {
	content := getParam(path, "?content=", "&filename=")
	filename := getParam(path, "&filename=", "")

	//Profiles should be saved in the HOME directory
	if _, err := os.Stat(getHome() + "/.kantpoll/profiles/" + currentCampaignId); os.IsNotExist(err) {
		os.Mkdir(getHome()+"/.kantpoll/profiles/"+currentCampaignId, 0700)
	}
	f, err := os.Create(getHome() + "/.kantpoll/profiles/" + currentCampaignId + "/" + filename)
	if err == nil {
		_, err := f.WriteString(content)
		if err == nil {
			printline("New file: "+currentCampaignId+"/"+filename, false, false)
		}
		f.Close()
		fprint(w, COMPLETE_STRING)
	} else {
		fprint(w, ERROR_STRING)
		printline("File was not created", false, false)
	}
}

/**
 * Similar to addProfile
 */
func addFile(filename, content string) {
	//Files should be saved in the HOME directory
	if _, err := os.Stat(getHome() + "/.kantpoll/profiles/" + currentCampaignId); os.IsNotExist(err) {
		os.Mkdir(getHome()+"/.kantpoll/profiles/"+currentCampaignId, 0700)
	}
	f, err := os.Create(getHome() + "/.kantpoll/profiles/" + currentCampaignId + "/" + filename)
	if err == nil {
		_, err := f.WriteString(content)
		if err == nil {
			printline("New file: "+currentCampaignId+"/"+filename, false, false)
		}
		f.Close()
	} else {
		printline("File was not created", false, false)
	}
}

/**
 * The error used to verify whether the profile was inserted in IPFS or not.
 */
type appError struct {
	Error   error
	Message string
}

/**
 * Before publishing we need to (try to) create a key (all OSs)
 * If a key already exists, IPFS will return a error message saying "refusing to overwrite".
 * It does not affects the ongoing procedures.
 */
func newIPFSKey(dir string) {
	cmd := exec.Command(executables[runtime.GOOS]["ipfs"], "key", "gen", "--type=rsa", "--size=2048", dir)
	cmd.Run()
	cmd.Process.Kill()
}

/**
 * Publish some page with some key(dir) (all OSs)
 */
func publishIPFS(address, dir string) *appError {
	cmd := exec.Command(executables[runtime.GOOS]["ipfs"], "name", "publish", "--key="+dir, address)
	err := cmd.Run()
	cmd.Process.Kill()

	if err == nil {
		printline("Directory "+dir+" was published on IPFS", false, false)
	} else {
		printline("Publishing error with the directory "+dir, true, false)
		printline(err.Error(), false, true)
		return &appError{err, "Publishing error with the directory " + dir}
	}
	return nil
}

/**
 * It stores an IPFS object from a given path locally to disk. (all OSs)
 */
func pinIPFS(address string) {
	cmd := exec.Command(executables[runtime.GOOS]["ipfs"], "pin", "add", "-r", address)
	err := cmd.Run()
	cmd.Process.Kill()

	if err == nil {
		printline(address+"was pinned to local storage", false, false)
	} else {
		printline("Pinning error with the address "+address, false, false)
	}
}

/**
 * Adding a new IPFS page
 */
func newIPFSPage(dir string) (string, *appError) {
	directory := getHome() + string(os.PathSeparator) + ".kantpoll" + string(os.PathSeparator) + "profiles" +
		string(os.PathSeparator) + dir

	cmd := exec.Command(executables[runtime.GOOS]["ipfs"], "add", "-r", directory)
	out, err := cmd.Output()
	outstr := string(out)
	cmd.Process.Kill()

	//Getting the ipfs address of the directory
	if len(outstr) >= 52 {
		begin := strings.LastIndex(outstr, "added")
		outstr = outstr[begin+6 : begin+6+46]
	}

	if err == nil {
		printline("IPFS page was inserted", false, false)
	} else {
		printline("IPFS insertion error", false, false)
		return "", &appError{err, "IPFS insertion error"}
	}

	return outstr, nil
}

/**
 * Returns the profile html content
 */
func getProfile(path string, w http.ResponseWriter) {
	fileStr := getParam(path, "?file=", "")

	completePath := getHome() + "/.kantpoll/profiles/" + fileStr
	content, err := ioutil.ReadFile(completePath)
	if err == nil {
		contentStr := string(content)
		fprint(w, contentStr)
		return
	}

	fprint(w, ERROR_STRING)
}

/**
 * In order to avoid this error: 'Failed to write genesis block: database already contains an incompatible genesis block'
 */
func removeBlockchainDir(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

/**
 * This function receives the same params received by the setBlockchain function,
 * then it compares the content of the genesis file and the params
 */
func verifyBlockchain(path string, w http.ResponseWriter) {
	chainid := getParam(path, "?chainid=", "&address=")
	address := getParam(path, "&address=", "&enode=")
	//enode := getParam(path, "&enode=", "&directory=")
	directory := getParam(path, "&directory=", "&nonce=")
	nonce := getParam(path, "&nonce=", "&delete=")
	deleteDirIfDifferent := getParam(path, "&delete=", "")

	//Composing the genesis.json file
	genesis := "{ \"config\": {" +
		"\"chainId\": " + chainid + "," +
		"\"homesteadBlock\": 1," +
		"\"eip150Block\": 2," +
		"\"eip150Hash\": \"0x0000000000000000000000000000000000000000000000000000000000000000\"," +
		"\"eip155Block\": 3," +
		"\"eip158Block\": 3," +
		"\"byzantiumBlock\": 4" +
		"}," +
		"\"difficulty\": \"0x1\"," +
		"\"gasLimit\": \"0x8000000\"," +
		"\"nonce\": \"" + nonce + "\"," +
		"\"alloc\": {" +
		"\"" + address + "\": { \"balance\": \"10000000000000000000000\" }" + //10000 ether
		"}}"

	content, err := ioutil.ReadFile(getHome() + "/.kantpoll/blockchains/" + directory + "/genesis.json")

	if err == nil {
		if string(content) == genesis {
			fprint(w, TRUE_STRING)
		} else {
			if deleteDirIfDifferent == TRUE_STRING {
				//Removing these directories to avoid this error: Failed to write genesis block:
				//database already contains an incompatible genesis block
				removeBlockchainDir(getHome() + "/.kantpoll/blockchains/" + directory + "/geth")
				removeBlockchainDir(getHome() + "/.kantpoll/blockchains/" + directory + "/keystore")
			}
			fprint(w, FALSE_STRING)
		}
	} else {
		fprint(w, ERROR_STRING)
	}
}

/**
 * It calls: geth --networkid "1151985..." etc
 */
func runBlockchain(path string, w http.ResponseWriter) {
	id := getParam(path, "?id=", "&address=")
	address := getParam(path, "&address=", "&dir=")
	dir := getParam(path, "&dir=", "&role=")
	theRole := getParam(path, "&role=", "")

	intId, _ := strconv.Atoi(id)
	//If the campaign has not changed, keep the geth process running
	if intId == lastId {
		printline("Old geth process was kept", false, false)
		fprint(w, COMPLETE_STRING)
	} else {
		if gethCmd != nil && gethCmd.Process != nil {
			err := gethCmd.Process.Kill()
			if err == nil {
				printline("Old geth process was killed", false, false)
			} else {
				printline("Old geth process was not killed", false, false)
			}
		}
		lastId = intId
		runGeth(w, id, address, dir, theRole)
	}
}

/**
 * It can be used by the creator of the campaign or another node (all OSs)
 */
func runGeth(w http.ResponseWriter, id, address, dir, theRole string) {
	if len(theRole) != 3 {
		return
	}
	dataDir := getHome() + string(os.PathSeparator) + ".kantpoll" + string(os.PathSeparator) + "blockchains" +
		string(os.PathSeparator) + dir
	pwdFile := getHome() + string(os.PathSeparator) + ".kantpoll" + string(os.PathSeparator) + "blockchains" +
		string(os.PathSeparator) + dir + string(os.PathSeparator) + "pwd"
	privkeyFile := getHome() + string(os.PathSeparator) + ".kantpoll" + string(os.PathSeparator) + "blockchains" +
		string(os.PathSeparator) + dir + string(os.PathSeparator) + "privkey"

	if theRole == "001" { //the creator
		gethCmd = exec.Command(executables[runtime.GOOS]["geth"], "--datadir", dataDir, "--networkid", id, "--nodiscover",
			"--port", "30001", "--maxpeers", HOW_MANY_ENODES, "--ipcdisable", "--rpc", "--rpcaddr", "localhost",
			"--rpcport", "8545", "--rpcapi", "admin,personal,net,eth,web3", "--rpccorsdomain",
			"http://localhost:1985,http://127.0.0.1:1985", "--mine", "--minerthreads=1",
			"--etherbase", address, "--unlock", address, "--password", pwdFile,
			"--gcmode", "archive", "--lightkdf", "console")
	} else {
		gethCmd = exec.Command(executables[runtime.GOOS]["geth"], "--datadir", dataDir, "--networkid", id, "--nodiscover",
			"--port", "30"+theRole, "--maxpeers", HOW_MANY_ENODES, "--ipcdisable", "--rpc", "--rpcaddr", "localhost",
			"--rpcport", "8545", "--rpcapi", "admin,personal,net,eth,web3", "--rpccorsdomain",
			"http://localhost:1985,http://127.0.0.1:1985", "--unlock", address, "--password", pwdFile,
			"--gcmode", "archive", "--lightkdf", "console")
	}

	stdinGeth, _ = gethCmd.StdinPipe()
	err := gethCmd.Start()

	if err == nil {
		printline("Geth running", false, false)
		fprint(w, COMPLETE_STRING)
	} else {
		printline("Geth not running", false, false)
		fprint(w, ERROR_STRING)
	}

	//Removing the privatekey and password files, and initializing the go bindings
	go func() {
		time.Sleep(SECONDS_TO_INIT_GETH * time.Second)
		os.Remove(privkeyFile)
		os.Remove(pwdFile)
	}()
}

/**
 * It initialize the variable mgz, which is used to interact with the Ethereum contract
 */
func initMgz(path string) {
	if !isMgzNull() {
		return
	}

	dir := getParam(path, "?dir=", "")
	keystoreDir := getHome() + string(os.PathSeparator) + ".kantpoll" + string(os.PathSeparator) + "blockchains" +
		string(os.PathSeparator) + dir + string(os.PathSeparator) + "keystore"
	fi, _ := ioutil.ReadDir(keystoreDir)

	var keyStr string
	if len(fi) > 0 {
		fileName := fi[0].Name()
		keyBytes, _ := ioutil.ReadFile(keystoreDir + string(os.PathSeparator) + fileName)
		keyStr = string(keyBytes)
	}

	//Binding to the Ethereum contract used to check pre-votes
	var err1 error
	web3Conn, err1 = ethclient.Dial("http://127.0.0.1:8545")
	if err1 == nil {
		_, err1 = web3Conn.NetworkID(context.TODO())
	}

	var err2 error
	ethTransactor, err2 = bind.NewTransactor(strings.NewReader(keyStr), ethPwd)

	if err1 != nil || err2 != nil {
		printline("Failed to connect to the Ethereum node.", false, false)
	} else {
		ethTransactor.GasPrice = getGasPrice()
		ethTransactor.GasLimit = getGasLimit()

		var err3 error
		switch currentContract.mgz {
		case "3":
			mgz3, err3 = NewMgz3(common.HexToAddress(currentContract.address), web3Conn)
			break
		case "27":
			mgz27, err3 = NewMgz27(common.HexToAddress(currentContract.address), web3Conn)
			break
		case "81":
			mgz81, err3 = NewMgz81(common.HexToAddress(currentContract.address), web3Conn)
		}

		if err3 != nil || isMgzNull() {
			printline("Failed to instantiate the contract.", false, false)
		} else {
			printline("Contract opened.", false, false)

			time.Sleep(SECONDS_TO_INIT_GETH * time.Second)
			updateNextPositionAvailable()
		}
	}
}

/**
 * Before calling runBlockchainLin/Win, create the password file to unlock the main account
 */
func createPwdFile(path string, w http.ResponseWriter) {
	dir := getParam(path, "?dir=", "&password=")
	password := getParam(path, "&password=", "")

	f, err := os.Create(getHome() + "/.kantpoll/blockchains/" + dir + "/pwd")
	if err == nil {
		_, err2 := f.WriteString(password)
		f.Close()
		if err2 == nil {
			fprint(w, COMPLETE_STRING)
		} else {
			fprint(w, ERROR_STRING)
		}
	} else {
		fprint(w, ERROR_STRING)
	}
}

/**
 * Create new file with a private key (and another with a password) to be imported with the command: geth account import
 */
func createPrivateKeyFile(directory, privkey, password string) bool {
	f, err := os.Create(getHome() + "/.kantpoll/blockchains/" + directory + "/privkey")
	if err == nil {
		_, err2 := f.WriteString(privkey)
		f.Close()
		if err2 == nil {
			f2, err3 := os.Create(getHome() + "/.kantpoll/blockchains/" + directory + "/pwd")
			if err3 == nil {
				_, err4 := f2.WriteString(password)
				f2.Close()
				if err4 == nil {
					return true
				}
			}
		}
	}
	return false
}

/**
 * Insert a new account with the command geth account import. In order to do that, create new privatekey and password file.
 */
func insertAccountIntoBlockchain(path string, w http.ResponseWriter) {
	dir := getParam(path, "?dir=", "&privkey=")
	privkey := getParam(path, "&privkey=", "&password=")
	password := getParam(path, "&password=", "")

	if createPrivateKeyFile(dir, privkey, password) {
		newAccount(w, dir)
	}
}

/**
 * Call the command: geth account import (all OSs)
 */
func newAccount(w http.ResponseWriter, dir string) {
	dataDir := getHome() + string(os.PathSeparator) + ".kantpoll" + string(os.PathSeparator) + "blockchains" +
		string(os.PathSeparator) + dir
	pwdFile := getHome() + string(os.PathSeparator) + ".kantpoll" + string(os.PathSeparator) + "blockchains" +
		string(os.PathSeparator) + dir + string(os.PathSeparator) + "pwd"
	privkeyFile := getHome() + string(os.PathSeparator) + ".kantpoll" + string(os.PathSeparator) + "blockchains" +
		string(os.PathSeparator) + dir + string(os.PathSeparator) + "privkey"

	cmd := exec.Command(executables[runtime.GOOS]["geth"], "--datadir", dataDir, "account", "import",
		privkeyFile, "--password", pwdFile)
	out, err := cmd.Output()
	outstr := string(out)
	cmd.Process.Kill()

	if verbose {
		printCommandResponse(outstr)
	}

	if err == nil {
		printline("Account inserted", false, false)
		fprint(w, COMPLETE_STRING)
	} else {
		printline("Account not inserted", false, false)
		fprint(w, ERROR_STRING)
	}
}

/**
 * Creating a key with the received name
 */
func addIPNS(path string, w http.ResponseWriter) {
	id := getParam(path, "?id=", "")
	newIPFSKey(id)
	fprint(w, COMPLETE_STRING)
}

/**
 * Call the "ipfs key list -l" to obtain the ipns address
 */
func getIPNS(path string, w http.ResponseWriter) {
	id := getParam(path, "?id=", "")

	if currentIpns.id == id {
		fprint(w, currentIpns.ipns)
		return
	}

	ipns := ""
	cmd := exec.Command(executables[runtime.GOOS]["ipfs"], "key", "list", "-l")
	out, err := cmd.Output()
	outstr := string(out)
	cmd.Process.Kill()

	if err == nil {
		lines := stringToLines(outstr)
		ipns = findIPNS(lines, id)

		currentIpns.id = id
		currentIpns.ipns = ipns
		if w != nil {
			fprint(w, ipns)
		}
	} else {
		if w != nil {
			fprint(w, ERROR_STRING)
		}
	}
}

/**
 * Check whether the .kantpoll directory was created. If it was not created, then create it
 */
func initDirectory() {
	userprofile := getHome()
	if _, err := os.Stat(userprofile + "/.kantpoll"); os.IsNotExist(err) {
		os.Mkdir(userprofile+"/.kantpoll", 0700)
	}
	if _, err := os.Stat(userprofile + "/.kantpoll/profiles"); os.IsNotExist(err) {
		os.Mkdir(userprofile+"/.kantpoll/profiles", 0700)
	}
	if _, err := os.Stat(userprofile + "/.kantpoll/blockchains"); os.IsNotExist(err) {
		os.Mkdir(userprofile+"/.kantpoll/blockchains", 0700)
	}
}

/**
 * Load the executables map
 */
func loadExecutables() {
	add(executables, "windows", "ipfs", "ipfs/go-ipfs/ipfs.exe")
	add(executables, "windows", "geth", "geth/"+gethFolderWin+"/geth.exe")
	add(executables, "windows", "tor", "tor/Tor/tor.exe")
	add(executables, "linux", "ipfs", "ipfs/go-ipfs/ipfs")
	add(executables, "linux", "geth", "geth/"+gethFolderLin+"/geth")
	add(executables, "linux", "tor", "tor/Tor/tor")
}

/**
 * This function builds the executables map
 */
func add(m map[string]map[string]string, os, cmd, value string) {
	mm, ok := m[os]
	if !ok {
		mm = make(map[string]string)
		m[os] = mm
	}
	mm[cmd] = value
}

/**
 * Transforming a string in an slice
 */
func stringToLines(s string) []string {
	var lines []string

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

/**
 * Look each line to find the ipns of a certain campaign
 */
func findIPNS(lines []string, id string) string {
	for i := 0; i < len(lines); i = i + 1 {
		if strings.Index(lines[i], id) > 0 {
			return lines[i][:46]
		}
	}
	return ERROR_STRING
}

/**
 * It calls mgzX.getVoter
 */
func getVoter(voterAddr string) (struct {
	Pubkey   [32]byte
	Prefix   *big.Int
	Group    *big.Int
	HasGroup bool
}, error) {

	ret := new(struct {
		Pubkey   [32]byte
		Prefix   *big.Int
		Group    *big.Int
		HasGroup bool
	})

	if isMgzNull() || !nsStarted || !isValidHex(voterAddr) {
		return *ret, nil
	}

	var err error
	switch currentContract.mgz {
	case "3":
		*ret, err = mgz3.GetVoter(ethCaller, common.HexToAddress(voterAddr))
	case "27":
		*ret, err = mgz27.GetVoter(ethCaller, common.HexToAddress(voterAddr))
	case "81":
		*ret, err = mgz81.GetVoter(ethCaller, common.HexToAddress(voterAddr))
	}
	return *ret, err
}

/**
 * Check this before interacting with the blockchain
 */
func isMgzNull() bool {
	return mgz3 == nil && mgz27 == nil && mgz81 == nil
}

/**
 * mgzX.howManyGroups()
 */
func getHowManyGroups() int {
	if isMgzNull() || !nsStarted {
		return 0
	}

	var bigN *big.Int
	var err error
	switch currentContract.mgz {
	case "3":
		bigN, err = mgz3.HowManyGroups(ethCaller)
		break
	case "27":
		bigN, err = mgz27.HowManyGroups(ethCaller)
		break
	case "81":
		bigN, err = mgz81.HowManyGroups(ethCaller)
	}

	if err != nil {
		return 0
	}
	return toInt(*bigN)
}

/**
 * It prints the voter's group index on the page
 */
func myGroupIndex(path string, w http.ResponseWriter, passphrase string) {
	if isMgzNull() || !nsStarted || !verifyParams(path, "?address=") {
		fprint(w, ERROR_STRING)
		return
	}

	voterAddr := getParam(path, "?address=", "")
	if len(voterAddr) != 42 {
		fprint(w, ERROR_STRING)
		return
	}

	voter, err := getVoter(voterAddr)
	if err == nil {
		if voter.HasGroup {
			if passphrase != ""{
				encrypted := AESEncrypt(passphrase, voter.Group.String())
				fprint(w, encrypted)
			} else{
				fprint(w, voter.Group.String())
			}
			return
		} else if !allZeros32(voter.Pubkey) {
			if passphrase != ""{
				encrypted := AESEncrypt(passphrase, "-2")
				fprint(w, encrypted)
			} else{
				fprint(w, "-2")
			}
			return
		}
	}

	if passphrase != ""{
		encrypted := AESEncrypt(passphrase, "-1")
		fprint(w, encrypted)
	} else{
		fprint(w, "-1")
	}
}

/**
 * It checks the candidate of the pre-vote in the blockchain
 */
func checkPreVote(path string, w http.ResponseWriter) {
	if isMgzNull() || (creator && !nsStarted) ||
		!verifyParams(path, "?ballot=", "&group=", "&fnumber=", "&candidate=") {
		fprint(w, ERROR_STRING)
		return
	}

	ballot := new(big.Int)
	ballotStr := getParam(path, "?ballot=", "&group=")
	ballot, ok := ballot.SetString(ballotStr, 10)
	if !ok || len(ballotStr) == 0 {
		fprint(w, ERROR_STRING)
		return
	}

	group := new(big.Int)
	groupStr := getParam(path, "&group=", "&fnumber=")
	group, ok2 := group.SetString(groupStr, 10)
	if !ok2 || len(groupStr) == 0 {
		fprint(w, ERROR_STRING)
		return
	}

	fnumberStr := getParam(path, "&fnumber=", "&candidate=")
	if len(fnumberStr) < 66 {
		fprint(w, ERROR_STRING)
		return
	}
	fnumberBytes := common.FromHex(fnumberStr)
	var fnumber [32]byte
	copy(fnumber[:], fnumberBytes[:32])

	candidate := new(big.Int)
	candidateStr := getParam(path, "&candidate=", "")
	candidate, ok3 := candidate.SetString(candidateStr, 10)
	if !ok3 || len(candidateStr) == 0 {
		fprint(w, ERROR_STRING)
		return
	}

	response := FALSE_STRING

	storedCandidate := new(big.Int)
	var storedVoter [20]uint8
	var preVote struct {
		Voter     [20]byte
		Candidate *big.Int
	}
	var err error

	switch currentContract.mgz {
	case "3":
		preVote, err = mgz3.GetPreVote(ethCaller, ballot, group, fnumber)
		break
	case "27":
		preVote, err = mgz27.GetPreVote(ethCaller, ballot, group, fnumber)
		break
	case "81":
		preVote, err = mgz81.GetPreVote(ethCaller, ballot, group, fnumber)
	}

	if err == nil {
		storedCandidate = preVote.Candidate
		storedVoter = preVote.Voter
	}

	if candidate.Cmp(storedCandidate) == 0 && !allZeros20(storedVoter) {
		response = TRUE_STRING
	}

	fprint(w, response)
}

/**
 *  it converts a big.Int into a int
 */
func toInt(bigInt big.Int) int {
	strInt := bigInt.Text(10)
	smallInt, _ := strconv.Atoi(strInt)

	return smallInt
}

/**
 *  It converts a int into a big.Int
 */
func toBig(smallInt int) *big.Int {
	strInt := strconv.Itoa(smallInt)
	var bigInt = new(big.Int)
	bigInt, _ = bigInt.SetString(strInt, 10)

	return bigInt
}

/**
 * It removes the 0s, then converts the slice into a string
 */
func toAscii(b32 [32]byte) string {
	ret := ""
	for i := 0; i < 32; i++ {
		if b32[i] != 0 {
			ret += string(b32[i])
		}
	}
	return ret
}

/**
 * It checks whether the voter is null
 */
func allZeros20(voter [20]uint8) bool {
	for i := 0; i < 20; i++ {
		if voter[i] != 0 {
			return false
		}
	}
	return true
}

/**
 * To receive requests and provide info about the campaign
 */
func startCampaign(path string) {
	if nsStarted == false {
		nsStarted = true
		hashcodesPrefix = getParam(path, "?prefix=", "&ecdsa=")

		ecdsaPrivkeyStr := getParam(path, "&ecdsa=", "&rsa=")
		var errPK error
		ecdsaPrivkey, errPK = crypto.HexToECDSA(ecdsaPrivkeyStr)
		if errPK != nil {
			printline("Cannot generate the EDCSA private key", true, true)
		}

		loadRSAPrivPem(string(common.FromHex(getParam(path, "&rsa=", "&ca="))))
		certificateAuthority = getParam(path, "&ca=", "&campaign=")
		currentCampaignId = getParam(path, "&campaign=", "&regexp=")
		regularExpression = string(common.FromHex(getParam(path, "&regexp=", "&pwd=")))
		ethPwd = getParam(path, "&pwd=", "")

		//Initializing dbs for votes and cancellations
		initDatabases()

		printline("The campaign has started", false, false)
	} else {
		printline("The campaign has already started", false, false)
	}
}

/**
 * It extracts the RSA private key from the PEM and stores it in rsaPrivKey
 */
func loadRSAPrivPem(privPem string) {
	rsaPrivPem = privPem
	block, _ := pem.Decode([]byte(rsaPrivPem))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		printline("Failed to decode PEM block containing private key", false, false)
	} else {
		var err error
		priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			printline("Failed to decode PEM block containing public key", false, false)
		} else {
			var ok bool
			rsaPrivKey, ok = priv.(*rsa.PrivateKey)
			if !ok {
				printline("Failed to decode PEM block containing public key", false, false)
			}
		}
	}
}

/**
 * It gets data from the blockchain and stores it in the memory
 */
func getCampaignInfo() {
	if isMgzNull() || !nsStarted {
		return
	}

	campaignInfoJSON := new(CampaignInfo)

	switch currentContract.mgz {
	case "3":
		var err error
		creatorAddress, err = mgz3.Chairperson(ethCaller)
		if err != nil {
			return
		}
		campaignInfoJSON.Chairperson = creatorAddress.String()

		mgz, _ := mgz3.Mgz(ethCaller)
		campaignInfoJSON.Mgz = toInt(*mgz)

		rounds, _ := mgz3.Rounds(ethCaller)
		campaignInfoJSON.Rounds = toInt(*rounds)

		remainingRounds, _ := mgz3.RemainingRounds(ethCaller)
		campaignInfoJSON.RemainingRounds = toInt(*remainingRounds)

		currentBallot, _ := mgz3.CurrentBallot(ethCaller)
		campaignInfoJSON.CurrentBallot = toInt(*currentBallot)

		currentMessage, _ := mgz3.CurrentMessage(ethCaller)
		campaignInfoJSON.CurrentMessage = toAscii(currentMessage)

		howManyBallots, _ := mgz3.HowManyBallots(ethCaller)
		campaignInfoJSON.HowManyBallots = toInt(*howManyBallots)

		howManyGroups, _ := mgz3.HowManyGroups(ethCaller)
		campaignInfoJSON.HowManyGroups = toInt(*howManyGroups)

		canCancel, _ := mgz3.CanCancel(ethCaller)
		campaignInfoJSON.CanCancel = canCancel

		stoppingAccessionToGroups, _ := mgz3.StoppingAccessionToGroups(ethCaller)
		campaignInfoJSON.StoppingAccessionToGroups = toAscii(stoppingAccessionToGroups)

		disableCandidateLink, _ := mgz3.DisableCandidateLink(ethCaller)
		campaignInfoJSON.DisableCandidateLink = disableCandidateLink

		var chairpersonTor ChairpersonTor

		if !allZeros20(creatorAddress) {
			onion, _ := mgz3.GetTor(ethCaller, creatorAddress, toBig(0))
			chairpersonTor.Onion = toAscii(onion)
			chairpersonTor.PublicKey = getChairpersonPubkey(creatorAddress)
		}

		campaignInfoJSON.ChairpersonTor = chairpersonTor
		campaignInfoJSON.Ballots = make([]Ballot, campaignInfoJSON.HowManyBallots)
		campaignInfoJSON.VotesPerBallotCandidateCategory = make([][]string, campaignInfoJSON.HowManyBallots)
		groupCategories := getGroupCategories()
		for i := 0; i < campaignInfoJSON.HowManyBallots; i++ {
			ballot := new(Ballot)
			b, _ := mgz3.Ballots(ethCaller, toBig(i))
			ballot.Id = toAscii(b.Id)
			ballot.Closed = b.Closed

			howMany, _ := mgz3.HowManyCandidatesInBallot(ethCaller, toBig(i))
			ballot.HowManyCandidates = toInt(*howMany)
			ballot.Candidates = make([]Candidate, ballot.HowManyCandidates)

			campaignInfoJSON.VotesPerBallotCandidateCategory[i] = make([]string, toInt(*howMany))
			for j := 0; j < ballot.HowManyCandidates; j++ {
				var candidate Candidate
				c, _ := mgz3.GetCandidate(ethCaller, toBig(i), toBig(j))
				candidate.Website = common.Bytes2Hex(c.Website[:32])
				candidate.Website = strings.Replace(candidate.Website, "0x", "", -1)
				candidate.Count = toInt(*c.Count)
				ballot.Candidates[j] = candidate

				campaignInfoJSON.VotesPerBallotCandidateCategory[i][j] = "{"
				for k := 0; k < len(groupCategories); k++ {
					votes, _ := mgz3.GetVotesPerBallotCandidateCategory(ethCaller, toBig(i), toBig(j), toBig(k))
					campaignInfoJSON.VotesPerBallotCandidateCategory[i][j] +=
						"\"" + groupCategories[k] + "\":" + votes.String() + ","
				}
				campaignInfoJSON.VotesPerBallotCandidateCategory[i][j] =
					strings.TrimSuffix(campaignInfoJSON.VotesPerBallotCandidateCategory[i][j], ",") + "}"
			}
			campaignInfoJSON.Ballots[i] = *ballot
		}
		break
	case "27":
		var err error
		creatorAddress, err = mgz27.Chairperson(ethCaller)
		if err != nil {
			return
		}
		campaignInfoJSON.Chairperson = creatorAddress.String()

		mgz, _ := mgz27.Mgz(ethCaller)
		campaignInfoJSON.Mgz = toInt(*mgz)

		rounds, _ := mgz27.Rounds(ethCaller)
		campaignInfoJSON.Rounds = toInt(*rounds)

		remainingRounds, _ := mgz27.RemainingRounds(ethCaller)
		campaignInfoJSON.RemainingRounds = toInt(*remainingRounds)

		currentBallot, _ := mgz27.CurrentBallot(ethCaller)
		campaignInfoJSON.CurrentBallot = toInt(*currentBallot)

		currentMessage, _ := mgz27.CurrentMessage(ethCaller)
		campaignInfoJSON.CurrentMessage = toAscii(currentMessage)

		howManyBallots, _ := mgz27.HowManyBallots(ethCaller)
		campaignInfoJSON.HowManyBallots = toInt(*howManyBallots)

		howManyGroups, _ := mgz27.HowManyGroups(ethCaller)
		campaignInfoJSON.HowManyGroups = toInt(*howManyGroups)

		canCancel, _ := mgz27.CanCancel(ethCaller)
		campaignInfoJSON.CanCancel = canCancel

		stoppingAccessionToGroups, _ := mgz27.StoppingAccessionToGroups(ethCaller)
		campaignInfoJSON.StoppingAccessionToGroups = toAscii(stoppingAccessionToGroups)

		disableCandidateLink, _ := mgz27.DisableCandidateLink(ethCaller)
		campaignInfoJSON.DisableCandidateLink = disableCandidateLink

		var chairpersonTor ChairpersonTor

		if !allZeros20(creatorAddress) {
			onion, _ := mgz27.GetTor(ethCaller, creatorAddress, toBig(0))
			chairpersonTor.Onion = toAscii(onion)
			chairpersonTor.PublicKey = getChairpersonPubkey(creatorAddress)
		}

		campaignInfoJSON.ChairpersonTor = chairpersonTor
		campaignInfoJSON.Ballots = make([]Ballot, campaignInfoJSON.HowManyBallots)
		campaignInfoJSON.VotesPerBallotCandidateCategory = make([][]string, campaignInfoJSON.HowManyBallots)
		groupCategories := getGroupCategories()
		for i := 0; i < campaignInfoJSON.HowManyBallots; i++ {
			ballot := new(Ballot)
			b, _ := mgz27.Ballots(ethCaller, toBig(i))
			ballot.Id = toAscii(b.Id)
			ballot.Closed = b.Closed

			howMany, _ := mgz27.HowManyCandidatesInBallot(ethCaller, toBig(i))
			ballot.HowManyCandidates = toInt(*howMany)
			ballot.Candidates = make([]Candidate, ballot.HowManyCandidates)

			campaignInfoJSON.VotesPerBallotCandidateCategory[i] = make([]string, toInt(*howMany))
			for j := 0; j < ballot.HowManyCandidates; j++ {
				var candidate Candidate
				c, _ := mgz27.GetCandidate(ethCaller, toBig(i), toBig(j))
				candidate.Website = common.Bytes2Hex(c.Website[:32])
				candidate.Website = strings.Replace(candidate.Website, "0x", "", -1)
				candidate.Count = toInt(*c.Count)
				ballot.Candidates[j] = candidate

				campaignInfoJSON.VotesPerBallotCandidateCategory[i][j] = "{"
				for k := 0; k < len(groupCategories); k++ {
					votes, _ := mgz27.GetVotesPerBallotCandidateCategory(ethCaller, toBig(i), toBig(j), toBig(k))
					campaignInfoJSON.VotesPerBallotCandidateCategory[i][j] +=
						"\"" + groupCategories[k] + "\":" + votes.String() + ","
				}
				campaignInfoJSON.VotesPerBallotCandidateCategory[i][j] =
					strings.TrimSuffix(campaignInfoJSON.VotesPerBallotCandidateCategory[i][j], ",") + "}"
			}
			campaignInfoJSON.Ballots[i] = *ballot
		}
		break
	case "81":
		var err error
		creatorAddress, err = mgz81.Chairperson(ethCaller)
		if err != nil {
			return
		}
		campaignInfoJSON.Chairperson = creatorAddress.String()

		mgz, _ := mgz81.Mgz(ethCaller)
		campaignInfoJSON.Mgz = toInt(*mgz)

		rounds, _ := mgz81.Rounds(ethCaller)
		campaignInfoJSON.Rounds = toInt(*rounds)

		remainingRounds, _ := mgz81.RemainingRounds(ethCaller)
		campaignInfoJSON.RemainingRounds = toInt(*remainingRounds)

		currentBallot, _ := mgz81.CurrentBallot(ethCaller)
		campaignInfoJSON.CurrentBallot = toInt(*currentBallot)

		currentMessage, _ := mgz81.CurrentMessage(ethCaller)
		campaignInfoJSON.CurrentMessage = toAscii(currentMessage)

		howManyBallots, _ := mgz81.HowManyBallots(ethCaller)
		campaignInfoJSON.HowManyBallots = toInt(*howManyBallots)

		howManyGroups, _ := mgz81.HowManyGroups(ethCaller)
		campaignInfoJSON.HowManyGroups = toInt(*howManyGroups)

		canCancel, _ := mgz81.CanCancel(ethCaller)
		campaignInfoJSON.CanCancel = canCancel

		stoppingAccessionToGroups, _ := mgz81.StoppingAccessionToGroups(ethCaller)
		campaignInfoJSON.StoppingAccessionToGroups = toAscii(stoppingAccessionToGroups)

		disableCandidateLink, _ := mgz81.DisableCandidateLink(ethCaller)
		campaignInfoJSON.DisableCandidateLink = disableCandidateLink

		var chairpersonTor ChairpersonTor

		if !allZeros20(creatorAddress) {
			onion, _ := mgz81.GetTor(ethCaller, creatorAddress, toBig(0))
			chairpersonTor.Onion = toAscii(onion)
			chairpersonTor.PublicKey = getChairpersonPubkey(creatorAddress)
		}

		campaignInfoJSON.ChairpersonTor = chairpersonTor
		campaignInfoJSON.Ballots = make([]Ballot, campaignInfoJSON.HowManyBallots)
		campaignInfoJSON.VotesPerBallotCandidateCategory = make([][]string, campaignInfoJSON.HowManyBallots)
		groupCategories := getGroupCategories()
		for i := 0; i < campaignInfoJSON.HowManyBallots; i++ {
			ballot := new(Ballot)
			b, _ := mgz81.Ballots(ethCaller, toBig(i))
			ballot.Id = toAscii(b.Id)
			ballot.Closed = b.Closed

			howMany, _ := mgz81.HowManyCandidatesInBallot(ethCaller, toBig(i))
			ballot.HowManyCandidates = toInt(*howMany)
			ballot.Candidates = make([]Candidate, ballot.HowManyCandidates)

			campaignInfoJSON.VotesPerBallotCandidateCategory[i] = make([]string, toInt(*howMany))
			for j := 0; j < ballot.HowManyCandidates; j++ {
				var candidate Candidate
				c, _ := mgz81.GetCandidate(ethCaller, toBig(i), toBig(j))
				candidate.Website = common.Bytes2Hex(c.Website[:32])
				candidate.Website = strings.Replace(candidate.Website, "0x", "", -1)
				candidate.Count = toInt(*c.Count)
				ballot.Candidates[j] = candidate

				campaignInfoJSON.VotesPerBallotCandidateCategory[i][j] = "{"
				for k := 0; k < len(groupCategories); k++ {
					votes, _ := mgz81.GetVotesPerBallotCandidateCategory(ethCaller, toBig(i), toBig(j), toBig(k))
					campaignInfoJSON.VotesPerBallotCandidateCategory[i][j] +=
						"\"" + groupCategories[k] + "\":" + votes.String() + ","
				}
				campaignInfoJSON.VotesPerBallotCandidateCategory[i][j] =
					strings.TrimSuffix(campaignInfoJSON.VotesPerBallotCandidateCategory[i][j], ",") + "}"
			}
			campaignInfoJSON.Ballots[i] = *ballot
		}
	}

	campaignInfoJSON.Signature = ""
	cijBytesBeforeSignature, _ := json.Marshal(campaignInfoJSON)
	campaignInfoJSON.Signature = sign(cijBytesBeforeSignature)

	cijBytes, _ := json.Marshal(campaignInfoJSON)
	campaignInfo = string(cijBytes)
}

/**
 * Sign a message with the received ECDSA private key
 */
func sign(message []byte) string {
	messageStr := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + string(message)
	hash := crypto.Keccak256Hash([]byte(messageStr))
	signature, _ := crypto.Sign(hash.Bytes(), ecdsaPrivkey)
	sigHex := hexutil.Encode(signature)
	return sigHex
}

/**
 * It verifies whether a address has signed a message
 */
func verify(message string, signatureHex string, addressStr string) bool {
	if len(message) == 0 || len(signatureHex) == 0 || len(addressStr) == 0 {
		return false
	}

	signature, err0 := hexutil.Decode(signatureHex)
	if err0 != nil {
		return false
	}

	if signature[64] != 27 && signature[64] != 28 {
		return false
	}
	signature[64] -= 27

	messageStr := "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + string(message)
	data := []byte(messageStr)
	hash := crypto.Keccak256Hash(data)
	sigPublicKey, err1 := crypto.Ecrecover(hash.Bytes(), signature)
	if err1 != nil {
		return false
	}

	pub, err2 := crypto.UnmarshalPubkey(sigPublicKey)
	if err2 != nil {
		return false
	}

	sigAddress := crypto.PubkeyToAddress(*pub)
	sigAddressStr := sigAddress.String()

	return sigAddressStr == addressStr
}

/**
 * It verifies whether a address has signed a message
 */
func checkVoterRequirements(user string) bool {
	if len(regularExpression) > 1 {
		return strings.HasPrefix(regularExpression, "^") && strings.HasPrefix(user, regularExpression[1:]) ||
			strings.HasSuffix(regularExpression, "$") &&
				strings.HasSuffix(user, regularExpression[:len(regularExpression)-1])
	}
	return true
}

/**
 * It returns a list/slice of group categories
 */
func getGroupCategories() []string {
	if isMgzNull() || !nsStarted {
		return make([]string, 0)
	}

	howMany := new(big.Int)
	switch currentContract.mgz {
	case "3":
		howMany, _ = mgz3.HowManyGroupCategories(ethCaller)
		break
	case "27":
		howMany, _ = mgz27.HowManyGroupCategories(ethCaller)
		break
	case "81":
		howMany, _ = mgz81.HowManyGroupCategories(ethCaller)
	}
	howManyInt := toInt(*howMany)

	var gcs = make([]string, howManyInt)
	var gc [32]byte
	var err error
	gc[0] = 1
	for i := 0; i < howManyInt; i++ {
		switch currentContract.mgz {
		case "3":
			gc, err = mgz3.GroupCategories(ethCaller, toBig(i))
			gcs[i] = toAscii(gc)
			break
		case "27":
			gc, err = mgz27.GroupCategories(ethCaller, toBig(i))
			gcs[i] = toAscii(gc)
			break
		case "81":
			gc, err = mgz81.GroupCategories(ethCaller, toBig(i))
			gcs[i] = toAscii(gc)
		}
	}

	if err != nil {
		return make([]string, 0)
	}

	groupCategoriesArray = gcs
	return gcs
}

/**
 * It returns the RSA public key of a chairperson
 */
func getChairpersonPubkey(chairperson common.Address) string {
	if isMgzNull() || !nsStarted {
		return ""
	}

	var pubkey [10][32]byte
	var b32 [32]byte
	var err error
	b32[0] = 1
	for i := 0; i < 10; i++ {
		switch currentContract.mgz {
		case "3":
			b32, err = mgz3.GetTor(ethCaller, chairperson, toBig(i+1))
			pubkey[i] = b32
			break
		case "27":
			b32, err = mgz27.GetTor(ethCaller, chairperson, toBig(i+1))
			pubkey[i] = b32
			break
		case "81":
			b32, err = mgz81.GetTor(ethCaller, chairperson, toBig(i+1))
			pubkey[i] = b32
		}
	}

	if err != nil {
		return ""
	}

	pubkeyArray := make([]byte, 320)
	for j := 0; j < 20; j++ {
		for k := 0; k < 32; k++ {
			if (j*32)+k > 319 {
				break
			}
			pubkeyArray[(j*32)+k] = pubkey[j][k]
		}
	}

	return common.Bytes2Hex(pubkeyArray)
}

/**
 * It gets data on the blockchain and stores it in the memory
 */
func getCampaignIPFSInfo() {
	if isMgzNull() || !nsStarted {
		return
	}

	var b32 [32]byte
	var err error
	b32[0] = 1
	campaignIPFSInfo = ""
	for i := 0; !allZeros32(b32); i++ {
		switch currentContract.mgz {
		case "3":
			b32, err = mgz3.GetCampaignIpfsInfo(ethCaller, toBig(i))
			if err != nil {
				return
			}
			campaignIPFSInfo += toAscii(b32)
			break
		case "27":
			b32, err = mgz27.GetCampaignIpfsInfo(ethCaller, toBig(i))
			if err != nil {
				return
			}
			campaignIPFSInfo += toAscii(b32)
			break
		case "81":
			b32, err = mgz81.GetCampaignIpfsInfo(ethCaller, toBig(i))
			if err != nil {
				return
			}
			campaignIPFSInfo += toAscii(b32)
		}
	}
}

/**
 * It gets data on the blockchain (about a group) and stores it in the memory
 */
func getGroupsInfo() {
	if isMgzNull() || !nsStarted {
		return
	}

	var hmb *big.Int
	var err error
	switch currentContract.mgz {
	case "3":
		hmb, err = mgz3.HowManyBallots(ethCaller)
		break
	case "27":
		hmb, err = mgz27.HowManyBallots(ethCaller)
		break
	case "81":
		hmb, err = mgz81.HowManyBallots(ethCaller)
	}

	if err != nil {
		return
	}

	hmgInt := getHowManyGroups()
	mgzInt, _ := strconv.Atoi(currentContract.mgz)
	hmbInt := toInt(*hmb)

	groupsInfo = make([]string, hmgInt)

	var chairperson string
	var groupInfo GroupInfo
	for i := 0; i < hmgInt; i++ {
		if isMgzNull() || !nsStarted {
			return
		}

		switch currentContract.mgz {
		case "3":
			group, err := mgz3.Groups(ethCaller, toBig(i))
			if err != nil {
				return
			}
			chairperson = hexutil.Encode(group.CPerson[:])
			if chairperson == "0x0000000000000000000000000000000000000000" {
				continue
			}
			groupInfo.Chairperson = chairperson
			groupInfo.Category = getCategory(group.Category)
			groupInfo.Index = i
			groupInfo.Size = toInt(*group.Size)

			if !allZeros20(group.CPerson) {
				onion, _ := mgz3.GetTor(ethCaller, group.CPerson, toBig(0))
				groupInfo.ChairpersonTor.Onion = toAscii(onion)
				groupInfo.ChairpersonTor.PublicKey = getChairpersonPubkey(group.CPerson)
			}

			isEmptyGroup := true
			groupInfo.Voters = make([]string, mgzInt)
			groupInfo.Pubkeys.Keys = make([]string, mgzInt)
			groupInfo.Pubkeys.Prefixes = make([]int, mgzInt)
			voters, _ := mgz3.GetGroupVoters(ethCaller, toBig(i))
			prefixes, pubkeys, _ := mgz3.GetGroupPubkeys(ethCaller, toBig(i))
			for j := 0; j < mgzInt; j++ {
				groupInfo.Voters[j] = voters[j].String()
				groupInfo.Pubkeys.Prefixes[j] = toInt(*prefixes[j])
				groupInfo.Pubkeys.Keys[j] = common.Bytes2Hex(pubkeys[j][:])
				if groupInfo.Pubkeys.Prefixes[j] != 0 {
					isEmptyGroup = false
				}
			}

			groupInfo.Ballots = make([]GroupBallot, hmbInt)
			for k := 0; k < hmbInt; k++ {
				groupInfo.Ballots[k].Committed = make([]bool, mgzInt)
				groupInfo.Ballots[k].Candidates = make([]int, mgzInt)
				groupInfo.Ballots[k].Numbers = make([]string, mgzInt)
				if !isEmptyGroup {
					fnumbers, candidates, err := mgz3.GetVotes(ethCaller, toBig(k), toBig(i))
					if err != nil {
						return
					}
					for l := 0; l < mgzInt; l++ {
						groupInfo.Ballots[k].Candidates[l] = toInt(*candidates[l])
						groupInfo.Ballots[k].Numbers[l] = common.Bytes2Hex(fnumbers[l][:])
						c, _ := mgz3.Committed(ethCaller, toBig(k), toBig(i), toBig(l))
						groupInfo.Ballots[k].Committed[l] = c
					}
				}
			}
			break
		case "27":
			group, err := mgz27.Groups(ethCaller, toBig(i))
			if err != nil {
				return
			}
			chairperson = hexutil.Encode(group.CPerson[:])
			if chairperson == "0x0000000000000000000000000000000000000000" {
				continue
			}
			groupInfo.Chairperson = chairperson
			groupInfo.Category = getCategory(group.Category)
			groupInfo.Index = i
			groupInfo.Size = toInt(*group.Size)

			if !allZeros20(group.CPerson) {
				onion, _ := mgz27.GetTor(ethCaller, group.CPerson, toBig(0))
				groupInfo.ChairpersonTor.Onion = toAscii(onion)
				groupInfo.ChairpersonTor.PublicKey = getChairpersonPubkey(group.CPerson)
			}

			isEmptyGroup := true
			groupInfo.Voters = make([]string, mgzInt)
			groupInfo.Pubkeys.Keys = make([]string, mgzInt)
			groupInfo.Pubkeys.Prefixes = make([]int, mgzInt)
			voters, _ := mgz27.GetGroupVoters(ethCaller, toBig(i))
			prefixes, pubkeys, _ := mgz27.GetGroupPubkeys(ethCaller, toBig(i))
			for j := 0; j < mgzInt; j++ {
				groupInfo.Voters[j] = voters[j].String()
				groupInfo.Pubkeys.Prefixes[j] = toInt(*prefixes[j])
				groupInfo.Pubkeys.Keys[j] = common.Bytes2Hex(pubkeys[j][:])
				if groupInfo.Pubkeys.Prefixes[j] != 0 {
					isEmptyGroup = false
				}
			}

			groupInfo.Ballots = make([]GroupBallot, hmbInt)
			for k := 0; k < hmbInt; k++ {
				groupInfo.Ballots[k].Committed = make([]bool, mgzInt)
				groupInfo.Ballots[k].Candidates = make([]int, mgzInt)
				groupInfo.Ballots[k].Numbers = make([]string, mgzInt)
				if !isEmptyGroup {
					fnumbers, candidates, err := mgz27.GetVotes(ethCaller, toBig(k), toBig(i))
					if err != nil {
						return
					}
					for l := 0; l < mgzInt; l++ {
						groupInfo.Ballots[k].Candidates[l] = toInt(*candidates[l])
						groupInfo.Ballots[k].Numbers[l] = common.Bytes2Hex(fnumbers[l][:])
						c, _ := mgz27.Committed(ethCaller, toBig(k), toBig(i), toBig(l))
						groupInfo.Ballots[k].Committed[l] = c
					}
				}
			}

			break
		case "81":
			group, err := mgz81.Groups(ethCaller, toBig(i))
			if err != nil {
				return
			}
			chairperson = hexutil.Encode(group.CPerson[:])
			if chairperson == "0x0000000000000000000000000000000000000000" {
				continue
			}
			groupInfo.Chairperson = chairperson
			groupInfo.Category = getCategory(group.Category)
			groupInfo.Index = i
			groupInfo.Size = toInt(*group.Size)

			if !allZeros20(group.CPerson) {
				onion, _ := mgz81.GetTor(ethCaller, group.CPerson, toBig(0))
				groupInfo.ChairpersonTor.Onion = toAscii(onion)
				groupInfo.ChairpersonTor.PublicKey = getChairpersonPubkey(group.CPerson)
			}

			isEmptyGroup := true
			groupInfo.Voters = make([]string, mgzInt)
			groupInfo.Pubkeys.Keys = make([]string, mgzInt)
			groupInfo.Pubkeys.Prefixes = make([]int, mgzInt)
			voters, _ := mgz81.GetGroupVoters(ethCaller, toBig(i))
			prefixes, pubkeys, _ := mgz81.GetGroupPubkeys(ethCaller, toBig(i))
			for j := 0; j < mgzInt; j++ {
				groupInfo.Voters[j] = voters[j].String()
				groupInfo.Pubkeys.Prefixes[j] = toInt(*prefixes[j])
				groupInfo.Pubkeys.Keys[j] = common.Bytes2Hex(pubkeys[j][:])
				if groupInfo.Pubkeys.Prefixes[j] != 0 {
					isEmptyGroup = false
				}
			}

			groupInfo.Ballots = make([]GroupBallot, hmbInt)
			for k := 0; k < hmbInt; k++ {
				groupInfo.Ballots[k].Committed = make([]bool, mgzInt)
				groupInfo.Ballots[k].Candidates = make([]int, mgzInt)
				groupInfo.Ballots[k].Numbers = make([]string, mgzInt)
				if !isEmptyGroup {
					fnumbers, candidates, err := mgz81.GetVotes(ethCaller, toBig(k), toBig(i))
					if err != nil {
						return
					}
					for l := 0; l < mgzInt; l++ {
						groupInfo.Ballots[k].Candidates[l] = toInt(*candidates[l])
						groupInfo.Ballots[k].Numbers[l] = common.Bytes2Hex(fnumbers[l][:])
						c, _ := mgz81.Committed(ethCaller, toBig(k), toBig(i), toBig(l))
						groupInfo.Ballots[k].Committed[l] = c
					}
				}
			}
		}

		groupInfo.Signature = ""
		giBytesBeforeSignature, _ := json.Marshal(groupInfo)
		groupInfo.Signature = sign(giBytesBeforeSignature)

		giBytes, _ := json.Marshal(groupInfo)
		groupsInfo[i] = string(giBytes)
	}
}

/**
 * It checks whether the bytes32 variable is null
 */
func allZeros32(b32 [32]uint8) bool {
	for i := 0; i < 32; i++ {
		if b32[i] != 0 {
			return false
		}
	}
	return true
}

/**
 * It returns the information of certain group
 */
func getGroupInfo(path string, w http.ResponseWriter) {
	if !verifyParams(path, "?index=") {
		fprint(w, ERROR_STRING)
		return
	}

	indexStr := getParam(path, "?index=", "")

	if len(indexStr) == 0 {
		fprint(w, ERROR_STRING)
	}

	index, err := strconv.Atoi(indexStr)

	if err != nil || index >= len(groupsInfo) {
		fprint(w, ERROR_STRING)
	} else if index >= 0 {
		fprint(w, groupsInfo[index])
	} else {
		fprint(w, ERROR_STRING)
	}
}

/**
 * It returns the information of the Campaign
 */
func printCampaignInfo(w http.ResponseWriter) {
	if campaignInfo == "" {
		fprint(w, ERROR_STRING)
	} else {
		fprint(w, campaignInfo)
	}
}

/**
 * It returns the Campaign's IPFS data
 */
func printCampaignIPFSInfo(w http.ResponseWriter) {
	if campaignIPFSInfo == "" {
		fprint(w, ERROR_STRING)
	} else {
		fprint(w, campaignIPFSInfo)
	}
}

/**
 * It verifies a signed vote and returns whether the signature is valid or not
 */
func verifyURS(message string, pubkeys string, signature string) bool {
	response := Execute(message, "", pubkeys, "", signature, false)
	return response == TRUE_STRING
}

/**
 * Terminating the program
 */
func exit() {
	stopRecevingRequests = true
	nsStarted = false

	if torCmd != nil && torCmd.Process != nil {
		torCmd.Process.Kill()
		printline("Tor killed", false, false)
	}

	if gethCmd != nil && gethCmd.Process != nil {
		args := "exit"
		io.WriteString(stdinGeth, args)
		gethCmd.CombinedOutput()

		go func() {
			stdinGeth.Close()
			time.Sleep(5 * time.Second)
			gethCmd.Process.Kill()

			//The last statements
			printline("Geth killed", false, false)
		}()
	}

	theCaptcha = nil

	if votesDB != nil {
		votesDB.Close()
	}
	if cancellationsDB != nil {
		cancellationsDB.Close()
	}

	time.Sleep(6 * time.Second)
	printline("Ciao", false, false)

	os.Exit(0)
}

/**
 * Print logs in a new line
 */
func printline(str string, start bool, end bool) {
	if start {
		fmt.Println("")
	}
	fmt.Println(str)
	if end {
		fmt.Println("")
	}
}

/**
 * It verifies whether all these parameters are present in the path string, and in the given order
 */
func verifyParams(path string, params ...string) bool {
	for i, param := range params {
		if strings.LastIndex(path, param) == -1 {
			return false
		}
		if i > 0 && strings.LastIndex(path, params[i-1]) > strings.LastIndex(path, param) {
			return false
		}
	}
	return true
}

/**
 * It is used in the verbose mode
 */
func printCommandResponse(outstr string) {
	printline("----------------------------------Command response ----------------------------------",
		true, false)
	printline(outstr, false, false)
	printline("------------------------------End of command response--------------------------------",
		false, true)
}

/**
 * Used to extract a parameter from a url
 */
func getParam(path, start, end string) string {
	ret := ""
	if end == "" {
		ret = path[strings.LastIndex(path, start)+len(start):]
	} else {
		ret = path[strings.LastIndex(path, start) +
			len(start) : strings.LastIndex(path, end)]
	}

	var err error
	ret, err = url.QueryUnescape(ret)

	if err != nil {
		return ""
	}
	return ret
}

/**
 * Print with the fmt.Fprint if w is not nil
 */
func fprint(w http.ResponseWriter, str string) {
	if w != nil {
		fmt.Fprint(w, str)
	}
}

/************************** AES tools ******************************/

func deriveKey(passphrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha256.New), salt
}

func AESEncrypt(passphrase, plaintext string) string {
	key, salt := deriveKey(passphrase, nil)
	iv := make([]byte, 12)
	rand.Read(iv)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data := aesgcm.Seal(nil, iv, []byte(plaintext), nil)
	return hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(data)
}

func AESDecrypt(passphrase, ciphertext string) string {
	arr := strings.Split(ciphertext, "-")

	if len(arr) < 3 {
		return ""
	}

	salt, _ := hex.DecodeString(arr[0])
	iv, _ := hex.DecodeString(arr[1])
	data, _ := hex.DecodeString(arr[2])
	key, _ := deriveKey(passphrase, salt)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data, _ = aesgcm.Open(nil, iv, data, nil)

	return string(data)
}

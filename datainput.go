// Copyright 2014 The Monero Developers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"strconv"
)

// readLength is the length of a hash compression read in bytes.
const readLength = 65536 - 1

// CmpPubKey compares two pubkeys and returns true if they are the same, else
// false. WARNING: Assumes the curves are equivalent!
func CmpPubKey(i, j *ecdsa.PublicKey) bool {
	if i.X.Cmp(j.X) != 0 {
		return false
	}

	if i.Y.Cmp(j.Y) != 0 {
		return false
	}

	return true
}

// keyInKeyRing checks if a pubkey exists in the keyring.
func (kr *PublicKeyRing) keyInKeyRing(k *ecdsa.PublicKey) bool {
	for _, key := range kr.Ring {
		if CmpPubKey(k, &key) {
			return true
		}
	}
	return false
}

// ReadKeyRing reads a key ring of public keys from a file in JSON object format,
// and also inserts the pubkey of a keypair if it's not already present (handles
// bug in URS implementation).
func ReadKeyRing(filestring string, kp *ecdsa.PrivateKey) (*PublicKeyRing, error) {

	//Replacing the file with a string, in order to make things easier later
	var f = []byte(filestring)

	// Unmarshall the loaded file into a map
	var keyMap = make(map[string]string)

	var err = json.Unmarshal(f, &keyMap)
	if err != nil {
		str := fmt.Sprintf("json error: Couldn't unmarshall keyring file.", err)
		jsonError := errors.New(str)
		return nil, jsonError
	}

	kr := NewPublicKeyRing(uint(len(keyMap)))

	// Stick the pubkeys into the keyring as long as it doesn't belong to the
	// keypair given.
	for i := 0; i < len(keyMap); i++ {
		pkBytes, errDecode := hex.DecodeString(keyMap[strconv.Itoa(i)])
		if errDecode != nil {
			decodeError := errors.New("decode error: Couldn't decode hex.")
			return nil, decodeError
		}

		pubkey, errParse := btcec.ParsePubKey(pkBytes, btcec.S256())
		if errParse != nil {
			return nil, errParse
		}

		ecdsaPubkey := ecdsa.PublicKey{pubkey.Curve, pubkey.X, pubkey.Y}

		if kp == nil || !CmpPubKey(&kp.PublicKey, &ecdsaPubkey) {
			kr.Add(ecdsaPubkey)
		} else {
			kr.Add(kp.PublicKey)
		}
	}

	// Stick the keypair in if it's missing.
	if kp != nil && !kr.keyInKeyRing(&kp.PublicKey) {
		kr.Add(kp.PublicKey)
	}

	return kr, nil
}

// ReadKeyPair reads an ECDSA keypair a file in JSON object format.
// Example JSON input:
//  {
//    "privkey": "..."
//  }
// It also checks if a pubkey is in the keyring and, if not, appends it to the
// keyring.
func ReadKeyPair(filestring string) (*ecdsa.PrivateKey, error) {

	//Replacing the file with a string, in order to make things easier later
	var f = []byte(filestring)

	// Unmarshall the loaded file into a map.
	var keyMap = make(map[string]string)
	var pubkey *ecdsa.PublicKey
	var privkey *ecdsa.PrivateKey

	var err = json.Unmarshal(f, &keyMap)
	if err != nil {
		jsonError := errors.New("json error: Couldn't unmarshall keypair file.")
		return nil, jsonError
	}

	privBytes, errDecode := hex.DecodeString(keyMap["privkey"])
	if errDecode != nil {
		decodeError := errors.New("decode error: Couldn't decode hex for privkey.")
		return nil, decodeError
	}

	// PrivKeyFromBytes doesn't return an error, so this could possibly be ugly.
	privkeyBtcec, _ := btcec.PrivKeyFromBytes(btcec.S256(), privBytes)

	pubBytes, errDecode := hex.DecodeString(keyMap["pubkey"])
	if errDecode != nil {
		decodeError := errors.New("decode error: Couldn't decode hex for privkey.")
		return nil, decodeError
	}

	pubkeyBtcec, errParse := btcec.ParsePubKey(pubBytes, btcec.S256())
	if errParse != nil {
		return nil, errParse
	}

	// Assign the things to return
	pubkey = &ecdsa.PublicKey{pubkeyBtcec.Curve,
		pubkeyBtcec.X,
		pubkeyBtcec.Y}

	privkey = &ecdsa.PrivateKey{*pubkey, privkeyBtcec.D}

	return privkey, nil
}

// StripTextFile removes line breaks from a text file to ensure that signatures
// of text files function correctly cross platform.
func StripTextFile(b []byte) []byte {
	input := bytes.NewBuffer(b)
	var output []byte

	rScanner := bufio.NewScanner(input)
	rScanner.Split(bufio.ScanLines)

	for rScanner.Scan() {
		str := rScanner.Text()
		bstr := []byte(str)
		output = append(output, bstr...)
	}

	return output
}

// GetTextFileData retrieves the []byte encoding of a text file from a path.
func GetTextFileData(filestring string) ([]byte, error) {
	//Replacing the file with a string, in order to make things easier later
	var f = []byte(filestring)

	return StripTextFile(f), nil
}

// GetSigFileData retrieves the []byte encoding of a signature file from a path.
func GetSigFileData(filestring string) ([]byte, error) {
	//Replacing the file with a string, in order to make things easier later
	var f = []byte(filestring)

	return StripTextFile(f), nil
}

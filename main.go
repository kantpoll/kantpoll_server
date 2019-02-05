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
	"flag"
	"os"
	"strings"
	"time"
)

/******************** Global variables ********************/
//Flags
var verbose bool
var creator bool
var observer bool

//Geth and IPFS folders
const gethFolderWinDefault = "geth-windows-amd64-1.8.20-24d727b6"
const gethFolderLinDefault = "geth-linux-amd64-1.8.20-24d727b6"

var gethFolderWin = gethFolderWinDefault
var gethFolderLin = gethFolderLinDefault

/******************** Functions ********************/

func main() {
	//Loading the flags
	verboseFlag := flag.Bool("verbose", false, "Activate the verbose mode")
	observerFlag := flag.Bool("observer", false, "Run as observer")
	contractFlag := flag.String("contract", "", "Insert the contract address")
	mgzFlag := flag.String("mgz", "", "Insert the maximum group size")
	gethFolderWinFlag := flag.String("gethFolderWin", "", "Insert the name of the Geth folder on Windows OS")
	gethFolderLinFlag := flag.String("gethFolderLin", "", "Insert the name of the Geth folder on Linux OS")
	helpFlag := flag.Bool("help", false, "Display the help")
	flag.Parse()

	if *helpFlag {
		printline("Command-line options:", false, false)
		flag.PrintDefaults()
		return
	}

	verbose = *verboseFlag
	observer = *observerFlag
	creator = !observer

	if len(*gethFolderWinFlag) > 0 {
		gethFolderWin = *gethFolderWinFlag
	}
	if len(*gethFolderLinFlag) > 0 {
		gethFolderLin = *gethFolderLinFlag
	}

	currentContract.address = *contractFlag
	currentContract.mgz = *mgzFlag

	//It creates the website directory and installs Tor, Electron, Geth, and IPFS.
	if !installPrograms() {
		printline("Installation error", false, false)
		time.Sleep(5 * time.Second)
		return
	}

	if creator {
		go func() {
			startServer()
		}()
		waitTheExit()
	} else if observer {
		go func() {
			startObserver()
		}()
		waitTheExit()
	}
}

/**
 * Wait for the user to type exit
 */
func waitTheExit() {
	reader := bufio.NewReader(os.Stdin)
	var sentence string

	for sentence != "exit" {
		sentenceBytes, _, _ := reader.ReadLine()
		sentence = string(sentenceBytes)
		strings.Replace(sentence, "\n", "", -1)
		if sentence == "exit" {
			exit()
		}
		time.Sleep(1 * time.Second)
	}
}

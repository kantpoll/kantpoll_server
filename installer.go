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
	"github.com/mholt/archiver"
	"os"
	"runtime"
	"time"
)

/******************** Global variables ********************/

/**
 * Create the website directory and install tor, geth, and ipfs.
 */
func installPrograms() bool {
	installSuccess := true

	if _, err := os.Stat("temp"); os.IsNotExist(err) {
		os.Mkdir("temp", 0770)
	}

	if _, err := os.Stat("dbs"); os.IsNotExist(err) {
		os.Mkdir("dbs", 0770)
	}

	if _, err := os.Stat(getHome() + "/.kantpoll"); os.IsNotExist(err) {
		os.Mkdir(getHome()+"/.kantpoll", 0700)
	}

	//Standalone installation
	if runtime.GOOS == "windows" {
		if _, err := os.Stat("geth.zip"); err == nil {
			os.Rename("geth.zip", "temp/geth.zip")
		}
		if _, err := os.Stat("ipfs.zip"); err == nil {
			os.Rename("ipfs.zip", "temp/ipfs.zip")
		}
		if _, err := os.Stat("tor.zip"); err == nil {
			os.Rename("tor.zip", "temp/tor.zip")
		}
	}

	//Checking if these directories exist
	_, errGeth := os.Stat("geth")
	_, errIpfs := os.Stat("ipfs")
	_, errTor := os.Stat("tor")
	_, errWebsite := os.Stat("website")

	if os.IsNotExist(errWebsite) {
		if _, err := os.Stat("website.zip"); err == nil {
			err2 := archiver.Zip.Open("website.zip", "temp")
			if err2 != nil {
				printline("Error unzipping website.zip", true, true)
				installSuccess = false
			} else {
				os.Rename("temp/website", "./website")
			}
		} else {
			printline("File website.zip not found", true, true)
			installSuccess = false
		}
	}

	if os.IsNotExist(errGeth) || os.IsNotExist(errIpfs) || os.IsNotExist(errTor) {
		printline("╔----------------------------------╗", true, false)
		printline("|           Installation           |", false, false)
		printline("╚----------------------------------╝", false, true)

		if os.IsNotExist(errGeth) {
			os.Mkdir("geth", 0770)
			printline("Installing Geth...", false, false)
			installSuccess = installSuccess && installGeth()
		}
		if os.IsNotExist(errIpfs) {
			os.Mkdir("ipfs", 0770)
			printline("Installing IPFS...", false, false)
			installSuccess = installSuccess && installIPFS()
		}
		if os.IsNotExist(errTor) {
			os.Mkdir("tor", 0770)
			printline("Installing Tor...", false, false)
			installSuccess = installSuccess && installTor()
		}

		if installSuccess {
			printline("╔----------------------------------╗ ", false, false)
			printline("|      Installation completed      | ", false, false)
			printline("╚----------------------------------╝ ", false, true)
		}
	}

	return installSuccess
}

/**
 * Download Geth and then save the file on the due directory
 */
func installGeth() bool {
	fileStr := ""

	if runtime.GOOS == "windows" {
		fileStr = "temp/geth.zip"
	} else if runtime.GOOS == "linux" {
		fileStr = "temp/geth.tar.gz"
	}

	var err error
	if runtime.GOOS == "windows" {
		err = archiver.Zip.Open(fileStr, "geth")
	} else if runtime.GOOS == "linux" {
		err = archiver.TarGz.Open(fileStr, "geth")
	}

	if err == nil {
		printline("Done.", false, false)
		if runtime.GOOS == "windows" {
			//Running geth for the first time in order that the firewall asks for permission
			go func() {
				duration := 15 * time.Second
				time.Sleep(duration)
				firstTimeGeth()
			}()
		}
		return true
	} else {
		printline("Failed to install Geth. Cannot unzip the file or the file does not exist.", true, true)
	}
	return false
}

/**
 * Download IPFS and then save the file on the due directory
 */
func installIPFS() bool {
	fileStr := ""

	if runtime.GOOS == "windows" {
		fileStr = "temp/ipfs.zip"
	} else if runtime.GOOS == "linux" {
		fileStr = "temp/ipfs.tar.gz"
	}

	var err4 error
	if runtime.GOOS == "windows" {
		err4 = archiver.Zip.Open(fileStr, "ipfs")
	} else if runtime.GOOS == "linux" {
		err4 = archiver.TarGz.Open(fileStr, "ipfs")
	}

	if err4 == nil {
		printline("Done.", false, false)
		return true
	} else {
		printline("Failed to install IPFS. Cannot unzip the file or the file does not exist.", true, true)
	}

	return false
}

/**
 * Download Tor and then save the file on the due directory
 * Also create a new torrc file
 */
func installTor() bool {
	fileStr := ""

	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
		fileStr = "temp/tor.zip"
	}

	//The executable fileStr should be placed with the main program
	var err4, err5 error
	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
		err4 = archiver.Zip.Open(fileStr, "tor")
		if _, err := os.Stat(getHome() + "/.kantpoll/tor"); os.IsNotExist(err) {
			err5 = archiver.Zip.Open(fileStr, getHome() + string(os.PathSeparator) + ".kantpoll" +
				string(os.PathSeparator) + "tor")
		}
	}

	err6 := configTor()

	if err4 == nil && err5 == nil && err6 == nil {
		printline("Done.", false, false)
		return true
	} else {
		printline("Failed to install Tor.", true, true)
		if err4 != nil{
			printline(err4.Error(), true, true)
		} else if err5 != nil{
			printline(err5.Error(), true, true)
		} else if err6 != nil{
			printline(err6.Error(), true, true)
		}
	}
	return false
}

/**
 * Configuring the hidden service
 */
func configTor() error {
	//The Data dir should be placed in the Home dir
	//to avoid being requested for admin privileges
	os.Mkdir(getHome()+"/.kantpoll/tor/Data/HiddenService/", 0700)

	torrcStr := ""
	if runtime.GOOS == "windows" {
		//Creating the torrc file
		torrcStr = "# Tor plus kantpoll hidden service \r\n" +
			"DataDirectory " + getHome() + "\\.kantpoll\\tor\\Data\\Tor \r\n" +
			"GeoIPFile " + getHome() + "\\.kantpoll\\tor\\Data\\Tor\\geoip \r\n" +
			"GeoIPv6File " + getHome() + "\\.kantpoll\\tor\\Data\\Tor\\geoip6 \r\n" +
			"HiddenServiceDir " + getHome() + "\\.kantpoll\\tor\\Data\\HiddenService \r\n" +
			"HiddenServicePort 80 127.0.0.1:1988 \r\n"
	} else if runtime.GOOS == "linux" {
		torrcStr = "# Tor plus kantpoll hidden service \r\n" +
			"DataDirectory " + getHome() + "/.kantpoll/tor/Data/Tor \r\n" +
			"GeoIPFile " + getHome() + "/.kantpoll/tor/Data/Tor/geoip \r\n" +
			"GeoIPv6File " + getHome() + "/.kantpoll/tor/Data/Tor/geoip6 \r\n" +
			"HiddenServiceDir " + getHome() + "/.kantpoll/tor/Data/HiddenService \r\n" +
			"HiddenServicePort 80 127.0.0.1:1988 \r\n"
	}

	f, _ := os.Create(getHome() + "/.kantpoll/tor/Data/Tor/torrc")
	defer f.Close()
	_, err := f.WriteString(torrcStr)

	return err
}

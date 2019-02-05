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
	crand "crypto/rand"
	"strconv"
)

//The main function was converted to the function execute in order to be browserified
func Execute(vts string, sts string, krs string, kps string, sigs string, blinds bool) string {
	// verify a text file
	if vts != "" && !blinds {
		if krs == "" {
			return "Error: verify text called but no keyring file given."

		}
		if sigs == "" {
			return "Error: verify text called but no signature file given."

		}

		m, err := GetTextFileData(vts)
		if err != nil {
			return "Text verification failed on message load."
		}

		kr, err := ReadKeyRing(krs, nil)
		if err != nil {
			return "Text verification failed on keyring load."
		}

		sig, err := GetSigFileData(sigs)
		if err != nil {
			return "Text verification failed on sig load."
		}

		decodedSig := &RingSign{nil, nil, nil, nil}

		err = decodedSig.FromBase58(string(sig))
		if err != nil {
			return "Text verification failed on sig parse."
		}

		isValid := Verify(kr, m, decodedSig)

		return strconv.FormatBool(isValid)

	}

	// sign a text file
	if sts != "" && !blinds {
		if krs == "" {
			return "Error: sign text called but no keyring file given."
		}
		if kps == "" {
			return "Error: sign text called but no keypair file given."
		}

		m, err := GetTextFileData(sts)
		if err != nil {
			return "Text signing failed on message load."
		}

		kp, err := ReadKeyPair(kps)
		if err != nil {
			return "Text signing failed on keypair load."
		}

		kr, err := ReadKeyRing(krs, kp)
		if err != nil {
			return "Text signing failed on keyring load."
		}

		ringsig, err := Sign(crand.Reader, kp, kr, m)
		if err != nil {
			return "Text signing failed on sig generation."
		}

		if Verify(kr, m, ringsig) {
			return ringsig.ToBase58()
		} else {
			return "Signature generated but failed to validate."
		}
	}

	// blind verify a text file
	if vts != "" && blinds {
		if krs == "" {
			return "Error: verify text called but no keyring file given."
		}
		if sigs == "" {
			return "Error: verify text called but no signature file given."
		}

		m, err := GetTextFileData(vts)
		if err != nil {
			return "Text verification failed on message load."
		}

		kr, err := ReadKeyRing(krs, nil)
		if err != nil {
			return "Text verification failed on keyring load."
		}

		sig, err := GetSigFileData(sigs)
		if err != nil {
			return "Text verification failed on sig load."
		}

		decodedSig := &BlindRingSign{nil, nil, nil, nil, nil, nil, nil, nil}

		err = decodedSig.FromBase58(string(sig))
		if err != nil {
			return "Text verification failed on sig parse."
		}

		isValid := BlindVerify(kr, m, decodedSig)

		return strconv.FormatBool(isValid)
	}

	// blind sign a text file
	if sts != "" && blinds {
		if krs == "" {
			return "Error: sign text called but no keyring file given."
		}
		if kps == "" {
			return "Error: sign text called but no keypair file given."
		}

		m, err := GetTextFileData(sts)
		if err != nil {
			return "Text signing failed on message load"
		}

		kp, err := ReadKeyPair(kps)
		if err != nil {
			return "Text signing failed on keypair load."
		}

		kr, err := ReadKeyRing(krs, kp)
		if err != nil {
			return "Text signing failed on keyring load."
		}

		ringsig, err := BlindSign(crand.Reader, kp, kr, m)
		if err != nil {
			return "Text signing failed on sig generation."
		}

		if BlindVerify(kr, m, ringsig) {
			return ringsig.ToBase58()
		} else {
			return "Signature generated but failed to validate."
		}
	}

	//New example text
	return "Function Execute receives the arguments: verify-text, sign-text, keyring, keypair, signature, blind."
}

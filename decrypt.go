/*
 * Copyright (c) 2025 - Nathanne Isip
 * This file is part of Shitkey.
 *
 * Shitkey is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published
 * by the Free Software Foundation, either version 3 of the License,
 * or (at your option) any later version.
 *
 * Shitkey is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Shitkey. If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/scrypt"
)

func runDecrypt(inputPath string) {
	printBanner()

	if !strings.HasSuffix(inputPath, skFileExt) {
		printError("Input file must have a '.sk' extension.")
		os.Exit(1)
	}

	outputPath := strings.TrimSuffix(inputPath, skFileExt)
	printPrompt("Enter password:")
	password := readPassword()

	printInfo(fmt.Sprintf(
		"Decrypting '%s' to '%s'...",
		colorBold(inputPath),
		colorBold(outputPath),
	))

	encryptedData, err := os.ReadFile(inputPath)
	if err != nil {
		printError("Error reading encrypted file: " + err.Error())
		os.Exit(1)
	}

	if len(encryptedData) < saltSize+nonceSize {
		printError("File is too short to be a valid encrypted file.")
		os.Exit(1)
	}

	salt := encryptedData[:saltSize]
	nonce := encryptedData[saltSize : saltSize+nonceSize]
	ciphertext := encryptedData[saltSize+nonceSize:]

	key, err := scrypt.Key(password, salt, 32768, 8, 1, keySize)
	if err != nil {
		printError("Error deriving key: " + err.Error())
		os.Exit(1)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		printError("Error creating cipher: " + err.Error())
		os.Exit(1)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		printError("Error creating GCM: " + err.Error())
		os.Exit(1)
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		printError(
			"Decryption failed. Incorrect password or corrupted file: " +
				err.Error(),
		)
		os.Exit(1)
	}

	if err := os.WriteFile(
		outputPath,
		plaintext,
		0644,
	); err != nil {
		printError("Error writing output file: " + err.Error())
		os.Exit(1)
	}

	printSuccess(
		"Decryption successful! File saved as '" +
			colorBold(outputPath) +
			"'.",
	)
}

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
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/scrypt"
)

func runEncrypt(inputPath string) {
	printBanner()
	outputPath := inputPath + skFileExt

	printPrompt("Enter password:")
	password := readPassword()

	printInfo(fmt.Sprintf(
		"Encrypting '%s' to '%s'...",
		colorBold(inputPath),
		colorBold(outputPath),
	))

	salt := make([]byte, saltSize)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		printError("Error generating salt: " + err.Error())
		os.Exit(1)
	}

	key, err := scrypt.Key(
		password,
		salt,
		32768,
		8, 1,
		keySize,
	)

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

	inputFile, err := os.ReadFile(inputPath)
	if err != nil {
		printError("Error reading input file: " + err.Error())
		os.Exit(1)
	}

	nonce := make([]byte, nonceSize)
	if _, err := io.ReadFull(
		rand.Reader,
		nonce,
	); err != nil {
		printError("Error generating nonce: " + err.Error())
		os.Exit(1)
	}

	ciphertext := gcm.Seal(
		nil,
		nonce,
		inputFile,
		nil,
	)

	outputFile, err := os.Create(outputPath)
	if err != nil {
		printError("Error creating output file: " + err.Error())
		os.Exit(1)
	}
	defer outputFile.Close()

	if _, err := outputFile.Write(salt); err != nil {
		printError("Error writing salt to file: " + err.Error())
		os.Exit(1)
	}

	if _, err := outputFile.Write(nonce); err != nil {
		printError("Error writing nonce to file: " + err.Error())
		os.Exit(1)
	}

	if _, err := outputFile.Write(ciphertext); err != nil {
		printError(
			"Error writing ciphertext to file: " +
				err.Error(),
		)
		os.Exit(1)
	}

	printSuccess(
		"Encryption successful! File saved as '" +
			colorBold(outputPath) +
			"'.",
	)
}

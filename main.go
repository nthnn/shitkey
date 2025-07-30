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
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "encrypt":
		if len(os.Args) != 3 {
			printUsage()
			printError(
				"Error: The encrypt command requires a filename.",
			)

			os.Exit(1)
		}

		filename := os.Args[2]
		runEncrypt(filename)

	case "decrypt":
		if len(os.Args) != 3 {
			printUsage()
			printError(
				"Error: The decrypt command requires a filename.",
			)

			os.Exit(1)
		}

		filename := os.Args[2]
		runDecrypt(filename)

	case "version":
		printBanner()
		fmt.Println("Version: " + colorBold(version))

	default:
		printUsage()
		printError("Invalid command.")
		os.Exit(1)
	}
}

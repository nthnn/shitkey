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
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func readPassword() []byte {
	fd := int(syscall.Stdin)
	bytePassword, err := terminal.ReadPassword(fd)
	fmt.Println()
	if err != nil {
		printError("Error reading password: " + err.Error())
		os.Exit(1)
	}
	return bytePassword
}

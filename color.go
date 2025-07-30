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

import "fmt"

func colorBold(s string) string {
	return fmt.Sprintf("\x1b[1m%s\x1b[0m", s)
}

func colorRed(s string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", s)
}

func colorGreen(s string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", s)
}

func colorYellow(s string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", s)
}

func colorCyan(s string) string {
	return fmt.Sprintf("\x1b[36m%s\x1b[0m", s)
}

func colorGray(s string) string {
	return fmt.Sprintf("\x1b[90m%s\x1b[0m", s)
}

func printError(msg string) {
	fmt.Println(colorRed(colorBold("Error: ")) + msg)
}

func printSuccess(msg string) {
	fmt.Println(colorGreen(colorBold("Success: ")) + msg)
}

func printInfo(msg string) {
	fmt.Println(colorYellow(colorBold("Info: ")) + msg)
}

func printPrompt(msg string) {
	fmt.Print(colorCyan(colorBold(msg + " ")))
}

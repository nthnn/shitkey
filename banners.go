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
	"strings"
)

func printBanner() {
	banner := "Shitkey - File Encryptor"
	startColor := [3]int{120, 0, 180}
	endColor := [3]int{0, 150, 255}

	var gradientBuilder strings.Builder
	for i, char := range banner {
		r := startColor[0] + (endColor[0]-startColor[0])*i/len(banner)
		g := startColor[1] + (endColor[1]-startColor[1])*i/len(banner)
		b := startColor[2] + (endColor[2]-startColor[2])*i/len(banner)

		gradientBuilder.WriteString(
			fmt.Sprintf("\x1b[38;2;%d;%d;%dm%c", r, g, b, char),
		)
	}

	gradientBuilder.WriteString("\x1b[0m")

	fmt.Println(gradientBuilder.String())
	fmt.Println(colorGray(strings.Repeat("-", len(banner))))
}

func printUsage() {
	printBanner()
	fmt.Println(colorCyan("Usage:"))
	fmt.Println(colorGreen("  " + colorBold("shitkey encrypt") + " <filename>"))
	fmt.Println(colorGreen("  " + colorBold("shitkey decrypt") + " <filename.sk>"))
	fmt.Println(colorGreen("  " + colorBold("shitkey version")))

	fmt.Println()
	fmt.Println(colorYellow("Encrypts or decrypts a file using a password."))
}

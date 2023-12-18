package util

import "fmt"

func PrintBytes(b []byte) {
	// print bytes with default line width (16)

	PrintBytesCustom(b, 16)
}

func PrintBytesCustom(bytes []byte, lineWidth int) {
	// print bytes with custom line width (16, 32, etc.)

	for i := 0; i < len(bytes); i += lineWidth {
		line := ""

		// add line number as hex with 4 padding
		line += fmt.Sprintf("%04x  ", i)

		// add hex values
		for j := 0; j < lineWidth; j++ {
			if i+j < len(bytes) {
				line += fmt.Sprintf("%02x ", bytes[i+j])
			} else {
				line += "   "
			}
		}

		// add ascii values

		line += " |"

		for j := 0; j < lineWidth; j++ {
			if i+j < len(bytes) {
				if bytes[i+j] >= 32 && bytes[i+j] <= 126 {
					line += fmt.Sprintf("%c", bytes[i+j])
				} else {
					line += "."
				}
			} else {
				line += " "
			}
		}

		line += "|"

		fmt.Println(line)
	}
}

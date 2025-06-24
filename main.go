package main

import (
	"bufio"
	"color/helpers"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if !(len(os.Args) >= 3 && len(os.Args) <= 4) {
		fmt.Println("Usage: ascii-art \"Your text here (use \\n for a vertical break)\"")
		return
	}
	color := os.Args[1]
	if strings.HasPrefix(color, "--color=") {
		color = strings.TrimPrefix(color, "--color=")
	} else {
		fmt.Println("Usage: ascii-art --color=<color> ")
		return
	}
	colorCode := helpers.GetColorCode(color)
	if colorCode == "" {
		fmt.Printf("Error: invalid color %q\n", color)
		return
	}
	reset := helpers.Reset
	target := ""
	raw := ""

	if len(os.Args) == 4 {
		target = os.Args[2]
		raw = os.Args[3]

	} else {
		target = ""
		raw = os.Args[2]
	}

	// Split on literal "\\n"
	segments := helpers.SplitSegments(raw)

	// If input is only newlines, print one blank per delimiter
	allEmpty := true
	for _, seg := range segments {
		if seg != "" {
			allEmpty = false
			break
		}
	}
	if allEmpty {
		for i := 0; i < len(segments)-1; i++ {
			fmt.Println()
		}
		return
	}

	// Load ASCII font
	file, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fontLines []string
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		fontLines = append(fontLines, scan.Text())
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	//Validate characters
	for _, seg := range segments {
		for _, r := range seg {
			idx := int(r) - 32
			if idx < 0 || idx >= len(fontLines)/9 {
				fmt.Printf("Error: invalid ASCII character %q (code %d)\n", r, r)
				os.Exit(1)
			}
		}
	}

	// Render segments and newline runs
	i := 0
	for i < len(segments) {
		if seg := segments[i]; seg != "" {
			// Render 8-line banner
			for row := 0; row < 8; row++ {
				var parts []string
				for _, r := range seg {
					idx := int(r) - 32
					if target == "" {
						colored := colorCode + fontLines[idx*9+(row+1)] + reset
						parts = append(parts, colored)
					} else {
						if strings.ContainsAny(target, string(r)) {
							colored := colorCode + fontLines[idx*9+(row+1)] + reset
							parts = append(parts, colored)
						} else {
							colored := fontLines[idx*9+(row+1)]
							parts = append(parts, colored)
						}
					}

				}
				fmt.Println(strings.Join(parts, ""))
			}
			i++

		} else {
			// Count consecutive empty segments

			j := i
			for j < len(segments) && segments[j] == "" {
				j++
			}
			count := j - i
			// Print one blank line per literal "\\n"
			for k := 0; k < count; k++ {
				fmt.Println()
			}
			i = j

		}
	}
}

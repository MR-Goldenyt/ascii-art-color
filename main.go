package main

import (
	"ascii-art/helpers"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ascii-art \"Your text here (use \\n for a vertical break)\"")
		os.Exit(1)
	}
	raw := os.Args[1]

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

	// Validate characters
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
					parts = append(parts, fontLines[idx*9+(row+1)])
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

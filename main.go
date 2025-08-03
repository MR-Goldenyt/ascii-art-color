package main

import (
	"fmt"
	"os"

	"color/helpers"
)

func main() {
	argsLength := len(os.Args) - 1
	if argsLength < 1 || argsLength > 4 {
		fmt.Println("Usage: ./ascii-art-color [STRING]" +
			"\nOR" +
			"\nUsage: ./ascii-art-color [STRING] [BANNER]" +
			"\nOR" +
			"\nUsage: ./ascii-art-color --color=<color> [STRING]" +
			"\nOR" +
			"\nUsage: ./ascii-art-color --color=<color> [STRING] [BANNER]" +
			"\nOR" +
			"\nUsage: ./ascii-art-color --color=<color> <substring to be colored> [STRING]" +
			"\nOR" +
			"\nUsage: ./ascii-art-color --color=<color> <substring to be colored> [STRING] [BANNER]")
		return
	}

	if argsLength == 1 {
		fontLines := helpers.LoadBanner("standard")
		segments := helpers.SplitSegments(os.Args[1])
		helpers.ValidateString(segments, fontLines)
		target := ""
		colorCode := ""
		helpers.RenderArt(segments, target, colorCode, fontLines)

	}
	// 	color := os.Args[1]
	// 	if strings.HasPrefix(color, "--color=") {
	// 		color = strings.TrimPrefix(color, "--color=")
	// 	} else {
	// 		color = ""
	// 	}
	// 	if color != "" {
	// 	}
	// 	colorCode := helpers.GetColorCode(color)
	// 	reset := helpers.Reset
	// 	target := ""
	// 	raw := ""

	// 	if len(os.Args) == 4 {
	// 		target = os.Args[2]
	// 		raw = os.Args[3]

	// 	} else {
	// 		target = ""
	// 		raw = os.Args[2]
	// 	}

	// 	// Split on literal "\\n"
	// 	segments := helpers.SplitSegments(raw)

	// 	// If input is only newlines, print one blank per delimiter
	// 	allEmpty := true
	// 	for _, seg := range segments {
	// 		if seg != "" {
	// 			allEmpty = false
	// 			break
	// 		}
	// 	}
	// 	if allEmpty {
	// 		for i := 0; i < len(segments)-1; i++ {
	// 			fmt.Println()
	// 		}
	// 		return
	// 	}

	// 	fontLines := helpers.LoadBanner()

	// 	// Render segments and newline runs
	// 	i := 0
	// 	for i < len(segments) {
	// 		if seg := segments[i]; seg != "" {
	// 			for row := 0; row < 8; row++ {
	// 				var parts []string
	// 				for pos := 0; pos < len(seg); pos++ {
	// 					r := seg[pos]
	// 					idx := int(r) - 32

	// 					if target != "" && len(seg)-pos >= len(target) && seg[pos:pos+len(target)] == target {
	// 						for j := 0; j < len(target); j++ {
	// 							ri := int(seg[pos+j]) - 32
	// 							colored := colorCode + fontLines[ri*9+(row+1)] + reset
	// 							parts = append(parts, colored)
	// 						}
	// 						pos += len(target) - 1
	// 					} else if target == "" {
	// 						colored := colorCode + fontLines[idx*9+(row+1)] + "\033[0m"
	// 						parts = append(parts, colored)
	// 					} else {
	// 						colored := fontLines[idx*9+(row+1)]
	// 						parts = append(parts, colored)
	// 					}
	// 				}
	// 				fmt.Println(strings.Join(parts, ""))
	// 			}
	// 			i++

	// 		} else {
	// 			// Count consecutive empty segments

	// 			j := i
	// 			for j < len(segments) && segments[j] == "" {
	// 				j++
	// 			}
	// 			count := j - i
	// 			// Print one blank line per literal "\\n"
	// 			for k := 0; k < count; k++ {
	// 				fmt.Println()
	// 			}
	// 			i = j

	//		}
	//	}
}

package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-color/helpers"
)

func main() {
	argsLength := len(os.Args) - 1
	if argsLength < 1 || argsLength > 4 {
		helpers.UsageMessage()
		return
	}

	colorCode := ""
	target := ""
	found := false
	var segments []string
	var fontLines []string

	if argsLength == 1 {
		fontLines = helpers.LoadBanner("standard")
		segments = helpers.SplitSegments(os.Args[1])
		helpers.ValidateString(segments, fontLines)
		helpers.RenderArt(segments, target, colorCode, fontLines)

	} else if argsLength == 2 {

		if strings.HasPrefix(os.Args[1], "--color=") {
			found = true
			colorCode = strings.TrimPrefix(os.Args[1], "--color=")
			colorCode = helpers.GetColorCode(colorCode)
		}

		if found {
			if colorCode == "" {
				fmt.Printf("Error: invalid color %q\n", strings.TrimPrefix(os.Args[1], "--color="))
				return
			}
			segments = helpers.SplitSegments(os.Args[2])
			fontLines = helpers.LoadBanner("standard")
			helpers.ValidateString(segments, fontLines)
		} else {
			if !(strings.ToLower(os.Args[2]) == "standard" || strings.ToLower(os.Args[2]) == "shadow" || strings.ToLower(os.Args[2]) == "thinkertoy") {
				fmt.Printf("Error: invalid Banner font %q\n", os.Args[2])
				return
			}
			segments = helpers.SplitSegments(os.Args[1])
			fontLines = helpers.LoadBanner(strings.ToLower(os.Args[2]))
			helpers.ValidateString(segments, fontLines)
		}
		helpers.RenderArt(segments, target, colorCode, fontLines)

	} else if argsLength == 3 {
		if strings.HasPrefix(os.Args[1], "--color=") {
			found = true
			colorCode = strings.TrimPrefix(os.Args[1], "--color=")
			colorCode = helpers.GetColorCode(colorCode)
		}

		if found {
			if colorCode == "" {
				fmt.Printf("Error: invalid color %q\n", strings.TrimPrefix(os.Args[1], "--color="))
				return
			}
			if !(strings.ToLower(os.Args[3]) == "standard" || strings.ToLower(os.Args[3]) == "shadow" || strings.ToLower(os.Args[3]) == "thinkertoy") {
				target = os.Args[2]
				segments = helpers.SplitSegments(os.Args[3])
				fontLines = helpers.LoadBanner("standard")
				helpers.ValidateString(segments, fontLines)
			} else {
				segments = helpers.SplitSegments(os.Args[2])
				fontLines = helpers.LoadBanner(strings.ToLower(os.Args[3]))
				helpers.ValidateString(segments, fontLines)
			}
		} else {
			helpers.UsageMessage()
			return
		}
		helpers.RenderArt(segments, target, colorCode, fontLines)
	} else if argsLength == 4 {
		if strings.HasPrefix(os.Args[1], "--color=") {
			found = true
			colorCode = strings.TrimPrefix(os.Args[1], "--color=")
			colorCode = helpers.GetColorCode(colorCode)
		}

		if found {
			if colorCode == "" {
				fmt.Printf("Error: invalid color %q\n", strings.TrimPrefix(os.Args[1], "--color="))
				return
			}
			if strings.ToLower(os.Args[4]) == "standard" || strings.ToLower(os.Args[4]) == "shadow" || strings.ToLower(os.Args[4]) == "thinkertoy" {
				target = os.Args[2]
				segments = helpers.SplitSegments(os.Args[3])
				fontLines = helpers.LoadBanner(strings.ToLower(os.Args[4]))
				helpers.ValidateString(segments, fontLines)
			} else {
				fmt.Printf("Error: invalid Banner font %q\n", os.Args[4])
				return
			}
		} else {
			helpers.UsageMessage()
			return
		}
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

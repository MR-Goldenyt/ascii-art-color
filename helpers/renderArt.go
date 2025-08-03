package helpers

import (
	"fmt"
	"strings"
)

func RenderArt(segments []string, target string, colorCode string, fontLines []string) {
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

	// Render segments and newline runs
	i := 0
	for i < len(segments) {
		if seg := segments[i]; seg != "" {
			for row := 0; row < 8; row++ {
				var parts []string
				for pos := 0; pos < len(seg); pos++ {
					r := seg[pos]
					idx := int(r) - 32

					if target != "" && len(seg)-pos >= len(target) && seg[pos:pos+len(target)] == target {
						for j := 0; j < len(target); j++ {
							ri := int(seg[pos+j]) - 32
							colored := colorCode + fontLines[ri*9+(row+1)] + Reset
							parts = append(parts, colored)
						}
						pos += len(target) - 1
					} else if target == "" {
						colored := colorCode + fontLines[idx*9+(row+1)] + "\033[0m"
						parts = append(parts, colored)
					} else {
						colored := fontLines[idx*9+(row+1)]
						parts = append(parts, colored)
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

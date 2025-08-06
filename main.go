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
}

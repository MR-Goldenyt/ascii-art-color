package helpers

import "fmt"

func UsageMessage() {
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
}

package helpers

import (
	"fmt"
	"os"
)

func ValidateString(segments []string, fontLines []string) {
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
}

package helpers

import (
	"strings"
)

// splitSegments chops a string like "Hello\nWorld"
// into ["Hello", "", "World"]
func SplitSegments(s string) []string {

	return strings.Split(s, "\\n")

}

package helpers

import (
	"strings"
)

// splitSegments chops a string like "Hello\nWorld"
// into ["Hello", "", "World"]
func SplitSegments(s string) []string {

	if !strings.ContainsAny(s, "\\\n"){
		return strings.Split(s, "\\n")
	}
	arr := []string{}
	return append(arr, s)

}

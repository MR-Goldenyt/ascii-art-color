package helpers


var ColorCodes = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	
}

const Reset = "\033[0m"

func GetColorCode(color string) string {
	if code, ok := ColorCodes[color]; ok {
		return code
	}
	return "" // default to no color
}

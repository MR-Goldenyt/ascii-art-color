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
	"orange" : "\033[38;5;208m",
	"purple":  "\033[38;5;129m",
	"pink":    "\033[38;5;213m",
	"gray":    "\033[90m",
	"lightred": "\033[91m",
	"lightgreen": "\033[92m",
	"lightyellow": "\033[93m",
	"lightblue": "\033[94m",
	"lightmagenta": "\033[95m",
	"lightcyan": "\033[96m",
	"lightwhite": "\033[97m",
	"brightblack": "\033[90m",
	"brightred": "\033[91m",
	"brightgreen": "\033[92m",
	"brightyellow": "\033[93m",
	"brightblue": "\033[94m",
	"brightmagenta": "\033[95m",
	"brightcyan": "\033[96m",
	"brightwhite": "\033[97m",
	"reset":   "\033[0m",
	"bold":    "\033[1m",
	"underline": "\033[4m",
	"blink":   "\033[5m",
	"reverse": "\033[7m",
	"hidden":  "\033[8m",
	"strikethrough": "\033[9m",
	"dim":     "\033[2m",
	"italic":  "\033[3m",
	"bright":  "\033[1m",
	"faint":   "\033[2m",
	"normal":  "\033[22m",
	"noitalic": "\033[23m",
	"nobold":  "\033[22m",
	"nounderline": "\033[24m",
	"noblink": "\033[25m",
	"noreverse": "\033[27m",
	"nohidden": "\033[28m",
	"nostrikethrough": "\033[29m",
	"nobright": "\033[22m",
	"nofaint": "\033[22m",
	
}

const Reset = "\033[0m"

func GetColorCode(color string) string {
	if code, ok := ColorCodes[color]; ok {
		return code
	}
	return "" // default to no color
}

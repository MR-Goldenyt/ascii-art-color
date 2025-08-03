package helpers

import (
	"bufio"
	"log"
	"os"
)

func LoadBanner(bannerFont string) []string {
	// Load ASCII font
	file, err := os.Open("banner/" + bannerFont + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fontLines []string
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		fontLines = append(fontLines, scan.Text())
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	return fontLines
}

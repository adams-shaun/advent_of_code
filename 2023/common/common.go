package common

import (
	"log"
	"os"
	"strings"
)

// DRY stuff
func ReadInput(infile string) []string {
	data, err := os.ReadFile(infile)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}

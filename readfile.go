package asciiart

import (
	"log"
	"os"
	"strings"
)

// ReadFile reads a given file and returns the body as a slice of string
func ReadFile() []string {
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	return strings.Split(string(data), "\r\n")
}

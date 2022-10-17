package asciiart

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func flag() (flag string, fileName string) {
	for i, v := range os.Args[1] {
		if string(v) == "=" {
			return os.Args[1][:i], os.Args[1][i+1:]
		}
	}
	return "", os.Args[1]
}

// ReadFile reads a given file and returns the body as a slice of string
// TO DO: make it always ready standard.txt as well
func ReadFile() ([]string, string) {
	flag, fileName := flag()
	fmt.Println(flag, fileName)
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	return strings.Split(string(data), "\r\n"), flag
}

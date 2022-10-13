package main

import (
	"asciiart"
	"os"
)

func main() {
	args := os.Args[1]
	f := asciiart.ReadFile()
	asciiart.PrintAscii(f, args)
}

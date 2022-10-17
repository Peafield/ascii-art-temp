package main

import (
	"asciiart"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Incorrect number of arguments given")
		os.Exit(1)
	}
	f := asciiart.ReadFile()
	asciiart.PrintAscii(f, args[1])
}

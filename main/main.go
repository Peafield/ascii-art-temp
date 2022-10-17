package main

import (
	"asciiart"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Incorrect number of arguments given")
		os.Exit(1)
	}
	input, flag := asciiart.ReadFile()
	standard, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	if flag == "--reverse" {
		asciiart.Reverse(input, strings.Split(string(standard), "\r\n"))
	} else {
		asciiart.PrintAscii(input, args[1])
	}

}

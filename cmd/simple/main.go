package main

import (
	"fmt"
	"os"

	"github.com/xyproto/bin"
)

func main() {
	filename := os.Args[0]
	isBinary, err := bin.BinaryFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s is a binary file: %v\n", filename, isBinary)
}

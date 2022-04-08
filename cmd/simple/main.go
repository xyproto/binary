package main

import (
	"fmt"
	"os"

	"github.com/xyproto/binary"
)

func main() {
	filename := os.Args[0]
	isBinary, err := binary.BinaryFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s is a binary file: %v\n", filename, isBinary)
}

package main

import (
	"bufio"
	"bytes"
)

func main() {
	// Create a buffer with some data
	data := []byte("Hello\n,\n World!")
	buffer := bytes.NewBuffer(data)

	// Call the getScanner function
	scanner := getScanner(buffer)

	// Verify that the returned reader is not nil
	if scanner == nil {
		panic("Expected non-nil reader, got nil")
	}
	for scanner.Scan() {
		println(scanner.Text())
	}
}

func getScanner(b *bytes.Buffer) *bufio.Scanner {
	scanner := bufio.NewScanner(b)
	return scanner
}

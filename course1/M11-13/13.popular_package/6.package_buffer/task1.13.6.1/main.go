package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	// Create a buffer for testing
	buffer := bytes.NewBufferString("Hello, World!")
	b := make([]byte, 13)
	r := getReader(buffer)
	r.Read(b)
	fmt.Println(string(b)) // Hello, World!
}

func getReader(b *bytes.Buffer) *bufio.Reader {
	reader := bufio.NewReader(b)
	return reader
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	// Create a new buffer
	buffer := bytes.NewBufferString("Hello, World!")

	// Call the getDataString function
	result := getDataString(buffer)

	// Check if the result matches the expected output
	expected := "Hello, World!"
	if result != expected {
		panic(fmt.Sprintf("Expected %s, but got %s", expected, result))
	}

}

func getDataString(b *bytes.Buffer) string {
	reader := bufio.NewReader(b)
	text, _ := reader.ReadString('\n')
	return text
}

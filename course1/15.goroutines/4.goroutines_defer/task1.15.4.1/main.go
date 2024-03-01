package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a, b := 8, 13
	fmt.Println(*testDefer(&a, &b))
}

func testDefer(a, b *int) *int {
	var c int
	defer func() {

	}()
	c = sum(*a, *b)
	return &c
}

func sum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func writeToFile(f *os.File, s string) error {
	writer := bufio.NewWriter(f)
	_, err := writer.Write([]byte(s))
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	defer f.Close()
	return err
}

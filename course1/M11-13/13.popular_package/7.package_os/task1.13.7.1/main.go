package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	err := WriteFile("/path/to/file.txt", []byte("Hello, World!"), os.FileMode(0644))
	if err != nil {
		// Обработка ошибки
	}
	fmt.Println(WriteFile("/path/to/file.txt", []byte("Hello, World!"), os.FileMode(0644)))
}

func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	// Ваш код здесь
	s := strings.Split(filePath, "/")
	r := make([]string, len(s)-1)
	for i := 0; i < len(r); i++ {
		r[i] = s[i]
	}
	str := strings.Join(r, "/") + "/"
	//mas := strings.Split(str, "")
	//mass := make([]string, len(mas)-1)
	//for i := 0; i < len(mass); i++ {
	//	mass[i] = mas[i+1]
	//}
	//res := strings.Join(mass, "")
	err := os.MkdirAll(str, perm)
	if err != nil {
		return err
	}
	os.Chdir(str)
	file, _ := os.Create(s[len(s)-1])
	writer := bufio.NewWriter(file)
	writer.Write(data)
	writer.Flush()
	return nil
}

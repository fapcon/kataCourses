package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	filePath := "course1/13.popular_package/7.package_os/task1.13.7.2/test_file.txt"

	err := WriteFile(filePath, strings.NewReader("Hello, World!"), os.Stdout)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
	}
}

func WriteFile(filePath string, data io.Reader, fd io.Writer) error {
	s := strings.Split(filePath, "/")
	r := make([]string, len(s)-1)
	for i := 0; i < len(r); i++ {
		r[i] = s[i]
	}
	str := strings.Join(r, "/") + "/"
	os.Chdir(str)
	f, err := os.Create(s[len(s)-1])
	if err != nil {
		return err
	}
	_, err = os.Open(s[len(s)-1])
	if err != nil {
		return err
	}

	//reader := bufio.NewReader(data)
	//w := bufio.NewWriter(fd)
	//buf := make([]byte, 1024)

	_, err = io.Copy(f, data)
	if err != nil {
		return err
	}
	return nil
}

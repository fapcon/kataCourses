package main

import (
	"io/ioutil"
)

func main() {

}

func ReadString(filePath string) string {
	// Ваш код для чтения файла и возврата его содержимого в виде строки
	fContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(fContent)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Write JSON data to file
	data := []map[string]interface{}{
		{
			"name": "Elliot",
			"age":  25,
		},
		{
			"name": "Fraser",
			"age":  30,
		},
	}
	err := writeJSON("users.json", data)
	if err != nil {
		panic(err)
	}
}

func writeJSON(filePath string, data interface{}) error {
	var fileName string
	s := strings.Split(filePath, "/")
	if len(s) == 1 {
		fileName = filePath
	} else {
		r := make([]string, len(s)-1)
		for i := 0; i < len(r); i++ {
			r[i] = s[i]
		}
		str := "/" + strings.Join(r, "/")
		fmt.Println(str)
		fmt.Println(s[len(s)-1])
		err := os.MkdirAll(str, 0644)
		if err != nil {
			return err
		}
		fileName = s[len(s)-1]
		os.Chdir(str)
	}
	_, err := os.Create(fileName)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

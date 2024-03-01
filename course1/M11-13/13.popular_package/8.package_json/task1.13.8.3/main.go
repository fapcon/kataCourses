package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}
type Comment struct {
	Text string `json:"text"`
}

func main() {
	user := make([]User, 3)
	filePath := "dir1/dir2/test_file.json"
	fmt.Println(writeJSON(filePath, user))
}

func writeJSON(filePath string, data []User) error {
	/*
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {

				var fileName string
				s := strings.Split(filePath, "/")
				if len(s) == 1 {
					fileName = filePath
				} else {
					r := make([]string, len(s)-1)
					for i := 0; i < len(r); i++ {
						r[i] = s[i]
					}
					str := strings.Join(r, "/")
					fmt.Println(str)
					fmt.Println(s[len(s)-1])
					err := os.MkdirAll(str, 0644)
					if err != nil {
						return err
					}
					fileName = s[len(s)-1]
					os.Chdir(fileName)
				}
				_, err := os.Create(fileName)
				if err != nil {
					return err
				}

		}
			var fileName string
			s := strings.Split(filePath, "/")
			if len(s) == 1 {
				fileName = filePath
			} else {
				r := make([]string, len(s)-1)
				for i := 0; i < len(r); i++ {
					r[i] = s[i]
				}
				str := strings.Join(r, "/")
				fmt.Println(str)
				fmt.Println(s[len(s)-1])
				err := os.MkdirAll(str, 0644)
				if err != nil {
					return err
				}
				fileName = s[len(s)-1]
				os.Chdir(str)
			}

	*/
	/*
		f := filepath.Dir(filePath)
		err := os.MkdirAll(f, 0644)
		if err != nil {
			return err
		}

	*/
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil && os.IsExist(err) {
		err.Error()
	}

	file, err := os.Create(filePath)
	if err != nil {
		err.Error()
	}
	defer file.Close()
	jsonData, err := json.Marshal(data)
	if err != nil {
		err.Error()
	}
	_, err = file.Write(jsonData)
	if err != nil {
		err.Error()
	}
	return nil
}

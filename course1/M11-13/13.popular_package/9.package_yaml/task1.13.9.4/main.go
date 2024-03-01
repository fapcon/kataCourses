package main

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func main() {

}
func writeYAML(filePath string, data interface{}) error {
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil && os.IsExist(err) {
		err.Error()
	}

	file, err := os.Create(filePath)
	if err != nil {
		err.Error()
	}
	defer file.Close()
	YAMLData, err := yaml.Marshal(data)
	if err != nil {
		err.Error()
	}
	_, err = file.Write(YAMLData)
	if err != nil {
		err.Error()
	}
	return nil
}
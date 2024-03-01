package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printTree(path string, prefix string, isLast bool, depth int) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, file := range files {
		fmt.Print(prefix)
		if isLast {
			fmt.Print("└── ")
			prefix += "    "
		} else {
			fmt.Print("├── ")
			prefix += "│   "
		}
		fmt.Println(file.Name())

		if file.IsDir() && depth > 0 {
			printTree(filepath.Join(path, file.Name()), prefix, i == len(files)-1, depth-1)
		}
		prefix = strings.TrimSuffix(prefix, "│   ")
		prefix = strings.TrimSuffix(prefix, "    ")
	}
}

func main() {
	var path string
	var depth int
	flag.IntVar(&depth, "n", 2, "depth of filepath")
	flag.Parse()
	flag.StringVar(&path, "path", "/usr/local/go/pkg/tes", "file path")
	flag.Parse()
	// получение флага
	if !strings.HasPrefix(path, "/") {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path = filepath.Join(wd, path)
	}
	printTree(path, "", false, depth)
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func run() ([]string, error) {
	searchDir := "l:/projetos/web"

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		fmt.Println(file)
	}

	return fileList, nil
}

func main() {
	run()
}

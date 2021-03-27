package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		i := strings.Index(path, "node_modules")
		if i > -1 {
			*files = append(*files, path)
		}
		return nil
	}
}

func main() {
	var files []string

	root := "L:\\Projetos\\web"
	err := filepath.Walk(root, visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}

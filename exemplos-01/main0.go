package main

import (
	"fmt"
	"path/filepath"
)

// func visit(files *[]string) filepath.WalkFunc {
// 	return func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		*files = append(*files, path)
// 		return nil
// 	}
// }

func main() {
	var files []string

	root := "L:\\Projetos\\web\\frontend-erp-golang"
	err := filepath.Walk(root, visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}

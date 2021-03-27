package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

func run() ([]string, error) {
	searchDir := "/home/farnetani/cursos/nextjs"

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(pathx string, f os.FileInfo, err error) error {
		if f.IsDir() {
			if path.Base(f.Name()) == "node_modules" {
				nm := regexp.MustCompile("node_modules")
				matches := nm.FindAllStringIndex(pathx, -1)
				if (len(matches)) == 1 {
					fileList = append(fileList, pathx)
				}
			}
		}
		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		fmt.Println(file)
		err := os.RemoveAll(file)
		if err != nil {
			fmt.Println(err)
		}

		// os.RemoveAll(file)
	}

	return fileList, nil
}

func main() {
	run()
}

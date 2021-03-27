package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := "L:\\Projetos\\web\\frontend-erp-golang"
	// file, err := os.Open(path)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	// names, err := file.Readdirnames(0)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(names)

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			i := strings.Index(path, "node_modules")
			if i > -1 {
				fmt.Println(path, info.Size())
				return nil
			} else {
				// fmt.Println(path, info.Size())
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

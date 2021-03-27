package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// dirname := "L:\\Projetos\\web\\frontend-erp-golang"
	dirname := "L:\\Projetos\\web"

	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// fmt.Println(file.Name())
		f2, err := os.Open(dirname + "\\" + file.Name())
		files2, err := f2.Readdir(-1)
		f2.Close()

		if err != nil {
			log.Fatal(err)
		}

		for _, file2 := range files2 {
			if file2.Name() == "node_modules" {
				fmt.Println(dirname + "\\" + file.Name() + "\\" + file2.Name())
			}
		}
	}
}

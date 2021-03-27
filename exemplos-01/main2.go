package main

import (
	"fmt"
	"os"
)

func main() {
	path := "L:\\Projetos\\web\\frontend-erp-golang"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	names, err := file.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(names)
}

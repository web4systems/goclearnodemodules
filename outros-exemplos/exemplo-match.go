package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	matches, _ := filepath.Glob("l:/projetos/web/**/**/*.json")
	for _, match := range matches {
		fmt.Println(match)
	}
}

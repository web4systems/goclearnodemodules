package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	err := ClearDir("tmp")
	if err != nil {
		fmt.Println(err)
	}
}

func ClearDir(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

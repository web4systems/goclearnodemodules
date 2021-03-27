package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	err := ClearDir("tmp")
	if err != nil {
		fmt.Println(err)
	}

}
func ClearDir(dir string) error {
	names, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entery := range names {
		os.RemoveAll(path.Join([]string{dir, entery.Name()}...))
	}
	return nil
}

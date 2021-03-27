package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := "L:\\Projetos\\web\\frontend-erp-golang"
	// files, err := WalkMatch(path, "*.js")
	files, err := WalkMatchDirectory(path, "*.js")
	if err != nil {
		fmt.Println("erro")
	}

	for _, f := range files {
		fmt.Println(f)
	}

}

func WalkMatchDirectory(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			i := strings.Index(path, "node_modules")
			if i > -1 {
				fmt.Println("achou" + path)
				matches = append(matches, path)
				return nil
			}
			// return nil
			return nil
		}
		// if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
		// 	return err
		// } else if matched {
		// 	// matches = append(matches, path)
		// }
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

package main

import (
	"fmt"
	"math"
	"os"
	"path"
	"path/filepath"
)

var (
	fileInfo *os.FileInfo
	err      error
)

type diretorio struct {
	nomeDiretorio string
	tamanho       float64
}

func main() {
	nomes, err := CapturarDiretorios("l:/projetos")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(nomes)
	for _, n := range nomes {
		fmt.Println(n)
	}
}

func CapturarDiretorios(pathx string) (diretorios []diretorio, err error) {

	var nomesIniciais, nomesIniciais2 []string

	file, err := os.Open(pathx)
	if err != nil {
		return diretorios, err
	}
	defer file.Close()

	nomesIniciais, err = file.Readdirnames(0)
	if err != nil {
		return diretorios, err
	}

	for _, n := range nomesIniciais {
		nomeComPathCompleto := pathx + "/" + n

		infoF, _ := os.Stat(nomeComPathCompleto)
		if infoF.IsDir() {
			nomesIniciais2 = append(nomesIniciais2, nomeComPathCompleto)
		}
	}

	for _, r := range nomesIniciais2 {
		info, err := os.Stat(r)
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			// fmt.Println(info.Name())
			// fmt.Println(r)

			// Se não chegou ainda no subdiretorio do node_modules, continua a busca recursivamente
			if path.Base(info.Name()) != "node_modules" {
				nomesRecursivos, err := CapturarDiretorios(r)

				if err != nil {
					return diretorios, err
				}
				for _, nx := range nomesRecursivos {
					if path.Base(nx.nomeDiretorio) == "node_modules" {
						var objDir diretorio
						objDir.nomeDiretorio = nx.nomeDiretorio
						objDir.tamanho, _ = DirSizeMB(nx.nomeDiretorio)
						diretorios = append(diretorios, objDir)
					}
				}
			} else {
				// Se achou o node_modules, aí adiciona o path completo no array names
				var objDir diretorio
				objDir.nomeDiretorio = r
				objDir.tamanho, _ = DirSizeMB(r)

				fmt.Println(objDir)

				diretorios = append(diretorios, objDir)
			}
		}
	}

	return diretorios, err

}

func DirSizeMB(path string) (float64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})

	// sizeMB := float64(size) / 1073741824 // converte para gigabytes
	sizeMB := float64(size) / 1048576 // converte para megabytes
	// sizeMB := float64(size) / 1024 // converte para kbytes
	sizeMBRound := math.Round(sizeMB)

	return sizeMBRound, err
}

// func DirSizeMB(path string) (float64, error) {
// 	var dirSize int64 = 0

// 	readSize := func(path string, file os.FileInfo, err error) error {
// 		if !file.IsDir() {
// 			dirSize += file.Size()
// 		}

// 		return nil
// 	}

// 	err := filepath.Walk(path, readSize)
// 	if err != nil {
// 		return 0, err
// 	}

// 	sizeMB := float64(dirSize) / 1024.0 / 1024.0

// 	return sizeMB, nil
// }

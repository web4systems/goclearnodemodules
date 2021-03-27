package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

var (
	fileInfo *os.FileInfo
	err      error
	q        int
)

type diretorio struct {
	nomeDiretorio    string
	tamanhoMegaBytes float64
	tamanhoBytes     float64
	qtdeArquivos     int
}

func main() {
	q = 1
	nomes, err := CapturarDiretorios("l:/projetos/web")
	if err != nil {
		panic(err)
	}
	// fmt.Println(nomes)
	var totalMegaBytes float64
	var i int64
	var x int
	for _, n := range nomes {
		fmt.Println(n.nomeDiretorio + " | " + fmt.Sprintf("%.0f", n.tamanhoMegaBytes) + "MB" + " | " + fmt.Sprintf("%.2f", n.tamanhoBytes) + " Arquivos: " + strconv.Itoa(n.qtdeArquivos))
		totalMegaBytes += n.tamanhoMegaBytes
		i += 1
		x += n.qtdeArquivos
	}
	fmt.Println(strconv.FormatInt(i, 10) + " diretórios: " + fmt.Sprintf("%.0f", totalMegaBytes) + "MB" + " (" + strconv.Itoa(x) + ")")
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
						objDir.tamanhoBytes, _ = DirSizeMB(nx.nomeDiretorio)
						objDir.tamanhoMegaBytes = objDir.tamanhoBytes / 1048576
						// objDir.qtdeArquivos, _ = fileCount(nx.nomeDiretorio)
						objDir.qtdeArquivos = q
						diretorios = append(diretorios, objDir)
					}
				}
			} else {
				// Se achou o node_modules, aí adiciona o path completo no array names
				var objDir diretorio
				objDir.nomeDiretorio = r
				objDir.tamanhoBytes, _ = DirSizeMB(r)
				objDir.tamanhoMegaBytes = objDir.tamanhoBytes / 1048576
				// objDir.qtdeArquivos, _ = fileCount(r)
				objDir.qtdeArquivos = q
				// fmt.Println(objDir)

				diretorios = append(diretorios, objDir)
			}
		} else {
			fmt.Println("aqui")
			q += 1
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
	// sizeMB := float64(size) / 1048576 // converte para megabytes
	// sizeMB := float64(size) / 1024 // converte para kbytes
	sizeMB := float64(size) // retorna em bytes
	// sizeMBRound := math.Round(sizeMB)

	return sizeMB, err
}

func fileCount(path string) (int, error) {
	fmt.Println(path)
	i := 0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return 0, err
	}
	for _, file := range files {
		if !file.IsDir() {
			i++
		}
	}
	return i, nil
}

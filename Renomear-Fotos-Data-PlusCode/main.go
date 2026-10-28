package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	olc "github.com/google/open-location-code/go"
	"github.com/rwcarlsen/goexif/exif"
)

var extensoes = []string{".jpg", ".jpeg", ".png", ".heic"}

func main() {
	// Determina o diretório de trabalho (argumento ou pasta atual)
	dirAtual, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter diretório atual:", err)
		return
	}
	if len(os.Args) > 1 {
		dirAtual = os.Args[1]
	}

	arquivos, _ := os.ReadDir(dirAtual)
	sort.Slice(arquivos, func(i, j int) bool {
		return arquivos[i].Name() < arquivos[j].Name()
	})

	// Limita o número de goroutines ao número de CPUs
	sem := make(chan struct{}, runtime.NumCPU())
	var wg sync.WaitGroup

	for _, arq := range arquivos {
		nomeArq := arq.Name()
		ext := strings.ToLower(filepath.Ext(nomeArq))

		if arq.IsDir() || !contem(extensoes, ext) {
			continue
		}

		wg.Add(1)
		sem <- struct{}{} // Adquire um slot para processar

		go func(caminhoOrigem, ext string) {
			defer wg.Done()
			defer func() { <-sem }() // Libera o slot

			// 1. Obter a data e o plus code da imagem
			data, plusCode, err := obterInfoGeo(caminhoOrigem)
			if err != nil {
				fmt.Printf("Erro ao ler metadados de %s: %v\n", filepath.Base(caminhoOrigem), err)
				return
			}

			// 2. Gerar o nome do novo arquivo e definir a extensão final
			novoNomeBase := data.Format("2006.01.02-15.04.05")
			if plusCode != "" {
				novoNomeBase += "--" + plusCode
			}

			novaExtensao := ext
			if ext == ".jpeg" {
				novaExtensao = ".jpg"
			}

			caminhoDestino := filepath.Join(filepath.Dir(caminhoOrigem), novoNomeBase+novaExtensao)

			// 3. Tratar colisões de nome
			for i := 1; ; i++ {
				if _, err := os.Stat(caminhoDestino); os.IsNotExist(err) {
					break
				}
				caminhoDestino = filepath.Join(filepath.Dir(caminhoOrigem), fmt.Sprintf("%s_%d%s", novoNomeBase, i, novaExtensao))
			}

			// 4. Renomear o arquivo original (sem conversão)
			if err := os.Rename(caminhoOrigem, caminhoDestino); err != nil {
				fmt.Printf("Erro ao renomear %s: %v\n", filepath.Base(caminhoOrigem), err)
			}

		}(filepath.Join(dirAtual, nomeArq), ext)
	}

	wg.Wait()
	fmt.Println("Processo concluído.")
}

// obterInfoGeo extrai a data e as coordenadas (convertidas para Plus Code).
func obterInfoGeo(caminho string) (time.Time, string, error) {
	var data time.Time
	var plusCode string

	arquivo, err := os.Open(caminho)
	if err != nil {
		return time.Time{}, "", err
	}
	defer arquivo.Close()

	x, err := exif.Decode(arquivo)
	if err == nil {
		// Tenta pegar data do EXIF
		if dt, err := x.DateTime(); err == nil {
			data = dt
		}
		// Tenta pegar coordenadas do EXIF
		if lat, long, err := x.LatLong(); err == nil {
			plusCode = olc.Encode(lat, long, 11)
		}
	}

	// Se não conseguiu data do EXIF, usa a data de modificação
	if data.IsZero() {
		info, err := os.Stat(caminho)
		if err != nil {
			return time.Time{}, "", err
		}
		data = info.ModTime()
	}

	return data, plusCode, nil
}

// contem verifica se um item existe em uma lista de strings.
func contem(lista []string, item string) bool {
	for _, v := range lista {
		if v == item {
			return true
		}
	}
	return false
}

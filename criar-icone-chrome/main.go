// main.go
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/disintegration/imaging"
)

func main() {
	sizes := []int{512, 256, 128, 64, 32, 16}

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatalf("Erro ao ler diretório: %v", err)
	}

	derivRegex := regexp.MustCompile(`(?i)^icon\d+\.(png|jpg|jpeg)$`)

	processed := false

	for _, entry := range files {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		ext := strings.ToLower(filepath.Ext(name))
		if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
			continue
		}

		// Ignorar derivações geradas pelo programa (ex: Icon512.png)
		if derivRegex.MatchString(name) {
			continue
		}

		// Ignorar se já for o arquivo principal renomeado para evitar processamento infinito dele mesmo
		// se houver outro arquivo candidato, mas ainda queremos processar o Icon.ext caso seja o único.
		// Na verdade, se for "Icon.png" ou "Icon.jpg", nós queremos processar, mas precisamos garantir
		// que não vamos re-processar eternamente se rodarmos de novo. Mas se rodarmos de novo,
		// ele vai processar o Icon.png e apenas sobrescrever os Icon512.png existentes, o que é aceitável.
		// Porém, para não re-processar à toa ou gerar conflito se existirem múltiplos arquivos,
		// vamos verificar se a imagem é quadrada primeiro.

		img, err := imaging.Open(name)
		if err != nil {
			log.Printf("Falha ao abrir %s: %v", name, err)
			continue
		}

		bounds := img.Bounds()
		width := bounds.Dx()
		height := bounds.Dy()
		if width != height {
			log.Printf("Ignorando %s: a imagem não é quadrada (%dx%d)", name, width, height)
			continue
		}

		// Casing padrão: Icon.png ou Icon.jpg ou Icon.jpeg
		targetName := "Icon" + filepath.Ext(name)
		if strings.ToLower(name) != strings.ToLower(targetName) {
			fmt.Printf("Renomeando %s para %s...\n", name, targetName)
			if err := os.Rename(name, targetName); err != nil {
				log.Printf("Falha ao renomear %s para %s: %v", name, targetName, err)
				continue
			}
			name = targetName
		}

		fmt.Printf("Processando %s...\n", name)
		for _, size := range sizes {
			dst := imaging.Resize(img, size, size, imaging.Lanczos)
			out := fmt.Sprintf("Icon%d.png", size)
			if err := imaging.Save(dst, out); err != nil {
				log.Printf("Falha ao salvar %s: %v", out, err)
			} else {
				fmt.Printf("Gerado: %s\n", out)
			}
		}
		processed = true
	}

	if !processed {
		fmt.Println("Nenhum arquivo de imagem quadrado válido (PNG ou JPG) foi encontrado para processamento.")
	} else {
		fmt.Println("Processamento concluído.")
	}
}


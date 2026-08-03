package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// processarPDF extrai a primeira página do PDF e substitui o arquivo original.
func processarPDF(inputFile string) error {
	ext := filepath.Ext(inputFile)
	base := strings.TrimSuffix(inputFile, ext)
	// Cria um nome temporário para o novo arquivo.
	tempOutput := base + "_first_temp" + ext

	// Extrai somente a primeira página e gera o PDF temporário.
	if err := api.TrimFile(inputFile, tempOutput, []string{"1"}, nil); err != nil {
		return err
	}

	// Remove o arquivo original.
	if err := os.Remove(inputFile); err != nil {
		return err
	}

	// Renomeia o arquivo temporário para o nome original.
	if err := os.Rename(tempOutput, inputFile); err != nil {
		return err
	}

	return nil
}

func main() {
	// Verifica se foi passado um argumento.
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	arg := os.Args[1]

	// Se o argumento for um diretório, utiliza-o; se for um arquivo, usa o diretório que o contém.
	info, err := os.Stat(arg)
	if err != nil {
		os.Exit(1)
	}
	var folder string
	if info.IsDir() {
		folder = arg
	} else {
		folder = filepath.Dir(arg)
	}

	// Lê os arquivos da pasta (não recursivamente).
	entries, err := os.ReadDir(folder)
	if err != nil {
		os.Exit(1)
	}

	// Processa cada arquivo PDF encontrado.
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.ToLower(filepath.Ext(entry.Name())) != ".pdf" {
			continue
		}
		inputFile := filepath.Join(folder, entry.Name())
		_ = processarPDF(inputFile)
	}
}

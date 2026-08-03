package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func processarPDF(inputFile string) error {
	ext := filepath.Ext(inputFile)
	base := strings.TrimSuffix(inputFile, ext)
	tempOutput := base + "_first_last_temp" + ext

	pageCount, err := api.PageCountFile(inputFile)
	if err != nil {
		return err
	}

	pageSelection := []string{"1", fmt.Sprintf("%d", pageCount)}

	if err := api.TrimFile(inputFile, tempOutput, pageSelection, nil); err != nil {
		return err
	}

	if err := os.Remove(inputFile); err != nil {
		return err
	}
	if err := os.Rename(tempOutput, inputFile); err != nil {
		return err
	}

	return nil
}

func main() {
	var folder string

	// Usa argumento se fornecido (ex: menu de contexto)
	if len(os.Args) > 1 {
		arg := os.Args[1]
		info, err := os.Stat(arg)
		if err != nil {
			return
		}
		if info.IsDir() {
			folder = arg
		} else {
			folder = filepath.Dir(arg)
		}
	} else {
		// Se nenhum argumento, usa o diretório do .exe (duplo clique)
		exePath, err := os.Executable()
		if err != nil {
			return
		}
		folder = filepath.Dir(exePath)
	}

	entries, err := os.ReadDir(folder)
	if err != nil {
		return
	}

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

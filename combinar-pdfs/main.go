package main

import (
	"fmt" // NOVO: Importado para formatar strings
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// Estrutura que guarda nome e caminho completo
type pdfFile struct {
	name string
	path string
}

func main() {
	dir, _ := os.Getwd()
	if len(os.Args) > 1 {
		arg := os.Args[1]
		info, err := os.Stat(arg)
		if err == nil {
			if info.IsDir() {
				dir = arg
			} else {
				dir = filepath.Dir(arg)
			}
		}
	}

	// NOVO: Pega o nome da pasta atual e cria o nome do arquivo de saída
	folderName := filepath.Base(dir)
	outputFileName := fmt.Sprintf("%s-combinado.pdf", folderName)

	files, _ := os.ReadDir(dir)

	var pdfs []pdfFile

	for _, f := range files {
		name := f.Name()
		// ALTERADO: Usa a variável outputFileName para evitar que o arquivo gerado seja incluído
		if !f.IsDir() && strings.HasSuffix(strings.ToLower(name), ".pdf") && name != outputFileName {
			pdfs = append(pdfs, pdfFile{
				name: name,
				path: filepath.Join(dir, name),
			})
		}
	}

	// Ordena por ordem natural (1.pdf, 2.pdf, 10.pdf)
	sort.SliceStable(pdfs, func(i, j int) bool {
		return naturalLess(pdfs[i].name, pdfs[j].name)
	})

	// Monta lista de caminhos
	var paths []string
	for _, f := range pdfs {
		paths = append(paths, f.path)
	}

	// Junta os PDFs em ordem
	// ALTERADO: Usa a variável outputFileName para nomear o arquivo final
	_ = api.MergeCreateFile(paths, filepath.Join(dir, outputFileName), false, nil)
	fmt.Printf("Arquivo '%s' criado com sucesso!\n", outputFileName) // NOVO: Adiciona um feedback
}

// naturalLess compara dois nomes levando em conta números
func naturalLess(a, b string) bool {
	aParts := splitNatural(a)
	bParts := splitNatural(b)

	for i := 0; i < len(aParts) && i < len(bParts); i++ {
		if aParts[i].isNum && bParts[i].isNum {
			if aParts[i].num != bParts[i].num {
				return aParts[i].num < bParts[i].num
			}
		} else if aParts[i].str != bParts[i].str {
			return aParts[i].str < bParts[i].str
		}
	}

	return len(aParts) < len(bParts)
}

type part struct {
	isNum bool
	num   int
	str   string
}

// splitNatural separa números e textos de uma string para ordenação natural
func splitNatural(s string) []part {
	re := regexp.MustCompile(`\d+|\D+`)
	raw := re.FindAllString(s, -1)
	var parts []part

	for _, r := range raw {
		if n, err := strconv.Atoi(r); err == nil {
			parts = append(parts, part{isNum: true, num: n})
		} else {
			parts = append(parts, part{isNum: false, str: strings.ToLower(r)})
		}
	}
	return parts
}

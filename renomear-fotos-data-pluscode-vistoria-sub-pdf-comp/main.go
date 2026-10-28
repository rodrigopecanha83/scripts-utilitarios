//go:build windows
// +build windows

package main

import (
	"bytes"
	"embed"
	"fmt"
	"image/jpeg"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/disintegration/imaging"
	olc "github.com/google/open-location-code/go"
	"github.com/jung-kurt/gofpdf"
	"github.com/rwcarlsen/goexif/exif"
	"golang.org/x/sys/windows"
)

//go:embed embed/*
var fs embed.FS

var (
	extensoesValidas = []string{".jpg", ".jpeg", ".png", ".heic"}
)

func hideConsole() {
	mod := windows.NewLazySystemDLL("kernel32.dll")
	getConsoleWindow := mod.NewProc("GetConsoleWindow")
	ret, _, _ := getConsoleWindow.Call()
	if ret != 0 {
		freeConsole := mod.NewProc("FreeConsole")
		freeConsole.Call()
	}
}

func ajustarCWD() {
	if len(os.Args) > 1 {
		if info, err := os.Stat(os.Args[1]); err == nil {
			if info.IsDir() {
				os.Chdir(os.Args[1])
			} else {
				os.Chdir(filepath.Dir(os.Args[1]))
			}
		}
	}
}

func perguntarDataFallback(pasta string) string {
	nomePasta := strings.ReplaceAll(filepath.Base(pasta), "\"", "")
	
	// Alteramos a instrução visual para o fiscal saber o que o botão Cancelar faz agora
	msg := fmt.Sprintf("Fotos antigas ou sem data na pasta:\n[%s]\n\nDigite a data da vistoria:\n1 = Hoje\n2 = Ontem\ndd/mm/aaaa = Outra data\n\n[ CLIQUE EM 'CANCELAR' PARA IGNORAR E MANTER A DATA ORIGINAL ]", nomePasta)

	msgVBS := strings.ReplaceAll(msg, "\n", `" & vbCrLf & "`)
	msgVBS = `"` + msgVBS + `"`

	tempDir := os.TempDir()
	vbsPath := filepath.Join(tempDir, "ask_date.vbs")
	outPath := filepath.Join(tempDir, "ask_date_out.txt")

	// O pulo do gato: Se resp for vazio (Cancelar ou 'X'), transformamos em "IGNORAR"
	vbsCode := fmt.Sprintf(`
Dim resp
resp = InputBox(%s, "Aviso de Data Antiga", "1")

If resp = "" Then
    resp = "IGNORAR"
End If

Dim fso, f
Set fso = CreateObject("Scripting.FileSystemObject")
Set f = fso.CreateTextFile("%s", True)
f.Write resp
f.Close
`, msgVBS, outPath)

	os.WriteFile(vbsPath, []byte(vbsCode), 0644)

	cmd := exec.Command("wscript.exe", "//nologo", vbsPath)
	cmd.Run()

	out, err := os.ReadFile(outPath)

	os.Remove(vbsPath)
	os.Remove(outPath)

	if err != nil {
		return "IGNORAR" // Fallback de segurança em caso de erro extremo
	}

	return strings.TrimSpace(string(out))
}

func converterRespostaParaData(resp string) time.Time {
	hoje := time.Now()
	switch resp {
	case "1", "":
		return hoje
	case "2":
		return hoje.AddDate(0, 0, -1)
	default:
		resp = strings.ReplaceAll(resp, "/", ".")
		t, err := time.Parse("02.01.2006", resp)
		if err == nil {
			return t
		}
		return hoje
	}
}

type infoArquivo struct {
	novoCaminho string
	data        time.Time
}

func renomearFotosComGeoInfo(dir string) ([]string, string, error) {
    arquivos, _ := os.ReadDir(dir)
    var imagens []os.DirEntry

    for _, arq := range arquivos {
        if !arq.IsDir() && contem(extensoesValidas, strings.ToLower(filepath.Ext(arq.Name()))) {
            imagens = append(imagens, arq)
        }
    }

    if len(imagens) == 0 {
        return nil, "", nil
    }

    hoje := time.Now()
    limitePassado := hoje.AddDate(0, 0, -90)
    limiteFuturo := hoje.AddDate(0, 0, 1)

    var fallbackDate time.Time
    var foundValid bool

    // PASSO 1: Busca o Consenso (MANTIDO!)
    for _, arq := range imagens {
        caminhoOrigem := filepath.Join(dir, arq.Name())
        data, _, err := obterInfoGeo(caminhoOrigem)

        if err == nil && !data.IsZero() && !data.Before(limitePassado) && !data.After(limiteFuturo) {
            fallbackDate = data
            foundValid = true
            break
        }
    }

    // NOSSA NOVA LÓGICA DE FALLBACK AQUI
    var resposta string
    if !foundValid {
        resposta = perguntarDataFallback(dir)
        if resposta != "IGNORAR" {
            fallbackDate = converterRespostaParaData(resposta)
        }
    }

    // PASSO 2: Renomeação Concorrente
    resultadosChan := make(chan infoArquivo, len(imagens))
    sem := make(chan struct{}, runtime.NumCPU())
    var wg sync.WaitGroup
    var renameMu sync.Mutex

    for _, arq := range imagens {
        wg.Add(1)
        sem <- struct{}{}

        go func(nomeArq string) {
            defer wg.Done()
            defer func() { <-sem }()

            caminhoOrigem := filepath.Join(dir, nomeArq)
            data, plusCode, _ := obterInfoGeo(caminhoOrigem)

            // Só aplica o fallback se o fiscal NÃO mandou ignorar
            if (data.IsZero() || data.Before(limitePassado) || data.After(limiteFuturo)) && resposta != "IGNORAR" {
                data = fallbackDate
            }
            
            // Prevenção de erro: Se a foto não tiver data nenhuma (0001-01-01) e o fiscal ignorou,
            // colocamos a data de 'hoje' para o nome do arquivo não ficar quebrado
            if data.IsZero() && resposta == "IGNORAR" {
                data = hoje
            }

            novoNomeBase := data.Format("2006.01.02-15.04.05")
            if plusCode != "" {
                novoNomeBase += "--" + plusCode
            }

            extOriginal := strings.ToLower(filepath.Ext(caminhoOrigem))
            finalExt := extOriginal
            if extOriginal == ".jpeg" {
                finalExt = ".jpg"
            }

            renameMu.Lock()

            var caminhoDestino string
            for i := 1; ; i++ {
                caminhoDestino = filepath.Join(dir, fmt.Sprintf("%s-%02d%s", novoNomeBase, i, finalExt))
                if _, err := os.Stat(caminhoDestino); os.IsNotExist(err) {
                    break
                }
            }

            errRename := os.Rename(caminhoOrigem, caminhoDestino)
            renameMu.Unlock()

            if errRename != nil {
                caminhoDestino = caminhoOrigem
            }

            resultadosChan <- infoArquivo{novoCaminho: caminhoDestino, data: data}
        }(arq.Name())
    }

    wg.Wait()
    close(resultadosChan)

    var resultados []infoArquivo
    for res := range resultadosChan {
        resultados = append(resultados, res)
    }

    sort.Slice(resultados, func(i, j int) bool {
        return resultados[i].data.Before(resultados[j].data)
    })

    var novosNomes []string
    var dataRelatorio string
    if len(resultados) > 0 {
        dataRelatorio = resultados[0].data.Format("02/01/2006")
        for _, res := range resultados {
            novosNomes = append(novosNomes, res.novoCaminho)
        }
    }

    return novosNomes, dataRelatorio, nil
}

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
		if dt, err := x.DateTime(); err == nil {
			data = dt
		}
		if lat, long, err := x.LatLong(); err == nil {
			plusCode = olc.Encode(lat, long, 11)
		}
	}
	return data, plusCode, nil
}

// --- ARQUITETURA DEFINITIVA: 100% NA MEMÓRIA RAM ---

func montarPDFRapido(fotos []string, dataRelatorio string, outPath string) error {
	pdf := gofpdf.New("P", "pt", "A4", "")

	// 1. Carrega os Assets
	fontBytes, _ := fs.ReadFile("embed/OpenSans-Bold.ttf")
	pdf.AddUTF8FontFromBytes("OpenSans", "B", fontBytes)

	templateBytes, _ := fs.ReadFile("embed/template.jpg")
	optTmpl := gofpdf.ImageOptions{ImageType: "JPEG"}
	pdf.RegisterImageOptionsReader("template", optTmpl, bytes.NewReader(templateBytes))

	// 2. Processa as Imagens Concorrentemente na Memória RAM
	memImages := make([][]byte, len(fotos))
	var wg sync.WaitGroup
	// Limita a concorrência ao número de núcleos da CPU para não estourar a RAM
	sem := make(chan struct{}, runtime.NumCPU())

	for i, foto := range fotos {
		wg.Add(1)
		sem <- struct{}{}

		go func(idx int, path string) {
			defer wg.Done()
			defer func() { <-sem }()

			// Abre a imagem e corrige a rotação automaticamente lendo o EXIF
			img, err := imaging.Open(path, imaging.AutoOrientation(true))
			if err == nil {
				// Recorta e redimensiona a imagem para a proporção exata de 3:2 na RAM
				img = imaging.Fill(img, 1800, 1200, imaging.Center, imaging.Lanczos)

				// Comprime para JPEG direto para um buffer de memória (sem gravar no HD)
				var buf bytes.Buffer
				jpeg.Encode(&buf, img, &jpeg.Options{Quality: 75})
				memImages[idx] = buf.Bytes()
			}
		}(i, foto)
	}
	wg.Wait() // Aguarda todas as fotos serem processadas na memória

	// 3. Monta o PDF a partir da Memória
	pares := (len(memImages) + 1) / 2
	for i := 0; i < pares; i++ {
		pdf.AddPage()

		// Fundo
		pdf.ImageOptions("template", 0, 0, 595.28, 841.89, false, optTmpl, 0, "")

		// Foto Superior
		idx1 := i * 2
		if idx1 < len(memImages) && memImages[idx1] != nil {
			imgName := fmt.Sprintf("img_%d", idx1)
			pdf.RegisterImageOptionsReader(imgName, gofpdf.ImageOptions{ImageType: "JPEG"}, bytes.NewReader(memImages[idx1]))
			pdf.ImageOptions(imgName, 81.85, 169.94, 432.06, 288.04, false, gofpdf.ImageOptions{ImageType: "JPEG"}, 0, "")
		}

		// Foto Inferior
		idx2 := i*2 + 1
		if idx2 < len(memImages) && memImages[idx2] != nil {
			imgName := fmt.Sprintf("img_%d", idx2)
			pdf.RegisterImageOptionsReader(imgName, gofpdf.ImageOptions{ImageType: "JPEG"}, bytes.NewReader(memImages[idx2]))
			pdf.ImageOptions(imgName, 81.85, 466.86, 432.06, 288.04, false, gofpdf.ImageOptions{ImageType: "JPEG"}, 0, "")
		}

		// Data
		pdf.SetFont("OpenSans", "B", 12)
		pdf.SetXY(307.0, 149.0)
		pdf.CellFormat(150, 15, dataRelatorio, "", 0, "L", false, 0, "")
	}

	// 4. Salva o PDF no HD (A única operação de gravação do processo inteiro!)
	return pdf.OutputFileAndClose(outPath)
}

func processarPasta(dir string) {
	novosNomes, dataRelatorio, err := renomearFotosComGeoInfo(dir)
	if err != nil || len(novosNomes) == 0 {
		return
	}

	pdfPath := filepath.Join(dir, "vistoria.pdf")
	montarPDFRapido(novosNomes, dataRelatorio, pdfPath)
}

// ==========================================
// PONTO DE ENTRADA OTIMIZADO (MULTITHREAD)
// ==========================================
func main() {
	hideConsole()
	ajustarCWD()

	dirAtual, _ := os.Getwd()

	var wg sync.WaitGroup
	// Semáforo Global para pastas: evita que o Go tente abrir 100 pastas de uma vez, 
	// o que poderia estourar a RAM do PC do usuário. Limita à quantidade de núcleos do processador.
	sem := make(chan struct{}, runtime.NumCPU())

	filepath.WalkDir(dirAtual, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		
		// Encontrou um diretório? Joga ele para uma thread!
		if d.IsDir() {
			wg.Add(1)
			sem <- struct{}{} // Ocupa uma vaga no processador

			go func(pastaAlvo string) {
				defer wg.Done()
				defer func() { <-sem }() // Libera a vaga ao terminar

				processarPasta(pastaAlvo)
			}(path)
		}
		return nil
	})

	// Segura o programa aberto até a última pasta (Thread) terminar o seu trabalho
	wg.Wait()
}

func contem(lista []string, item string) bool {
	for _, v := range lista {
		if v == item {
			return true
		}
	}
	return false
}
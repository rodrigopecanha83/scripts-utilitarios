# Renomear Fotos por Data e Plus Code

Utilitário de linha de comando concorrente em Go para renomear lotes de imagens (`.jpg`, `.jpeg`, `.png`, `.heic`) utilizando a data/hora do metadado EXIF e a localização geográfica convertida em Plus Code (Open Location Code).

## Funcionalidades Principais

* **Extração de Data EXIF:** Obtém a data/hora original da foto a partir dos metadados EXIF (com fallback para a data de modificação do arquivo).
* **Geolocalização em Plus Code:** Converte coordenadas GPS da imagem em um Plus Code (Google Open Location Code, precisão 11) formatado como `--[PlusCode]`.
* **Renomeação Padronizada:** Formata os arquivos como `AAAA.MM.DD-HH.MM.SS--[PlusCode].ext` ou `AAAA.MM.DD-HH.MM.SS.ext`.
* **Processamento Concorrente:** Utiliza Goroutines com limite ajustado ao número de núcleos de CPU do sistema para alto desempenho.
* **Prevenção de Colisão:** Adiciona sufixos numéricos (`_1`, `_2`) caso já existam arquivos com o mesmo nome gerado.

## Pré-requisitos

* Go 1.20+ (caso vá compilar a partir do código-fonte `main.go`).
* Dependências Go: `github.com/google/open-location-code/go/olc` e `github.com/rwcarlsen/goexif/exif`.
* Executável pré-compilado `renomear-fotos-data-pluscode.exe` para Windows.
* Suporte aos formatos de imagem: `.jpg`, `.jpeg`, `.png` e `.heic`.

## Como Instalar e Configurar

1. Clone ou baixe a pasta `renomear-fotos-data-pluscode`.
2. Para compilar manualmente (opcional):
   ```bash
   go build -o renomear-fotos-data-pluscode.exe main.go
   ```
3. Opcional: Execute o arquivo `renomear-fotos-data.reg` para adicionar o atalho ao menu do Windows Explorer.

## Como Usar

Exemplo rápido de uso via linha de comando ou arrastar-e-soltar:

```bash
# Executar no diretório atual
.\renomear-fotos-data-pluscode.exe

# Executar especificando uma pasta de fotos
.\renomear-fotos-data-pluscode.exe "C:\Caminho\Para\Fotos"
```

Alternativamente, você pode arrastar e soltar a pasta com fotos sobre o executável `renomear-fotos-data-pluscode.exe`.

## Desenvolvedor:
Rodrigo Peçanha
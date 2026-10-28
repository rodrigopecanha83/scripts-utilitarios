# Criar Ícones Chrome

Utilitário de linha de comando para gerar automaticamente ícones em múltiplos tamanhos (512, 256, 128, 64, 32 e 16 pixels) a partir de uma imagem quadrada, ideal para uso em extensões do Google Chrome e Microsoft Edge.

## Funcionalidades Principais

* **Detecção Automática de Imagem:** Localiza arquivos de imagem válidos (PNG, JPG ou JPEG) no diretório atual.
* **Validação de Proporção:** Verifica se a imagem de entrada é estritamente quadrada antes de efetuar o redimensionamento.
* **Padronização do Arquivo de Origem:** Renomeia o arquivo original de forma automática para `Icon.png` (ou `.jpg`/`.jpeg`).
* **Geração Multi-tamanho:** Cria os arquivos redimensionados `Icon512.png`, `Icon256.png`, `Icon128.png`, `Icon64.png`, `Icon32.png` e `Icon16.png` usando reamostragem de alta qualidade (filtro Lanczos).

## Pré-requisitos

* Go 1.20+ (caso vá compilar a partir do código-fonte `main.go`).
* Dependência Go: `github.com/disintegration/imaging`.
* Executável pré-compilado `criar-icone-chrome.exe` para execução direta no Windows.
* Imagem de entrada no formato PNG, JPG ou JPEG com proporção 1:1 (quadrada).

## Como Instalar e Configurar

1. Clone ou baixe a pasta `criar-icone-chrome`.
2. Para compilar manualmente (opcional):
   ```bash
   go build -o criar-icone-chrome.exe main.go
   ```

## Como Usar

Exemplo rápido de uso via linha de comando:

1. Coloque o executável `criar-icone-chrome.exe` na mesma pasta onde está a sua imagem quadrada (PNG ou JPG).
2. Execute o comando:

```bash
.\criar-icone-chrome.exe
```

O programa renomeará a imagem de origem para `Icon.png` e gerará todas as variações de tamanho necessárias no mesmo diretório.

## Desenvolvedor:
Rodrigo Peçanha

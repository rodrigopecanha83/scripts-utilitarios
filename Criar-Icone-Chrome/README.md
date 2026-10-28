# Criar-Icone-Chrome

Utilitário de linha de comando para gerar automaticamente o pacote completo de ícones em múltiplos tamanhos para extensões do Google Chrome e Chromium.

## Funcionalidades Principais

* **Detecção Automática:** Localiza imagens válidas nos formatos PNG, JPG ou JPEG no diretório.
* **Validação de Proporção:** Confirma se a imagem de entrada possui proporção estritamente quadrada (1:1).
* **Geração Multi-resolução:** Gera ícones nas resoluções `512x512`, `256x256`, `128x128`, `64x64`, `32x32` e `16x16` pixels com reamostragem Lanczos.
* **Padronização do Arquivo Original:** Renomeia o arquivo original automaticamente para `Icon.png` (ou formato correspondente).

## Pré-requisitos

* Go 1.20+ (caso vá compilar a partir do código-fonte `main.go`).
* Dependência Go: `github.com/disintegration/imaging`.
* Executável pré-compilado `criar-icone-chrome.exe` para Windows.
* Imagem de origem quadrada (1:1) em PNG ou JPG.

## Instalação e Configuração

1. Clone ou baixe este repositório para o seu computador.
2. (Opcional) Para compilar o executável a partir do código Go:
   ```bash
   go build -o criar-icone-chrome.exe main.go
   ```

## Como Usar

Coloque o executável na mesma pasta da imagem de origem e execute:

```bash
.\criar-icone-chrome.exe
```

## Contribuindo

Sugestões e correções são bem-vindas. Abra uma Issue ou envie um Pull Request.

## Licença

TODO: definir licença. Veja o arquivo `LICENSE` para mais detalhes.

## Autor

Rodrigo Peçanha

# Separar-PDF-Limites

Utilitário de linha de comando para recortar arquivos PDF em lote mantendo exclusivamente a primeira e a última página de cada documento.

## Funcionalidades Principais

* **Recorte de Páginas Limite:** Extrai apenas a página 1 e a última página (página final) de cada arquivo PDF.
* **Processamento em Lote:** Varre o diretório alvo e processa todos os arquivos com extensão `.pdf`.
* **Substituição Direta:** Atualiza os documentos diretamente sem gerar duplicatas.
* **Integração com Menu de Contexto:** Arquivo `.reg` para execução rápida com o botão direito no Windows Explorer.

## Pré-requisitos

* Go 1.20+ (caso vá compilar a partir do código-fonte `main.go`).
* Dependência Go: `github.com/pdfcpu/pdfcpu`.
* Executável pré-compilado `separar-pdf-limites.exe` para Windows.

## Instalação e Configuração

1. Clone ou baixe este repositório para o seu computador.
2. (Opcional) Para registrar no menu de contexto do Windows Explorer:
   - Copie os arquivos da pasta para `C:\MenuContexto\Separar-PDF-Limites\`.
   - Execute o arquivo `separar-pdf-limites.reg`.
3. (Opcional) Para compilar a partir do código Go:
   ```bash
   go build -o separar-pdf-limites.exe main.go
   ```

## Como Usar

Execute o utilitário passando o caminho da pasta ou no diretório atual:

```bash
# Execução na pasta atual
.\separar-pdf-limites.exe

# Ou informando o caminho do diretório
.\separar-pdf-limites.exe "C:\Caminho\Para\Pasta"
```

## Contribuindo

Sugestões e correções são bem-vindas. Abra uma Issue ou envie um Pull Request.

## Licença

Este projeto é distribuído sob os termos da licença **GNU General Public License v3.0 (GNU GPLv3)**. Consulte o arquivo [LICENSE](LICENSE) para mais detalhes.

## Autor

Rodrigo Peçanha


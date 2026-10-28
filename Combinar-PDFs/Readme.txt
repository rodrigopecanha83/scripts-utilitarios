# Combinar PDFs

Ferramenta de linha de comando para combinar múltiplos arquivos PDF em um único documento de forma automática, preservando a ordenação natural dos nomes dos arquivos.

## Funcionalidades Principais

* **Junção Automática de PDFs:** Une todos os arquivos `.pdf` presentes na pasta em um único documento.
* **Ordenação Natural:** Ordena os arquivos de forma inteligente (ex: `1.pdf`, `2.pdf`, `10.pdf`) antes da junção.
* **Nomeação Dinâmica:** Nomeia o arquivo gerado automaticamente como `[nome-da-pasta]-combinado.pdf`.
* **Integração com Menu de Contexto:** Arquivo `.reg` para atalho no clique com botão direito do mouse no Windows Explorer.

## Pré-requisitos

* Go 1.20+ (caso vá compilar a partir do código-fonte `main.go`).
* Dependência Go: `github.com/pdfcpu/pdfcpu`.
* Executável pré-compilado `combinar-pdfs.exe` para execução direta no Windows.

## Como Instalar e Configurar

1. Clone ou baixe a pasta `combinar-pdfs`.
2. Para adicionar ao menu de contexto do Windows (opcional):
   - Mova a pasta ou os arquivos `combinar-pdfs.exe` e `Icon.ico` para `C:\MenuContexto\combinar-pdfs\`.
   - Execute o arquivo `combinar-pdfs.reg` para registrar o atalho no menu do botão direito do Windows Explorer.
3. Para compilar manualmente (opcional):
   ```bash
   go build -o combinar-pdfs.exe main.go
   ```

## Como Usar

Exemplo rápido de uso via linha de comando ou menu de contexto:

```bash
# Execução direta na pasta dos PDFs
.\combinar-pdfs.exe

# Ou passe o caminho da pasta como argumento
.\combinar-pdfs.exe "C:\Caminho\Para\Pasta"
```

Alternativamente, clique com o botão direito sobre a pasta ou dentro dela no Windows Explorer e selecione **Combinar PDFs**.

## Desenvolvedor:
Rodrigo Peçanha
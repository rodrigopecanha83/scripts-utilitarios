# Separar PDF Página 1

Utilitário de linha de comando para recortar automaticamente todos os arquivos PDF de uma pasta, mantendo exclusivamente a primeira página de cada documento (substituindo os arquivos originais in-place).

## Funcionalidades Principais

* **Extração da Primeira Página:** Mantém apenas a 1ª página de cada arquivo PDF da pasta.
* **Processamento em Lote:** Processa todos os arquivos com extensão `.pdf` presentes na pasta informada.
* **Substituição Direta (In-place):** Substitui os arquivos originais diretamente pelos novos arquivos reduzidos.
* **Integração com Menu de Contexto:** Arquivo `.reg` para atalho no menu de contexto do Windows Explorer.

## Pré-requisitos

* Go 1.20+ (caso vá compilar a partir do código-fonte `main.go`).
* Dependência Go: `github.com/pdfcpu/pdfcpu`.
* Executável pré-compilado `separar-pdf-pag1.exe` para Windows.

## Como Instalar e Configurar

1. Clone ou baixe a pasta `separar-pdf-pag1`.
2. Para adicionar ao menu de contexto do Windows (opcional):
   - Mova a pasta e os arquivos para `C:\MenuContexto\separar-pdf-pag1\`.
   - Execute o arquivo `separar-pdf-pag1.reg` para cadastrar o atalho no menu de contexto.
3. Para compilar manualmente (opcional):
   ```bash
   go build -o separar-pdf-pag1.exe main.go
   ```

## Como Usar

Exemplo rápido de uso via linha de comando ou menu de contexto:

```bash
# Executar passando a pasta desejada como argumento
.\separar-pdf-pag1.exe "C:\Caminho\Para\Pasta"
```

Alternativamente, clique com o botão direito sobre a pasta no Windows Explorer e selecione **Cortar PDF (Apenas 1ª pág)**.

## Desenvolvedor:
Rodrigo Peçanha

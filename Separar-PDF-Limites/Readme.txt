# Separar PDF Limites (1ª e Última Página)

Utilitário de linha de comando para recortar automaticamente todos os arquivos PDF de uma pasta, mantendo apenas a primeira e a última página de cada documento (substituindo os arquivos originais in-place).

## Funcionalidades Principais

* **Extração de Páginas Limite:** Mantém exclusivamente a página 1 e a página final (N) de cada arquivo PDF da pasta.
* **Processamento em Lote:** Varre o diretório alvo e aplica a redução a todos os arquivos com extensão `.pdf`.
* **Substituição Direta (In-place):** Atualiza os arquivos diretamente sem gerar arquivos duplicados com outros nomes.
* **Integração com Menu de Contexto:** Registro do Windows Explorer para executar a ação com o botão direito do mouse.

## Pré-requisitos

* Go 1.20+ (caso vá compilar a partir do código-fonte `main.go`).
* Dependência Go: `github.com/pdfcpu/pdfcpu`.
* Executável pré-compilado `separar-pdf-limites.exe` para Windows.

## Como Instalar e Configurar

1. Clone ou baixe a pasta `separar-pdf-limites`.
2. Para adicionar ao menu de contexto do Windows (opcional):
   - Mova a pasta e os arquivos para `C:\MenuContexto\separar-pdf-limites\`.
   - Execute o arquivo `separar-pdf-limites.reg` para instalar o atalho no menu de contexto.
3. Para compilar manualmente (opcional):
   ```bash
   go build -o separar-pdf-limites.exe main.go
   ```

## Como Usar

Exemplo rápido de uso via linha de comando ou menu de contexto:

```bash
# Executar no diretório contendo os PDFs
.\separar-pdf-limites.exe

# Executar informando o caminho da pasta
.\separar-pdf-limites.exe "C:\Caminho\Para\Pasta"
```

Alternativamente, clique com o botão direito sobre a pasta no Windows Explorer e selecione **Cortar PDF (1ª e última)**.

## Desenvolvedor:
Rodrigo Peçanha

# Combinar-PDFs

Utilitário de linha de comando para combinar múltiplos arquivos PDF em um único documento mantendo a ordenação natural dos arquivos.

## Funcionalidades Principais

* **Junção Automática de PDFs:** Une todos os arquivos `.pdf` presentes na pasta em um único documento.
* **Ordenação Natural dos Arquivos:** Ordena os arquivos numericamente (ex.: `1.pdf`, `2.pdf`, `10.pdf`) antes da mesclagem.
* **Nomeação Dinâmica:** Gera automaticamente o documento consolidado com o padrão `[nome-da-pasta]-combinado.pdf`.
* **Integração com Menu de Contexto:** Arquivo `.reg` incluso para execução rápida pelo botão direito no Windows Explorer.

## Pré-requisitos

* Go 1.20+ (caso vá compilar a partir do código-fonte `main.go`).
* Dependência Go: `github.com/pdfcpu/pdfcpu`.
* Executável pré-compilado `combinar-pdfs.exe` para Windows.

## Instalação e Configuração

1. Clone ou baixe este repositório para o seu computador.
2. (Opcional) Para registrar no menu de contexto do Windows Explorer:
   - Copie os arquivos da pasta para `C:\MenuContexto\Combinar-PDFs\`.
   - Execute o arquivo `combinar-pdfs.reg`.
3. (Opcional) Para compilar a partir do fonte:
   ```bash
   go build -o combinar-pdfs.exe main.go
   ```

## Como Usar

Execute o utilitário pelo terminal ou através do menu de contexto:

```bash
# Execução direta dentro da pasta com os PDFs
.\combinar-pdfs.exe

# Ou informando o caminho do diretório
.\combinar-pdfs.exe "C:\Caminho\Para\Pasta"
```

## Contribuindo

Sugestões e correções são bem-vindas. Abra uma Issue ou envie um Pull Request.

## Licença

TODO: definir licença. Veja o arquivo `LICENSE` para mais detalhes.

## Autor

Rodrigo Peçanha

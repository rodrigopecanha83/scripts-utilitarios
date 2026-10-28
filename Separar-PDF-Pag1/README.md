# Separar-PDF-Pag1

Utilitário de linha de comando para recortar arquivos PDF em lote mantendo exclusivamente a primeira página de cada documento.

## Funcionalidades Principais

* **Extração da Primeira Página:** Extrai unicamente a página 1 de cada arquivo PDF do diretório.
* **Processamento em Lote:** Varre e processa automaticamente todos os arquivos `.pdf` da pasta selecionada.
* **Substituição Direta:** Atualiza os arquivos originais sem gerar arquivos adicionais com nomes diferentes.
* **Integração com Menu de Contexto:** Arquivo `.reg` para acionamento direto pelo botão direito no Windows Explorer.

## Pré-requisitos

* Go 1.20+ (caso vá compilar a partir do código-fonte `main.go`).
* Dependência Go: `github.com/pdfcpu/pdfcpu`.
* Executável pré-compilado `separar-pdf-pag1.exe` para Windows.

## Instalação e Configuração

1. Clone ou baixe este repositório para o seu computador.
2. (Opcional) Para registrar no menu de contexto do Windows Explorer:
   - Copie os arquivos da pasta para `C:\MenuContexto\Separar-PDF-Pag1\`.
   - Execute o arquivo `separar-pdf-pag1.reg`.
3. (Opcional) Para compilar a partir do código Go:
   ```bash
   go build -o separar-pdf-pag1.exe main.go
   ```

## Como Usar

Execute o utilitário passando o diretório alvo ou diretamente na pasta dos arquivos:

```bash
# Execução informando o caminho do diretório
.\separar-pdf-pag1.exe "C:\Caminho\Para\Pasta"
```

## Contribuindo

Sugestões e correções são bem-vindas. Abra uma Issue ou envie um Pull Request.

## Licença

TODO: definir licença. Veja o arquivo `LICENSE` para mais detalhes.

## Autor

Rodrigo Peçanha

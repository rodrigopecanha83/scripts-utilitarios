# Renomear-Fotos-Data-PlusCode-Vistoria-Sub-PDF-Comp

Ferramenta concorrente em Go para renomeação de imagens por EXIF/PlusCode e compilação de relatórios de vistoria em PDF processados em memória RAM.

## Funcionalidades Principais

* **Renomeação Concorrente:** Extrai data e coordenadas GPS dos metadados EXIF e renomeia imagens com data e Plus Code.
* **Processamento Recursivo:** Percorre e processa automaticamente todas as subpastas da pasta alvo em paralelo.
* **Geração Automática de Relatório:** Cria um documento `vistoria.pdf` em cada pasta contendo as fotos em alta resolução (duas fotos por página).
* **Arquitetura em Memória RAM:** Executa redimensionamento e rotação de imagens diretamente na memória sem arquivos temporários no disco.

## Pré-requisitos

* Go 1.24+ (caso vá compilar a partir do código-fonte `main.go`).
* Dependências Go declaradas em `go.mod`.
* Executável pré-compilado `renomear-fotos-data-pluscode-vistoria-sub-pdf-comp.exe` para Windows.

## Instalação e Configuração

1. Clone ou baixe este repositório para o seu computador.
2. (Opcional) Para registrar no menu de contexto do Windows Explorer:
   - Copie os arquivos da pasta para `C:\MenuContexto\Renomear-Fotos-Data-PlusCode-Vistoria-Sub-PDF-Comp\`.
   - Execute o arquivo `renomear-fotos-data-pluscode-vistoria-sub-pdf-comp.reg`.
3. (Opcional) Para compilar a partir do código Go:
   ```bash
   go build -o renomear-fotos-data-pluscode-vistoria-sub-pdf-comp.exe main.go
   ```

## Como Usar

Execute o utilitário passando o diretório com as subpastas como argumento ou no diretório atual:

```bash
# Execução na pasta atual
.\renomear-fotos-data-pluscode-vistoria-sub-pdf-comp.exe

# Ou passando o caminho do diretório
.\renomear-fotos-data-pluscode-vistoria-sub-pdf-comp.exe "C:\Caminho\Para\Pasta"
```

## Contribuindo

Sugestões e correções são bem-vindas. Abra uma Issue ou envie um Pull Request.

## Licença

TODO: definir licença. Veja o arquivo `LICENSE` para mais detalhes.

## Autor

Rodrigo Peçanha

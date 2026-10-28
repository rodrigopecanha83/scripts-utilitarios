# Renomear-Fotos-Data-PlusCode

Utilitário concorrente em Go para renomear imagens utilizando metadados EXIF de data/hora e geolocalização convertida em Plus Code.

## Funcionalidades Principais

* **Extração de Data EXIF:** Lê a data/hora original da fotografia gravada nos metadados da câmera.
* **Conversão em Plus Code:** Converte coordenadas de GPS em Plus Codes (Open Location Code) de 11 dígitos.
* **Processamento Concorrente:** Executa a renomeação em paralelo utilizando goroutines atreladas aos núcleos de CPU.
* **Tratamento de Conflitos:** Insere sufixos ordenados (`_1`, `_2`) caso já existam arquivos com o mesmo nome gerado.

## Pré-requisitos

* Go 1.20+ (caso vá compilar a partir do código-fonte `main.go`).
* Dependências Go:
  * `github.com/google/open-location-code/go/olc`
  * `github.com/rwcarlsen/goexif/exif`
* Executável pré-compilado `renomear-fotos-data-pluscode.exe` para Windows.

## Instalação e Configuração

1. Clone ou baixe este repositório para o seu computador.
2. (Opcional) Para registrar no menu de contexto do Windows Explorer:
   - Copie os arquivos da pasta para `C:\MenuContexto\Renomear-Fotos-Data-PlusCode\`.
   - Execute o arquivo `renomear-fotos-data-pluscode.reg`.
3. (Opcional) Para compilar o executável a partir do código Go:
   ```bash
   go build -o renomear-fotos-data-pluscode.exe main.go
   ```

## Como Usar

Execute o utilitário informando o caminho da pasta de imagens ou diretamente no diretório atual:

```bash
# Execução no diretório atual
.\renomear-fotos-data-pluscode.exe

# Ou informando o caminho do diretório
.\renomear-fotos-data-pluscode.exe "C:\Caminho\Para\Fotos"
```

## Contribuindo

Sugestões e correções são bem-vindas. Abra uma Issue ou envie um Pull Request.

## Licença

TODO: definir licença. Veja o arquivo `LICENSE` para mais detalhes.

## Autor

Rodrigo Peçanha

# Convert-PDF

Utilitário em Python para converter arquivos de imagem em documentos PDF individuais de forma automatizada, organizando os arquivos originais em pasta de backup.

## Funcionalidades Principais

* **Conversão de Imagens em PDF:** Converte arquivos de imagem para documentos PDF individuais utilizando a biblioteca img2pdf.
* **Formatos de Imagem Suportados:** Compatível com imagens nos formatos `.jpg`, `.jpeg`, `.png`, `.heic`, `.webp` e `.avif`.
* **Normalização de Extensões:** Padroniza as extensões de todos os arquivos do diretório alvo para letras minúsculas antes da conversão.
* **Validação de Integridade:** Verifica a integridade de cada imagem com Pillow antes de iniciar o processo de conversão.
* **Organização dos Originais:** Move as imagens originais processadas para a subpasta `Lixo-Conversão-PDF`, preservando os arquivos.

## Pré-requisitos

* Python 3.8+
* Dependências listadas em `requirements.txt`:
  * `img2pdf==0.5.1`
  * `pillow==11.0.0`
  * `pikepdf==9.4.2`
  * `lxml==5.3.0`
  * `packaging==24.2`
  * `wrapt==1.17.0`
  * `Deprecated==1.2.15`

## Instalação e Configuração

1. Clone ou baixe este repositório para o seu computador.
2. Navegue até o diretório do projeto e instale as dependências:
   ```bash
   pip install -r requirements.txt
   ```

## Como Usar

Execute o script passando o diretório desejado como argumento ou executando diretamente na pasta alvo:

```bash
# Executar no diretório atual
python Convert-PDF.py

# Ou informando o caminho do diretório
python Convert-PDF.py "C:\Caminho\Para\Imagens"
```

## Contribuindo

Sugestões e correções são bem-vindas. Abra uma Issue ou envie um Pull Request.

## Licença

Este projeto é distribuído sob os termos da licença **GNU General Public License v3.0 (GNU GPLv3)**. Consulte o arquivo [LICENSE](LICENSE) para mais detalhes.

## Autor

Rodrigo Peçanha


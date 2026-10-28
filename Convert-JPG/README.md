# Convert-JPG

Utilitário em Python para converter arquivos de imagem de múltiplos formatos para JPG padronizado, gerenciando nomes duplicados e preservando arquivos originais.

## Funcionalidades Principais

* **Conversão Multiformato para JPG:** Converte imagens nos formatos `.png`, `.jpeg`, `.heic` e `.webp` para `.jpg` em modo RGB.
* **Normalização de Extensões:** Converte previamente todas as extensões dos arquivos da pasta para letras minúsculas.
* **Geração de Nomes Únicos:** Adiciona sufixo numérico sequencial (`_cv_1.jpg`, `_cv_2.jpg`) caso já exista um arquivo com o mesmo nome.
* **Organização dos Originais:** Move as imagens originais processadas para a subpasta `Lixo-Conversão-JPG`.

## Pré-requisitos

* Python 3.8+
* Biblioteca Pillow:
  * `pillow`

## Instalação e Configuração

1. Clone ou baixe este repositório para o seu computador.
2. Instale a biblioteca necessária:
   ```bash
   pip install pillow
   ```

## Como Usar

Execute o script informando o diretório das imagens ou diretamente na pasta desejada:

```bash
# Executar no diretório atual
python Convert-JPG.py

# Ou informando o caminho do diretório
python Convert-JPG.py "C:\Caminho\Para\Imagens"
```

## Contribuindo

Sugestões e correções são bem-vindas. Abra uma Issue ou envie um Pull Request.

## Licença

TODO: definir licença. Veja o arquivo `LICENSE` para mais detalhes.

## Autor

Rodrigo Peçanha

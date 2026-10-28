# Scripts Utilitários — Central de Automação

Repositório central de ferramentas e scripts utilitários desenvolvidos para automação de tarefas, processamento de imagens em lote, manipulação de documentos PDF e atalhos de produtividade para o Windows Explorer.

## Utilitários Disponíveis

* **`Combinar-PDFs`:** Ferramenta em Go para combinar múltiplos arquivos PDF em um único documento respeitando ordenação natural.
* **`Convert-JPG`:** Utilitário Python para converter imagens de diversos formatos (`.png`, `.jpeg`, `.heic`, `.webp`) para JPG.
* **`Convert-PDF`:** Utilitário Python para converter imagens avulsas em documentos PDF individuais.
* **`Criar-Icone-Chrome`:** Gerador em Go para criar o pacote completo de ícones (`16`, `32`, `64`, `128`, `256`, `512`) para extensões Chrome a partir de uma imagem quadrada.
* **`Insert-Data`:** Script Python para adicionar prefixo de data (`AAAA.MM.DD - `) no nome dos arquivos.
* **`Processar-Fotos`:** Automação em Python com ImageMagick para renomear, redimensionar e aplicar marca d'água em fotos.
* **`Renomear-Dir-Atual`:** Script Python para renomeação sequencial de fotos utilizando o nome da pasta atual.
* **`Renomear-Dir-Sup`:** Script Python para renomeação sequencial de fotos utilizando o nome da pasta superior.
* **`Renomear-Fotos-Base-Diretorio-Atual`:** Utilitário em Go com integração ao menu de contexto para renomear fotos pela pasta atual.
* **`Renomear-Fotos-Data-PlusCode`:** Utilitário concorrente em Go para renomear fotos com data EXIF e Plus Code (Open Location Code).
* **`Renomear-Fotos-Data-PlusCode-Vistoria-Sub-PDF-Comp`:** Processador em Go multithread em memória RAM para renomear fotos de vistoria e gerar relatórios em PDF.
* **`Separar-PDF-Limites`:** Recorta arquivos PDF mantendo apenas a primeira e a última página.
* **`Separar-PDF-Pag1`:** Recorta arquivos PDF mantendo exclusivamente a primeira página.

## Pré-requisitos Gerais

* Sistema Operacional: Windows 10/11 (para atalhos `.reg` do menu de contexto).
* Go 1.20+ (para compilação dos utilitários em Go).
* Python 3.8+ (para execução dos scripts Python).

## Instalação e Configuração

1. Clone este repositório para o seu computador:
   ```bash
   git clone https://github.com/rodrigopecanha83/scripts-utilitarios.git "C:\Projetos\Scripts"
   ```
2. Acesse a pasta do script desejado para instruções específicas de execução ou compilação.
3. Para utilitários com arquivos `.reg` de menu de contexto, siga as orientações descritas no [instructions-reg.txt](instructions-reg.txt).

## Como Usar

Cada subdiretório possui seu próprio `README.md` com instruções detalhadas de linha de comando e configuração.

## Contribuindo

Sugestões e correções são bem-vindas. Abra uma Issue ou envie um Pull Request.

## Licença

TODO: definir licença. Veja o arquivo `LICENSE` para mais detalhes.

## Autor

Rodrigo Peçanha

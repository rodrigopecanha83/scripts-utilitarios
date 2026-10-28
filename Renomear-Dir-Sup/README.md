# Renomear-Dir-Sup

Utilitário em Python para renomear em lote arquivos de imagem utilizando o nome da pasta pai (um nível acima) como prefixo sequencial.

## Funcionalidades Principais

* **Prefixo com Nome da Pasta Superior:** Identifica automaticamente o nome do diretório pai (`pasta.parent.name`) para nomear os arquivos.
* **Numeração Sequencial Ordenada:** Aplica numeração de três dígitos (`001`, `002`, `003`) em cada imagem suportada.
* **Formatos de Imagem Suportados:** Filtra e processa arquivos `.jpg`, `.jpeg`, `.png`, `.webp` e `.heic`.
* **Proteção contra Sobrescrita:** Verifica a existência prévia do arquivo gerado para evitar sobrescrita acidental.

## Pré-requisitos

* Python 3.8+
* Módulos nativos da biblioteca padrão do Python (`os`, `sys`, `pathlib`).

## Instalação e Configuração

1. Clone ou baixe este repositório para o seu computador.
2. Não há dependências externas a serem instaladas.

## Como Usar

Execute o script diretamente na pasta das imagens ou passando o caminho do diretório:

```bash
# Executar no diretório atual
python Renomear-Dir-Sup.py

# Ou passando o caminho do diretório
python Renomear-Dir-Sup.py "C:\Caminho\Para\Fotos"
```

## Contribuindo

Sugestões e correções são bem-vindas. Abra uma Issue ou envie um Pull Request.

## Licença

TODO: definir licença. Veja o arquivo `LICENSE` para mais detalhes.

## Autor

Rodrigo Peçanha

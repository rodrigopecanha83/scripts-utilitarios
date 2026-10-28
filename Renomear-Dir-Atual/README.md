# Renomear-Dir-Atual

Utilitário em Python para renomear em lote arquivos de imagem utilizando o nome do diretório atual como prefixo numérico sequencial.

## Funcionalidades Principais

* **Prefixo com Nome da Pasta:** Adiciona o nome do diretório atual como prefixo nos arquivos de imagem.
* **Numeração Sequencial Ordenada:** Aplica numeração de três dígitos (`001`, `002`, `003`) em cada imagem.
* **Formatos de Imagem Suportados:** Filtra e processa arquivos `.jpg`, `.jpeg`, `.png`, `.webp` e `.heic`.
* **Proteção contra Sobrescrita:** Verifica a existência de nomes duplicados antes de renomear para evitar perda de dados.

## Pré-requisitos

* Python 3.8+
* Módulos nativos da biblioteca padrão do Python (`os`, `sys`, `pathlib`).

## Instalação e Configuração

1. Clone ou baixe este repositório para o seu computador.
2. Não há dependências externas a serem instaladas.

## Como Usar

Execute o script diretamente na pasta com as fotos ou passando o caminho do diretório como parâmetro:

```bash
# Executar no diretório atual
python Renomear-Dir-Atual.py

# Ou passando o caminho do diretório
python Renomear-Dir-Atual.py "C:\Caminho\Para\Fotos"
```

## Contribuindo

Sugestões e correções são bem-vindas. Abra uma Issue ou envie um Pull Request.

## Licença

Este projeto é distribuído sob os termos da licença **GNU General Public License v3.0 (GNU GPLv3)**. Consulte o arquivo [LICENSE](LICENSE) para mais detalhes.

## Autor

Rodrigo Peçanha


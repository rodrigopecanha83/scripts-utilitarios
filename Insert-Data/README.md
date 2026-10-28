# Insert-Data

Utilitário em Python para inserir prefixo com a data de criação no nome de arquivos em lote, compatível com sistemas Windows e Unix.

## Funcionalidades Principais

* **Prefixo de Data Padronizado:** Adiciona o prefixo no formato `AAAA.MM.DD - ` no início do nome de cada arquivo.
* **Detecção Multiplataforma de Criação:** Identifica a data real de criação via API do Windows (`ctime`) e `stat` no Linux/Unix, com fallback para data de modificação.
* **Prevenção de Duplicação:** Utiliza expressão regular para identificar arquivos que já possuem data no nome e evitar prefixos duplicados.

## Pré-requisitos

* Python 3.8+
* Módulos nativos da biblioteca padrão do Python (`os`, `re`, `datetime`, `pathlib`, `platform`, `subprocess`).

## Instalação e Configuração

1. Clone ou baixe este repositório para o seu computador.
2. Não há dependências externas a serem instaladas via `pip`.

## Como Usar

Navegue até a pasta que contém os arquivos que deseja renomear e execute o script:

```bash
python Insert-Data.py
```

## Contribuindo

Sugestões e correções são bem-vindas. Abra uma Issue ou envie um Pull Request.

## Licença

TODO: definir licença. Veja o arquivo `LICENSE` para mais detalhes.

## Autor

Rodrigo Peçanha

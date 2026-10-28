# Processar-Fotos

Utilitário em Python para automação de imagens, realizando renomeação sequencial, redimensionamento em lote e aplicação de marca d'água.

## Funcionalidades Principais

* **Renomeação Padronizada:** Organiza imagens em uma pasta `original` com numeração sequencial de três dígitos baseada no nome da pasta pai.
* **Redimensionamento e Otimização:** Converte e redimensiona imagens para JPG com resolução máxima de 1500x1500 e qualidade 80.
* **Aplicação de Marca d'Água:** Sobrepõe imagem de marca d'água na borda inferior das fotos processadas com ajuste geométrico via ImageMagick.
* **Limpeza Automática:** Remove arquivos temporários e formatos convertidos após o processamento.

## Pré-requisitos

* Python 3.8+
* ImageMagick (`mogrify` e `convert` disponíveis no PATH do sistema).

## Instalação e Configuração

1. Clone ou baixe este repositório para o seu computador.
2. Certifique-se de que o ImageMagick está instalado e acessível via terminal.
3. Configure o caminho do arquivo de imagem da sua marca d'água no script (variável `marca_dagua`).

## Como Usar

Execute o script informando a pasta das fotos ou diretamente na pasta atual:

```bash
# Executar no diretório atual
python Processar-Fotos.py

# Ou informando o caminho do diretório
python Processar-Fotos.py "C:\Caminho\Para\Fotos"
```

## Contribuindo

Sugestões e correções são bem-vindas. Abra uma Issue ou envie um Pull Request.

## Licença

TODO: definir licença. Veja o arquivo `LICENSE` para mais detalhes.

## Autor

Rodrigo Peçanha

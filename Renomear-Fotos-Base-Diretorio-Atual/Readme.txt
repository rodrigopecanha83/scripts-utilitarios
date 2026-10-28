# Renomear Fotos Base Diretorio Atual

Utilitário para renomear automaticamente todas as imagens de um diretório utilizando o próprio nome da pasta como prefixo sequencial.

## Funcionalidades Principais

* **Nomenclatura Baseada na Pasta:** Utiliza o nome do diretório pai como prefixo para os arquivos (ex: pasta `Viagem` gera `Viagem_1.jpg`, `Viagem_2.jpg`).
* **Renomeação Sequencial Automática:** Identifica os arquivos de imagem suportados e aplica numeração ordenada.
* **Integração com Menu de Contexto:** Arquivo `.reg` incluso para adicionar o comando diretamente ao botão direito no Windows Explorer.

## Pré-requisitos

* Sistema Operacional: Windows.
* Executável `Renomear-Fotos-Base-Diretorio-Atual.exe`.
* Nenhuma outra biblioteca ou dependência externa necessária.

## Como Instalar e Configurar

1. Clone ou baixe a pasta `Renomear-Fotos-Base-Diretorio-Atual`.
2. Para adicionar ao menu de contexto do Windows Explorer (opcional):
   - Mova a pasta com o executável e o ícone para `C:\MenuContexto\Renomear-Fotos-Base-Diretorio-Atual\`.
   - Execute o arquivo `Renomear-Fotos-Base-Diretorio-Atual.reg` para registrar os atalhos.

## Como Usar

Exemplo rápido de uso via linha de comando ou menu de contexto:

```bash
# Execução direta dentro da pasta com as fotos
.\Renomear-Fotos-Base-Diretorio-Atual.exe

# Ou passe o caminho da pasta como argumento
.\Renomear-Fotos-Base-Diretorio-Atual.exe "C:\Caminho\Para\Fotos"
```

Alternativamente, clique com o botão direito na pasta no Windows Explorer e selecione **Renomear Fotos Base Diretorio Atual**.

## Desenvolvedor:
Rodrigo Peçanha

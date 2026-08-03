================================
JUNTADOR DE ARQUIVOS PDF
================================
Versão: 1.0
Data: 21 de Agosto de 2025


DESCRIÇÃO
--------------------------------
Este é um programa utilitário para combinar múltiplos arquivos .pdf em um único documento de forma rápida.

O programa foi projetado para ser o mais simples possível: basta colocá-lo na pasta com os PDFs e executá-lo. Ele nomeia o arquivo final automaticamente com o nome da pasta.


REQUISITOS
--------------------------------
1. Sistema Operacional: Windows, macOS ou Linux.
2. Nenhuma outra dependência ou software é necessário.


COMO USAR
--------------------------------
1. Copie o arquivo executável `combinar-pdfs.exe` para a mesma pasta onde estão os seus arquivos .pdf.
2. Dê um duplo clique no executável (no Windows) ou execute-o pelo terminal.
3. O programa rodará e exibirá uma mensagem de sucesso no terminal ao concluir.
4. Um novo arquivo chamado `[nome-da-pasta]-combinado.pdf` aparecerá na mesma pasta.


FUNCIONAMENTO INTERNO
--------------------------------
- O programa busca todos os arquivos com a extensão `.pdf` no diretório onde foi executado.
- Ele ordena os arquivos encontrados usando "ordenação natural" para que `10.pdf` venha depois de `2.pdf`.
- Para a junção dos arquivos, o programa utiliza a biblioteca `pdfcpu`, uma ferramenta robusta para manipulação de PDFs em Go.
- Ao final do processo, uma mensagem de confirmação é exibida na tela.


SOLUÇÃO DE PROBLEMAS (TROUBLESHOOTING)
--------------------------------
- Em caso de sucesso, o programa exibirá a mensagem: `Arquivo '[nome-da-pasta]-combinado.pdf' criado com sucesso!`

- Se ocorrer qualquer erro durante a combinação, uma mensagem de erro será exibida diretamente na janela do terminal, e o programa será encerrado.

- Causas comuns de erro:
  - Um dos arquivos PDF de origem está corrompido ou protegido por senha.
  - O programa não possui permissão para escrever arquivos na pasta atual.

--------------------------------
Autor: Rodrigo
Atibaia/SP
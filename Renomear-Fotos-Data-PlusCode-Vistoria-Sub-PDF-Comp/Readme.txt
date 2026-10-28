# Renomear Fotos Data PlusCode Vistoria Subpastas PDF

Ferramenta de alta performance desenvolvida em Go para renomear automaticamente imagens com a data da vistoria e localização (PlusCode via dados EXIF), organizando e gerando relatórios em PDF de forma recorrente em subpastas com processamento 100% em memória RAM e multithread.

## Funcionalidades Principais

* **Renomeação Automática por EXIF:** Extrai a data da foto e coordenadas GPS do EXIF para gerar o Plus Code e renomear imagens (.jpg, .jpeg, .png, .heic).
* **Processamento Recursivo em Subpastas:** Processa concorrentemente todas as subpastas da pasta alvo utilizando múltiplas threads atreladas ao número de núcleos da CPU.
* **Geração Automática de PDF de Vistoria:** Cria automaticamente um relatório `vistoria.pdf` em cada pasta contendo as fotos em alta qualidade (2 por página) e a data da vistoria.
* **Arquitetura 100% em Memória RAM:** Redimensiona e ajusta a orientação das imagens sem gerar arquivos temporários em disco, garantindo velocidade máxima e menor desgaste do SSD/HD.
* **Integração com Menu de Contexto:** Suporte a execução via clique com o botão direito do mouse no Windows Explorer.

## Pré-requisitos

* Go 1.24+ (caso vá compilar a partir do código-fonte `main.go`).
* Executável pré-compilado `renomear-fotos-data-pluscode-vistoria-sub-pdf-comp.exe` para execução direta no Windows.

## Como Instalar e Configurar

1. Clone ou baixe a pasta `renomear-fotos-data-pluscode-vistoria-sub-pdf-comp`.
2. Para adicionar ao menu de contexto do Windows (opcional):
   - Mova a pasta ou os arquivos para `C:\MenuContexto\renomear-fotos-data-pluscode-vistoria-sub-pdf-comp\`.
   - Execute o arquivo `renomear-fotos-data-pluscode-vistoria-sub-pdf-comp.reg` para registrar o atalho no menu do botão direito do Windows Explorer.
3. Para compilar manualmente (opcional):
   ```bash
   go build -o renomear-fotos-data-pluscode-vistoria-sub-pdf-comp.exe main.go
   ```

## Como Usar

Exemplo rápido de uso via linha de comando ou menu de contexto:

```bash
# Execução direta na pasta atual
.\renomear-fotos-data-pluscode-vistoria-sub-pdf-comp.exe

# Ou passe o caminho da pasta como argumento
.\renomear-fotos-data-pluscode-vistoria-sub-pdf-comp.exe "C:\Caminho\Para\Pasta"
```

Alternativamente, clique com o botão direito sobre a pasta no Windows Explorer e selecione **Renomear Fotos Data PlusCode Vistoria Subpastas PDF**.

## Desenvolvedor:
Rodrigo Peçanha

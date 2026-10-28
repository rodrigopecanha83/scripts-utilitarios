import os
import sys
from pathlib import Path

def renomear_arquivos(pasta: Path):
    if not pasta.is_dir():
        print(f"Erro: {pasta} não é um diretório válido.")
        return

    contador = 1
    for arquivo in pasta.iterdir():
        if arquivo.is_file():
            extensao = arquivo.suffix.lower()
            if extensao in {".jpg", ".jpeg", ".png", ".webp", ".heic"}:
                novo_nome = pasta / f"{pasta.name} {contador:03d}{extensao}"

                # Evita sobrescrever arquivos existentes
                while novo_nome.exists():
                    contador += 1
                    novo_nome = pasta / f"{pasta.name} {contador:03d}{extensao}"

                arquivo.rename(novo_nome)
                contador += 1

    print("Arquivos renomeados com sucesso!")

if __name__ == "__main__":
    # Diretório passado como argumento ou diretório atual
    pasta = Path(sys.argv[1]) if len(sys.argv) > 1 else Path.cwd()
    renomear_arquivos(pasta)

# Inspect Page Tool

Tool per ispezionare il contenuto completo di una pagina specifica e vedere tutti i widget contenuti.

## Compilazione

```powershell
$env:PATH = "C:\msys64\mingw64\bin;$env:PATH"
$env:CGO_ENABLED = "1"
cd examples/inspect_page
go build -o inspect_page.exe
```

## Uso

```bash
.\inspect_page.exe <percorso_file_mpr> <nome_pagina>
```

### Esempi

Ispeziona la pagina StateMachine_Details:
```bash
.\inspect_page.exe "C:\Workspaces\Mendix\MDUI\System_Mendix_CLI\OC EX System.mpr" StateMachine_Details
```

## Output

Il tool mostra:
1. Informazioni sulla pagina (nome, ID, tipo)
2. Struttura completa in formato JSON
3. Lista di tutti i widget/controlli trovati con i loro tipi

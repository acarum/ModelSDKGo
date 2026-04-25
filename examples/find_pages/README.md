# Find Pages Tool

Tool per cercare e listare tutte le pagine in un progetto Mendix.

## Compilazione

```bash
cd examples/find_pages
go build -o find_pages.exe
```

Oppure con CGO abilitato:

```powershell
$env:PATH = "C:\msys64\mingw64\bin;$env:PATH"
$env:CGO_ENABLED = "1"
cd examples/find_pages
go build -o find_pages.exe
```

## Uso

```bash
.\find_pages.exe <percorso_file_mpr>
```

### Esempi

Lista tutte le pagine del progetto:
```bash
.\find_pages.exe "C:\Workspaces\Mendix\MDUI\System_Mendix_CLI\OC EX System.mpr"
```

Output mostra:
- Nome della pagina
- ID univoco
- Tipo del documento (Forms$Page)
- URL (se disponibile)
- ID del modulo contenitore

## Output

Il tool stampa:
1. Informazioni sul file MPR (versione, versione Mendix)
2. Lista completa di tutte le pagine con dettagli
3. Conteggio totale delle pagine trovate

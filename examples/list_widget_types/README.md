# List All Widget Types Tool

Tool per elencare tutti i tipi di widget (built-in e custom) usati in un progetto Mendix.

## Compilazione

```powershell
$env:PATH = "C:\msys64\mingw64\bin;$env:PATH"
$env:CGO_ENABLED = "1"
cd examples/list_widget_types
go build -o list_widget_types.exe
```

## Uso

```bash
.\list_widget_types.exe <percorso_file_mpr>
```

### Esempio

```bash
.\list_widget_types.exe "C:\Workspaces\Mendix\MDUI\System_Mendix_CLI\OC EX System.mpr"
```

## Output

Il tool mostra:
1. Conteggio di tutte le pagine e layout scansionati
2. **CUSTOM WIDGETS**: Widget personalizzati/community (es. DataGrid2, Charts, ecc.)
3. **SUMMARY**: Riepilogo con conteggio totale di tipi unici, custom widgets e widget Mendix built-in

Ordina i widget per numero di occorrenze (dal più usato al meno usato).

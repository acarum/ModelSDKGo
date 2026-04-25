# Find Widgets Tool

Questo tool cerca l'utilizzo di un widget specifico in tutte le pagine e layout di un progetto Mendix e stampa le sue proprietà.

## Uso

```bash
find_widgets.exe <path-to-mpr-file> <widget-type>
```

### Parametri

- `<path-to-mpr-file>`: Percorso al file .mpr del progetto Mendix
- `<widget-type>`: Tipo di widget da cercare (case-insensitive, supporta match parziali)

### Esempi

Cerca tutti i widget DataGrid:
```bash
find_widgets.exe "C:\Projects\MyApp.mpr" DataGrid
```

Cerca tutti i Button:
```bash
find_widgets.exe "C:\Projects\MyApp.mpr" Button
```

Cerca tutti i TextBox:
```bash
find_widgets.exe "C:\Projects\MyApp.mpr" TextBox
```

Cerca widget personalizzati:
```bash
find_widgets.exe "C:\Projects\MyApp.mpr" CustomWidget
```

## Output

Per ogni pagina/layout che contiene il widget cercato, il tool mostra:

- Nome della pagina/layout
- Modulo di appartenenza
- URL (per le pagine)
- Numero di widget trovati
- Per ogni widget:
  - Tipo esatto
  - Tutte le proprietà configurate

### Esempio di output

```
Page: AccountOverview
Module: Administration
URL: /admin/accounts
Found 2 widget(s):

  Widget #1:
    Type: Pages$DataGrid2
    Properties:
      Name: dataGrid1
      EntityPath: Administration.Account
      ShowPaging: true
      PageSize: 20
      Columns: [5 items]
      
  Widget #2:
    Type: Pages$DataGrid2
    Properties:
      Name: dataGrid2
      EntityPath: Administration.Role
      ShowPaging: false

--------------------------------------------------------------------------------
```

## Widget comuni da cercare

- **DataGrid** / **DataGrid2** - Griglie dati
- **Button** - Pulsanti
- **TextBox** - Campi di testo
- **DatePicker** - Selettore data
- **DropDown** / **ReferenceSelector** - Menu a discesa
- **ListView** - Liste
- **DataView** - Vista dati
- **LayoutGrid** - Griglie di layout
- **Container** - Contenitori
- **Image** - Immagini
- **Label** - Etichette

## Note

- La ricerca è case-insensitive
- Supporta match parziali (es: "Grid" trova "DataGrid", "LayoutGrid", ecc.)
- Funziona sia con progetti MPR v1 che v2
- Scansiona sia pagine che layout

## Compilazione

Per ricompilare da sorgente:

```bash
$env:PATH = "C:\msys64\mingw64\bin;$env:PATH"
$env:CGO_ENABLED = "1"
go build -o find_widgets.exe
```

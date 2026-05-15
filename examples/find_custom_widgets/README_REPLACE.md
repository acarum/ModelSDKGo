# Mendix Widget ID Replacement Tools

Tools per cercare e sostituire Widget ID nei file Mendix MPR usando **binary byte replacement** per preservare la struttura BSON.

## 🎯 Problema Risolto

Quando si modifica un Widget ID in un file Mendix MPR:
- ❌ **Metodo sbagliato**: Deserializzare BSON → Modificare → Serializzare
  - Corrompe la struttura BSON
  - Altera GUIDs e metadata
  - Causa `System.InvalidCastException` in Mendix Studio

- ✅ **Metodo corretto**: Binary byte replacement
  - Sostituisce solo i byte della stringa
  - Preserva esattamente la struttura BSON
  - Mendix Studio legge il file senza problemi

## 📦 Strumenti Disponibili

### 1. PowerShell Script (`replace-widgetid.ps1`)

Script PowerShell completo con validazione e conferma.

**Utilizzo:**
```powershell
.\replace-widgetid.ps1 `
    -MprPath "path\to\file.mpr" `
    -SourceWidgetId "siemens.DependencyGraph.DependencyGraph" `
    -DestWidgetId "siemens.dependencyGraph.DependencyGraph"
```

**Opzioni:**
- `-DryRun` - Mostra file da modificare senza applicare modifiche
- Richiede conferma prima di procedere
- Mostra offset e file modificati

**Esempio:**
```powershell
# Dry run per vedere cosa verrà modificato
.\replace-widgetid.ps1 -MprPath "MyApp.mpr" `
    -SourceWidgetId "old.widget.id" `
    -DestWidgetId "new.widget.id" `
    -DryRun

# Esegui la modifica effettiva
.\replace-widgetid.ps1 -MprPath "MyApp.mpr" `
    -SourceWidgetId "old.widget.id" `
    -DestWidgetId "new.widget.id"
```

### 2. Go Tool (`main.go`)

Tool Go con tre modalità: find, dump JSON, e replace.

**Compilazione:**
```bash
go build -o find_custom_widgets.exe .
```

**Modalità Find:**
```bash
# Cerca widget per ID (case-insensitive substring match)
go run main.go "path/to/file.mpr" "dependencygraph"

# Cerca e salva JSON delle pagine
go run main.go "path/to/file.mpr" "dependencygraph" --dump-json
```

**Modalità Replace:**
```bash
go run main.go "path/to/file.mpr" "old.widget.id" --replace "new.widget.id"
```

**Esempio completo:**
```bash
# 1. Trova widget con ID corrente
go run main.go "MyApp.mpr" "siemens.DependencyGraph.DependencyGraph"

# 2. Esporta JSON per analisi
go run main.go "MyApp.mpr" "DependencyGraph" --dump-json

# 3. Sostituisci con nuovo ID
go run main.go "MyApp.mpr" "siemens.DependencyGraph.DependencyGraph" `
    --replace "siemens.dependencyGraph.DependencyGraph"
```

## ⚠️ Requisiti Importanti

### Lunghezza Stringhe
**Le stringhe source e destination devono avere la stessa lunghezza!**

✅ Valido:
```
Source: "siemens.DependencyGraph.DependencyGraph" (41 caratteri)
Dest:   "siemens.dependencyGraph.DependencyGraph" (41 caratteri)
```

❌ Invalido:
```
Source: "old.widget" (10 caratteri)
Dest:   "new.id" (6 caratteri)  # Lunghezze diverse!
```

### Backup
**Crea sempre un backup prima di modificare!**

```bash
# Git backup
cd path/to/mendix/project
git add -A
git commit -m "Before widget ID replacement"

# Oppure copia manuale
cp "MyApp.mpr" "MyApp.mpr.backup"
cp -r "mprcontents" "mprcontents.backup"
```

### Ripristino
Se qualcosa va storto:

```bash
# Con Git
git checkout -- mprcontents/

# Oppure
cp -r "mprcontents.backup" "mprcontents"
```

## 🧪 Testing

Usa lo script di test per verificare gli strumenti:

```powershell
.\test-replace.ps1 -MprPath "path\to\file.mpr"
```

## 🔍 Come Funziona

### Binary Byte Replacement

Entrambi gli strumenti usano lo stesso approccio:

```
1. Leggi file come byte array
   [0x73, 0x69, 0x65, 0x6D, ...] 
   
2. Cerca pattern UTF-8
   "siemens.DependencyGraph.DependencyGraph"
   
3. Sostituisci byte per byte
   Old: [0x44, 0x65, 0x70] (Dep)
   New: [0x64, 0x65, 0x70] (dep)
   
4. Salva byte array modificato
```

Questo preserva:
- ✓ Struttura BSON originale
- ✓ GUIDs e metadata
- ✓ Offset e allineamento
- ✓ Tutti gli altri dati binari

## 📊 Output

### PowerShell
```
=== Mendix Widget ID Replacer ===

MPR File: MyApp.mpr
Source:   siemens.DependencyGraph.DependencyGraph
Target:   siemens.dependencyGraph.DependencyGraph

Found 3 file(s) to modify:
  • 03\ab\03ab4e3f-3c2d-4fd5-ba42-b7a93db2a96a.mxunit
    Offset: 53594
  • 42\51\425177f5-ae7a-41d5-84cf-1f845744e07a.mxunit
    Offset: 50714
  • ef\fa\effaf390-5e7c-43df-9026-3831525e3b43.mxunit
    Offset: 795824

=== Summary ===
  Success: 3 file(s)

✓ Replacement completed!
```

### Go
```
=== REPLACE MODE ===
Searching for widgets with widgetId: siemens.DependencyGraph.DependencyGraph
Will replace with: siemens.dependencyGraph.DependencyGraph

=== Found Widgets to Replace ===
  Page: OperatorLanding (1 widget(s))
  Page: Routing_Popup (1 widget(s))
  Page: Routing (1 widget(s))

Total: 3 widget(s) in 3 unit(s)

Replace? (yes/no): yes

=== Performing Replacement ===
Processing Page: OperatorLanding... ✓ Replaced 1 occurrence(s)
Processing Page: Routing_Popup... ✓ Replaced 1 occurrence(s)
Processing Page: Routing... ✓ Replaced 1 occurrence(s)

✓ Replacement completed!
```

## 🚀 Workflow Consigliato

1. **Backup del progetto**
   ```bash
   git commit -am "Before widget replacement"
   ```

2. **Trova widget correnti**
   ```bash
   go run main.go "MyApp.mpr" "OldWidgetId"
   ```

3. **Test con dry-run**
   ```powershell
   .\replace-widgetid.ps1 -MprPath "MyApp.mpr" `
       -SourceWidgetId "old.id" -DestWidgetId "new.id" -DryRun
   ```

4. **Esegui sostituzione**
   ```powershell
   .\replace-widgetid.ps1 -MprPath "MyApp.mpr" `
       -SourceWidgetId "old.id" -DestWidgetId "new.id"
   ```

5. **Verifica**
   ```bash
   go run main.go "MyApp.mpr" "new.id"
   ```

6. **Test in Mendix Studio**
   - Apri il progetto in Mendix Studio Pro
   - Verifica che le pagine si aprano senza errori
   - Testa la funzionalità dei widget

## 📝 Note Tecniche

### BSON Structure
I file `.mxunit` sono BSON binario con strutture speciali:
- GUIDs: `{$ID: {Subtype: 0, Data: "base64..."}}`
- Widget Type: `{$Type: "CustomWidgets$CustomWidget"}`
- Widget ID: `{Type: {WidgetId: "string"}}`

La deserializzazione/serializzazione BSON può alterare queste strutture.

### File Location
```
MyApp.mpr                           # Database file
mprcontents/
  XX/
    YY/
      UUID.mxunit                   # Binary BSON files
```

### UTF-8 Encoding
Le stringhe nei file BSON sono UTF-8 encoded:
```
"siemens" = [0x73, 0x69, 0x65, 0x6D, 0x65, 0x6E, 0x73]
```

## 🐛 Troubleshooting

### "Pattern lengths must match"
Le stringhe devono avere la stessa lunghezza in byte.

**Soluzione:** Usa stringhe della stessa lunghezza:
```
✓ "siemens.DependencyGraph.DependencyGraph" → "siemens.dependencyGraph.DependencyGraph"
✗ "old.id" → "new.widget.id"
```

### "System.InvalidCastException" in Mendix Studio
Il file è stato corrotto.

**Soluzione:** 
```bash
# Ripristina dal backup
git checkout -- mprcontents/

# Usa binary byte replacement, non BSON unmarshal/marshal
```

### "No files found with pattern"
Il pattern potrebbe essere già stato sostituito o non esistere.

**Soluzione:**
```bash
# Cerca con substring match
go run main.go "MyApp.mpr" "dependency"

# Controlla con case-insensitive
go run main.go "MyApp.mpr" "DependencyGraph"
```

## 📚 Riferimenti

- **Mendix MPR Format**: Binary BSON database
- **BSON Specification**: http://bsonspec.org/
- **Go BSON Driver**: go.mongodb.org/mongo-driver/bson
- **Mendix Model SDK**: github.com/anthropics/modelsdk-go

## ✅ Checklist Finale

Prima di usare in produzione:

- [ ] Backup completo del progetto
- [ ] Test con dry-run
- [ ] Verifica lunghezze stringhe match
- [ ] Test su copia del progetto
- [ ] Verifica apertura in Mendix Studio
- [ ] Test funzionalità widget
- [ ] Commit modifiche a Git

---

**Autori:** Script sviluppati per gestire Widget ID Mendix
**Data:** Maggio 2026
**License:** MIT

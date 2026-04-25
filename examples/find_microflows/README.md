# Finds Microflows Tool

## Descrizione

Tool per cercare tutti i microflow in un progetto Mendix che:
- Chiamano un microflow specifico tramite "Call Microflow Activity"
- Chiamano una Java Action specifica tramite "Java Action Call"
- Chiamano External Actions tramite "Call External Action" (OData services)

## Compilazione

```powershell
$env:PATH = "C:\msys64\mingw64\bin;$env:PATH"
$env:CGO_ENABLED = "1"
go build -o examples/finds_microflows/finds_microflows.exe examples/finds_microflows/main.go
```

**Nota**: Richiede CGO abilitato per il supporto SQLite3.

## Utilizzo

```powershell
.\examples\finds_microflows\finds_microflows.exe <mpr_file_path> [target_microflow_name] [java_action_name] [external_action_name]
```

### Parametri

- `mpr_file_path`: Percorso completo del file .mpr del progetto Mendix
- `target_microflow_name`: (opzionale) Nome del microflow da cercare. Default: "EXFN_ServiceLayer.CallCommand_MF"
- `java_action_name`: (opzionale) Nome della Java Action da cercare. Default: "CallCommandAction"
- `external_action_name`: (opzionale) Nome dell'External Action da cercare. Default: "" (tutte)

### Esempi

Cerca microflow, Java Action e tutte le External Actions con valori default:
```powershell
.\examples\finds_microflows\finds_microflows.exe "C:\path\to\project.mpr"
```

Cerca solo External Actions (disabilita gli altri):
```powershell
.\examples\finds_microflows\finds_microflows.exe "C:\path\to\project.mpr" "xxxNone" "xxxNone"
```

Cerca una specifica External Action:
```powershell
.\examples\finds_microflows\finds_microflows.exe "C:\path\to\project.mpr" "xxxNone" "xxxNone" "UpdateStatus"
```

Cerca microflow e Java Action specifici (tutte le External Actions):
```powershell
.\examples\finds_microflows\finds_microflows.exe "C:\path\to\project.mpr" "MyModule.MyMicroflow" "MyJavaAction"
```

Cerca usando nomi parziali:
```powershell
.\examples\finds_microflows\finds_microflows.exe "C:\path\to\project.mpr" "CallCommand" "CommandAction" "Status"
```

## Output

Il tool stampa:
- Nome del microflow che contiene la chiamata
- Modulo a cui appartiene il microflow
- Numero di chiamate trovate
- Per ogni chiamata:
  - Tipo (MicroflowCall, JavaAction, o ExternalAction)
  - Nome del microflow chiamato (se MicroflowCall)
  - Nome della Java Action chiamata (se JavaAction)
  - Nome dell'External Action chiamata (se ExternalAction)
  - Nome dell'activity
  - Caption dell'activity
  - **AppName** (sempre mostrato, anche se vuoto)
    - Per MicroflowCall/JavaAction: estratto dai parametri
    - Per ExternalAction: nome dell'action (campo Name)
  - **CommandName** (sempre mostrato, anche se vuoto)
    - Per MicroflowCall/JavaAction: estratto dai parametri
    - Per ExternalAction: valore del parametro "command"

**Nota**: I valori possono essere:
- Variabili: `$VariableName` (valori dinamici passati al runtime)
- Stringhe letterali: `'Value'` (valori hardcoded nel codice)
- Espressioni: qualsiasi espressione Mendix valida
- Vuoto: se il parametro non è configurato

### Esempio Output

```
Opening MPR: C:\Workspaces\Mendix\MDUI\System_Mendix_CLI\OC EX System.mpr
Searching for microflows calling: CallCommand_MF
Searching for Java Action: CallCommandAction

Found 124 microflows. Searching for calls to 'CallCommand_MF' and Java Action 'CallCommandAction'...

✓ Microflow: SUB_HideAndUnhideCommand
  Module: Documents
  Calls found: 2
  [1] Type: MicroflowCall
      Microflow: EXFN_ServiceLayer.CallCommand_MF
      Caption: Activity
      AppName: $AppName
      CommandName: $CommandName
  [2] Type: JavaAction
      JavaAction: EXFN_ServiceLayer.CallCommandAction
      Caption: Activity
      AppName: 'AppName'
      CommandName: 'CommandName'

✓ Microflow: SUB_CallCommand
  Module: Documents
  Calls found: 1
  [1] Type: MicroflowCall
      Microflow: EXFN_ServiceLayer.CallCommand_MF
      Caption: Activity
      AppName: $AppName
      CommandName: $CommandName

=== Summary ===
Total microflows scanned: 124
Microflows calling 'CallCommand_MF' or Java Action 'CallCommandAction': 5
```

**Interpretazione dei valori**:
- `$AppName`, `$CommandName`: Variabili Mendix (valori passati dinamicamente)
- `'AppName'`, `'CommandName'`: Stringhe letterali hardcoded nel codice
- Vuoto: Parametro non configurato o non trovato

## Note Tecniche

- Supporta Mendix MPR v2 (con folder mprcontents/)
- Cerca ricorsivamente in tutti gli oggetti del microflow
- La ricerca è case-sensitive
- Usa pattern matching con `strings.Contains()` per supportare ricerche parziali
- Estrae automaticamente i parametri AppName e CommandName dai ParameterMappings
- Supporta MicroflowCall, JavaActionCall e CallExternalAction nello stesso microflow
- **Mostra sempre AppName e CommandName** per ogni chiamata trovata

### Tipi di chiamate supportati

| Tipo | Mendix Type | AppName da | CommandName da |
|------|-------------|------------|----------------|
| **MicroflowCall** | Microflows$MicroflowCallAction | Parametro "appname" | Parametro "commandname" |
| **JavaAction** | Microflows$JavaActionCallAction | Parametro "appname" | Parametro "commandname" |
| **ExternalAction** | Microflows$CallExternalAction | Campo "Name" dell'action | Parametro "command" |

### Tipi di valori dei parametri

| Formato | Esempio | Significato |
|---------|---------|-------------|
| `$VariableName` | `$AppName` | Variabile locale o parametro del microflow |
| `'String'` | `'AppName'` | Stringa letterale hardcoded |
| Espressione | `$Object/Name` | Espressione Mendix (navigazione, funzioni, etc.) |
| Vuoto | | Parametro non configurato |

## Struttura Interna

Il tool cerca `ActionActivity` con tipo "Microflows$ActionActivity" che contengono:

### MicroflowCall
1. Un `Action` di tipo "Microflows$MicroflowCallAction"
2. Con un `MicroflowCall` che ha il campo `Microflow` uguale o contenente il target
3. Estrae i parametri dai `ParameterMappings`:
   - Cerca parametri con nome contenente "appname" (case-insensitive)
   - Cerca parametri con nome contenente "commandname" (case-insensitive)
   - Mostra il valore dell'`Argument` per ciascun parametro trovato

### JavaAction
1. Un `Action` di tipo "Microflows$JavaActionCallAction"
2. Con campo `JavaAction` uguale o contenente il target
3. Estrae i parametri dai `ParameterMappings`:
   - Legge il campo `Parameter` per il nome del parametro
   - Estrae l'`Argument` da dentro l'oggetto `Value` (struttura: `Value.Argument`)
   - Rimuove caratteri di formattazione (newline, spazi) dall'Argument
   - Identifica AppName e CommandName (case-insensitive nel nome del parametro)
   - Mostra sempre i valori (anche se vuoti)

### ExternalAction (NEW!)
1. Un `Action` di tipo "Microflows$CallExternalAction"
2. Con campo `Name` uguale o contenente il target (se specificato)
3. Estrae i parametri:
   - `AppName`: campo `Name` dell'action (es. "UpdateStatus", "CreateStatusDefinition")
   - `CommandName`: valore del parametro con `ParameterName` = "command"
   - Legge `Argument` direttamente dai `ParameterMappings` (tipo `ExternalActionParameterMapping`)
4. Se `external_action_name` è vuoto, trova TUTTE le External Actions nel progetto

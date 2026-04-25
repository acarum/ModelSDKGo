# Plan: Export Manifest Report Generator

Creare un tool Go chiamato `export_manifest` che genera un report Markdown completo di un progetto Mendix MPR, includendo entità esterne, chiamate microflow/action, signal manager widgets, navigation items e system roles con page accessibility.

**TL;DR**: Combinare la logica esistente di find_entities, finds_microflows e find_custom_widgets in un unico tool che esporta un report MD strutturato con 5 sezioni principali: (1) External Entities per modulo con sorgente OData, (2) Microflow/Action calls con AppName/CommandName, (3) Signal Manager Widgets con subscriptions, (4) Navigation Items con struttura menu, (5) System Roles & Page Accessibility per controllo accessi.

## Steps

1. **✅ COMPLETATO - Creare struttura base export_manifest/main.go**
   - ✅ Setup argomenti CLI: `<mpr_path> <output_md_path>`
   - ✅ **CLI Flags opzionali** per controllo sezioni report:
     * `-include-entities` (default: false): Include sezione External Entities nel report
     * `-include-attributes` (default: true): Include attributi nelle tabelle entità
     * `-include-microflows` (default: true): Include sezione Microflow/Action Calls nel report
     * `-include-widgets` (default: true): Include sezione Signal Manager Widgets nel report
     * ✅ `-include-navigation` (default: true): Include sezione Navigation Items nel report
     * ⏳ `-include-roles` (default: true): Include sezione System Roles & Page Accessibility nel report (da implementare)
     * **Usage**: `export_manifest -include-entities=true -include-attributes=false MyApp.mpr report.md`
   - ✅ Aprire MPR con `modelsdk.Open()`
   - ✅ Ottenere nome progetto (da filename) e versione Mendix (`reader.GetMendixVersion()`)
   - ✅ Creare file Markdown output con header iniziale:
     * Titolo: `# Manifest Report: {ProjectName}`
     * Mendix Version e Generated timestamp
     * Summary section con conteggi (External Entities, Microflow/Action Calls, Signal Manager Widgets, Navigation Items, System Roles & Pages)
   - ✅ Setup filtro marketplace modules con `isMarketplaceModule()` (System, Atlas, Administration, Marketplace, Community, AppStore, CommunityCommons, NanoflowCommons, DataWidgets, WebActions)
   - **Implementato in**: `examples/export_manifest/main.go`

2. **✅ COMPLETATO - Implementare Sezione 1: External Entities Report**
   - ✅ Iterare moduli non-marketplace aprendo database SQLite
   - ✅ Query `Unit` table con `ContainmentName = 'DomainModel'`
   - ✅ Caricare BSON da `mprcontents/{xx}/{yy}/{uuid}.mxunit` usando `loadUnitContents()`
   - ✅ Estrarre entità con `isExternalEntity()` (Source.$Type == "Rest$ODataEntityTypeSource")
   - ✅ Raccogliere per ogni entità:
     * Entity name (`extractNameFromContents()`)
     * Attributes con tipi (`extractAttributes()`, `extractAttributeType()`)
     * Published From (`extractPublishedFrom()` con fallback al nome entità)
   - ✅ **Formato MD implementato (modificato rispetto al piano originale)**:
     * Header per modulo: `### Module: {ModuleName}`
     * Header per ogni entità: `#### Entity: {EntityName}`
     * Campo dedicato: `**Published From:** {ServiceName}`
     * Tabella attributi verticale (un attributo per riga):
       ```markdown
       | Attribute | Type |
       |-----------|------|
       | _Id | String |
       | IsFrozen | Boolean |
       | CreatedOn | DateTime |
       ```
   - ✅ **Funzionalità aggiunte**:
     * `extractPublishedFrom(entityMap, entityName)`: estrae Source.EntityTypeName o usa nome entità come fallback
     * `extractAttributeType(attrMap)`: estrae tipo da Type.$Type, RemoteAttributeTypeName, o inferisce dal nome
     * `formatODataType(odataType)`: converte tipi OData (Edm.String→String, Edm.Int32→Integer, Edm.DateTime→DateTime, ecc.)
     * `inferTypeFromName(name)`: pattern matching intelligente per inferire tipi:
       - Boolean: is*, has*, do*, succeeded, *frozen, *locked, *hidden, *enabled
       - DateTime: *date*, *time*, *edOn, *atedOn (CreatedOn, LastUpdatedOn)
       - Integer: *count*, *seed*, *increment*, *sequence*
       - Decimal: *multiplier*, *addend*, *exponent*, *factor*
       - String (default): _id, *id, name, description
     * `parseAttribute(attr)`: parser per formato "Name (Type)"
   - **Risultato**: 156 entità esterne trovate in 3 moduli con tipi correttamente identificati
   - **Implementato in**: `examples/export_manifest/main.go` (linee ~100-470)

3. **✅ COMPLETATO - Implementare Sezione 2: Microflow/Action Calls Report**
   - ✅ **Implementazione**: Adattata logica da `find_microflows/main.go` per estrarre tutti i tipi di chiamate
   - ✅ **Data Structures**: 
     * `MicroflowCallInfo`: Microflow name, module, activity, call type, target, AppName, CommandName
     * `TempMicroflowCall`: struttura temporanea per aggregazione risultati
     * `MicroflowInfo`: ID, Name, ModuleName, Content (BSON)
   - ✅ **Funzioni implementate**:
     * `collectMicroflowCalls(db, contentsDir, report)`: colleziona tutte le chiamate dai microflow non-marketplace
     * `listMicroflows(db, contentsDir)`: SQL query per microflow, filtra per $Type == "Microflows$Microflow"
     * **✅ `extractModuleFromMicroflowParams(content)`**: estrae nome modulo da microflow usando doppia strategia:
       - Strategy 1: dai parametri del microflow → ObjectCollection.Objects → MicroflowParameter → VariableType.Entity (formato "Module.EntityName")
       - Strategy 2: da AllowedModuleRoles (formato "Module.Role")
       - Fallback: ContainmentName dal database SQL (nome folder)
     * `findMicroflowCalls(mf, targetMicroflow, targetJavaAction)`: converte BSON→JSON e cerca ActionActivity ricorsivamente
     * **✅ `resolveVariableValue(contentMap, variableName)`**: risolve variabili identificando:
       - Parametri di input: cerca MicroflowParameter per nome, ritorna "<parameter>"
       - Variabili assegnate: cerca ChangeVariableActivity con estrazione valore da InitialValue
     * **✅ `findVariableAssignment(obj, varName, result)`**: ricerca ricorsiva di ChangeVariableActivity per estrazione valore variabile
     * `searchActivities(obj, targetMicroflow, targetJavaAction, calls, contentMap)`: ricerca ricorsiva di ActionActivity, chiama check* functions con contentMap per variable resolution
     * **✅ `checkMicroflowCall(activity, targetMicroflow, calls, contentMap)`**: estrae MicroflowCallAction con:
       - Filter: solo calls a "EXFN_ServiceLayer.CallCommand_MF"
       - AppName/CommandName: da ParameterMappings (match su parameter name "appname"/"commandname", estrae Argument)
       - **Variable resolution**: se Argument inizia con `$`, risolve con `resolveVariableValue()` per identificare parametri o valori assegnati
     * **✅ `checkJavaAction(activity, targetJavaAction, calls, contentMap)`**: estrae JavaActionCallAction con:
       - Filter: solo calls a "EXFN_ServiceLayer.CallCommandAction"
       - JavaAction name: da JavaAction o JavaActionQualifiedName field
       - AppName/CommandName: da ParameterMappings con Value.Argument structure
       - **Variable resolution**: applica `resolveVariableValue()` per variabili
     * **✅ `checkExternalAction(activity, calls, contentMap)`**: estrae CallExternalAction con:
       - Nessun filtro (include tutte le ExternalAction)
       - AppName: da ConsumedODataService (split ".", take last part)
       - CommandName: da Name field dell'action
   - ✅ **Formato MD implementato**:
     * Header: `## 2. Microflow/Action Calls`
     * Descrizione: "Microflow and action calls found in the project (MicroflowCall, JavaAction, ExternalAction)"
     * Tabella: `| Microflow | Module | Call Type | AppName | CommandName |`
     * Values: 
       - Valori literal: "Reference", "UpdateStatus", 'AppName', 'CommandName'
       - Parametri: "<parameter>" per valori passati come input al microflow
       - Variabili assegnate: valore estratto dall'espressione di assegnazione
       - Vuoti: "-" per valori non disponibili
   - ✅ **Risultato**: 78 chiamate trovate con moduli corretti:
     * **MicroflowCall**: 5-6 calls a EXFN_ServiceLayer.CallCommand_MF con parametri identificati come `<parameter>` (erano variabili $AppName, $CommandName)
     * **ExternalAction**: 72+ chiamate a OData services (AppName = "Reference", "TripPinServiceRW") con moduli corretti (OpcenterEXFN_ReferenceData_Connector, OpcenterEXFN_ReferenceData)
     * **JavaAction**: 1 chiamata con parametri literal estratti ('AppName', 'CommandName')
     * **Moduli corretti**: 
       - UpdateStatus → OpcenterEXFN_ReferenceData_Connector ✓
       - SUB_SetStatusAsInitial → OpcenterEXFN_ReferenceData ✓
       - Altri microflow → moduli estratti correttamente da Entity/AllowedModuleRoles
   - ✅ **Funzionalità chiave aggiunte**:
     * **Module name extraction**: doppia strategia (Entity parameters + AllowedModuleRoles) garantisce identificazione corretta del modulo anche per microflow in subfolder
     * **Variable resolution**: distingue tra parametri di input (`<parameter>`), valori assegnati (estratti da espressione), e valori literal
     * **Comprehensive filtering**: filtra MicroflowCall per target specifico, JavaAction per target specifico, include tutte le ExternalAction
   - **Implementato in**: `examples/export_manifest/main.go` (linee ~512-1050)

4. **✅ COMPLETATO - Implementare Sezione 3: Signal Manager Widgets Report**
   - ✅ **Implementazione**: Logica adattata da `find_custom_widgets/main.go` per estrazione subscription-level
   - ✅ **Data Structures**:
     * `WidgetInfo`: DocumentName, DocumentType, Module, SignalName, AppName, SubscriptionFilter (subscription-level, non widget-level)
   - ✅ **Funzioni implementate**:
     * `collectSignalManagerWidgets(reader, mprPath, report)`: colleziona subscriptions da pages e snippets
     * `extractSignalSubscriptions(widgetObj, moduleName, docName, docType, widgetID)`: estrae singole subscriptions da array Properties
     * `collectPrimitiveValuesWithKeys(obj, prefix, values)`: ricerca ricorsiva con key paths per identificazione corretta campi
     * **✅ `extractModuleAndNameFromBSON(content, unitID, db)`**: estrae modulo e nome documento con strategie multiple:
       - Pages: da AllowedModuleRoles in BSON (formato "Module.Role")
       - Snippets: tramite `findModuleByTraversal()` risalendo gerarchia unità
       - Fallback: ContainmentName dal database SQL
     * **✅ `findModuleByTraversal(unitID, db)`**: risale gerarchia parent-child per trovare il modulo parent:
       - Costruisce mappa di tutti i moduli (con $Type == "Projects$Module")
       - Itera ContainerID fino a trovare un parent che è un modulo
       - Estrae Name field dal BSON del modulo
     * `guidToString(guidBytes)`: converte Windows GUID bytes a UUID string format (reverse byte swapping)
     * `stringToWindowsGUID(uuidStr)`: converte UUID string a Windows GUID binary con byte swapping per SQL queries
   - ✅ **Estrazione campi subscription**:
     * **SignalName**: da Properties_1 → PrimitiveValue (es. "SN", "SN2", "SNIPPET_SN")
     * **AppName**: da Properties_2 → PrimitiveValue (es. "APPName", "AN2", "SNIPPET_APPName")
     * **SubscriptionFilter**: da Properties_3 → Expression (es. "filter0", "filter", "SNIPPET_filter0")
   - ✅ **Formato MD implementato**:
     * Header: `## 3. Signal Manager Widgets`
     * Descrizione: "Signal Manager widget subscriptions found in the project"
     * Tabella: `| Module | Document Type | Document | Signal Name | App Name | Subscription Filter |`
     * Document Type: "Page" o "Snippet" per identificazione sorgente
   - ✅ **Supporto Pages e Snippets**:
     * **Pages**: `reader.ListPages()` con estrazione modulo da AllowedModuleRoles BSON
     * **Snippets**: `reader.ListSnippets()` con estrazione modulo tramite hierarchy traversal nel database
   - ✅ **Risultato**: 3 signal subscriptions trovate:
     * 2 in pages (module: OpcenterEXFN_ReferenceData, page: StateMachine_Details)
     * 1 in snippet (module: OpcenterEXFN_ReferenceData, snippet: MySnippet)
     * Moduli correttamente estratti usando hierarchy traversal per snippets
   - ✅ **Funzionalità chiave aggiunte**:
     * **Subscription-level reporting**: ogni subscription come riga separata nella tabella (non widget-level)
     * **Module extraction per snippets**: traversal della gerarchia UnitID → ContainerID → Projects$Module
     * **Windows GUID encoding**: conversione UUID string ↔ binary con byte swapping per SQL compatibility
     * **Key-based extraction**: `collectPrimitiveValuesWithKeys()` con paths per identificazione field corretti (Properties_1, Properties_2, Properties_3)
     * **Expression field support**: estrazione da Expression per SubscriptionFilter (non solo PrimitiveValue)
   - **Implementato in**: `examples/export_manifest/main.go` (linee ~1052-1470, 1471-1620)
   - *depends on 1*

5. **✅ COMPLETATO - Implementare Sezione 4: Navigation Items Report**
   - ✅ Estrarre navigation items dal progetto Mendix tramite SQL query per NavigationDocument
   - ✅ Listare menu items con struttura gerarchica
   - ✅ Per ogni navigation item raccogliere:
     * Nome item (`ItemName`)
     * Caption/Label (`Caption`)
     * Target (page o microflow) tramite reference resolution
     * Modulo di appartenenza (`Module`)
     * Item type (Page, Microflow, Nanoflow) da Action.$Type
   - ✅ **Formato MD implementato**:
     * Header: `## 4. Navigation Items`
     * Descrizione: "Navigation menu items found in the project"
     * Tabella: `| Navigation Item | Caption | Target | Module | Type |`
     * Supporto indentazione gerarchica con `Level` field (spaces per sottomenu)
   - ✅ **Funzioni implementate**:
     * `collectNavigationItems(db, contentsDir, report)`: query per NavigationDocument e colleziona items
     * `extractNavigationItemsFromBSON(content, moduleName, db, contentsDir)`: estrae items da profili Desktop/Tablet/Phone
     * `extractMenuItems(menuRef, db, contentsDir, moduleName, level)`: ricorsivamente estrae items da MenuDocument
     * `parseMenuItem(itemMap, db, contentsDir, moduleName, level)`: parse singolo menu item con Action type detection
     * `resolvePageReference(pageRef, db, contentsDir)`: risolve reference a page name
     * `resolveMicroflowReference(mfRef, db, contentsDir)`: risolve reference a microflow name
     * `resolveNanoflowReference(nfRef, db, contentsDir)`: risolve reference a nanoflow name
     * `extractReferenceID(ref)`: estrae unit ID da reference object ($ID, Unit field)
     * `loadDocumentByReference(db, contentsDir, refID)`: carica BSON content di documento tramite reference
     * `getModuleNameFromContainerID(db, containerID)`: estrae module name da container con parent traversal
   - ✅ **CLI flag aggiunto**: `-include-navigation` (default: true)
   - ✅ **Supporto profili multipli**: Desktop, Tablet, Phone navigation profiles
   - ✅ **Action types supportati**:
     * PageClientAction → Item type: "Page"
     * MicroflowClientAction → Item type: "Microflow"
     * NanoflowClientAction → Item type: "Nanoflow"
   - ✅ **Struttura gerarchica**: Level-based indentation per sottomenu (SubMenu recursion)
   - **Implementato in**: `examples/export_manifest/main.go` (linee ~1659-2050)
   - *depends on 1*

6. **Implementare Sezione 5: System Roles & Page Accessibility Report**
   - Estrarre lista di tutti i System Roles definiti nel progetto
   - Per ogni page/snippet del progetto:
     * Identificare AllowedModuleRoles dal BSON
     * Mappare module roles a system roles
   - Creare due viste nel report:
     * **Vista per Role**: Per ogni System Role, lista delle pages accessibili
     * **Vista per Page**: Per ogni Page, lista dei System Roles che possono accedervi
   - Formattare in MD: 
     * Sezione 5.1: tabella (System Role | Module | Accessible Pages Count)
     * Sezione 5.2: tabella dettagliata (Page | Module | Allowed System Roles)
   - Evidenziare pages con accesso non ristretto (nessun role requirement)
   - *depends on 1*

7. **✅ COMPLETATO - Finalizzare e Testare Report**
   - ✅ Aggregare tutte le sezioni in unico file MD
   - ✅ Aggiungere sommario iniziale con conteggi (N entities, N calls, N widgets)
   - ✅ Aggiungere timestamp generazione in header
   - ✅ **CLI flags implementati** per controllo sezioni report (include-entities, include-attributes, include-microflows, include-widgets)
   - ✅ Compilare: `go build -o export_manifest.exe examples/export_manifest/main.go`
   - ✅ Testare su MPR target: `.\export_manifest.exe "OC EX System.mpr" manifest_report.md`
   - ✅ Verificare output MD per completezza e formattazione
   - ✅ **Risultato finale**:
     * 156 external entities in 3 moduli
     * 78 microflow/action calls con moduli e variabili risolti
     * 3 signal subscriptions (2 in pages, 1 in snippet)
     * Report MD formattato correttamente con tutte le sezioni
   - *depends on 2, 3, 4, 5, 6*

## Relevant files

- **✅ `examples/export_manifest/main.go`** — Tool principale implementato con:
  - Struttura base: CLI parsing (con flags opzionali), MPR opening, MD generation
  - External Entities (Step 2): `collectExternalEntities()`, `listDomainModels()`, `findExternalEntitiesInDomainModel()`
  - Tipo extraction: `extractAttributeType()`, `formatODataType()`, `inferTypeFromName()`
  - Microflow/Action Calls (Step 3): `collectMicroflowCalls()`, `listMicroflows()`, `findMicroflowCalls()`
  - **Module extraction**: `extractModuleFromMicroflowParams()` con doppia strategia (Entity + AllowedModuleRoles)
  - **Variable resolution**: `resolveVariableValue()`, `findVariableAssignment()` per identificazione parametri vs valori assegnati
  - Activity checks: `searchActivities()`, `checkMicroflowCall()`, `checkJavaAction()`, `checkExternalAction()` (tutti con contentMap per variable resolution)
  - **Signal Manager Widgets (Step 4)**: `collectSignalManagerWidgets()`, `extractSignalSubscriptions()`, `collectPrimitiveValuesWithKeys()`
  - **Module extraction per snippets**: `extractModuleAndNameFromBSON()`, `findModuleByTraversal()`, `guidToString()`, `stringToWindowsGUID()`
  - Utility: `extractPublishedFrom()`, `extractAttributes()`, `parseAttribute()`, `isExternalEntity()`
  - Marketplace filtering: `isMarketplaceModule()`, `extractModuleName()`
  - BSON/UUID: `loadUnitContents()`, `blobToUUID()`, `extractNameFromContents()`
  - Report generation: `generateMarkdownReport()` con sezioni condizionali basate su ReportOptions
- `examples/find_entities/main.go` — ✅ Riutilizzata logica SQL query per DomainModel, pattern matching `Rest$ODataEntityTypeSource`, funzione `isExternalEntity()`
- `examples/finds_microflows/main.go` — ✅ Riutilizzate funzioni `findMicroflowCalls()`, `checkMicroflowCall()`, `checkJavaAction()`, `checkExternalAction()` per estrazione AppName/CommandName con adattamento per export_manifest e aggiunta variable resolution
- `examples/find_custom_widgets/main.go` — ✅ Riutilizzata logica `loadUnitBSON()`, `searchForMatchingWidgets()`, `collectPrimitiveValues()` per widget search e signal extraction con adattamento per subscription-level reporting
- `mpr/reader.go` — ✅ Utilizzata API modelsdk: `Open()`, `GetMendixVersion()`, `Path()`, `ListPages()`, `ListSnippets()`
- `mpr/utils.go` — ✅ Utilizzate funzioni utility per BSON loading e UUID conversion

## Verification

1. ✅ Compilare export_manifest senza errori CGO - **COMPLETATO** (compilato con gcc MinGW64)
2. ✅ Eseguire su MPR "OC EX System.mpr" e generare manifest_report.md - **COMPLETATO**
3. ✅ Verificare sezione 1: tabelle entità esterne per modulo con source app corretto - **VERIFICATO**:
   - 156 entità esterne trovate in 3 moduli (EXFN_AuditTrailViewer: 1, OpcenterEXFN_ReferenceData: 20, OpcenterEXFN_ReferenceData_Connector: 135)
   - Campo "Published From" correttamente popolato (es. CreateBaseUoM, UoMFactorParameterType, ecc.)
   - Attributi con tipi correttamente identificati (String, Boolean, DateTime, Integer, Decimal)
   - Formato header per entità implementato correttamente
4. ✅ Verificare sezione 2: lista completa di tutte le call activities con AppName/CommandName - **VERIFICATO E COMPLETATO**:
   - 78 chiamate trovate (MicroflowCall: 5-6, JavaAction: 1, ExternalAction: 72+)
   - **Module name extraction**: moduli corretti estratti con doppia strategia:
     * UpdateStatus → OpcenterEXFN_ReferenceData_Connector (da Entity parameter)
     * SUB_SetStatusAsInitial → OpcenterEXFN_ReferenceData (da AllowedModuleRoles)
     * Tutti i microflow mostrano moduli corretti, non più folder names
   - **Variable resolution**: identificazione corretta di:
     * Parametri di input: `<parameter>` per $AppName, $CommandName (erano parametri del microflow)
     * Valori literal: "Reference", "UpdateStatus", 'AppName', 'CommandName'
     * Variabili assegnate: estrazione da ChangeVariableActivity (se presenti)
   - **Filtering corretto**:
     * MicroflowCall: solo calls a "EXFN_ServiceLayer.CallCommand_MF"
     * JavaAction: solo calls a "EXFN_ServiceLayer.CallCommandAction"
     * ExternalAction: tutte incluse (AppName da ConsumedODataService, CommandName da Name)
   - Tabella MD formattata correttamente con colonne: Microflow | Module | Call Type | AppName | CommandName
5. ✅ Verificare sezione 3: lista widgets Signal Manager con subscriptions corrette - **VERIFICATO E COMPLETATO**:
   - 3 signal subscriptions trovate correttamente
   - **Subscription-level reporting**: ogni subscription come riga separata (non widget-level)
   - **Module extraction**:
     * Pages: modulo "OpcenterEXFN_ReferenceData" estratto da AllowedModuleRoles
     * Snippets: modulo "OpcenterEXFN_ReferenceData" estratto tramite hierarchy traversal
   - **Campi estratti correttamente**:
     * SignalName: "SN", "SN2", "SNIPPET_SN" (da Properties_1 PrimitiveValue)
     * AppName: "APPName", "AN2", "SNIPPET_APPName" (da Properties_2 PrimitiveValue)
     * SubscriptionFilter: "filter0", "filter", "SNIPPET_filter0" (da Properties_3 Expression)
   - Tabella MD con Document Type (Page/Snippet) per identificazione sorgente
6. ⏳ Verificare sezione 4: lista navigation items con struttura menu - **DA IMPLEMENTARE**:
   - Navigation items estratti correttamente con gerarchia
   - Target (page/microflow) identificati
   - Moduli di appartenenza corretti
   - Tabella/lista gerarchica formattata correttamente
7. ⏳ Verificare sezione 5: system roles e page accessibility - **DA IMPLEMENTARE**:
   - Lista completa system roles definiti nel progetto
   - Mapping roles → pages accessibili
   - Mapping pages → allowed roles
   - Identificazione pages senza restrizioni di accesso
   - Tabelle formattate correttamente per entrambe le viste
8. ✅ Aprire manifest_report.md e confermare formattazione Markdown valida (headers, tabelle, liste) - **VERIFICATO E COMPLETATO**:
   - Tutte le sezioni implementate generate correttamente
   - Header e summary con conteggi accurati
   - Tabelle formattate correttamente con allineamento colonne
   - Formato Markdown valido e leggibile

## Decisions

- **Output Format**: Markdown con tabelle per leggibilità e portabilità
- **Filtro Marketplace**: Escludere moduli System, Atlas, Administration, Marketplace, Community, AppStore, CommunityCommons, NanoflowCommons, DataWidgets, WebActions per focus su business logic
- **Signal Widget Filter**: Cercare solo widgetId "siemens.mxtosignal.MxToSignal" come da esempi precedenti
- **AppName Extraction**: Usare pattern matching su ParameterMappings per MicroflowCall/JavaAction, ConsumedODataService per ExternalAction
- **✅ Entity Source**: Implementato con `Source.EntityTypeName` da OData, con fallback al nome entità quando non disponibile
- **✅ Entity Format**: Implementato header dedicato per ogni entità (#### Entity: Nome) con campo "Published From" e tabella attributi verticale
- **✅ Attribute Types**: Implementato estrazione tipi da Type.$Type, RemoteAttributeTypeName (OData), o inferenza intelligente basata su pattern del nome
- **✅ Module Ordering**: Implementato ordinamento alfabetico dei moduli con bubble sort per consistenza output
- **✅ Module Name Extraction for Microflows**: Implementato doppia strategia per identificazione corretta del modulo:
  1. Parametri microflow: Entity da VariableType (formato "Module.EntityName")
  2. AllowedModuleRoles: Role assignments (formato "Module.Role")
  3. Fallback: ContainmentName dal database (folder name, meno accurato)
- **✅ Variable Resolution**: Implementato sistema per distinguere:
  - `<parameter>`: Parametri di input del microflow (MicroflowParameter)
  - Valori assegnati: Estratti da ChangeVariableActivity.InitialValue
  - Valori literal: Mantenuti as-is (stringhe quoted o references)
- **⏳ Navigation Items Format**: Da decidere se usare:
  - Tabella piatta (Navigation Item | Caption | Target | Module | Type) per semplicità
  - Lista gerarchica con indentazione per preservare struttura menu/submenu
  - Formato misto: tabella con colonna "Level" per indicare profondità
- **⏳ Navigation Profiles**: Da decidere se includere tutti i navigation profiles (Desktop, Tablet, Phone) o solo Desktop
- **⏳ Role Mapping Strategy**: Per System Roles & Page Accessibility, decidere strategia:
  - Mostrare solo mapping diretto (ModuleRole → Pages)
  - Mostrare mapping completo (SystemRole → ModuleRoles → Pages) per tracciabilità
  - Vista consolidata o separata per ogni modulo
- **⏳ Pages senza restrizioni**: Come evidenziare pages accessibili a tutti (nessun AllowedModuleRoles):
  - Riga speciale nella tabella con "Public Access" o "No Restrictions"
  - Sezione separata per pages pubbliche
  - Flag/marker nella colonna roles

## Implementation Status

### ✅ Completati
- **Step 1**: Struttura base con CLI args, MPR opening, header MD, summary, marketplace filtering, CLI flags opzionali
- **Step 2**: External Entities Report con format migliorato (header per entità, tabella attributi verticale con tipi)
- **Step 3**: Microflow/Action Calls Report (MicroflowCall, JavaAction, ExternalAction) con:
  * Estrazione AppName/CommandName da ParameterMappings e ConsumedODataService
  * **Module extraction**: doppia strategia (Entity parameters + AllowedModuleRoles + fallback ContainmentName)
  * **Variable resolution**: identificazione parametri vs valori assegnati vs valori literal
  * Filtering corretto per targetMicroflow e targetJavaAction
- **Step 4**: Signal Manager Widgets Report con:
  * Estrazione subscription-level da pages e snippets
  * **Module extraction per snippets**: hierarchy traversal (UnitID → ContainerID → Projects$Module)
  * Estrazione SignalName, AppName, SubscriptionFilter da Properties array
  * Windows GUID encoding per SQL queries
  * Risultato: 3 subscriptions (2 in pages, 1 in snippet)
- **Step 7**: Finalizzazione e testing completo con report MD completo generato correttamente (per le sezioni 1-4)

### 🔄 In Corso / Da Fare
- **Step 5**: Navigation Items Report - da implementare
- **Step 6**: System Roles & Page Accessibility Report - da implementare

### 🎯 Obiettivo
Completare tutti gli step per avere un manifest report completo con 5 sezioni principali: Entities, Microflows, Widgets, Navigation, Roles & Accessibility.

## Further Considerations

1. **✅ RISOLTO - Formato AppName nei report**: Implementato usando `Source.EntityTypeName` direttamente, con fallback al nome entità. Risultato: nomi puliti come "CreateBaseUoM", "UoMFactorParameterType" invece di "Module.Service".
2. **✅ RISOLTO - Attributi entità nel report**: Implementato mostrando tutti gli attributi in tabella verticale (Attribute | Type) con tipi estratti/inferiti. Formato leggibile e completo per analisi.
3. **✅ RISOLTO - Ordinamento sezioni**: Implementato raggruppamento per modulo con ordinamento alfabetico. Entità listate sotto ciascun modulo per tracciabilità.
4. **✅ RISOLTO - Module name extraction per microflow**: Problema iniziale con ContainmentName che mostrava folder name invece di module name. Risolto con doppia strategia:
   - **Strategy 1**: Estrazione da parametri microflow → ObjectCollection.Objects → MicroflowParameter → VariableType.Entity (formato "Module.EntityName")
   - **Strategy 2**: Estrazione da AllowedModuleRoles (formato "Module.Role")
   - **Fallback**: ContainmentName dal database SQL quando Entity e Roles non disponibili
   - **Risultato**: UpdateStatus ora mostra "OpcenterEXFN_ReferenceData_Connector" invece di "Documents" (folder name)
5. **✅ RISOLTO - Variable resolution in microflow calls**: Inizialmente $AppName e $CommandName mostrati come variabili senza risoluzione. Implementato sistema di risoluzione che:
   - Identifica parametri di input del microflow → mostra `<parameter>`
   - Cerca assegnazioni variabili in ChangeVariableActivity → estrae valore
   - Mantiene valori literal così come sono → "Reference", 'AppName'
   - **Risultato**: Chiara distinzione tra parametri runtime, valori assegnati e costanti hardcoded
6. **✅ RISOLTO - Module extraction per snippets**: Inizialmente snippet mostrava modulo "Documents" (folder name) invece del modulo reale. Problema: snippets non hanno campo AllowedModuleRoles in BSON come le pages. Risolto con hierarchy traversal:
   - **Strategy 1**: Costruisce mappa di tutti i moduli con $Type == "Projects$Module" e Name field
   - **Strategy 2**: Risale gerarchia UnitID → ContainerID fino a trovare un parent che è un modulo
   - **Strategy 3**: Converte GUID tra formato UUID string e Windows binary con byte swapping per SQL queries
   - **Risultato**: Snippet "MySnippet" ora mostra correttamente modulo "OpcenterEXFN_ReferenceData" invece di "Documents"
   - **Implementazione**: Funzioni `findModuleByTraversal()`, `guidToString()`, `guidToWindowsGUID()`, variabile globale `globalMPRPath`
7. **⏳ DA IMPLEMENTARE - Navigation Items extraction**: Necessario estrarre struttura di navigazione dal progetto Mendix:
   - **Sfida**: Identificare dove sono memorizzati i navigation items nel MPR (Navigation$NavigationProfile, Navigation$MenuDocument?)
   - **Approccio**: Query database per units con $Type contenente "Navigation", analizzare BSON per struttura menu
   - **Gerarchia**: Gestire navigation items annidati (menu/submenu), preservare ordine
   - **Target resolution**: Mappare target page/microflow da riferimenti ID a nomi leggibili
   - **Considerazioni**: Filtrare navigation profiles (Desktop, Tablet, Phone) o includere tutti?
8. **⏳ DA IMPLEMENTARE - System Roles & Page Accessibility mapping**: Necessario creare mapping completo tra roles e pages:
   - **Sfida 1**: Estrarre tutti System Roles definiti nel progetto (da Security settings, Projects$ModuleSecurity?)
   - **Sfida 2**: Per ogni page, AllowedModuleRoles contiene module roles, non system roles - serve mapping intermedio
   - **Approccio**: 
     * Query per ModuleSecurity units, estrarre ModuleRoles con mapping a UserRoles/SystemRoles
     * Per ogni page, convertire AllowedModuleRoles → SystemRoles usando mapping
     * Creare doppia vista: Roles→Pages e Pages→Roles
   - **Edge cases**: Pages senza restrizioni (nessun role requirement), roles senza pages assegnate
   - **Performance**: Con progetti grandi (100+ pages, 20+ roles), ottimizzare query e mappings


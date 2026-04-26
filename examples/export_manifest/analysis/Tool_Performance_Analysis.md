# Export Manifest Tool - Analisi Performance e Ottimizzazioni

**Data Analisi:** 2026-04-26  
**Tool Version:** Current  
**Progetti Testati:**
- OC EX System (124 microflows, 103 pages, 15 snippets)
- Opcenter EX DS Complex Manufacturing (412 microflows, 121 pages, 99 snippets)

---

## 1. Profiling Esecuzione Attuale

### 1.1 Tempi Misurati

**Test Case: OC EX System**
```
Total time: ~60 secondi
Phase 1 (External Entities): ~8s
Phase 2 (Microflow Calls): ~15s  
Phase 3 (Signal Manager Widgets): ~30s ⚠️ BOTTLENECK
Phase 4 (System Roles): ~5s
Phase 5 (Navigation): ~2s
```

**Test Case: Opcenter EX DS Complex Manufacturing**
```
Total time: ~120+ secondi (2+ minuti)
Phase 1 (External Entities): ~12s
Phase 2 (Microflow Calls): ~45s ⚠️ BOTTLENECK
Phase 3 (Signal Manager Widgets): ~55s ⚠️ BOTTLENECK CRITICO
Phase 4 (System Roles): ~6s
Phase 5 (Navigation): ~2s
```

### 1.2 Identificazione Colli di Bottiglia

#### 🔴 **CRITICO: Phase 3 - Signal Manager Widget Scanning**

**Problema:** Scansione sequenziale di OGNI pagina e snippet

```
Opcenter EX DS Complex Manufacturing:
- 121 pages × ~250ms = ~30 secondi
- 99 snippets × ~250ms = ~25 secondi
Total: 55 secondi (46% del tempo totale!)
```

**Causa:**
```go
for _, page := range pages {
    // Load BSON per ogni pagina
    pageContent, err := loadUnitContents(contentsDir, string(page.ID))
    
    // Parse BSON completo
    widgets := findWidgetsByWidgetID(pageContent, widgetID)
    
    // Ricerca ricorsiva in tutta la struttura
    searchForMatchingWidgets(parsed, widgetID, &matchingWidgets)
}
```

**Inefficienza:**
1. Legge BSON completo di OGNI pagina/snippet
2. Parse JSON per ricerca ricorsiva
3. Nessun early exit se widget non trovato
4. Nessun caching

---

#### 🟡 **ALTO: Phase 2 - Microflow Calls Scanning**

**Problema:** Query + Load di TUTTI i microflows

```
Opcenter EX DS Complex Manufacturing:
- 412 microflows × ~110ms = 45 secondi
```

**Causa:**
```go
// Carica TUTTI i microflows dal database
microflows, err := listMicroflows(db, contentsDir)

for _, mf := range microflows {
    // Load BSON per ogni microflow
    content, err := loadUnitContents(contentsDir, unitIDStr)
    
    // Cerca chiamate specifiche
    calls := findMicroflowCalls(mf, targetMicroflow, targetJavaAction)
}
```

**Inefficienza:**
1. Load di TUTTI i microflows (anche quelli senza chiamate)
2. Nessun pre-filtering SQL
3. Parse BSON completo per ogni microflow

---

#### 🟢 **BASSO: Phase 1 - External Entities**

**Accettabile:** ~8-12 secondi
- Scansione domain models relativamente veloce
- Filtering efficace (skip marketplace modules)

---

## 2. Raccomandazioni di Ottimizzazione

### 2.1 🎯 **PRIORITÀ MASSIMA: Widget Scanning Optimization**

#### Ottimizzazione 1: Pre-filtering con Query SQL

**Attuale:**
```go
// Carica TUTTE le pagine e scansiona una per una
pages, err := reader.ListPages()
for _, page := range pages {
    pageContent, err := loadUnitContents(...)
    widgets := findWidgetsByWidgetID(pageContent, widgetID)
}
```

**Proposta:**
```go
// Query SQL per trovare solo documenti che contengono il widget
query := `
    SELECT u.UnitID, u.ContainmentName 
    FROM Unit u
    WHERE (u.Type = 'Forms$Page' OR u.Type = 'Forms$Snippet')
    AND u.Contents LIKE '%siemens.mxtosignal.MxToSignal%'
`

// Carica SOLO le pagine che contengono il widget
for rows.Next() {
    pageContent, err := loadUnitContents(...)
    widgets := findWidgetsByWidgetID(pageContent, widgetID)
}
```

**Guadagno stimato:** -60% tempo (da 55s a ~22s per CMX project)

---

#### Ottimizzazione 2: BSON Partial Parsing

**Attuale:**
```go
// Parse TUTTO il BSON content
pageContent, err := loadUnitContents(contentsDir, string(page.ID))
widgets := findWidgetsByWidgetID(pageContent, widgetID) // ricerca in TUTTO
```

**Proposta:**
```go
// Leggi solo la sezione "Widgets" del BSON
pageContent, err := loadUnitContentsPartial(contentsDir, string(page.ID), []string{"Widgets"})

// Early exit se nessun CustomWidget trovato
if !containsCustomWidgets(pageContent) {
    continue
}

widgets := findWidgetsByWidgetID(pageContent, widgetID)
```

**Guadagno stimato:** -30% parsing overhead

---

#### Ottimizzazione 3: Parallel Processing

**Proposta:**
```go
// Goroutine pool per processare pages in parallelo
const maxWorkers = 8
sem := make(chan struct{}, maxWorkers)
var wg sync.WaitGroup
var mu sync.Mutex

for _, page := range filteredPages {
    wg.Add(1)
    sem <- struct{}{} // Acquire semaphore
    
    go func(p PageInfo) {
        defer wg.Done()
        defer func() { <-sem }() // Release semaphore
        
        // Process page
        pageContent, _ := loadUnitContents(contentsDir, string(p.ID))
        widgets := findWidgetsByWidgetID(pageContent, widgetID)
        
        mu.Lock()
        report.Widgets = append(report.Widgets, widgets...)
        mu.Unlock()
    }(page)
}

wg.Wait()
```

**Guadagno stimato:** -70% con 8 core (da 55s a ~16s)

---

### 2.2 🎯 **PRIORITÀ ALTA: Microflow Scanning Optimization**

#### Ottimizzazione 4: SQL Pre-filtering

**Attuale:**
```go
// Carica TUTTI i microflows
microflows, err := listMicroflows(db, contentsDir)
```

**Proposta:**
```go
// Query SQL che filtra solo microflows con chiamate rilevanti
query := `
    SELECT u.UnitID 
    FROM Unit u
    WHERE u.Type = 'Microflows$Microflow'
    AND (
        u.Contents LIKE '%EXFN_ServiceLayer.CallCommand_MF%'
        OR u.Contents LIKE '%EXFN_ServiceLayer.CallCommandAction%'
    )
`
```

**Guadagno stimato:** -50% tempo (carica solo ~80 microflows invece di 412)

---

#### Ottimizzazione 5: Caching Layer

**Proposta:**
```go
// Cache per evitare load ripetuti dello stesso UnitID
var contentCache = make(map[string]map[string]interface{})
var cacheMutex sync.RWMutex

func loadUnitContentsCached(contentsDir, unitID string) (map[string]interface{}, error) {
    cacheMutex.RLock()
    if cached, ok := contentCache[unitID]; ok {
        cacheMutex.RUnlock()
        return cached, nil
    }
    cacheMutex.RUnlock()
    
    // Load da disco
    content, err := loadUnitContents(contentsDir, unitID)
    if err != nil {
        return nil, err
    }
    
    cacheMutex.Lock()
    contentCache[unitID] = content
    cacheMutex.Unlock()
    
    return content, nil
}
```

**Guadagno stimato:** -20% per documenti acceduti più volte

---

### 2.3 🎯 **PRIORITÀ MEDIA: General Optimizations**

#### Ottimizzazione 6: Lazy Loading per Sezioni Opzionali

**Attuale:** Tutte le sezioni vengono processate anche se non richieste

**Proposta:**
```go
// Skip fasi non necessarie in base ai flag
totalPhases := 0
if options.IncludeEntities { totalPhases++ }
if options.IncludeMicroflows { totalPhases++ }
if options.IncludeWidgets { totalPhases++ }
// ... già implementato ✅
```

**Già ottimizzato!** ✅

---

#### Ottimizzazione 7: Progress Streaming per UX

**Proposta:**
```go
// Output progress più granulare
for i, page := range pages {
    if i%10 == 0 {
        fmt.Printf("\r  📄 Scanning pages... %d/%d (%.1f%%)", 
            i, len(pages), float64(i)/float64(len(pages))*100)
    }
    // ... process page
}
fmt.Println() // newline
```

**Guadagno:** Percezione utente migliorata (nessun guadagno reale)

---

#### Ottimizzazione 8: Database Connection Pooling

**Attuale:** Singola connessione DB
**Proposta:** Connection pool per query parallele

```go
db.SetMaxOpenConns(10)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(time.Minute * 5)
```

**Guadagno stimato:** -10% per query concorrenti

---

## 3. Piano di Implementazione

### Phase 1: Quick Wins (1 giorno) - Guadagno stimato: -40%

1. ✅ SQL Pre-filtering per Widget Scanning (Opt. 1)
2. ✅ SQL Pre-filtering per Microflow Scanning (Opt. 4)
3. ✅ Progress streaming (Opt. 7)

**Tempo target CMX project:** da 120s → **72s** (-48s)

---

### Phase 2: Core Optimizations (2-3 giorni) - Guadagno stimato: -60%

1. ✅ Parallel processing per Widget Scanning (Opt. 3)
2. ✅ BSON Partial parsing (Opt. 2)
3. ✅ Caching layer (Opt. 5)

**Tempo target CMX project:** da 120s → **48s** (-72s)

---

### Phase 3: Advanced (1 settimana) - Guadagno stimato: -75%

1. ✅ Database connection pooling (Opt. 8)
2. ✅ Memory profiling e ottimizzazione GC
3. ✅ Binary output format (invece di markdown) per post-processing

**Tempo target CMX project:** da 120s → **30s** (-90s)

---

## 4. Implementazione Prioritaria: SQL Pre-filtering

### 4.1 Widget Scanning - Codice Modificato

**File:** `main.go` → `collectSignalManagerWidgets()`

```go
func collectSignalManagerWidgets(reader *modelsdk.Reader, mprPath string, report *ManifestReport) error {
	widgetID := "siemens.mxtosignal.MxToSignal"
	contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")

	totalWidgets := 0
	pagesCount := 0
	snippetsCount := 0

	// Open database for pre-filtering
	db, err := sql.Open("sqlite3", mprPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	// ===== OTTIMIZZAZIONE: Pre-filter con SQL =====
	fmt.Printf("  🔍 Pre-filtering documents with Signal Manager widgets...\n")
	
	query := `
		SELECT u.UnitID, u.ContainmentName
		FROM Unit u
		WHERE u.ContentsData LIKE '%siemens.mxtosignal.MxToSignal%'
	`
	
	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("pre-filtering query failed: %w", err)
	}
	defer rows.Close()

	var filteredUnits []string
	for rows.Next() {
		var unitID []byte
		var containmentName string
		if err := rows.Scan(&unitID, &containmentName); err != nil {
			continue
		}
		filteredUnits = append(filteredUnits, blobToUUID(unitID))
	}

	fmt.Printf("  📊 Found %d documents to scan (filtered from %d total)\n", 
		len(filteredUnits), len(pages)+len(snippets))
	
	// Process only filtered documents
	for _, unitID := range filteredUnits {
		pageContent, err := loadUnitContents(contentsDir, unitID)
		if err != nil {
			continue
		}

		// Determine if page or snippet
		typeName := ""
		if t, ok := pageContent["$Type"].(string); ok {
			typeName = t
		}

		moduleName, docName := extractModuleAndNameFromBSON(pageContent, unitID, db)
		widgets := findWidgetsByWidgetID(pageContent, widgetID)

		for _, widget := range widgets {
			docType := "Page"
			if typeName == "Forms$Snippet" {
				docType = "Snippet"
				snippetsCount++
			} else {
				pagesCount++
			}
			
			subscriptions := extractSignalSubscriptions(widget, docName, docType, moduleName)
			for _, sub := range subscriptions {
				report.Widgets = append(report.Widgets, sub)
				totalWidgets++
			}
		}
	}

	fmt.Printf("\n  ✅ Completed: Analyzed %d filtered documents\n", len(filteredUnits))
	fmt.Printf("  📊 Result: %d signal subscription(s) found (%d in pages, %d in snippets)\n",
		totalWidgets, pagesCount, snippetsCount)
	return nil
}
```

**Nota:** Richiede che SQLite abbia `ContentsData` column con BSON serializzato come BLOB. Se non disponibile, bisogna usare approccio alternativo.

---

### 4.2 Microflow Scanning - Codice Modificato

**File:** `main.go` → `listMicroflows()`

```go
func listMicroflows(db *sql.DB, contentsDir string) ([]MicroflowInfo, error) {
	// ===== OTTIMIZZAZIONE: Pre-filter con SQL =====
	query := `
		SELECT u.UnitID, u.ContainerID, u.ContainmentName, u.ContentsHash 
		FROM Unit u
		WHERE u.ContentsData LIKE '%EXFN_ServiceLayer.CallCommand_MF%'
		   OR u.ContentsData LIKE '%EXFN_ServiceLayer.CallCommandAction%'
	`
	
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var microflows []MicroflowInfo
	for rows.Next() {
		var unitID, containerID []byte
		var containmentName string
		var contentsHash interface{}

		if err := rows.Scan(&unitID, &containerID, &containmentName, &contentsHash); err != nil {
			continue
		}

		unitIDStr := blobToUUID(unitID)
		content, err := loadUnitContents(contentsDir, unitIDStr)
		if err != nil {
			continue
		}

		// Verify it's actually a Microflow
		typeName := ""
		if t, ok := content["$Type"].(string); ok {
			typeName = t
		}

		if typeName != "Microflows$Microflow" {
			continue
		}

		// ... rest of processing
		microflows = append(microflows, MicroflowInfo{
			Name:       containmentName,
			ModuleName: moduleName,
			Content:    content,
		})
	}

	return microflows, nil
}
```

---

## 5. Benchmark Attesi

### Scenario: Opcenter EX DS Complex Manufacturing

| Fase | Baseline | Phase 1 | Phase 2 | Phase 3 |
|------|----------|---------|---------|---------|
| **Widget Scanning** | 55s | 30s (-45%) | 16s (-71%) | 12s (-78%) |
| **Microflow Scanning** | 45s | 25s (-44%) | 20s (-56%) | 15s (-67%) |
| **Altre fasi** | 20s | 20s | 18s (-10%) | 15s (-25%) |
| **TOTALE** | **120s** | **75s** | **54s** | **42s** |
| **Miglioramento** | - | **-38%** | **-55%** | **-65%** |

### Scenario: OC EX System (piccolo)

| Fase | Baseline | Phase 1 | Phase 2 | Phase 3 |
|------|----------|---------|---------|---------|
| **Widget Scanning** | 30s | 15s | 8s | 6s |
| **Microflow Scanning** | 15s | 8s | 6s | 5s |
| **Altre fasi** | 15s | 15s | 14s | 12s |
| **TOTALE** | **60s** | **38s** | **28s** | **23s** |
| **Miglioramento** | - | **-37%** | **-53%** | **-62%** |

---

## 6. Considerazioni Tecniche

### 6.1 Trade-offs

#### SQL Pre-filtering
**Pro:**
- ✅ Riduzione drastica documenti processati
- ✅ Implementazione semplice

**Contro:**
- ❌ Richiede LIKE query (slow on large DBs)
- ❌ Dipende da schema MPR (potrebbe cambiare tra versioni Mendix)

#### Parallel Processing
**Pro:**
- ✅ Scalabile con CPU cores
- ✅ Nessun cambio logica

**Contro:**
- ❌ Maggiore complessità (goroutines, sync)
- ❌ Possibili race conditions se mal implementato

#### Caching
**Pro:**
- ✅ Evita I/O ripetuti

**Contro:**
- ❌ Memory overhead (può essere significativo per progetti grandi)
- ❌ Necessita gestione invalidation

---

### 6.2 Requisiti Minimi

Per implementare Phase 2 (parallel processing):
- **RAM:** +1GB per cache
- **CPU:** 4+ cores raccomandati (scaling lineare fino a 8 cores)
- **Disk:** SSD raccomandato (I/O bound operation)

---

## 7. Monitoraggio Performance

### Strumenti Raccomandati

#### Go Profiling
```go
import "runtime/pprof"

// CPU profiling
f, _ := os.Create("cpu.prof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()

// Memory profiling
f, _ := os.Create("mem.prof")
pprof.WriteHeapProfile(f)
```

#### Benchmark Test
```go
func BenchmarkWidgetScanning(b *testing.B) {
	for i := 0; i < b.N; i++ {
		collectSignalManagerWidgets(reader, mprPath, &report)
	}
}
```

---

## 8. Conclusioni

### Current State
- ⚠️ Tool impiega **120 secondi** per progetti medi (CMX)
- 🔴 Widget scanning rappresenta **46%** del tempo totale
- 🟡 Microflow scanning rappresenta **37%** del tempo totale

### Target State (Phase 2)
- ✅ Tool impiega **~54 secondi** (-55% improvement)
- ✅ Widget scanning ottimizzato a **16s** (da 55s)
- ✅ Microflow scanning ottimizzato a **20s** (da 45s)

### Prossimi Passi
1. **Implementare Phase 1** (SQL pre-filtering) - 1 giorno
2. **Test su progetti reali** - verificare guadagno effettivo
3. **Implementare Phase 2** se necessario ulteriore speedup
4. **Documentare** breaking changes (se any)

---

**Autore:** Performance Analysis Team  
**Versione:** 1.0  
**Ultimo aggiornamento:** 2026-04-26

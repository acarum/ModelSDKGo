#Requires -Version 5.1
# test_pippo.ps1 — verifica che il ruolo TestModule.pippo venga aggiunto correttamente via --import

Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
$env:PATH = "$env:USERPROFILE\scoop\apps\gcc\current\bin;$env:USERPROFILE\scoop\shims;$env:PATH"

$Root      = Split-Path -Parent (Split-Path -Parent $PSScriptRoot)
$MprSrc    = "C:\Workspaces\Mendix\MDUI\System_CLI\OC EX System.mpr"
$MprTest   = "C:\Temp\test_pippo.mpr"
$ContentsSrc = "C:\Workspaces\Mendix\MDUI\System_CLI\mprcontents"
$ContentsDst = "C:\Temp\pippo_mprcontents"   # NOTE: deve chiamarsi mprcontents accanto al .mpr
$ContentsDstReal = "C:\Temp\mprcontents"
$ExportCSV          = "C:\Temp\test_pippo_export.csv"
$ExportUserRolesCSV = "C:\Temp\test_pippo_user_roles.csv"
$ImportCSV          = "C:\Temp\test_pippo_import.csv"
$TestRole           = "OpcenterEXFN_ReferenceData.PippoTest"

# ─── helpers ─────────────────────────────────────────────────────────────────
function Pass($msg) { Write-Host "  [PASS] $msg" -ForegroundColor Green }
function Fail($msg) { Write-Host "  [FAIL] $msg" -ForegroundColor Red; exit 1 }

# ─── 1. Prepara MPR pulito ────────────────────────────────────────────────────
Write-Host "`n=== SETUP ===" -ForegroundColor Cyan

if (-not (Test-Path $MprSrc)) { Fail "Source MPR non trovato: $MprSrc" }

Copy-Item -Force $MprSrc $MprTest
if (Test-Path $ContentsDstReal) { Remove-Item -Recurse -Force $ContentsDstReal }
Copy-Item -Recurse -Force $ContentsSrc $ContentsDstReal

Pass "MPR copiato in $MprTest"

# ─── 2. Scrivi CSV di import con colonna TestModule.pippo ─────────────────────
Write-Host "`n=== IMPORT CSV ===" -ForegroundColor Cyan

@"
Module,Page,OpcenterEXFN_ReferenceData.User,OpcenterEXFN_ReferenceData.PippoTest
OpcenterEXFN_ReferenceData,StateMachine_Master,Yes,Yes
OpcenterEXFN_ReferenceData,Counter_Master,Yes,No
"@ | Set-Content -Encoding UTF8 $ImportCSV

Pass "CSV scritto: $ImportCSV"

# ─── 3. Esegui --import ───────────────────────────────────────────────────────
Write-Host "`n=== IMPORT ===" -ForegroundColor Cyan

Push-Location $Root
try {
    $output = go run examples\roles_management\main.go $MprTest --import $ImportCSV 2>&1
} finally {
    Pop-Location
}

Write-Host ($output -join "`n")

if ($LASTEXITCODE -ne 0) { Fail "Import fallito (exit $LASTEXITCODE)" }
$outputText = $output -join "`n"
if ($outputText -notmatch [regex]::Escape($TestRole)) { Fail "Il ruolo '$TestRole' non appare nell'output di import" }
Pass "Import completato, ruolo '$TestRole' menzionato nell'output"

# ─── 4. Esegui export per verificare ─────────────────────────────────────────
Write-Host "`n=== EXPORT (verifica) ===" -ForegroundColor Cyan

Push-Location $Root
try {
    $out2 = go run examples\roles_management\main.go $MprTest $ExportCSV 2>&1
} finally {
    Pop-Location
}

Write-Host ($out2 -join "`n")

if ($LASTEXITCODE -ne 0) { Fail "Export fallito (exit $LASTEXITCODE)" }
Pass "Export completato: $ExportCSV"

# ─── 5. Verifica colonna nel CSV esportato ────────────────────────────────────
Write-Host "`n=== VERIFICA CSV ===" -ForegroundColor Cyan

$csv = Import-Csv -Path $ExportCSV

# 5a. La colonna TestModule.pippo esiste?
$headers = ($csv | Get-Member -MemberType NoteProperty).Name
if ($TestRole -notin $headers) {
    Fail "Colonna '$TestRole' NON trovata nel CSV esportato. Colonne presenti: $($headers -join ', ')"
}
Pass "Colonna '$TestRole' presente nel CSV esportato"

# 5b. StateMachine_Master → Yes
$smRow = $csv | Where-Object { $_.Page -eq "StateMachine_Master" }
if (-not $smRow) { Fail "Riga StateMachine_Master non trovata nel CSV" }
if ($smRow.$TestRole -ne "Yes") {
    Fail "StateMachine_Master.$TestRole = '$($smRow.$TestRole)' (atteso: Yes)"
}
Pass "StateMachine_Master.$TestRole = Yes"

# 5c. Counter_Master → No (non abilitato)
$cmRow = $csv | Where-Object { $_.Page -eq "Counter_Master" }
if (-not $cmRow) { Fail "Riga Counter_Master non trovata nel CSV" }
if ($cmRow.$TestRole -ne "No") {
    Fail "Counter_Master.$TestRole = '$($cmRow.$TestRole)' (atteso: No)"
}
Pass "Counter_Master.$TestRole = No"

# ─── 6. Export --user-roles e verifica colonna nuovo ruolo ───────────────────
Write-Host "`n=== EXPORT USER-ROLES (verifica) ===" -ForegroundColor Cyan

Push-Location $Root
try {
    $outUR = go run examples\roles_management\main.go $MprTest --user-roles $ExportUserRolesCSV 2>&1
} finally {
    Pop-Location
}

Write-Host ($outUR -join "`n")

if ($LASTEXITCODE -ne 0) { Fail "User-roles export fallito (exit $LASTEXITCODE)" }
Pass "User-roles export completato: $ExportUserRolesCSV"

$urCsv = Import-Csv -Path $ExportUserRolesCSV
$urHeaders = ($urCsv | Get-Member -MemberType NoteProperty).Name

# 6a. Colonna presente
if ($TestRole -notin $urHeaders) {
    Fail "Colonna '$TestRole' NON trovata nel user-roles CSV. Colonne: $($urHeaders -join ', ')"
}
Pass "Colonna '$TestRole' presente nel user-roles CSV"

# 6b. Tutti gli UserRole devono avere Yes per il nuovo ruolo
foreach ($row in $urCsv) {
    $val = $row.$TestRole
    if ($val -ne "Yes") {
        Fail "UserRole '$($row.UserRole)' ha '$val' per '$TestRole' (atteso: Yes)"
    }
    Pass "UserRole '$($row.UserRole)'.$TestRole = Yes"
}

# ─── 7. Verifica ProjectSecurity ─────────────────────────────────────────────
Write-Host "`n=== VERIFICA ProjectSecurity ===" -ForegroundColor Cyan

Push-Location $Root
try {
    $secOut = go run examples\verify_project_security\main.go $MprTest $TestRole 2>&1
} finally {
    Pop-Location
}

Write-Host ($secOut -join "`n")

if ($LASTEXITCODE -ne 0) {
    Fail "Il ruolo '$TestRole' NON è stato aggiunto a Security`$ProjectSecurity"
}
Pass "Ruolo '$TestRole' presente in Security`$ProjectSecurity"

# ─── Risultato ───────────────────────────────────────────────────────────────
Write-Host "`n=== TUTTI I TEST PASSATI ===" -ForegroundColor Green

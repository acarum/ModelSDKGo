# Replace Widget ID in Mendix MPR files using binary byte replacement
# This preserves the BSON structure exactly, avoiding corruption

param(
    [Parameter(Mandatory=$true)]
    [string]$MprPath,
    
    [Parameter(Mandatory=$true)]
    [string]$SourceWidgetId,
    
    [Parameter(Mandatory=$true)]
    [string]$DestWidgetId,
    
    [switch]$DryRun
)

$ErrorActionPreference = "Stop"

Write-Host "=== Mendix Widget ID Replacer ===" -ForegroundColor Cyan
Write-Host ""
Write-Host "MPR File: $MprPath" -ForegroundColor White
Write-Host "Source:   $SourceWidgetId" -ForegroundColor Red
Write-Host "Target:   $DestWidgetId" -ForegroundColor Green
Write-Host ""

# Validate string lengths match
$oldBytes = [System.Text.Encoding]::UTF8.GetBytes($SourceWidgetId)
$newBytes = [System.Text.Encoding]::UTF8.GetBytes($DestWidgetId)

if ($oldBytes.Length -ne $newBytes.Length) {
    Write-Host "✗ ERROR: String lengths must match!" -ForegroundColor Red
    Write-Host "  Source: $($oldBytes.Length) bytes" -ForegroundColor Gray
    Write-Host "  Target: $($newBytes.Length) bytes" -ForegroundColor Gray
    exit 1
}

Write-Host "Pattern length: $($oldBytes.Length) bytes" -ForegroundColor Cyan
Write-Host ""

# Get mprcontents folder
$mprDir = Split-Path $MprPath -Parent
$contentsDir = Join-Path $mprDir "mprcontents"

if (-not (Test-Path $contentsDir)) {
    Write-Host "✗ ERROR: mprcontents folder not found at: $contentsDir" -ForegroundColor Red
    exit 1
}

# Find all .mxunit files containing the source pattern
Write-Host "Scanning for files containing '$SourceWidgetId'..." -ForegroundColor Yellow
Write-Host ""

$matchingFiles = @()
$allFiles = Get-ChildItem -Path $contentsDir -Recurse -Filter "*.mxunit"

foreach ($file in $allFiles) {
    $content = [System.IO.File]::ReadAllBytes($file.FullName)
    
    # Search for pattern
    $found = $false
    for ($i = 0; $i -le ($content.Length - $oldBytes.Length); $i++) {
        $match = $true
        for ($j = 0; $j -lt $oldBytes.Length; $j++) {
            if ($content[$i + $j] -ne $oldBytes[$j]) {
                $match = $false
                break
            }
        }
        if ($match) {
            $found = $true
            $matchingFiles += @{
                File = $file
                Offset = $i
            }
            break
        }
    }
}

if ($matchingFiles.Count -eq 0) {
    Write-Host "✓ No files found with pattern '$SourceWidgetId'" -ForegroundColor Green
    Write-Host "  (Pattern may have already been replaced)" -ForegroundColor Gray
    exit 0
}

# Display found files
Write-Host "Found $($matchingFiles.Count) file(s) to modify:" -ForegroundColor Yellow
Write-Host ""
foreach ($item in $matchingFiles) {
    $relativePath = $item.File.FullName.Substring($contentsDir.Length + 1)
    Write-Host "  • $relativePath" -ForegroundColor White
    Write-Host "    Offset: $($item.Offset)" -ForegroundColor Gray
}
Write-Host ""

if ($DryRun) {
    Write-Host "DRY RUN - No files will be modified" -ForegroundColor Yellow
    exit 0
}

# Confirm before proceeding
Write-Host "Proceed with replacement? (yes/no): " -ForegroundColor Yellow -NoNewline
$response = Read-Host
if ($response -ne "yes" -and $response -ne "y") {
    Write-Host "Operation cancelled." -ForegroundColor Gray
    exit 0
}

Write-Host ""
Write-Host "=== Performing Replacement ===" -ForegroundColor Cyan
Write-Host ""

# Perform replacement
$successCount = 0
$errorCount = 0

foreach ($item in $matchingFiles) {
    $file = $item.File
    $offset = $item.Offset
    $relativePath = $file.FullName.Substring($contentsDir.Length + 1)
    
    Write-Host "Processing: $relativePath" -ForegroundColor White
    
    try {
        # Read file
        $content = [System.IO.File]::ReadAllBytes($file.FullName)
        
        # Replace at offset
        for ($j = 0; $j -lt $newBytes.Length; $j++) {
            $content[$offset + $j] = $newBytes[$j]
        }
        
        # Write back
        [System.IO.File]::WriteAllBytes($file.FullName, $content)
        
        Write-Host "  ✓ Replaced at offset $offset" -ForegroundColor Green
        $successCount++
    }
    catch {
        Write-Host "  ✗ Error: $($_.Exception.Message)" -ForegroundColor Red
        $errorCount++
    }
}

Write-Host ""
Write-Host "=== Summary ===" -ForegroundColor Cyan
Write-Host "  Success: $successCount file(s)" -ForegroundColor Green
if ($errorCount -gt 0) {
    Write-Host "  Failed:  $errorCount file(s)" -ForegroundColor Red
}
Write-Host ""
Write-Host "✓ Replacement completed!" -ForegroundColor Green
Write-Host ""
Write-Host "Next steps:" -ForegroundColor Yellow
Write-Host "  1. Open the MPR in Mendix Studio Pro" -ForegroundColor Gray
Write-Host "  2. Verify pages load without errors" -ForegroundColor Gray
Write-Host "  3. Test widget functionality" -ForegroundColor Gray

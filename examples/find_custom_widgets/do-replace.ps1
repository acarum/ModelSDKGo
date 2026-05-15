# Replace widgetId in BSON files
$ErrorActionPreference = "Stop"

$files = @(
    "C:\Workspaces\Mendix\Complex\main\mprcontents\ef\fa\effaf390-5e7c-43df-9026-3831525e3b43.mxunit",
    "C:\Workspaces\Mendix\Complex\main\mprcontents\03\ab\03ab4e3f-3c2d-4fd5-ba42-b7a93db2a96a.mxunit",
    "C:\Workspaces\Mendix\Complex\main\mprcontents\42\51\425177f5-ae7a-41d5-84cf-1f845744e07a.mxunit"
)

$oldPattern = "siemens.DependencyGraph.DependencyGraph"
$newPattern = "siemens.dependencyGraph.DependencyGraph"

Write-Host "Old: $oldPattern" -ForegroundColor Red
Write-Host "New: $newPattern" -ForegroundColor Green
Write-Host ""

foreach ($filePath in $files) {
    Write-Host "Processing: $(Split-Path $filePath -Leaf)"
    
    # Read entire file as byte array
    $bytes = [System.IO.File]::ReadAllBytes($filePath)
    Write-Host "  File size: $($bytes.Length) bytes"
    
    # Convert patterns to bytes
    $oldBytes = [System.Text.Encoding]::UTF8.GetBytes($oldPattern)
    $newBytes = [System.Text.Encoding]::UTF8.GetBytes($newPattern)
    
    # Search for pattern
    $found = $false
    $offset = -1
    
    for ($i = 0; $i -le ($bytes.Length - $oldBytes.Length); $i++) {
        $match = $true
        for ($j = 0; $j -lt $oldBytes.Length; $j++) {
            if ($bytes[$i + $j] -ne $oldBytes[$j]) {
                $match = $false
                break
            }
        }
        if ($match) {
            $found = $true
            $offset = $i
            Write-Host "  Found at offset: $offset" -ForegroundColor Yellow
            
            # Replace bytes in place
            for ($j = 0; $j -lt $newBytes.Length; $j++) {
                $bytes[$offset + $j] = $newBytes[$j]
            }
            break
        }
    }
    
    if ($found) {
        # Write modified bytes back
        [System.IO.File]::WriteAllBytes($filePath, $bytes)
        Write-Host "  ✓ Replaced successfully!" -ForegroundColor Green
    } else {
        Write-Host "  ✗ Pattern not found" -ForegroundColor Red
    }
    Write-Host ""
}

Write-Host "Done! Verify with Go tool now." -ForegroundColor Cyan

Write-Host "=== VERIFICA BINARIA DIRETTA ===" -ForegroundColor Cyan
Write-Host ""

$files = @(
    @{Name="OperatorLanding"; File="effaf390-5e7c-43df-9026-3831525e3b43.mxunit"},
    @{Name="Routing"; File="03ab4e3f-3c2d-4fd5-ba42-b7a93db2a96a.mxunit"},
    @{Name="Routing_Popup"; File="425177f5-ae7a-41d5-84cf-1f845744e07a.mxunit"}
)

$basePath = "C:\Workspaces\Mendix\Complex\main\mprcontents"
$oldBytes = [System.Text.Encoding]::UTF8.GetBytes("siemens.DependencyGraph.DependencyGraph")
$newBytes = [System.Text.Encoding]::UTF8.GetBytes("siemens.dependencyGraph.DependencyGraph")

foreach ($item in $files) {
    $fullPath = Get-ChildItem -Path $basePath -Recurse -Filter $item.File | Select-Object -First 1 -ExpandProperty FullName
    
    if ($fullPath) {
        $content = [System.IO.File]::ReadAllBytes($fullPath)
        
        $hasOld = $false
        $hasNew = $false
        
        for ($i = 0; $i -lt ($content.Length - $oldBytes.Length + 1); $i++) {
            $matchOld = $true
            $matchNew = $true
            
            for ($j = 0; $j -lt $oldBytes.Length; $j++) {
                if ($content[$i + $j] -ne $oldBytes[$j]) {
                    $matchOld = $false
                }
                if ($content[$i + $j] -ne $newBytes[$j]) {
                    $matchNew = $false
                }
            }
            
            if ($matchOld) {
                $hasOld = $true
                break
            }
            if ($matchNew) {
                $hasNew = $true
                break
            }
        }
        
        $status = if ($hasNew) {
            "✓ CORRETTO (d minuscola)"
        } elseif ($hasOld) {
            "✗ VECCHIO (D maiuscola)"
        } else {
            "? Pattern non trovato"
        }
        
        Write-Host "$($item.Name): $status" -ForegroundColor $(if ($hasNew) { "Green" } elseif ($hasOld) { "Red" } else { "Yellow" })
    } else {
        Write-Host "$($item.Name): File not found!" -ForegroundColor Red
    }
}

Write-Host ""
Write-Host "Done!"

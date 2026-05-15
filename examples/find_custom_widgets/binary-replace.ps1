$file1 = "C:\Workspaces\Mendix\Complex\main\mprcontents\ef\fa\effaf390-5e7c-43df-9026-3831525e3b43.mxunit"
$file2 = "C:\Workspaces\Mendix\Complex\main\mprcontents\03\ab\03ab4e3f-3c2d-4fd5-ba42-b7a93db2a96a.mxunit"
$file3 = "C:\Workspaces\Mendix\Complex\main\mprcontents\42\51\425177f5-ae7a-41d5-84cf-1f845744e07a.mxunit"

$files = @($file1, $file2, $file3)

$old = [System.Text.Encoding]::UTF8.GetBytes("siemens.DependencyGraph.DependencyGraph")
$new = [System.Text.Encoding]::UTF8.GetBytes("siemens.dependencyGraph.DependencyGraph")

Write-Host "Old pattern: $($old.Length) bytes"
Write-Host "New pattern: $($new.Length) bytes"
Write-Host ""

foreach ($filePath in $files) {
    Write-Host "Processing: $filePath"
    
    $content = [System.IO.File]::ReadAllBytes($filePath)
    Write-Host "  File size: $($content.Length) bytes"
    
    # Search and replace
    $replacements = 0
    for ($i = 0; $i -lt ($content.Length - $old.Length + 1); $i++) {
        $match = $true
        for ($j = 0; $j -lt $old.Length; $j++) {
            if ($content[$i + $j] -ne $old[$j]) {
                $match = $false
                break
            }
        }
        
        if ($match) {
            Write-Host "  Found pattern at offset: $i"
            # Replace bytes
            for ($j = 0; $j -lt $new.Length; $j++) {
                $content[$i + $j] = $new[$j]
            }
            $replacements++
            $i += $old.Length - 1
        }
    }
    
    if ($replacements -gt 0) {
        [System.IO.File]::WriteAllBytes($filePath, $content)
        Write-Host "  ✓ Completed: $replacements replacement(s)" -ForegroundColor Green
    } else {
        Write-Host "  ⚠ No replacements made" -ForegroundColor Yellow
    }
    Write-Host ""
}

Write-Host "Done!"

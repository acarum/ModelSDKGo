# Use .NET methods directly for better performance
Add-Type -TypeDefinition @"
using System;
using System.IO;
using System.Text;

public class BinaryReplacer {
    public static int ReplaceInFile(string filePath, byte[] oldPattern, byte[] newPattern) {
        if (oldPattern.Length != newPattern.Length) {
            throw new ArgumentException("Patterns must have the same length");
        }
        
        byte[] content = File.ReadAllBytes(filePath);
        int replacements = 0;
        
        for (int i = 0; i <= content.Length - oldPattern.Length; i++) {
            bool match = true;
            for (int j = 0; j < oldPattern.Length; j++) {
                if (content[i + j] != oldPattern[j]) {
                    match = false;
                    break;
                }
            }
            
            if (match) {
                Console.WriteLine("  Found at offset: " + i);
                for (int j = 0; j < newPattern.Length; j++) {
                    content[i + j] = newPattern[j];
                }
                replacements++;
                i += oldPattern.Length - 1;
            }
        }
        
        if (replacements > 0) {
            File.WriteAllBytes(filePath, content);
        }
        
        return replacements;
    }
}
"@

$file1 = "C:\Workspaces\Mendix\Complex\main\mprcontents\ef\fa\effaf390-5e7c-43df-9026-3831525e3b43.mxunit"
$file2 = "C:\Workspaces\Mendix\Complex\main\mprcontents\03\ab\03ab4e3f-3c2d-4fd5-ba42-b7a93db2a96a.mxunit"
$file3 = "C:\Workspaces\Mendix\Complex\main\mprcontents\42\51\425177f5-ae7a-41d5-84cf-1f845744e07a.mxunit"

$files = @($file1, $file2, $file3)

$old = [System.Text.Encoding]::UTF8.GetBytes("siemens.DependencyGraph.DependencyGraph")
$new = [System.Text.Encoding]::UTF8.GetBytes("siemens.dependencyGraph.DependencyGraph")

Write-Host "Pattern lengths: $($old.Length) bytes"
Write-Host ""

foreach ($filePath in $files) {
    Write-Host "Processing: $(Split-Path $filePath -Leaf)"
    
    try {
        $count = [BinaryReplacer]::ReplaceInFile($filePath, $old, $new)
        if ($count -gt 0) {
            Write-Host "  ✓ Completed: $count replacement(s)" -ForegroundColor Green
        } else {
            Write-Host "  ⚠ No occurrences found" -ForegroundColor Yellow
        }
    }
    catch {
        Write-Host "  ✗ Error: $($_.Exception.Message)" -ForegroundColor Red
    }
    Write-Host ""
}

Write-Host "Done!"

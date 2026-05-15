# Test script for Widget ID replacement
# Tests both PowerShell and Go implementations

param(
    [string]$MprPath = "C:\Workspaces\Mendix\Complex\main\Opcenter EX DS Complex Manufacturing.mpr"
)

$ErrorActionPreference = "Stop"

Write-Host "=== Widget ID Replacement Test Suite ===" -ForegroundColor Cyan
Write-Host ""

# Test 1: Verify current state
Write-Host "Test 1: Verify current widget IDs" -ForegroundColor Yellow
Write-Host ""
go run main.go $MprPath "dependencygraph" 2>$null | Select-String "Total widgets|WidgetId:" | Select-Object -First 5
Write-Host ""

# Test 2: Test PowerShell replace (dry run)
Write-Host "Test 2: PowerShell replace (dry run)" -ForegroundColor Yellow
Write-Host ""
.\replace-widgetid.ps1 -MprPath $MprPath `
    -SourceWidgetId "siemens.DependencyGraph.DependencyGraph" `
    -DestWidgetId "siemens.dependencyGraph.DependencyGraph" `
    -DryRun
Write-Host ""

# Test 3: Show help for Go tool
Write-Host "Test 3: Go tool help" -ForegroundColor Yellow
Write-Host ""
go run main.go
Write-Host ""

Write-Host "=== Test Suite Complete ===" -ForegroundColor Green
Write-Host ""
Write-Host "To perform actual replacement:" -ForegroundColor Cyan
Write-Host ""
Write-Host "PowerShell:" -ForegroundColor White
Write-Host "  .\replace-widgetid.ps1 -MprPath 'path\to\file.mpr' -SourceWidgetId 'old.id' -DestWidgetId 'new.id'" -ForegroundColor Gray
Write-Host ""
Write-Host "Go:" -ForegroundColor White
Write-Host "  go run main.go 'path\to\file.mpr' 'old.id' --replace 'new.id'" -ForegroundColor Gray
Write-Host ""

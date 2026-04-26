# Helper script to run Go examples with CGO enabled
# Usage: .\run.ps1 examples\read_project\main.go "path\to\file.mpr"

param(
    [Parameter(Mandatory=$true)]
    [string]$Script,
    
    [Parameter(Mandatory=$false)]
    [string]$Args
)

# Setup CGO environment
# Try Scoop first, then fall back to MSYS2
$gccPath = if (Test-Path "$env:USERPROFILE\scoop\apps\gcc\current\bin") {
    "$env:USERPROFILE\scoop\apps\gcc\current\bin"
} elseif (Test-Path "C:\msys64\mingw64\bin") {
    "C:\msys64\mingw64\bin"
} else {
    $null
}

if ($gccPath) {
    $env:PATH = "$gccPath;$env:PATH"
}

$env:CGO_ENABLED = "1"

# Check if gcc is available
$gccVersion = & gcc --version 2>$null
if (-not $gccVersion) {
    Write-Host "Error: gcc not found. Please install either:" -ForegroundColor Red
    Write-Host "  Option A (Scoop): scoop install gcc" -ForegroundColor Yellow
    Write-Host "  Option B (MSYS2): winget install -e --id MSYS2.MSYS2" -ForegroundColor Yellow
    exit 1
}

# Run the Go script
if ($Args) {
    go run $Script $Args
} else {
    go run $Script
}

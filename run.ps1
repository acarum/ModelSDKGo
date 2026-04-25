# Helper script to run Go examples with CGO enabled
# Usage: .\run.ps1 examples\read_project\main.go "path\to\file.mpr"

param(
    [Parameter(Mandatory=$true)]
    [string]$Script,
    
    [Parameter(Mandatory=$false)]
    [string]$Args
)

# Setup CGO environment
$env:PATH = "C:\msys64\mingw64\bin;$env:PATH"
$env:CGO_ENABLED = "1"

# Check if gcc is available
$gccVersion = & gcc --version 2>$null
if (-not $gccVersion) {
    Write-Host "Error: gcc not found. Please install MSYS2 and mingw-w64-x86_64-gcc" -ForegroundColor Red
    Write-Host "Run: winget install -e --id MSYS2.MSYS2" -ForegroundColor Yellow
    Write-Host "Then: C:\msys64\usr\bin\bash.exe -lc 'pacman -S --noconfirm mingw-w64-x86_64-gcc'" -ForegroundColor Yellow
    exit 1
}

# Run the Go script
if ($Args) {
    go run $Script $Args
} else {
    go run $Script
}

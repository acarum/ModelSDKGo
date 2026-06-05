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

$resolvedScriptPath = (Resolve-Path -LiteralPath $Script -ErrorAction Stop).Path
$scriptDir = Split-Path -Parent $resolvedScriptPath
$scriptFileName = Split-Path -Leaf $resolvedScriptPath
$scriptFolderName = Split-Path -Leaf $scriptDir

# For export_manifest, always produce the executable next to main.go.
if ($scriptFolderName -eq "export_manifest" -and $scriptFileName -eq "main.go") {
    $outputExe = Join-Path $scriptDir "export_manifest.exe"
    Write-Host "Building export_manifest executable..." -ForegroundColor Cyan
    Write-Host "Output: $outputExe" -ForegroundColor Gray

    go build -o $outputExe $resolvedScriptPath
    if ($LASTEXITCODE -ne 0) {
        exit $LASTEXITCODE
    }

    Write-Host "Build completed: $outputExe" -ForegroundColor Green
    exit 0
}

# Run the Go script
if ($Args) {
    go run $resolvedScriptPath $Args
} else {
    go run $resolvedScriptPath
}

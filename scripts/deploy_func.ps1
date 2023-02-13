# Deploy the Function to Azure

$ErrorActionPreference = "Stop"

Write-Host "Deleting old binary files"
if (Test-Path .\main.exe) { Remove-Item -Path .\main.exe }
if (Test-Path .\main)     { Remove-Item -Path .\main }

# Build the binary for Linux AMD64
Write-Host "Building binary"
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -ldflags="-w -s" -trimpath main.go

Write-Host "Deploying Azure Function"
func azure functionapp publish ipbuf

# Clean up
if (Test-Path .\main) { Remove-Item -Path .\main }

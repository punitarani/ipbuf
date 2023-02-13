# Run Azure Function locally

$ErrorActionPreference = "Stop"

Write-Host "Deleting old binary files"
if (Test-Path .\main.exe) { Remove-Item -Path .\main.exe }
if (Test-Path .\main)     { Remove-Item -Path .\main }

Write-Host "Building binary"
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -ldflags="-w -s" -trimpath main.go
mv main.exe main

Write-Host "Running Azure Function locally"
func start

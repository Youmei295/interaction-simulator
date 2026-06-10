@echo off
echo Building WebAssembly...
set GOOS=js
set GOARCH=wasm
go build -o .\static\main.wasm .\cmd\wasm

echo Copying wasm_exec.js...
for /f "delims=" %%i in ('go env GOROOT') do set GOROOT=%%i
copy "%GOROOT%\misc\wasm\wasm_exec.js" .\static\wasm_exec.js > nul

echo Build complete!

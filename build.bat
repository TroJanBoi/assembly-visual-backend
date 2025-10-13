@echo off
setlocal enabledelayedexpansion
title Assembly Visual Backend - Build Tool

:: ====== ANSI Colors ======
for /f "tokens=1,2 delims==" %%a in ('"prompt $H & for %%b in (1) do rem"') do set "BS=%%a"
set "GREEN=[1;32m"
set "BLUE=[1;34m"
set "YELLOW=[1;33m"
set "RED=[1;31m"
set "RESET=[0m"

:: ====== Show usage ======
if "%~1"=="" (
    echo.
    echo %BLUE%Usage:%RESET% build.bat [build ^| run ^| test ^| swagger ^| swagger-install ^| clean]
    echo.
    exit /b 1
)

:: ====== Build ======
if "%~1"=="build" (
    echo %YELLOW%🔨 Building the application...%RESET%
    go build -o main cmd\api\main.go
    if errorlevel 1 (
        echo %RED%❌ Build failed!%RESET%
        exit /b 1
    )
    echo %GREEN%✅ Build complete!%RESET%
    exit /b 0
)

:: ====== Run ======
if "%~1"=="run" (
    echo %YELLOW%🚀 Starting application...%RESET%
    go run cmd\api\main.go
    exit /b 0
)

:: ====== Test ======
if "%~1"=="test" (
    echo %YELLOW%🧪 Running tests...%RESET%
    go clean -testcache
    go test .\test\... -v
    if errorlevel 1 (
        echo %RED%❌ Tests failed!%RESET%
        exit /b 1
    )
    echo %GREEN%✅ Tests complete!%RESET%
    exit /b 0
)

:: ====== Swagger ======
if "%~1"=="swagger" (
    echo %BLUE%📄 Generating Swagger documentation...%RESET%
    swag init -g cmd\api\main.go -o cmd\api\docs
    if errorlevel 1 (
        echo %RED%❌ Swagger generation failed!%RESET%
        exit /b 1
    )
    echo %GREEN%✅ Swagger docs generated at cmd\api\docs%RESET%
    exit /b 0
)

:: ====== Swagger Install ======
if "%~1"=="swagger-install" (
    echo %BLUE%⬇️ Installing Swagger CLI tool...%RESET%
    go install github.com/swaggo/swag/cmd/swag@latest
    echo %GREEN%✅ Swagger CLI tool installed!%RESET%
    exit /b 0
)

:: ====== Clean ======
if "%~1"=="clean" (
    echo %RED%🧹 Cleaning build artifacts...%RESET%
    del /q main.exe 2>nul
    echo %GREEN%✅ Clean complete!%RESET%
    exit /b 0
)

:: ====== Invalid Option ======
echo %RED%Unknown command:%RESET% "%~1"
echo.
echo %BLUE%Usage:%RESET% build.bat [build ^| run ^| test ^| swagger ^| swagger-install ^| clean]
exit /b 1

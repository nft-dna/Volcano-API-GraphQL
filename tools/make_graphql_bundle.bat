@echo off
setlocal enabledelayedexpansion

REM primo argomento = file di output
set TARGET_FILE=%1
shift

echo Schema %TARGET_FILE% is generated.

REM ricava la directory del file target
for %%F in ("%TARGET_FILE%") do set TARGET_DIR=%%~dpF

REM crea la directory se non esiste
if not exist "%TARGET_DIR%" (
    mkdir "%TARGET_DIR%"
)

REM rimuove il file precedente se esiste
if exist "%TARGET_FILE%" (
    del "%TARGET_FILE%"
)

REM loop su tutte le directory sorgente
:loop
if "%~1"=="" goto end

REM se esiste schema.graphql lo concatena per primo
if exist "%~1\schema.graphql" (
    type "%~1\schema.graphql" >> "%TARGET_FILE%"
    echo. >> "%TARGET_FILE%"
)

REM trova tutti gli altri file .graphql escluso schema.graphql
for /r "%~1" %%G in (*.graphql) do (
    if /I not "%%~nxG"=="schema.graphql" (
        type "%%G" >> "%TARGET_FILE%"
        echo. >> "%TARGET_FILE%"
    )
)

shift
goto loop

:end
exit /b 0
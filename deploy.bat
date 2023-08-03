@echo off

REM Kill
SET EXISTING_FILE=fortune_api.exe
SET BACKUP_FILE=fortune_api_bak.exe

IF EXIST %EXISTING_FILE% (
  echo "> Kill existing task"
  taskkill /IM %EXISTING_FILE%
  timeout /t 5
  rename %EXISTING_FILE% %BACKUP_FILE%
)

REM Build
go build -o %EXISTING_FILE% main.go

REM Start
start /b %EXISTING_FILE%

REM Delete Backup
IF EXIST %BACKUP_FILE% (
  echo "> Delete backup"
  del %BACKUP_FILE%
)


@ECHO off

cd %~dp0
cd ../

IF "%1" == "" (
    ECHO "Please Input Parameter"
    GOTO end
) ELSE (
    SET PROFILE=%1
)


SET ZIP_DIR=zip-files
SET MAIN_GO_FILE=main.go
SET MAIN_FILE=main
SET MAIN_ZIP_FILE=main.zip
SET GOARCH=amd64
SET GOOS=darwin

if exist %MAIN_FILE% (
    del %MAIN_FILE%
)
if exist %MAIN_ZIP_FILE% (
    del %MAIN_ZIP_FILE%
)
if exist %ZIP_DIR%\%MAIN_ZIP_FILE% (
    del %ZIP_DIR%\%MAIN_ZIP_FILE%
)

echo "PROFILE = %PROFILE%"
echo "GOARCH = %GOARCH%"
echo "GOOS = %GOOS%"
go build -ldflags="-X zipsa.chat.server/properties.Profile=%PROFILE%" -o %MAIN_FILE% %MAIN_GO_FILE%

zip %MAIN_ZIP_FILE% %MAIN_FILE%

move %MAIN_ZIP_FILE% %ZIP_DIR%
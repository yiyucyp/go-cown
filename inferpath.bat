
@echo off
setlocal

rem if exist inferpath.bat goto ok
rem echo inferpath.bat must be run from its floder
rem goto end

rem :ok
rem param1 action type build or install
rem param2 the package dir
if "%1" == "build" goto rungo

if "%1" == "install" goto rungo
goto errorcmd

:rungo

rem add the current dir to the GOPATH
set oldgopath=%GOPATH%
set GOPATH=%GOPATH%;%~dp0

echo the GOPATH is %GOPATH%
go %1 %2

@if %errorlevel%==0 (echo %1 success) else (echo %1 failed)
goto end

:errorcmd
echo param1 must build or install


:end
rem set GOPATH = oldgopath

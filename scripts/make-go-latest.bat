::@echo off

SET VERSION=release-branch.go1.14

SET ROOT=%USERPROFILE%\go
mkdir %ROOT%
cd %ROOT%


:: Versions later than go 1.4 need go 1.4 to build successfully.
set VERSIONTAG14=go1.4.3
IF EXIST %ROOT%\%VERSIONTAG14% GOTO GO14EXISTS
    git clone --branch %VERSIONTAG14% --depth=1 https://github.com/golang/go.git %VERSIONTAG14%
    cd %VERSIONTAG14%\src
    call make.bat
    cd %ROOT%
:GO14EXISTS

:: Make sure we build the new version from scratch
IF NOT EXIST %ROOT%\%VERSION% GOTO AFTERDELETEVERSION
    del /S /F /Q %ROOT%\%VERSION%
:AFTERDELETEVERSION

git clone --branch %VERSION% --depth=1 https://github.com/golang/go.git %VERSION%
cd %VERSION%\src
set GOROOT_BOOTSTRAP=%ROOT%\%VERSIONTAG14%
call make.bat
cd ..\..

IF NOT EXIST src GOTO AFTERDELETE
   del /S /F /Q src
   del /S /F /Q bin
   del /S /F /Q pkg
:AFTERDELETE

MKLINK /D src %VERSION%\src
MKLINK /D bin %VERSION%\bin
MKLINK /D pkg %VERSION%\pkg

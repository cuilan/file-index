@echo off
REM go build this application
REM @auther zhang.yan
REM @date 2021-10-20
TITLE build

go build -o file-index.exe .

rd /q /s dist
mkdir dist
mkdir dist\conf\
mkdir dist\logs\
mkdir dist\db\

copy file-index.exe .\dist\
copy .\bat\start.bat .\dist\
copy .\bat\clear.bat .\dist\
copy .\conf\ .\dist\conf\

REM PAUSE
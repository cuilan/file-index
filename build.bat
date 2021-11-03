@echo off
REM go build this application
REM @auther zhang.yan
REM @date 2021-10-20
TITLE build

go build -o main.exe .

rd /q /s dist
mkdir dist
mkdir dist\conf\
mkdir dist\logs\

copy main.exe .\dist\
copy .\bat\start.bat .\dist\
copy .\bat\clear.bat .\dist\
copy .\conf\ .\dist\conf\

cd dist
ren main.exe FileIndex.exe

REM PAUSE
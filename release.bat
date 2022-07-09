@echo off
go build -ldflags -H=windowsgui
ResourceHacker -open alpha-trainer.exe -save alpha-trainer.exe -action addskip -res assets/images/icon.ico -mask ICONGROUP,MAIN,
xcopy /s /y assets\* D:\Programo\custom\alpha-trainer\assets\
xcopy /y alpha-trainer.exe D:\Programos\custom\alpha-trainer\
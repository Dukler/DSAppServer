@ECHO OFF
:: Assign all Path variables
SET CGO_ENABLED=0
SET GOOS=windows
set output=C:\Users\iarwa\Workspace\Servers\DSAppServer\
set input=C:\Users\iarwa\Workspace\Go\DSAppServer\

go build -a -installsuffix cgo -o server.exe .

ECHO build finished
ECHO copying executable
xcopy "%input%server.exe" %output% /K /D /H /Y
ECHO copying utils
xcopy "%input%utils\connection.env" "%output%utils" /K /D /H /Y
ECHO copying migrations
xcopy "%input%migrations" "%output%migrations" /E /H /C /I
ECHO done.
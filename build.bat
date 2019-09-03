@ECHO OFF
:: Assign all Path variables
SET CGO_ENABLED=0
SET GOOS=linux
set output=C:\Users\iarwa\Workspace\Docker\dsapp\web\
set input=C:\Users\iarwa\Workspace\Go\DSAppServer\

go build -a -installsuffix cgo -o main .

ECHO build finished
ECHO copying executable
xcopy "%input%main" %output% /K /D /H /Y
ECHO copying utils
xcopy "%input%utils\connection.env" "%output%utils" /K /D /H /Y
ECHO copying migrations
xcopy "%input%migrations" "%output%" /K /D /H /Y
ECHO done.
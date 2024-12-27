setlocal
SET GOOS=windows
go build -ldflags="-H=windowsgui" -o water.exe .
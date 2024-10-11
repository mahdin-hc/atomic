set GOARCH=arm64
set GOOS=android
go build -ldflags="-s -w"
pause
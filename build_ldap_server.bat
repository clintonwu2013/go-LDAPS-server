mkdir linuxExec
copy .\ssl.crt .\linuxExec
copy .\ssl.key .\linuxExec
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

go build -ldflags "-s -w" -o linuxExec\ldaps_server_X64

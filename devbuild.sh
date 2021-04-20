rm Gamedb.db
GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOARCH=amd64 go build .
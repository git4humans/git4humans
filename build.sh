GOOS=linux
GOARCH=386
go build -o dist/linux-x86/git4humans/gh cmd/gh/main.go

GOOS=windows
GOARCH=386
go build -o dist/win-x86/git4humans/gh.exe cmd/gh/main.go

GOOS=darwin
GOARCH=amd64
go build -o dist/osx/git4humans/gh cmd/gh/main.go

GOOS=linux
GOARCH=amd64
go build -o dist/linux/git4humans/gh cmd/gh/main.go

GOOS=windows
GOARCH=amd64
go build -o dist/win/git4humans/gh.exe cmd/gh/main.go
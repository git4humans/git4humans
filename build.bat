set GOOS=linux
set GOARCH=386
go build -o dist/linux-x86/git4humans/gh cmd/gh/main.go

set GOOS=windows
set GOARCH=386
go build -o dist/win-x86/git4humans/gh.exe cmd/gh/main.go

set GOOS=darwin
set GOARCH=amd64
go build -o dist/osx/git4humans/gh cmd/gh/main.go

set GOOS=linux
set GOARCH=amd64
go build -o dist/linux/git4humans/gh cmd/gh/main.go

set GOOS=windows
set GOARCH=amd64
go build -o dist/win/git4humans/gh.exe cmd/gh/main.go
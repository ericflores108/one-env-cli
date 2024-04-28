
.PHONY: all test clean

all: clean test

build-all: build-darwin-free build-darwin-pro build-darwin-pro-profile build-linux-free build-linux-pro build-linux-pro-profile build-windows-free build-windows-pro build-windows-pro-profile

build-darwin-amd64-free:
	GOOS=darwin GOARCH=amd64 go build -tags "darwin free" -o builds/free/darwin/one-env-cli main.go
	chmod +x builds/free/darwin/one-env-cli

install-linux-amd64-free:
    GOOS=linux GOARCH=amd64 go install -tags "linux free" github.com/ericflores108/one-env-cli 

build-darwin-free:
	GOOS=darwin GOARCH=amd64 go build -tags "darwin free" -o builds/free/darwin/one-env-cli main.go
	chmod +x builds/free/darwin/one-env-cli

build-darwin-pro:
	GOOS=darwin GOARCH=amd64 go build -tags "darwin pro" -o builds/pro/darwin/one-env-cli main.go
	chmod +x builds/pro/darwin/one-env-cli

build-darwin-pro-profile:
	GOOS=darwin GOARCH=amd64 go build -tags "darwin pro profile" -o builds/profile/darwin/one-env-cli main.go
	chmod +x builds/profile/darwin/one-env-cli

build-linux-free:
	GOOS=linux go build -tags "linux free" -o builds/free/linux/one-env-cli main.go
	chmod +x builds/free/linux/one-env-cli

build-linux-pro:
	GOOS=linux go build -tags "linux pro" -o builds/pro/linux/one-env-cli main.go
	chmod +x builds/pro/linux/one-env-cli

build-linux-pro-profile:
	GOOS=linux go build -tags "linux pro profile" -o builds/profile/linux/one-env-cli main.go
	chmod +x builds/profile/linux/one-env-cli

build-windows-free:
	GOOS=windows go build -tags "windows free" -o builds/free/windows/one-env-cli.exe main.go

build-windows-pro:
	GOOS=windows go build -tags "windows pro" -o builds/pro/windows/one-env-cli.exe main.go

build-windows-pro-profile:
	GOOS=windows go build -tags "windows pro profile" -o builds/profile/windows/one-env-cli.exe main.go

install-darwin-free:
	go install -tags "darwin free" github.com/ericflores108/one-env-cli

install-darwin-pro:
	go install -tags "darwin pro" github.com/ericflores108/one-env-cli

install-darwin-pro-profile:
	go install -tags "darwin pro profile" github.com/ericflores108/one-env-cli

install-linux-free:
	GOOS=linux go install -tags "linux free" github.com/ericflores108/one-env-cli

install-linux-pro:
	GOOS=linux go install -tags "linux pro" github.com/ericflores108/one-env-cli

install-linux-pro-profile:
	GOOS=linux go install -tags "linux pro profile" github.com/ericflores108/one-env-cli

install-windows-free:
	GOOS=windows go install -tags "windows free" github.com/ericflores108/one-env-cli

install-windows-pro:
	GOOS=windows go install -tags "windows pro" github.com/ericflores108/one-env-cli

install-windows-pro-profile:
	GOOS=windows go install -tags "windows pro profile" github.com/ericflores108/one-env-cli

test:
	go test ./cmd -tags pro

test-verbose:
	go test -v ./cmd -tags pro

manpages:
	mkdir -p pages
	go run documentation/main.go

clean:
	go clean -cache -testcache -modcache
	rm -rf bin/
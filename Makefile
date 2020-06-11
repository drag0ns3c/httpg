bin := httpg
version := $(shell git rev-parse --abbrev-ref HEAD)

.PHONY: build
build: clean
	packr2
	mkdir build
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o build/$(bin)-$(version)-macos64
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/$(bin)-$(version)-linux64
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o build/$(bin)-$(version)-win64.exe
	packr2 clean

.PHONY: clean
clean:
	packr2 clean
	rm -rf build

.PHONY: init
init:
	go get ./...

.PHONY: fmt
fmt:
	go fmt ./...
bin := httpg
version := $(shell git rev-parse --abbrev-ref HEAD)

.PHONY: build
build: clean
	mkdir build
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o build/$(bin)-$(version)-macos64
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/$(bin)-$(version)-linux64
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o build/$(bin)-$(version)-win64.exe

.PHONY: clean
clean:
	rm -rf build

.PHONY: init
init:
	go get ./...

.PHONY: fmt
fmt:
	go fmt ./...
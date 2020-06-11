.PHONY: build
build: clean
	packr2
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o httpg-darwin-amd64
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o httpg-linux-amd64
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o httpg-windows-amd64

.PHONY: clean
clean:
	packr2 clean

.PHONY: init
init:
	go get ./...
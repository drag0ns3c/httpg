.PHONY: build
build: clean
	packr2

.PHONY: clean
clean:
	packr2 clean

.PHONY: init
init:
	go get ./...
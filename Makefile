.PHONY: all
all: build

BIN_DIR = $(PWD)/bin

.PHONY: build
build: dependencies
	go build -o bin/devicesvc cmd/main.go

clean:
	rm -rf bin/*

dependencies:
	go mod download

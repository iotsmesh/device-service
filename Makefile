.PHONY: all
all: build

BIN_DIR = $(PWD)/bin

.PHONY: build
build: dependencies build

clean:
	rm -rf bin/*
	
dependencies:
	go mod download

build:
	go build -o ./bin/devicesvc cmd/main.go
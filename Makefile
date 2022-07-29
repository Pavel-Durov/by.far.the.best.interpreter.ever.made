CWD := $(shell cd -P -- '$(shell dirname -- "$0")' && pwd -P)
SHELL := /bin/bash
export GO111MODULE := on
export GOBIN := $(CWD)/.bin

repl:
	go run src/main.go 

test:
	go test ./src/...

.PHONY:
build: build-amd64 build-arm64

build-%:
	GOARCH=$* go build -o ./bin/$*/ByBarTheBestInterpreterEverMade ./src
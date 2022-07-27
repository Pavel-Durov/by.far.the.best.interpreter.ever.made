CWD := $(shell cd -P -- '$(shell dirname -- "$0")' && pwd -P)
SHELL := /bin/bash
export GO111MODULE := on
export GOBIN := $(CWD)/.bin

run:
	go run src/main.go 

test:
	go test ./src/...
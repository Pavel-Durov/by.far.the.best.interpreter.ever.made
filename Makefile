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

clean:
	test -f ./waiig.zip && rm ./waiig.zip
	test -f ./waiig && rm -fr ./waiig

download-book-code-samples:
	curl https://interpreterbook.com/waiig_code_1.3.zip > waiig.zip && unzip -d ./waiig waiig.zip && rm ./waiig.zip
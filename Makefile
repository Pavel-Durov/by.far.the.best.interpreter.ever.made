CWD := $(shell cd -P -- '$(shell dirname -- "$0")' && pwd -P)
CGO_ENABLED=0

SHELL := /bin/bash
export GOBIN := $(CWD)/.bin

repl:
	go run src/main.go 

test:
	go test ./src/... --race

test-trace:
	TRACE=true go test -v ./src... 

.PHONY:
build: build-amd64 build-arm64

build-%:
	GOARCH=$* go build -trimpath -o ./bin/$*/ByBarTheBestInterpreterEverMade ./src 

clean:
	test -f ./waiig.zip && rm ./waiig.zip
	test -f ./waiig && rm -fr ./waiig

download-book-code-samples:
	curl https://interpreterbook.com/waiig_code_1.3.zip > waiig.zip && unzip -d ./waiig waiig.zip && rm ./waiig.zip
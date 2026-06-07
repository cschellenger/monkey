.PHONY: test build generate

EXECUTABLES = java go
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))

generate: grammar/MonkeyLexer.g4 grammar/MonkeyParser.g4
	go generate ./...

test: generate
	go test ./...

build: generate
	go build .

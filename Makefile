.PHONY: build test install binaries help

help:
	@echo make build test install
	@echo make binaries

build:
	go build

test:
	go test

install:
	go install

dist:
	mkdir $@

DEPS:=*.go go.sum Makefile
GOBUILDFLAGS:=-ldflags "-s -w"

dist/echo-json.exe: $(DEPS)
	GOOS=windows GOARCH=amd64 go build $(GOBUILDFLAGS) -o $@

dist/echo-json.%: $(DEPS)
	GOOS=$* GOARCH=amd64 go build $(GOBUILDFLAGS) -o $@

binaries: dist $(addprefix dist/echo-json.,linux darwin exe)
	@ls -l dist

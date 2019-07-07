.PHONY: build test install binaries help image push clean

help:
	@echo make build test install
	@echo make binaries
	@echo make image push
	@echo make clean

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

image:
	docker build --pull -t filex/echo-json .

push: image
	docker push filex/echo-json:latest

clean:
	rm -rf dist/
	docker rmi -f $$(docker images -f 'reference=filex/echo-json' --format '{{.ID}}')

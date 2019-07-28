.PHONY: build test install binaries help image push clean

all: test build

help:
	@echo make test build install
	@echo make binaries
	@echo make image push
	@echo make clean

TRAVIS_TAG?=$(shell git describe --tags HEAD)

TAG:=$(TRAVIS_TAG)
gobuild_args := -ldflags "-s -w -X main.Version=$(TAG)"

build: echo-json

echo-json:
	go build $(gobuild_args)

test:
	go test

$$GOPATH/bin/echo-json: echo-json
	cp $< $@

install: $$GOPATH/bin/echo-json

dist:
	mkdir $@

DEPS:=*.go go.sum Makefile

dist/echo-json.exe: $(DEPS)
	GOOS=windows GOARCH=amd64 go build $(gobuild_args) -o $@

dist/echo-json.%: $(DEPS)
	GOOS=$* GOARCH=amd64 go build $(gobuild_args) -o $@

binaries: dist $(addprefix dist/echo-json.,linux darwin exe)
	@ls -l dist

image:
	docker build --build-arg TAG="$(TAG)" --pull -t filex/echo-json:$(TAG) .

push: image
	docker push filex/echo-json:$(TAG)
	docker tag filex/echo-json:$(TAG) filex/echo-json:latest
	docker push filex/echo-json:latest

clean:
	rm -rf dist/ echo-json
	- docker rmi -f $$(docker images -f 'reference=filex/echo-json' --format '{{.ID}}')

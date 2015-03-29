prefix ?= /usr/local

.PHONY: all build assets test install clean

all: build

build: muxt

assets: assets.go

muxt: *.go main/*.go
	cd main && go build -v -o ../muxt

assets.go: assets/**/*
	go-bindata -pkg=muxt -o assets.go -prefix=assets/ assets/...

test: assets
	go test

install: all
	install -d $(prefix)/bin
	install muxt $(prefix)/bin

clean:
	rm -f assets.go muxt


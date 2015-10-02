prefix ?= /usr/local

.PHONY: all build assets test install clean

all: build

build: bin/muxt

assets: src/muxt/assets/assets.go

bin/muxt: bin assets src/muxt/**/*.go
	gb build

bin:
	mkdir bin

src/muxt/assets/assets.go: vendor/bin/go-bindata src/muxt/assets/**/*
	vendor/bin/go-bindata -pkg=assets -o src/muxt/assets/assets.go -prefix=src/muxt/assets/ src/muxt/assets/...

test: assets
	gb test

install: all
	install -d $(prefix)/bin
	install bin/muxt $(prefix)/bin

clean:
	rm -rf src/muxt/assets/assets.go bin pkg

vendor/bin:
	mkdir -p vendor/bin

vendor/bin/go-bindata: vendor/bin
	GOPATH=$(PWD)/vendor go build -o vendor/bin/go-bindata github.com/jteeuwen/go-bindata/go-bindata

prefix ?= /usr/local

.PHONY: all build assets dependencies test install clean

all: dependencies test build

build: assets src/muxt/**/*.go
	gb build

assets: src/muxt/assets/assets.go

dependencies:
	gb vendor restore
	cd vendor && gb build

src/muxt/assets/assets.go: src/muxt/assets/**/*
	vendor/bin/go-bindata -pkg=assets -o src/muxt/assets/assets.go -prefix=src/muxt/assets/ src/muxt/assets/...

test: assets
	gb test

install: all
	install -d $(prefix)/bin
	install bin/muxt $(prefix)/bin

clean:
	rm -rf src/muxt/assets/assets.go bin pkg


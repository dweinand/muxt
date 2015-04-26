prefix ?= /usr/local

.PHONY: all build assets test install clean

all: build

build: bin/muxt

assets: assets/assets.go

bin/muxt: bin *.go cmd/muxt.go assets/assets.go
	cd cmd && go build -v -o ../bin/muxt

bin:
	mkdir bin

assets/assets.go: assets/**/*
	go-bindata -pkg=assets -o assets/assets.go -prefix=assets/ assets/...

test: assets
	go test

install: all
	install -d $(prefix)/bin
	install bin/muxt $(prefix)/bin

clean:
	rm -rf assets/assets.go bin

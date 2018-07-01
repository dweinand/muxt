prefix ?= /usr/local

.PHONY: all build assets test install clean

all: test build

build: assets **/*.go
	vgo build github.com/dweinand/muxt/cmd/muxt -o bin/muxt
	packr clean

assets:
	go get -u github.com/gobuffalo/packr/packr
	packr

test: assets
	vgo test

install: all
	install -d $(prefix)/bin
	install bin/muxt $(prefix)/bin

clean:
	rm -rf bin pkg

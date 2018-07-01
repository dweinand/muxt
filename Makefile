prefix ?= /usr/local
OSES ?= darwin

.PHONY: all build assets test install clean

all: test build

build: assets **/*.go $(OSES)
	packr clean

$(OSES): **/*.go
	GOOS=$@ GOARCH=amd64 vgo build -o bin/muxt-$@-amd64 ./cmd/muxt

assets:
	packr

dependencies:
	go get -u golang.org/x/vgo
	go get -u github.com/gobuffalo/packr/packr

test: assets
	vgo test -race -coverprofile=coverage.txt -covermode=atomic

coverage:
	bash <(curl -s https://codecov.io/bash)

install: all
	install -d $(prefix)/bin
	install bin/muxt $(prefix)/bin

clean:
	rm -rf bin pkg

release:
	- ghr -t $(GITHUB_TOKEN) -u $(GITHUB_USERNAME) -r $(GITHUB_REPONAME) $(GITHUB_TAG) bin/

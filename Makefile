prefix ?= /usr/local
OSES ?= darwin

.PHONY: all build test install clean

all: test build

build: **/*.go $(OSES)

$(OSES): **/*.go
	mkdir -p bin/$@-amd64
	GOOS=$@ GOARCH=amd64 go build -o bin/$@-amd64/muxt ./cmd/muxt

dependencies:
	go get -u github.com/gobuffalo/packr/packr

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic

install: all
	install -d $(prefix)/bin
	install bin/muxt $(prefix)/bin

clean:
	rm -rf bin pkg

release:
	mkdir -p pkg
	$(foreach os,$(OSES),tar -C bin/$(os)-amd64 -cvzf pkg/muxt_$(GITHUB_TAG)_$(os)-amd64.tar.gz muxt;)
	ghr -t $(GITHUB_TOKEN) -u $(GITHUB_USERNAME) -r $(GITHUB_REPONAME) $(GITHUB_TAG) pkg/

prefix ?= /usr/local
OSES ?= darwin

.PHONY: all build assets test install clean

all: test build

build: assets **/*.go $(OSES)
	packr clean

$(OSES): **/*.go
	mkdir -p bin/$@-amd64
	GOOS=$@ GOARCH=amd64 go build -o bin/$@-amd64/muxt ./cmd/muxt

assets:
	packr

dependencies:
	go get -u github.com/gobuffalo/packr/packr

test: assets
	go test -race -coverprofile=coverage.txt -covermode=atomic

install: all
	install -d $(prefix)/bin
	install bin/muxt $(prefix)/bin

clean:
	rm -rf bin pkg

release:
	mkdir -p pkg
	$(foreach var,$(OSES),tar -C bin/$(var)-amd64 -cvzf pkg/muxt_$(GITHUB_TAG)_$(var)-amd64.tar.gz muxt;)
	ghr -t $(GITHUB_TOKEN) -u $(GITHUB_USERNAME) -r $(GITHUB_REPONAME) $(GITHUB_TAG) pkg/

default: main/main.go assets.go
	cd main && go build -v -o ../muxt

test: assets.go
	go test

clean:
	rm assets.go muxt

assets.go: assets/tmux/command.sh
	go-bindata -pkg=muxt -o assets.go -prefix=assets/ assets/...


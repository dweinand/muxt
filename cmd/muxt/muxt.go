package main

import (
	"github.com/dweinand/muxt"
	"github.com/dweinand/muxt/shell"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) != 2 {
		usage()
	}

	start(os.Args[1])
}

func start(name string) {
	session, err := muxt.Load(name)
	if err != nil {
		log.Fatal(err)
	}

	err = session.Start(shell.New("sh"))
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	log.Fatal("Usage: muxt [NAME]")
}

package main

import (
	"log"
	"muxt"
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

	err = session.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	log.Fatal("Usage: muxt [NAME]")
}
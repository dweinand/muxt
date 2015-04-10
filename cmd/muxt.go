package main

import (
	"fmt"
	"github.com/dweinand/muxt"
	"os"
)

func main() {
	args := os.Args
	numArgs := len(args)
	if numArgs < 1 {
		usage()
	}

	start(args[1])
}

func start(name string) {
	session, err := muxt.Load(name)
	exitOnError(err)

	err = session.Start()
	exitOnError(err)
}

func usage() {
	fmt.Println("muxt [NAME]")
	os.Exit(1)
}

func exitOnError(err error) {
	if err != nil {
		fmt.Printf("[error] %v\n", err)
		os.Exit(1)
	}
}

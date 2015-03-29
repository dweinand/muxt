package main

import (
	"fmt"
	"github.com/dweinand/muxt"
	"os"
	"path/filepath"
	"strings"
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
	session, err := loadConfig(name)
	assertNoError(err)

	err = muxt.Start(session)
	assertNoError(err)
}

func usage() {
	fmt.Print("muxt [NAME]\n")
	os.Exit(1)
}

func loadConfig(name string) (*muxt.Session, error) {
	configDir := filepath.Join(os.Getenv("HOME"), ".muxt")
	configFile := strings.Join([]string{name, ".toml"}, "")
	configPath := filepath.Join(configDir, configFile)
	return muxt.Load(configPath)
}

func assertNoError(err error) {
	if err != nil {
		panic(err)
	}
}

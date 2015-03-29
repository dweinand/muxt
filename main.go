package main

import (
	"fmt"
	"github.com/dweinand/muxt/pkg/config"
	"github.com/dweinand/muxt/pkg/tmux"
	"os"
	"path/filepath"
	"strings"
)

func loadConfig(name string) (*config.Session, error) {
	configDir := filepath.Join(os.Getenv("HOME"), ".muxt")
	confiFile := strings.Join([]string{name, ".toml"}, "")
	configPath := filepath.Join(configDir, confiFile)
	return config.Load(configPath)
}

func start(name string) {
	session, err := loadConfig(name)
	if err != nil {
		panic(err)
	}
	err = tmux.StartSession(session)
	if err != nil {
		panic(err)
	}
}

func usage() {
	fmt.Print("muxt [NAME]\n")
	os.Exit(1)
}

func main() {
	args := os.Args
	numArgs := len(args)
	if numArgs < 1 {
		usage()
	}

	start(args[1])
}

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
	tmux.StartSession(session)
}

func usage() {
	fmt.Print("muxt start [NAME]\n")
	os.Exit(1)
}

func main() {
	args := os.Args
	numArgs := len(args)
	if numArgs < 2 {
		usage()
	}
	command := args[1]
	switch command {
	case "start":
		if numArgs < 3 {
			usage()
		}
		start(args[2])
	default:
		usage()
	}
}

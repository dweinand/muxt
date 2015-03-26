package config

import (
	"github.com/naoina/toml"
	"io/ioutil"
	"os"
)

func Parse(buf []byte) (*Session, error) {
	var config Session
	err := toml.Unmarshal(buf, &config)
	return &config, err
}

func Load(path string) (*Session, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return Parse(buf)
}

type Session struct {
	Name   string
	Root   string
	Window []Window
}

type Window struct {
	Name    string
	Root    string
	Layout  string
	Command string
	Pane    []Pane
}

type Pane struct {
	Name    string
	Command string
}

package muxt

import (
	"bytes"
	"fmt"
	"github.com/dweinand/shell"
	"github.com/naoina/toml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Session struct {
	Name      string
	Root      string
	Window    []Window
	Pre       string
	PreWindow string
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

var ConfigDir string
var homeDir string
var shellExec func(string) error

func init() {
	homeDir = os.Getenv("HOME")
	ConfigDir = filepath.Join(homeDir, ".muxt")
	shellExec = shell.Exec
}

func Parse(buf []byte) (*Session, error) {
	var session Session
	err := toml.Unmarshal(buf, &session)

	return &session, err
}

func Load(name string) (*Session, error) {
	path := PathFromName(name)

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	session, err := Parse(buf)
	if err != nil {
		return nil, err
	}
	if session.Name == "" {
		session.Name = NameFromPath(path)
	}

	return session, nil
}

func PathFromName(name string) string {
	_, err := os.Stat(name)
	if err != nil && os.IsNotExist(err) {
		if filepath.Ext(name) != ".toml" {
			name = fmt.Sprintf("%v.toml", name)
		}
		return filepath.Join(ConfigDir, name)
	}

	return name
}

func NameFromPath(path string) string {
	path = filepath.Base(path)
	return strings.Replace(path, filepath.Ext(path), "", 1)
}

func (s *Session) Start() error {
	command, err := s.Script()
	if err != nil {
		return err
	}

	err = shellExec(command)
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) Script() (string, error) {
	var b bytes.Buffer

	command, err := Asset("tmux/command.sh")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("command").Parse(string(command))
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&b, s)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

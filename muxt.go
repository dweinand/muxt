package muxt

import (
	"bytes"
	"github.com/dweinand/shell"
	"github.com/naoina/toml"
	"io/ioutil"
	"os"
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
func Start(sessionConfig *Session) error {
	command, err := StartUpScript(sessionConfig)
	if err != nil {
		return err
	}
	err = shell.Exec(command)
	if err != nil {
		return err
	}
	return nil
}

func StartUpScript(sessionConfig *Session) (string, error) {
	var b bytes.Buffer
	command, err := Asset("tmux/command.sh")
	if err != nil {
		return "", err
	}
	tmpl, err := template.New("command").Parse(string(command))
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(&b, sessionConfig)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

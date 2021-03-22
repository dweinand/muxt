package muxt

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/dweinand/muxt/shell"
	"github.com/naoina/toml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed assets/tmux/command.sh
var tmuxCommand embed.FS

// ConfigDir is the location where muxt will look for config files
var ConfigDir = filepath.Join(os.Getenv("HOME"), ".muxt")

// Session is a session configuration
type Session struct {
	Name       string
	Root       string
	Window     []Window
	Pre        string
	PreWindow  string
	execScript func(string) error
}

// Window is a window configuration
type Window struct {
	Name     string
	Root     string
	Layout   string
	Commands []string
	Pane     []Pane
}

// Pane is a pane configuration
type Pane struct {
	Name     string
	Commands []string
}

// NewSession creates a new empty Session
func NewSession() *Session {
	return &Session{}
}

// Parse converts toml data into a Session
func Parse(buf []byte) (*Session, error) {
	session := NewSession()

	if err := toml.Unmarshal(buf, session); err != nil {
		return nil, err
	}

	return session, nil
}

// Load looks for a toml config file and then parses it into a Session
func Load(name string) (*Session, error) {
	path := pathFromName(name)

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
		session.Name = nameFromPath(path)
	}

	return session, nil
}

func pathFromName(name string) string {
	_, err := os.Stat(name)
	if err != nil && os.IsNotExist(err) {
		if filepath.Ext(name) != ".toml" {
			name = fmt.Sprintf("%v.toml", name)
		}

		return filepath.Join(ConfigDir, name)
	}

	return name
}

func nameFromPath(path string) string {
	path = filepath.Base(path)

	return strings.Replace(path, filepath.Ext(path), "", 1)
}

// Start launches tmux and loads the Session
func (s *Session) Start(sh shell.Execer) error {
	script, err := s.Script()
	if err != nil {
		return err
	}

	return sh.Exec(script)
}

// Script returns the generated shell commands without executing them
func (s *Session) Script() (string, error) {
	var b bytes.Buffer

	command, err := tmuxCommand.ReadFile("assets/tmux/command.sh")
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

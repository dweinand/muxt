package muxt

import (
	"bytes"
	"fmt"
	"github.com/naoina/toml"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"text/template"
)

type Session struct {
	Name       string
	Root       string
	Window     []Window
	Pre        string
	PreWindow  string
	execScript func(string) error
}

type Window struct {
	Name     string
	Root     string
	Layout   string
	Commands []string
	Pane     []Pane
}

type Pane struct {
	Name     string
	Commands []string
}

var (
	ConfigDir = filepath.Join(os.Getenv("HOME"), ".muxt")
)

func NewSession() *Session {
	return &Session{
		execScript: func(script string) error {
			sh, err := exec.LookPath("sh")
			if err != nil {
				return err
			}

			return syscall.Exec(sh, []string{"sh", "-c", script}, os.Environ())
		},
	}
}

func Parse(buf []byte) (*Session, error) {
	session := NewSession()

	if err := toml.Unmarshal(buf, session); err != nil {
		return nil, err
	}

	return session, nil
}

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

func (s *Session) Start() error {
	script, err := s.Script()
	if err != nil {
		return err
	}

	return s.execScript(script)
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

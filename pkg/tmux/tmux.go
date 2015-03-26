package tmux

import (
	"bytes"
	"github.com/dweinand/muxt/pkg/assets"
	"github.com/dweinand/muxt/pkg/config"
	"github.com/dweinand/muxt/pkg/shell"
	"text/template"
)

func StartSession(sessionConfig *config.Session) error {
	command, err := renderCommand(sessionConfig)
	if err != nil {
		return err
	}
	err = shell.Exec(command)
	if err != nil {
		return err
	}
	return nil
}

func renderCommand(sessionConfig *config.Session) (string, error) {
	var b bytes.Buffer
	command, err := assets.Asset("tmux/command.sh")
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

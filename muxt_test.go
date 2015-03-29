package muxt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func init() {
	shellExec = func(s string) error {
		fmt.Print(s)
		return nil
	}
}

func testSessionConfig() *Session {
	session := &Session{
		Name: "test",
		Root: "~/",
		Window: []Window{
			{
				Name:   "editor",
				Layout: "main-vertical",
				Pane: []Pane{
					{
						Command: "vim",
					},
					{
						Command: "guard",
					},
				},
			},
			{
				Name:    "server",
				Command: "bundle exec rails s",
			},
			{
				Name:    "logs",
				Command: "tail -f log/development.log",
			},
		},
	}
	return session
}

func TestLoad(t *testing.T) {
	path := "assets/config/test.toml"
	expect := testSessionConfig()
	actual, err := Load(path)
	assert.Nil(t, err)
	assert.Equal(t, actual, expect)
}

func TestLoadName(t *testing.T) {
	ConfigDir = "assets/config"
	name := "test"
	expect := testSessionConfig()
	actual, err := Load(name)
	assert.Nil(t, err)
	assert.Equal(t, actual, expect)
}

func TestLoadNameWithExt(t *testing.T) {
	oldConfigDir := ConfigDir
	ConfigDir = "assets/config"
	name := "test.toml"
	expect := testSessionConfig()
	actual, err := Load(name)
	assert.Nil(t, err)
	assert.Equal(t, actual, expect)
	ConfigDir = oldConfigDir
}

func TestLoadEmpty(t *testing.T) {
	path := "assets/config/empty.toml"
	expect := &Session{Name: "empty"}
	actual, err := Load(path)
	assert.Nil(t, err)
	assert.Equal(t, actual, expect)
}

func TestStart(t *testing.T) {
	session := &Session{Name: "snaggletooth"}
	shellExec = func(s string) error {
		pattern := "has-session -t snaggletooth"
		matched, err := regexp.MatchString(pattern, s)
		assert.Nil(t, err)
		assert.True(t, matched)
		return nil
	}
	err := session.Start()
	assert.Nil(t, err)
}

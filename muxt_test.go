package muxt

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

type fakeShell struct {
	script string
}

var oldConfigDir string

func (f *fakeShell) Exec(script string) error {
	f.script = script
	return nil
}

func TestLoad(t *testing.T) {
	path := "assets/config/test.toml"

	expected := testSessionConfig()
	actual, err := Load(path)

	assert.Nil(t, err)
	assert.Equal(t, actual, expected)
}

func TestLoadName(t *testing.T) {
	oldConfigDir, ConfigDir = ConfigDir, "assets/config"
	defer func() {
		ConfigDir = oldConfigDir
	}()

	name := "test"

	expected := testSessionConfig()
	actual, err := Load(name)

	assert.Nil(t, err)
	assert.Equal(t, actual, expected)
}

func TestLoadNameWithExt(t *testing.T) {
	oldConfigDir, ConfigDir = ConfigDir, "assets/config"
	defer func() {
		ConfigDir = oldConfigDir
	}()

	name := "test.toml"

	expected := testSessionConfig()
	actual, err := Load(name)

	assert.Nil(t, err)
	assert.Equal(t, actual, expected)
}

func TestLoadEmpty(t *testing.T) {
	path := "assets/config/empty.toml"

	expected := NewSession()
	expected.Name = "empty"
	actual, err := Load(path)

	assert.Nil(t, err)
	assert.Equal(t, actual, expected)
}

func TestStart(t *testing.T) {
	pattern := "has-session -t snaggletooth"
	session := &Session{Name: "snaggletooth"}

	sh := &fakeShell{}

	session.Start(sh)
	matched, err := regexp.MatchString(pattern, sh.script)

	assert.Nil(t, err)
	assert.True(t, matched)
}

func TestScript(t *testing.T) {
	pattern := "has-session -t snaggletooth"
	session := &Session{Name: "snaggletooth"}

	actual, err := session.Script()
	assert.Nil(t, err)
	matched, err := regexp.MatchString(pattern, actual)

	assert.Nil(t, err)
	assert.True(t, matched)
}

func testSessionConfig() *Session {
	session := NewSession()
	session.Name = "test"
	session.Root = "~/"
	session.Window = []Window{
		{
			Name:   "editor",
			Layout: "main-vertical",
			Pane: []Pane{
				{
					Commands: []string{"vim"},
				},
				{
					Commands: []string{"guard"},
				},
			},
		},
		{
			Name:     "server",
			Commands: []string{"bundle exec rails s"},
		},
		{
			Name:     "logs",
			Commands: []string{"tail -f log/development.log"},
		},
	}

	return session
}

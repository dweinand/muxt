package muxt

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

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

func TestScript(t *testing.T) {
	session := &Session{Name: "snaggletooth"}

	script, err := session.Script()
	assert.Nil(t, err)

	pattern := "has-session -t snaggletooth"
	matched, err := regexp.MatchString(pattern, script)
	assert.Nil(t, err)
	assert.True(t, matched)
}

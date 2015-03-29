package muxt

import (
	"github.com/stretchr/testify/assert"
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
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, actual, expect)
}

func TestLoadName(t *testing.T) {
	ConfigDir = "assets/config"
	name := "test"
	expect := testSessionConfig()
	actual, err := Load(name)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, actual, expect)
}

func TestLoadNameWithExt(t *testing.T) {
	oldConfigDir := ConfigDir
	ConfigDir = "assets/config"
	name := "test.toml"
	expect := testSessionConfig()
	actual, err := Load(name)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, actual, expect)
	ConfigDir = oldConfigDir
}

func TestLoadEmpty(t *testing.T) {
	path := "assets/config/empty.toml"
	expect := &Session{Name: "empty"}
	actual, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, actual, expect)
}

package shell

import (
	"os"
	"os/exec"
	"syscall"
)

// Execer is an interface that wraps the Exec method.
type Execer interface {
	Exec(string) error
}

// Sh implements Execer.
type Sh struct {
	name string
}

// New creates a new shell.
func New(name string) *Sh {
	return &Sh{name: name}
}

// Exec runs a script
func (s *Sh) Exec(script string) error {
	path, err := exec.LookPath(s.name)
	if err != nil {
		return err
	}

	return syscall.Exec(path, []string{s.name, "-c", script}, os.Environ())
}

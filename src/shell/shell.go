package shell

import (
	"os"
	"os/exec"
	"syscall"
)

type Execer interface {
	Exec(string) error
}
type sh struct {
	name string
}

func New(name string) *sh {
	return &sh{name: name}
}

func (s *sh) Exec(script string) error {
	path, err := exec.LookPath(s.name)
	if err != nil {
		return err
	}

	return syscall.Exec(path, []string{s.name, "-c", script}, os.Environ())
}

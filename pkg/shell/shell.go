package shell

import (
	"os"
	"os/exec"
	"syscall"
)

func Exec(script string) error {
	bin, err := exec.LookPath("sh")
	if err != nil {
		panic(err)
	}
	return syscall.Exec(bin, []string{"sh", "-c", script}, os.Environ())
}

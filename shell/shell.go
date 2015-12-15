package shell

import (
	"os"
	"os/exec"
)

// Run is a no muss, no fuss way to run a command in a shell.
func Run(command ...string) (string, error) {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	return string(output), err
}

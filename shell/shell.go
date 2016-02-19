// +build !windows
package shell

import (
	"os"
	"os/exec"
	"strings"
)

var tmpDir = ""

// Run is a no muss, no fuss way to run a command in a shell.
func Run(command ...string) (string, error) {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	return string(output), err
}

// TempDir makes and returns a temporary directory. Any provided arguments will
// be evaluated as successive directories within that temporary directory.
// Unix-only.
func TempDir(subdir ...string) (string, error) {
	if tmpDir == "" {
		// Try to mktmp the base directory
		dirPrefix := "temp"
		if len(subdir) > 0 {
			dirPrefix = subdir[0]
		}

		dir, err := Run("mktemp", "-d", "-t", dirPrefix)
		if err != nil {
			return "", err
		}
		tmpDir = strings.Trim(dir, "\n\t ")
	}

	// No subdirectory
	if len(subdir) < 2 {
		return tmpDir, nil
	}

	// Make the directories underneath the base directory
	fullPath := tmpDir + strings.Join(subdir[1:], "/")
	if err := os.MkdirAll(fullPath, os.ModeDir); err != nil {
		return "", err
	}

	return fullPath, nil
}

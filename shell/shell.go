package shell

import (
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// Run is a no muss, no fuss way to run a command in a shell.
func Run(command ...string) (string, error) {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	return string(output), err
}

// TempDir gets the system temporary directory and makes a random directory
// underneath that for this probram to use. It may be the developer's
// responsibility to clean up this directory before the application closes.
// Consider using cleanup.Defer() in this repository.
func TempDir() (string, error) {
	length := 8
	base := os.TempDir()
	var dirName string

	for {
		dirName = RandLowerChars(length)
		if !FileExists(base + dirName) {
			break
		}
	}

	if err := os.MkdirAll(base+dirName, os.ModeDir|os.ModePerm); err != nil {
		return "", err
	}

	return base + dirName, nil
}

func FileExists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// TODO move this elsewhere
const lowerChars = "0123456789abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandLowerChars(length int) string {
	subdir := make([]byte, length)

	for i := 0; i < length; i++ {
		subdir[i] = lowerChars[rand.Intn(len(lowerChars))]
	}

	return string(subdir)
}

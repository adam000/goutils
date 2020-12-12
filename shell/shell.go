package shell

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// Run is a no muss, no fuss way to run a command in a shell.
// Deprecated: use RunInDir() instead
func Run(command ...string) (string, error) {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	return string(output), err
}

// A DirMatch is a predicate to determine whether a given directory (as a string)
// matches a set of rules
type DirMatch func(dir string) (bool, error)

// ParseDirectory checks if `dir` fits a predicate `matches`. If it encounters an error it
// cannot proceed from, it returns `stopErr`. Otherwise, it returns all `subdirs` for the
// parent to recurse into.
func ParseDirectory(matches DirMatch, dir string) (isRoot bool, subdirs []string, stopErr error) {
	if stat, err := os.Stat(dir); err != nil {
		stopErr = fmt.Errorf("Could not read root directory: %w", err)
		return
	} else if !stat.IsDir() {
		stopErr = fmt.Errorf("'%s' is not a directory", dir)
		return
	}

	isMatch, err := matches(dir)
	if err != nil {
		stopErr = fmt.Errorf("Running command in directory '%s': %w", dir, err)
		return
	} else if !isMatch {
		// Recurse within
		files, dirErr := ioutil.ReadDir(dir)
		if dirErr != nil {
			stopErr = fmt.Errorf("Reading directory '%s': %w", dir, dirErr)
			return
		}

		subdirs := make([]string, 0)
		for _, fileInfo := range files {
			if fileInfo.IsDir() {
				nextDir := filepath.Join(dir, fileInfo.Name())
				subdirs = append(subdirs, nextDir)
			}
		}
		return false, subdirs, nil
	}

	return true, nil, nil
}

// RunInDir runs `command` from `directory`, and returns any text from the `command` or `err` that occurred.
func RunInDir(directory string, command ...string) (stdoutText string, stderrText string, err error) {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Dir = directory

	// Create Pipes
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", "", fmt.Errorf("Failed to create stdout pipe: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", "", fmt.Errorf("Failed to create stderr pipe: %w", err)
	}

	// Start Command
	if err = cmd.Start(); err != nil {
		return "", "", fmt.Errorf("Failed to start command: %w", err)
	}

	// Read From Pipes
	stdoutBytes, stdoutErr := ioutil.ReadAll(stdout)
	if stdoutErr != nil {
		return "", "", fmt.Errorf("Failed reading stdout (%w) (perhaps another error: %v)", stdoutErr, err)
	}
	outputStr := string(stdoutBytes)

	stderrBytes, stderrErr := ioutil.ReadAll(stderr)
	if stderrErr != nil {
		return outputStr, "", fmt.Errorf("Failed reading stderr (%w) while trying to deal with command failure: %v", stderrErr, err)
	}

	// Finish Command
	err = cmd.Wait()
	return outputStr, string(stderrBytes), err
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

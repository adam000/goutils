// +build windows

package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func Open(filename string) (*os.File, error) {
	dir := os.Getenv("LOCALAPPDATA")
	if dir == "" {
		return nil, fmt.Errorf("Could not find %%LOCALAPPDATA%% envar")
	}

	f, err := os.Open(filepath.Join(dir, filename))
	return f, err
}

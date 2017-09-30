// +build windows

// Package config wraps xdgdir.Config to allow for OS-independent configuration.
package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func Open(filename string) (*os.File, error) {
	dir := os.Getenv("USERPROFILE")
	if dir == "" {
		return nil, fmt.Errorf("Could not find %%USERPROFILE%% envar")
	}

	f, err := os.Open(filepath.Join(dir, filename))
	return f, err
}

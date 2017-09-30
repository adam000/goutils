// +build !windows

// Package config wraps xdgdir.Config to allow for OS-independent configuration.
package config

import (
	"os"

	"go4.org/xdgdir"
)

func Open(filename string) (*os.File, error) {
	file, err := xdgdir.Config.Open(filename)
	return file, err
}

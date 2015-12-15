package git

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/adam000/goutils/shell"
)

// Try to get the root directory of this repo
func DetectRoot() (string, error) {
	output, err := shell.Run("git", "rev-parse", "--show-toplevel")
	if err != nil {
		return "", fmt.Errorf("Error: %s\n(Hint: this error may happen if you run `sync` without " +
			"any params and outside of any git repository)\n")
	}

	return strings.Trim(filepath.Base(output), "\n "), nil
}

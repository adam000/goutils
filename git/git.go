package git

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/adam000/goutils/shell"
)

// Try to get the root directory of this repo
// Deprecated: use IsGitRoot() instead
func DetectRoot() (string, error) {
	output, err := shell.Run("git", "rev-parse", "--show-toplevel")
	if err != nil {
		return "", fmt.Errorf("Error: %v\n(Hint: this error may happen if you run `sync` without "+
			"any params and outside of any git repository)\n", err)
	}

	return strings.Trim(filepath.Base(output), "\n "), nil
}

// A predicate to determine if `dir` is the root of a git repository.
// Combos well with `shell.ParseDirectory`, for example.
func IsGitRoot(dir string) (bool, error) {
	// Try `git rev-parse --show-toplevel` at our current directory
	result, stderrText, err := shell.RunInDir(dir, "git", "rev-parse", "--show-toplevel")
	if err != nil {
		if !strings.Contains(stderrText, "not a git repository") {
			return false, fmt.Errorf("Failed to run git rev-parse: %w", err)
		}
		return false, nil
	}

	// Make sure we're in the base -- if we ever hit this error, we have... big problems
	// Not 100% sure this is an exactly accurate way to find this out, might need os.SameFile
	if strings.TrimSpace(result) != dir {
		return false, fmt.Errorf("Job directory '%s' is not the base git directory '%s'", dir, result)
	}

	return true, nil
}

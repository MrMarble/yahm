package git

import (
	"errors"
	"os/exec"
	"strings"
)

var ErrGitNotFound = errors.New(".git folder not found")

// GetGitRoot returns the root of the git repository using git rev-parse --get-git-dir
func GetGitRoot() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--git-dir").Output()
	if err != nil {
		return "", err
	}
	gitDir := strings.TrimSpace(string(out))
	return gitDir, nil
}

package git

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGitProjectRoot(t *testing.T) {
	_, currentFile, _, _ := runtime.Caller(0)
	currentPath := filepath.Join(filepath.Dir((filepath.Dir(currentFile))))
	gitProjectRoot := filepath.Dir(currentPath) + "/.git"

	tests := []struct {
		name string
		dir  string
		want string
	}{
		{"works from .git directory", gitProjectRoot, gitProjectRoot},
		{"works from any directory", currentPath, gitProjectRoot},
		{"works from any file", currentFile, gitProjectRoot},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGitRoot()
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

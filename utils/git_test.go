package utils

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestGetGitProjectRoot(t *testing.T) {
	_, currentFile, _, _ := runtime.Caller(0)
	currentPath := filepath.Join(filepath.Dir(currentFile))
	gitProjectRoot := filepath.Join(filepath.Dir(currentPath), ".git")

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
			got, err := GetGitProjectRoot(tt.dir)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGitProjectRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

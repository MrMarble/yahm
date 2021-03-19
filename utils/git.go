package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// GetGitProjectRoot returns de closest .git directory
func GetGitProjectRoot(directory string) (string, error) {
	start := filepath.Clean(directory)
	return getGitProjectRoot(strings.Split(start, string(filepath.Separator)))
}

func getGitProjectRoot(directory []string) (string, error) {
	if len(directory) < 2 {
		return "", errors.New("no more paths to traverse")
	}
	pop(&directory)
	dir := filepath.Join(directory...)
	fullPath := absolute(filepath.Join(dir, ".git"))
	if info, err := os.Stat(fullPath); !os.IsNotExist(err) {
		if !info.IsDir() {
			buff, err := ioutil.ReadFile(fullPath)
			if err != nil {
				return "", err
			}
			match, err := regexp.Match(`(?m)^gitdir: (.*)\s*$`, buff)
			if err != nil {
				return "", err
			}
			if match {
				return filepath.Clean(fullPath), nil
			}
		}
		return filepath.Clean(fullPath), nil
	} else {
		return getGitProjectRoot(directory)
	}
}

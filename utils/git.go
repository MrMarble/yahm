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
	dir := filepath.Join(directory...)
	fullPath := absolute(filepath.Join(dir, ".git"))
	info, err := os.Stat(fullPath)
	_, pathErr := err.(*os.PathError)
	if os.IsNotExist(err) || pathErr {
		pop(&directory)
		return getGitProjectRoot(directory)
	} else {

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
	}
}

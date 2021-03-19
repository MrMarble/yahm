// +build !windows

package utils

import "os"

func absolute(path string) string {
	if string(path[0]) == string(os.PathSeparator) {
		return path
	}
	return string(os.PathSeparator) + path
}

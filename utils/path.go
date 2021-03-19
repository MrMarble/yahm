// +build !windows

package utils

import "os"

func absolute(path string) string {
	return string(os.PathSeparator) + path
}

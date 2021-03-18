package utils

import "os"

func Absolute(path string) string {
	return string(os.PathSeparator) + path
}

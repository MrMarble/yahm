// +build windows

package utils

import (
	"path/filepath"
	"strings"
)

func absolute(path string) string {
	volume := filepath.VolumeName(path)
	pathList := strings.Split(path, string(filepath.Separator))
	if strings.Contains(pathList[0], volume) {
		pathList[0] = pathList[0][len(volume):]
	}
	volume = volume + string(filepath.Separator)
	return filepath.Join(volume, filepath.Join(pathList...))
}

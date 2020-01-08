package goutils

import "strings"

// SubPath returns subpath of file from path.
func SubPath(path, pattern string) string {
	if pattern == "" {
		return path
	}

	index := strings.LastIndex(path, pattern)
	if index == -1 {
		return path
	}

	return path[index+len(pattern)+1:]
}

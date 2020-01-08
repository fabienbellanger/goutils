package goutils

import (
	"testing"
)

// TestSubPath tests with correct parameters.
func TestSubPath(t *testing.T) {
	path := "/home/toto/project/dir/file.go"
	pattern := "project"
	get := SubPath(path, pattern)
	wanted := "dir/file.go"

	if get != wanted {
		t.Errorf("SubPath with many patterns - get: %v, want: %v\n", get, wanted)
	}
}

// TestSubPathManyPatterns tests with correct parameters.
func TestSubPathManyPatterns(t *testing.T) {
	path := "/home/toto/project/project/dir/file.go"
	pattern := "project"
	get := SubPath(path, pattern)
	wanted := "dir/file.go"

	if get != wanted {
		t.Errorf("SubPath with many patterns - get: %v, want: %v\n", get, wanted)
	}
}

// TestSubPathEmptyPattern tests with correct parameters.
func TestSubPathEmptyPattern(t *testing.T) {
	path := "/home/toto/project/dir/file.go"
	pattern := ""
	get := SubPath(path, pattern)
	wanted := "/home/toto/project/dir/file.go"

	if get != wanted {
		t.Errorf("SubPath with empty pattern - get: %v, want: %v\n", get, wanted)
	}
}

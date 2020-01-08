package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSubPath tests with correct parameters.
func TestSubPath(t *testing.T) {
	actual := SubPath("/home/toto/project/dir/file.go", "project")
	expected := "dir/file.go"

	assert.Equal(t, expected, actual)
}

// TestSubPathManyPatterns tests with correct parameters.
func TestSubPathManyPatterns(t *testing.T) {
	actual := SubPath("/home/toto/project/project/dir/file.go", "project")
	expected := "dir/file.go"

	assert.Equal(t, expected, actual)
}

// TestSubPathEmptyPattern tests with correct parameters.
func TestSubPathEmptyPattern(t *testing.T) {
	actual := SubPath("/home/toto/project/dir/file.go", "")
	expected := "/home/toto/project/dir/file.go"

	assert.Equal(t, expected, actual)
}

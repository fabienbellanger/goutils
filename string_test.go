package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestUcfirst tests Ucfirst function.
func TestUcfirst(t *testing.T) {
	actual := "Test chaîne Avec et sans majuscule"
	expected := Ucfirst("test chaîne Avec et sans majuscule")
	assert.Equal(t, expected, actual)

	actual = "Test chaîne Avec et sans majuscule"
	expected = Ucfirst("Test chaîne Avec et sans majuscule")
	assert.Equal(t, expected, actual)

	actual = "Été"
	expected = Ucfirst("été")
	assert.Equal(t, expected, actual)
}

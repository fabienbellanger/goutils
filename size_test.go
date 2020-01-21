package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHumanSize tests HumanSize.
func TestHumanSize(t *testing.T) {
	assert.Equal(t, HumanSize(1000), "1.0kB")
	assert.Equal(t, HumanSize(1024), "1.0kB")
	assert.Equal(t, HumanSize(1000000), "1.0MB")
	assert.Equal(t, HumanSize(1048576), "1.0MB")
	assert.Equal(t, HumanSize(2*mb), "2.0MB")
	assert.Equal(t, HumanSize(float64(3.42*gb)), "3.4GB")
	assert.Equal(t, HumanSize(float64(5.372*tb)), "5.4TB")
	assert.Equal(t, HumanSize(float64(2.22*pb)), "2.2PB")
	assert.Equal(t, HumanSize(float64(10000000000000*pb)), "10000.0YB")
}

// TestHumanSizeWithPrecision tests HumanSizeWithPrecision with custom precision.
func TestHumanSizeWithPrecision(t *testing.T) {
	assert.Equal(t, HumanSizeWithPrecision(1000, 0), "1kB")
	assert.Equal(t, HumanSizeWithPrecision(1024, 4), "1.0240kB")
	assert.Equal(t, HumanSizeWithPrecision(1000000, 3), "1.000MB")
	assert.Equal(t, HumanSizeWithPrecision(1048576, 6), "1.048576MB")
	assert.Equal(t, HumanSizeWithPrecision(2*mb, 1), "2.0MB")
}

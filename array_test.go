package goutils

import (
	"testing"
)

// TestInArray
func TestInArray(t *testing.T) {
	tInt := []int{10, 56, 23, 85}
	found, index := InArray(56, tInt)
	foundWanted, indexWanted := true, 1

	if found != foundWanted || indexWanted != indexWanted {
		t.Errorf("Ucfirst - got: %t, %d, want: %t, %d.", found, index, foundWanted, indexWanted)
	}

	tInt = []int{10, 56, 23, 85}
	found, index = InArray(589, tInt)
	foundWanted, indexWanted = true, -1

	if found == foundWanted || indexWanted > -1 {
		t.Errorf("Ucfirst - got: %t, %d, want: %t, %d.", found, index, foundWanted, indexWanted)
	}

	tString := []string{"45", "ghgh", "kl7878"}
	found, index = InArray("kl7878", tString)
	foundWanted, indexWanted = true, 2

	if found != foundWanted || index != indexWanted {
		t.Errorf("Ucfirst - got: %t, %d, want: %t, %d.", found, index, foundWanted, indexWanted)
	}

	tString = []string{"45", "ghgh", "kl7878"}
	found, index = InArray(589, tString)
	foundWanted, indexWanted = false, -1

	if found != foundWanted || index != indexWanted {
		t.Errorf("Ucfirst - got: %t, %d, want: %t, %d.", found, index, foundWanted, indexWanted)
	}
}

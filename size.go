package goutils

import (
	"fmt"
)

const (
	// Decimal
	kb = 1000
	mb = 1000 * kb
	gb = 1000 * mb
	tb = 1000 * gb
	pb = 1000 * tb
)

type unitMap map[string]int64

var (
	decimalMap   = unitMap{"k": kb, "m": mb, "g": gb, "t": tb, "p": pb}
	decimapAbbrs = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
)

// getSizeAndUnit returns the size in the correct unit.
func getSizeAndUnit(size float64, base float64, _map []string) (float64, string) {
	i := 0
	unitsLimit := len(_map) - 1
	for size >= base && i < unitsLimit {
		size = size / base
		i++
	}
	return size, _map[i]
}

// HumanSizeWithPrecision allows the size to be in any precision, instead of 1 digit precision used in HumanSize.
func HumanSizeWithPrecision(size float64, precision int) string {
	size, unit := getSizeAndUnit(size, 1000.0, decimapAbbrs)
	return fmt.Sprintf("%.*f%s", precision, size, unit)
}

// HumanSize returns a human-readable approximation of a size capped at 1 valid number (eg. "2.7MB", "796.0KB").
func HumanSize(size float64) string {
	return HumanSizeWithPrecision(size, 1)
}

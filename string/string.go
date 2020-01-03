package string

import (
	"strings"
)

// Ucfirst makes a string's first character uppercase.
func Ucfirst(s string) string {
	// Tableau de caractères Unicode pour gérér les caractères accentués
	sToUnicode := []rune(s)

	return strings.ToUpper(string(sToUnicode[0])) + string(sToUnicode[1:])
}

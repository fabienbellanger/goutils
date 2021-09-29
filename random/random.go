package random

import (
	"math/rand"
	"time"
)

const (
	AlphaNumCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890123456789"
	AlphaCharset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// GenerateBytes generates a random []byte of custom length from a charset.
func GenerateBytes(length int, charset string) []byte {
	b := make([]byte, length)
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return b
}

// GenerateString generates a random string of custom length from a charset.
func GenerateString(length int, charset string) string {
	return string(GenerateBytes(length, charset))
}

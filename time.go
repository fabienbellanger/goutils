package goutils

import "time"

// StartOfDay returns a the first time of a day.
func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// EndOfDay returns a the last time of a day.
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())
}

// Yesterday returns yesterday time.
func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

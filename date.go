package goutils

import (
	"time"
)

// SQLDateToTime converts SQL date to time.Time.
func SQLDateToTime(d string) (time.Time, error) {
	return time.Parse("2006-01-02", d)
}

// SQLDatetimeToTime converts SQL date to time.Time.
func SQLDatetimeToTime(d string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", d)
}

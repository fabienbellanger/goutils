package goutils

import (
	"testing"
	"time"
)

// TestSQLDateToTime
func TestSQLDateToTime(t *testing.T) {
	d := "2019-10-23"

	get, err := SQLDateToTime(d)
	wanted := time.Date(2019, 10, 23, 0, 0, 0, 0, time.UTC)

	if err != nil {
		t.Errorf("SQLDateToTime - got: error, want: no error\n")
	}
	if get != wanted {
		t.Errorf("SQLDateToTime - gt: %v, want: %v\n", get, wanted)
	}
}

// TestSQLDatetimeToTime
func TestSQLDatetimeToTime(t *testing.T) {
	d := "2019-10-23 12:10:56"

	get, err := SQLDatetimeToTime(d)
	wanted := time.Date(2019, 10, 23, 12, 10, 56, 0, time.UTC)
	if err != nil {
		t.Errorf("SQLDateToTime - got: error, want: no error\n")
	}
	if get != wanted {
		t.Errorf("SQLDateToTime - gt: %v, want: %v\n", get, wanted)
	}
}

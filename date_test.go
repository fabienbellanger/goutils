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

// TestTimeToSQLDate tests time to SQL date format (YYYY-MM-DD).
func TestTimeToSQLDate(t *testing.T) {
	d := time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC)
	wanted := "2020-01-05"
	get := TimeToSQLDate(d)
	if get != wanted {
		t.Errorf("TimeToSQLDate - gt: %v, want: %v\n", get, wanted)
	}
}

// TestTimeToSQLDatetime tests time to SQL datetime format (YYYY-MM-DD HH:MM:SS).
func TestTimeToSQLDatetime(t *testing.T) {
	d := time.Date(2020, 1, 5, 23, 12, 9, 0, time.UTC)
	wanted := "2020-01-05 23:12:09"
	get := TimeToSQLDatetime(d)
	if get != wanted {
		t.Errorf("TimeToSQLDatetime - gt: %v, want: %v\n", get, wanted)
	}
}

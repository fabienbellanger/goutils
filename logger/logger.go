package logger

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/fabienbellanger/goutils"
)

// DefaultWriter defines the writer for logs.
var DefaultWriter io.Writer

// Log displays custom logs.
func Log(v ...interface{}) {
	if DefaultWriter == nil {
		DefaultWriter = os.Stderr
	}
	log.SetOutput(DefaultWriter)

	// Remove logs timestamp
	log.SetFlags(0)

	log.Printf("ERROR  | %s | %+v\n", goutils.TimeToSQLDatetime(time.Now()), v)
}

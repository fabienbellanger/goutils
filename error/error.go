package error

import (
	"log"
	"os"
	"strconv"

	"github.com/logrusorgru/aurora"
)

// color enables colors in console.
var color aurora.Aurora = aurora.NewAurora(true)

// Check manages errors.
func Check(err error, exitCode int) {
	if err != nil {
		if exitCode != 0 {
			log.Printf(color.Sprintf(color.Red("Error(%v): %v"), strconv.Itoa(exitCode), err.Error()))

			os.Exit(exitCode)
		} else {
			log.Printf(color.Sprintf(color.Red("%v"), err.Error()))
		}
	}
}

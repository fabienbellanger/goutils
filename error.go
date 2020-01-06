package goutils

import (
	"log"
	"os"
	"strconv"

	"github.com/fabienbellanger/goutils/logger"
	"github.com/logrusorgru/aurora"
)

// color enables colors in console.
var color aurora.Aurora = aurora.NewAurora(true)

// CheckError manages errors.
func CheckError(err error, exitCode int) {
	if err != nil {
		if exitCode != 0 {
			log.Printf(color.Sprintf(color.Red("Error(%v): %v"), strconv.Itoa(exitCode), err.Error()))

			os.Exit(exitCode)
		} else {
			log.Printf(color.Sprintf(color.Red("%v"), err.Error()))
		}
	}
}

// CheckError2 manages errors.
func CheckError2(err error, exitCode int) {
	if err != nil {
		if exitCode != 0 {
			logger.CustomLog("Error(" + strconv.Itoa(exitCode) + "): " + err.Error())

			os.Exit(exitCode)
		} else {
			logger.CustomLog(err.Error())
		}
	}
}

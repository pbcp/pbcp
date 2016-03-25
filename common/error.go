package common

import (
	"os"
	"log"
)

var logger = log.New(os.Stderr, "", 0)

// Fatal ends execution with an error if it is not nil
func Fatal(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}
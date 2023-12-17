package loggers

import (
	"log"
	"os"
)

var InfoLogger *log.Logger

func NewInfoLogger(file *os.File) *log.Logger {
	return log.New(file, "INFO: ", log.Ldate|log.Ltime)
}

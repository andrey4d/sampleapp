package loggers

import (
	"log"
	"os"
)

var ErrorLogger *log.Logger

func NewErrorLogger(file *os.File) *log.Logger {
	return log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

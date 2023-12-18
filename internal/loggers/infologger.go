/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package loggers

import (
	"log"
	"os"
)

var InfoLogger *log.Logger

func NewInfoLogger(file *os.File) *log.Logger {
	return log.New(file, "INFO: ", log.Ldate|log.Ltime)
}

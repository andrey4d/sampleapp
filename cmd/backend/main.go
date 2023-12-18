/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package main

import (
	"os"
	"sampleapp/internal/loggers"
	"sampleapp/internal/servers"
)

func main() {
	loggerError := loggers.NewErrorLogger(os.Stderr)
	loggerInfo := loggers.NewInfoLogger(os.Stdout)

	port := "8070"
	if err := servers.Run(port, loggerInfo); err != nil {
		loggerError.Fatalf("start backend server %+v\n", err)
	}
}

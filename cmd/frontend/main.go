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

	port := "8080"
	backendURL := os.Getenv("BACKEND_URL")

	if backendURL == "" {
		loggerError.Fatalf("backend server address is not defined\n")
	}

	if err := servers.FrontRun(port, backendURL, loggerInfo); err != nil {
		loggerError.Fatalf("start frontend server %+v\n", err)
	}

}

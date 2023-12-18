/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package servers

import (
	"log"
	"net/http"
	"sampleapp/internal/handlers/back"
)

func Run(port string, loggerInfo *log.Logger) error {

	http.HandleFunc("/about", back.About(loggerInfo))
	http.HandleFunc("/health", back.Health(loggerInfo))
	http.HandleFunc("/info", back.Info(loggerInfo))
	http.HandleFunc("/", back.Hello(loggerInfo))

	loggerInfo.Printf("Run backend server on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)

	return err
}

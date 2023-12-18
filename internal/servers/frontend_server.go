/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package servers

import (
	"fmt"
	"log"
	"net/http"
	"sampleapp/internal/handlers/front"
)

func FrontRun(port, backendURL string, loggerInfo *log.Logger) error {

	http.HandleFunc("/", front.Index(backendURL, loggerInfo))
	http.HandleFunc("/info", front.Info(backendURL, loggerInfo))
	http.HandleFunc("/about", front.About(backendURL, loggerInfo))
	http.HandleFunc("/log", Log(loggerInfo))

	loggerInfo.Printf("Run frontend server on port %s\n", port)
	loggerInfo.Printf("Wait backend on %s\n", backendURL)
	err := http.ListenAndServe(":"+port, nil)

	return err
}

func Log(log *log.Logger) http.HandlerFunc {
	log.Println("logger logHandler enabled")

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Method", r.Method)
		fmt.Fprint(w, "Logger....")
	}

}

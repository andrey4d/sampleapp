/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package servers

import (
	"html/template"
	"log"
	"net/http"
	"sampleapp/internal/handlers"
	"sampleapp/internal/handlers/web"
)

func FrontRun(port, backendURL string, loggerInfo *log.Logger) error {

	handlerOptions := handlers.HandlerOptions{
		BaseTemplate: template.Must(template.ParseFiles(web.StyleTemplate, web.FooterTemplate, web.BaseTemplate)),
		Logger:       loggerInfo,
		BackendUrl:   backendURL,
	}

	http.HandleFunc("/", web.Index(handlerOptions))
	http.HandleFunc("/info", web.Info(handlerOptions))
	http.HandleFunc("/about", web.About(handlerOptions))

	loggerInfo.Printf("Run frontend server on port %s\n", port)
	loggerInfo.Printf("Wait backend on %s\n", backendURL)
	err := http.ListenAndServe(":"+port, nil)

	return err
}

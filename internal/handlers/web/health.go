/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package web

import (
	"encoding/json"
	"html/template"
	"net/http"
	"sampleapp/internal/handlers"
	"sampleapp/internal/libs"
)

type HtmlHealth struct {
	Title  string
	Health string `json:"health"`
}

func Health(options handlers.HandlerOptions) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		options.Logger.Printf("host: %s, path: %s , method: %s, agent: %s", r.RemoteAddr, r.URL.Path, r.Method, r.UserAgent())

		tmplBase, err := options.BaseTemplate.Clone()
		if err != nil {
			options.Logger.Println(err)
		}
		tmplBody := template.Must(template.ParseFiles(HealthTemplate))
		tmplBase.AddParseTree("content", tmplBody.Tree)

		htmlHealth := &HtmlHealth{
			Title: title,
		}

		out, err := libs.GetJsonFromApi(options.BackendUrl + "/health")
		if err != nil {
			options.Logger.Println(err)
			htmlHealth.Health = "Unhealthy"
		} else {
			json.Unmarshal(out, htmlHealth)
		}

		w.Header().Add("Content-Type", "text/html")
		if err := tmplBase.ExecuteTemplate(w, "base", htmlHealth); err != nil {
			options.Logger.Println(err)
		}
	}
}

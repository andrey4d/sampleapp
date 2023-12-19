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

type htmlOk struct {
	Title string
	Ok    string `json:"ok"`
}

func Index(options handlers.HandlerOptions) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		options.Logger.Printf("host: %s, path: %s , method: %s, agent: %s", r.RemoteAddr, r.URL.Path, r.Method, r.UserAgent())
		if r.URL.Path != "/" {
			ErrorHandler(w, r, http.StatusNotFound)
			return
		}
		tmplBase, err := options.BaseTemplate.Clone()
		if err != nil {
			options.Logger.Println(err)
		}
		tmplBody := template.Must(template.ParseFiles(IndexTemplate))
		tmplBase.AddParseTree("content", tmplBody.Tree)

		out, err := libs.GetJsonFromApi(options.BackendUrl)
		if err != nil {
			options.Logger.Println(err)
			return
		}

		htmlOk := &htmlOk{
			Title: title,
		}

		json.Unmarshal(out, htmlOk)
		w.Header().Add("Content-Type", "text/html")
		if err := tmplBase.ExecuteTemplate(w, "base", htmlOk); err != nil {
			options.Logger.Println(err)
		}
	}
}

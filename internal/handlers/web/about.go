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

type HtmlAbout struct {
	Title   string
	Caption string
	About   map[string]string
}

func About(options handlers.HandlerOptions) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		options.Logger.Printf("host: %s, path: %s , method: %s, agent: %s", r.RemoteAddr, r.URL.Path, r.Method, r.UserAgent())
		tmplBase, err := options.BaseTemplate.Clone()
		if err != nil {
			options.Logger.Println(err)
		}

		tmpl := template.Must(template.ParseFiles(AboutTemplate))
		tmplBase.AddParseTree("content", tmpl.Tree)

		data, err := libs.GetJsonFromApi(options.BackendUrl + "/about")
		if err != nil {
			options.Logger.Println(err)
			return
		}
		jsonAbout := new(map[string]string)
		json.Unmarshal(data, jsonAbout)

		htmlAbout := &HtmlAbout{
			Title:   title,
			Caption: "About Backend API",
			About:   *jsonAbout,
		}

		w.Header().Add("Content-Type", "text/html")
		if err := tmplBase.ExecuteTemplate(w, "base", htmlAbout); err != nil {
			options.Logger.Println(err)
		}

	}
}

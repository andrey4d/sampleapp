/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package front

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"sampleapp/internal/libs"
)

type htmlOk struct {
	Style  string
	Title  string
	Footer string
	Ok     string `json:"ok"`
}

func Index(backendURL string, log *log.Logger) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("host: %s, path: %s , method: %s, agent: %s", r.RemoteAddr, r.URL.Path, r.Method, r.UserAgent())
		if r.URL.Path != "/" {
			ErrorHandler(w, r, http.StatusNotFound)
			return
		}
		tmpl := template.Must(template.ParseFiles(indexTemplate))

		w.Header().Add("Content-Type", "text/html")

		out, err := libs.GetJsonFromApi(backendURL)
		if err != nil {
			log.Println(err)
			return
		}
		htmlOk := &htmlOk{
			Title: title,
			Style: style,
		}

		json.Unmarshal(out, htmlOk)
		tmpl.Execute(w, htmlOk)
	}
}

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

type HtmlAbout struct {
	Style   template.CSS
	Title   string
	Caption string
	About   map[string]string
}

func About(backendURL string, log *log.Logger) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("host: %s, path: %s , method: %s, agent: %s", r.RemoteAddr, r.URL.Path, r.Method, r.UserAgent())

		tmpl := template.Must(template.ParseFiles(aboutTemplate))
		footerTmpl := template.Must(template.New("footer").Parse(footer))
		tmpl.AddParseTree("footer", footerTmpl.Tree)

		w.Header().Add("Content-Type", "text/html")

		data, err := libs.GetJsonFromApi(backendURL + "/about")
		if err != nil {
			log.Println(err)
			return
		}
		jsonAbout := new(map[string]string)
		json.Unmarshal(data, jsonAbout)

		htmlAbout := &HtmlAbout{
			Title:   title,
			Style:   template.CSS(style),
			Caption: "About Backend API",
			About:   *jsonAbout,
		}
		tmpl.Execute(w, htmlAbout)

	}
}

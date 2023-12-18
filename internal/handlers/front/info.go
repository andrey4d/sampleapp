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

type HtmlInfo struct {
	Style   string
	Title   string
	Caption string
	Footer  string
	HostEnv HostEnv
}

type HostEnv struct {
	Hostname    string            `json:"hostname"`
	Environment map[string]string `json:"environment"`
}

func Info(backendURL string, log *log.Logger) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("host: %s, path: %s , method: %s, agent: %s", r.RemoteAddr, r.URL.Path, r.Method, r.UserAgent())

		tmpl := template.Must(template.ParseFiles(infoTemplate))
		w.Header().Add("Content-Type", "text/html")
		data, err := libs.GetJsonFromApi(backendURL)
		if err != nil {
			log.Println(err)
			return
		}
		hostEnv := &HostEnv{}
		json.Unmarshal(data, hostEnv)

		htmlInfo := &HtmlInfo{
			Title:   title,
			Style:   style,
			HostEnv: *hostEnv,
		}
		tmpl.Execute(w, htmlInfo)
	}
}

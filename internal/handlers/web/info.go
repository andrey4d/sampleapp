/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package web

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"sampleapp/internal/handlers"
	"sampleapp/internal/libs"
)

type HtmlInfo struct {
	Title       string
	Caption     string
	WebHostname string
	HostEnv     HostEnv
}

type HostEnv struct {
	ApiHostname string            `json:"hostname"`
	Environment map[string]string `json:"environment"`
}

func Info(options handlers.HandlerOptions) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		options.Logger.Printf("host: %s, path: %s , method: %s, agent: %s", r.RemoteAddr, r.URL.Path, r.Method, r.UserAgent())

		tmplBase, err := options.BaseTemplate.Clone()
		if err != nil {
			options.Logger.Println(err)
		}

		tmpl := template.Must(template.ParseFiles(InfoTemplate))
		tmplBase.AddParseTree("content", tmpl.Tree)

		data, err := libs.GetJsonFromApi(options.BackendUrl + "/info")
		if err != nil {
			options.Logger.Println(err)
			return
		}
		hostEnv := &HostEnv{}
		json.Unmarshal(data, hostEnv)

		webHostname, _ := os.Hostname()

		htmlInfo := &HtmlInfo{
			Title:       "API Server host environment.",
			WebHostname: webHostname,
			Caption:     "Server environment: " + hostEnv.ApiHostname,
			HostEnv:     *hostEnv,
		}
		w.Header().Add("Content-Type", "text/html")
		if err := tmplBase.ExecuteTemplate(w, "base", htmlInfo); err != nil {
			options.Logger.Println(err)
		}
	}
}

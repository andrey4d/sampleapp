/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package web

import (
	"encoding/json"
	"log"
	"net/http"
)

type health struct {
	Health string `json:"health"`
}

func Health(log *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("host: %s, path: %s , method: %s, agent: %s", r.RemoteAddr, r.URL.Path, r.Method, r.UserAgent())
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(
			health{
				Health: "healthy",
			})
	}
}

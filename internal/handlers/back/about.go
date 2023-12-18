/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package back

import (
	"encoding/json"
	"log"
	"net/http"
)

func NewAbout() *map[string]string {
	ad := map[string]string{}
	ad["/"] = "Return Ok"
	ad["/info"] = "Return OS environment"
	ad["/health"] = "Return healthy"
	ad["/about"] = "Return this info"
	return &ad
}

func About(log *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("host: %s, path: %s , method: %s, agent: %s", r.RemoteAddr, r.URL.Path, r.Method, r.UserAgent())
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(NewAbout())
	}
}

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

func Hello(log *log.Logger) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("host: %s, path: %s , method: %s, agent: %s", r.RemoteAddr, r.URL.Path, r.Method, r.UserAgent())
		if r.URL.Path != "/" {
			ErrorHandler(w, r, http.StatusNotFound)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"ok": "Ok"})
	}
}

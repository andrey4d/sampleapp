/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package back

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func Info(log *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("host: %s, path: %s , method: %s, agent: %s", r.RemoteAddr, r.URL.Path, r.Method, r.UserAgent())
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(NewHostEnv())
	}
}

type HostEnv struct {
	Hostname    string            `json:"hostname"`
	Environment map[string]string `json:"environment"`
}

func NewHostEnv() *HostEnv {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown_host"
	}

	env := make(map[string]string)
	for _, v := range os.Environ() {
		pair := strings.SplitN(v, "=", 2)
		env[pair[0]] = pair[1]
	}
	return &HostEnv{
		Hostname:    hostname,
		Environment: env,
	}
}

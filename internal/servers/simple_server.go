/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package servers

import (
	"net/http"
	"sampleapp/internal/handlers"
)

func SimpleServer() {
	http.HandleFunc("/", handlers.Hello)

	http.ListenAndServe(":8080", nil)
}

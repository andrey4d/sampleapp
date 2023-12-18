/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */

package front

import (
	"fmt"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Page not found 404")
	}
}

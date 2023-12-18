/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package libs

import (
	"io"
	"net/http"
)

func GetJsonFromApi(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	out, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return out, nil
}

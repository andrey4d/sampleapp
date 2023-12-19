/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"html/template"
	"log"
)

type HandlerOptions struct {
	BackendUrl   string
	Logger       *log.Logger
	BaseTemplate *template.Template
}

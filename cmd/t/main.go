/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package main

import (
	"bytes"
	"fmt"
	"log"
	"sampleapp/internal/handlers/web"
	"text/template"
)

func main() {
	buffer := new(bytes.Buffer)
	// tmplBase := template.Must(template.New("").ParseFiles("web/templates/"+"index.html", "web/templates/"+"base.html"))
	log.Println("111")
	tmplBase := template.Must(template.ParseFiles(web.StyleTemplate, web.FooterTemplate, web.BaseTemplate))
	tmplBody := template.Must(template.ParseFiles(web.IndexTemplate))
	// fmt.Println("++", tmplBody.Tree)
	// fmt.Println(tmplBase.Tree.Root.Nodes)
	tmplBase, _ = tmplBase.Clone()

	_, err := tmplBase.AddParseTree("content", tmplBody.Tree)
	fmt.Println("Error: ", err, tmplBody.Name())

	tmplBase.ExecuteTemplate(buffer, "base", nil)

	fmt.Println(buffer)
}

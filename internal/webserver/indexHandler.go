package webserver

import (
	"html/template"
	"log"
	"net/http"
)

func (wsc *WSConfig) indexHandler(resp http.ResponseWriter, req *http.Request) {
	// 404
	if req.URL.Path != "/" {
		tmpl := template.Must(template.ParseFiles("html/page-not-found.html"))
		tmpl.Execute(resp, nil)
		return
	}

	files := []string{
		"./html/index.html",
		"./html/head.html",
		"./html/header.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(resp, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(resp, "index", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(resp, "Internal server error", http.StatusInternalServerError)
	}

	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(resp, nil)
}

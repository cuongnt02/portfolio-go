package main

import (
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }

    files := []string {
        "./ui/html/base.html",
        "./ui/html/pages/home.html",
    }

    ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

    ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        app.serverError(w, err)
    }
}

func (app *application) notes(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/notes" {
        app.notFound(w)
        return
    }
    w.Write([]byte("Notes Page"))
}

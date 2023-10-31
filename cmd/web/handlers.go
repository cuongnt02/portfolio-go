package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    files := []string {
        "./ui/html/base.html",
        "./ui/html/pages/home.html",
    }

    ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Print(err.Error());
        http.Error(w, "Internal Server Error", 500)
        return
    }

    ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func notes(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/notes" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Notes Page"))
}

package main

import (
	"flag"
	"net/http"
)

func (app *application) routes() *http.ServeMux {
    staticDir := flag.String("staticDir", "./ui/static", "Application file server static directory")

    flag.Parse()

    mux := http.NewServeMux()
    fileServer := http.FileServer(http.Dir(*staticDir))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/about", app.about)
    mux.HandleFunc("/threads", app.createThread)
    return mux
}


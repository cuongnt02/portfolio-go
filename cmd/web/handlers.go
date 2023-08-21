package main;

import (
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("HELLO WORLD"))
}

func about(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("ABOUT PAGE"))
}

func threads(w http.ResponseWriter, r*http.Request) {
    if r.Method != "POST" {
        w.Header().Set("Allow", "POST")
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
    w.Write([]byte("THREADS"))
}

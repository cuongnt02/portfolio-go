package main


import (
    "log"
    "net/http"
)


func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/about", about)
    mux.HandleFunc("/threads", threads)
    log.Print("Starting server on: 4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}

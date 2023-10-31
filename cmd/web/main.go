package main

import (
	"log"
	"net/http"
	"os"
)


func main() {
    
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/notes", notes)
    port := ":" + os.Getenv("PORT")
    if port == ":" {
        log.Println("Using default port: 8000")
        port = ":8000"
    }

    log.Print("Starting server on: 8000")
    err := http.ListenAndServe(port, mux)
    log.Fatal(err)

}



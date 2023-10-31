package main

import (
	"log"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello From Heroku"))
}

func main() {
    
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    port := os.Getenv("PORT")
    if port == "" {
        log.Println("Using default port: 8000")
        port = "8000"
    }

    log.Print("Starting server on: 8000")
    err := http.ListenAndServe(port, mux)
    log.Fatal(err)

}



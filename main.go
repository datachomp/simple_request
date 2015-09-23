package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    log.Println("hit on %s",  r.URL.Path[1:])
    fmt.Fprintf(w, "Welcome to: %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Println("Listening on localhost:8282")
    http.ListenAndServe(":8282", nil)
}

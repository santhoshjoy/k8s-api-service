package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "OK")
    })
    http.HandleFunc("/api/v1/data", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "This is the API response from service.")
    })
    http.ListenAndServe(":8080", nil)
}

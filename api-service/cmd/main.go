package main

import (
    "log"
    "net/http"
    "api-service/internal/handler"
)

func main() {
    http.HandleFunc("/users", handler.UserHandler)
    log.Println("API Service started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
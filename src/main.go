package main

import (
    "io"
    "net/http"
)

func main() {
    http.HandleFunc("/", HandleIndex)
    http.ListenAndServe(":8082", nil)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello World")
}


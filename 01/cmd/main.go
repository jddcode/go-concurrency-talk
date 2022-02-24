package main

import (
	"github.com/jddcode/go-concurrency-talk/01/internal/work"
	"net/http"
)

func main() {
	http.HandleFunc("/", work.New().Handler)
	http.ListenAndServe(":8080", nil)
}

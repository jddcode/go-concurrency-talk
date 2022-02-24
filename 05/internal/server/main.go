package server

import (
	"github.com/jddcode/go-concurrency-talk/05/internal/work"
	"net/http"
)

func Serve(slowSysChannel, mediumSysChannel chan string) {
	http.HandleFunc("/", work.New(slowSysChannel, mediumSysChannel).Handler)
	http.ListenAndServe(":8080", nil)
}

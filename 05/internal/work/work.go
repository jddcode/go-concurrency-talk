package work

import (
	"net/http"
)

type Work interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type work struct {
	SlowSysChannel chan string
	MediumSysChannel chan string
}

func (wk work) Handler(w http.ResponseWriter, r *http.Request) {
	wk.SlowSysChannel <- "some task"
	wk.MediumSysChannel <- "some task"
	w.WriteHeader(http.StatusOK)
}

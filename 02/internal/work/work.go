package work

import (
	"net/http"
	"time"
)

type Work interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type work struct {}

func (wk work) Handler(w http.ResponseWriter, r *http.Request) {
	go wk.callSlowSys()
	go wk.callMediumSys()
	w.WriteHeader(http.StatusOK)
}

func (wk work) callSlowSys() {
	time.Sleep(time.Minute)
}

func (wk work) callMediumSys() {
	time.Sleep(time.Second * 30)
}

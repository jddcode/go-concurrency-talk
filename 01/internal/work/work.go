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
	wk.callSlowSys()
	wk.callMediumSys()
	w.WriteHeader(http.StatusOK)
}

func (w work) callSlowSys() {
	time.Sleep(time.Minute)
}

func (w work) callMediumSys() {
	time.Sleep(time.Second * 30)
}

package main

import (
	mediumSys "github.com/jddcode/go-concurrency-talk/03/internal/medium-sys"
	slowSys "github.com/jddcode/go-concurrency-talk/03/internal/slow-sys"
	"github.com/jddcode/go-concurrency-talk/03/internal/work"
	"net/http"
)

func main() {
	slowSysLib, channel := slowSys.New()
	go slowSysLib.Run()

	mediumSysLib, channelTwo := mediumSys.New()
	go mediumSysLib.Run()

	http.HandleFunc("/", work.New(channel, channelTwo).Handler)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	mediumSys "github.com/jddcode/go-concurrency-talk/05/internal/medium-sys"
	"github.com/jddcode/go-concurrency-talk/05/internal/server"
	slowSys "github.com/jddcode/go-concurrency-talk/05/internal/slow-sys"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	slowSysLib, channel := slowSys.New()
	go func() {
		wg.Add(1)
		slowSysLib.Run()
		wg.Done()
	}()

	mediumSysLib, channelTwo := mediumSys.New()
	go func() {
		wg.Add(1)
		mediumSysLib.Run()
		wg.Done()
	}()

	go server.Serve(channel, channelTwo)
}

package slowSys

import (
	"fmt"
	"time"
)

type SlowSys interface {
	Run()
}

type slowSys struct {
	Work chan string
}

func (s slowSys) Run() {
	for {
		task := <- s.Work
		fmt.Println("Sending to SlowSys:", task)
		time.Sleep(time.Minute)
	}
}

package mediumSys

import (
	"fmt"
	"time"
)

type MediumSys interface {
	Run()
}

type mediumSys struct {
	Work chan string
}

func (m mediumSys) Run() {
	for {
		task := <- m.Work
		fmt.Println("Sending to MediumSys:", task)
		time.Sleep(time.Minute)
	}
}

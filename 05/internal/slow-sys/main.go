package slowSys

func New() (SlowSys, chan string) {
	channel := make(chan string, 100)
	return slowSys{
		Work: channel,
	}, channel
}

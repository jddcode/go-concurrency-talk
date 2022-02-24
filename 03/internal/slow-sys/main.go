package slowSys

func New() (SlowSys, chan string) {
	channel := make(chan string)
	return slowSys{
		Work: channel,
	}, channel
}

package mediumSys

func New() (MediumSys, chan string) {
	channel := make(chan string, 100)
	return mediumSys{
		Work: channel,
	}, channel
}

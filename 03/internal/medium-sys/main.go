package mediumSys

func New() (MediumSys, chan string) {
	channel := make(chan string)
	return mediumSys{
		Work: channel,
	}, channel
}

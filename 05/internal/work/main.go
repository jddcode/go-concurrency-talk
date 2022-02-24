package work

func New(slowSysChannel, mediumSysChannel chan string) Work {
	return work{
		SlowSysChannel: slowSysChannel,
		MediumSysChannel: mediumSysChannel,
	}
}

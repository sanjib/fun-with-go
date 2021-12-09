package main

type sammysangAdapter struct {
	// add a field for the SammysangTV reference
	sstv *SammysangTV
}

func (ss *sammysangAdapter) turnOn() {
	ss.sstv.setOnState(true)
}

func (ss *sammysangAdapter) turnOff() {
	ss.sstv.setOnState(false)
}

func (ss *sammysangAdapter) volumeUp() int {
	ss.sstv.currentVolume++
	ss.sstv.setVolume(ss.sstv.currentVolume)
	return ss.sstv.currentVolume
}

func (ss *sammysangAdapter) volumeDown() int {
	ss.sstv.currentVolume--
	ss.sstv.setVolume(ss.sstv.currentVolume)
	return ss.sstv.currentVolume
}

func (ss *sammysangAdapter) channelUp() int {
	ss.sstv.currentChan++
	ss.sstv.setVolume(ss.sstv.currentChan)
	return ss.sstv.currentChan
}

func (ss *sammysangAdapter) channelDown() int {
	ss.sstv.currentChan--
	ss.sstv.setVolume(ss.sstv.currentChan)
	return ss.sstv.currentChan
}

func (ss *sammysangAdapter) goToChannel(ch int) {
	ss.sstv.setChannel(ch)
}

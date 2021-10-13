package models

type YoutubeChannel struct {
	allObservers []Observer
}

func (y *YoutubeChannel) GetAllObservers() []Observer {
	return y.allObservers
}

func (y *YoutubeChannel) Attach(newObserver Observer) {
	y.allObservers=append(y.allObservers, newObserver)
}

func (y *YoutubeChannel) DeAttach(deleteObserver Observer) {
	for i:=0 ; i< len(y.allObservers) ; i++ {
		if y.allObservers[i] == deleteObserver {
			y.allObservers = append(y.allObservers[:i], y.allObservers[i+1:]...)
			return
		}
	}
}

func (y *YoutubeChannel) NotifyAllObservers() {
	for i:=0 ; i< len(y.allObservers) ; i++ {
		y.allObservers[i].GetNotified()
	}
}


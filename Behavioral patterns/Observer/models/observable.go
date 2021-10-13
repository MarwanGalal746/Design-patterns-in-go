package models

type Observable interface {
	GetAllObservers() []Observer
	Attach(Observer)
	DeAttach(Observer)
	NotifyAllObservers()
}

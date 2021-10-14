package models

type State interface {
	DoAction(*Movie)
	GetName() string
}

package models

type Position interface {
	HandleRequest(Request)
	SetSuccessor(Position)
	GetSuccessor() Position
}


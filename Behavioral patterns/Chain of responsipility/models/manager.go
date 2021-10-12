package models

import "fmt"

type Manager struct {
	successor Position
}

func (m *Manager) HandleRequest(request Request) {
	if request == Task {
		fmt.Println("approved by manager")
	} else {
		m.successor.HandleRequest(request)
	}
}

func (m *Manager) SetSuccessor(position Position){
	m.successor = position
}

func (m *Manager) GetSuccessor() Position {
	return m.successor
}



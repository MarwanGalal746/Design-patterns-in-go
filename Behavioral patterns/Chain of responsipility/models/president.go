package models

import "fmt"

type President struct {
	successor Position
}

func (p *President) HandleRequest(request Request) {
	fmt.Println("Anything can be approved by president")
}

func (p *President) SetSuccessor(position Position){
	p.successor = position
}

func (p *President) GetSuccessor() Position {
	return p.successor
}

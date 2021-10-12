package models

import "fmt"

type VicePresident struct {
	successor Position
}

func (vp *VicePresident) HandleRequest(request Request) {
	if request == Vacation {
		fmt.Println("approved by vice president")
	} else {
		vp.successor.HandleRequest(request)
	}
}

func (vp *VicePresident) SetSuccessor(position Position){
	vp.successor = position
}

func (vp *VicePresident) GetSuccessor() Position {
	return vp.successor
}


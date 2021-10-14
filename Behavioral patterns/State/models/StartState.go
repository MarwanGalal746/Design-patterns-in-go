package models

import "fmt"

type StartState struct {}

func (s *StartState) DoAction(movie *Movie) {
	fmt.Println("In start state")
	movie.SetState(s)
}

func (s *StartState) GetName() string {
	return "Start State"
}
package models

import "fmt"

type StopState struct {}

func (s *StopState) DoAction(movie *Movie) {
	fmt.Println("In stop state")
	movie.SetState(s)
}

func (s *StopState) GetName() string {
	return "Stop State"
}

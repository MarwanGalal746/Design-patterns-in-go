package models

import "fmt"

type Subscriber struct {
	name string
}

func (s *Subscriber) GetNotified() {
	fmt.Println(s.name + " has been notified that a video was uploaded.")
}

func (s *Subscriber) SetName(newName string) {
	s.name = newName
}


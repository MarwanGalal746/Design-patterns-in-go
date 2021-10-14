package main

import (
	"State/models"
	"fmt"
)

func main() {
	var movie models.Movie
	var state1 models.StartState
	var state2 models.StopState
	state1.DoAction(&movie)
	fmt.Println("Current state of the movie: ", movie.GetState())
	state2.DoAction(&movie)
	fmt.Println("Current state of the movie: ", movie.GetState())
}

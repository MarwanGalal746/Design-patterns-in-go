package main

import "Observer/models"

func main() {
	var user1 models.Subscriber
	user1.SetName("Mohamed")
	var user2 models.Subscriber
	user2.SetName("Amr")
	var channel models.YoutubeChannel
	channel.Attach(&user1)
	channel.Attach(&user2)
	channel.NotifyAllObservers()
	channel.DeAttach(&user1)
	channel.NotifyAllObservers()
}

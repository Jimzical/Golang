package main

import (
	"time"
)

func main() {
	messageChannel := make(chan string)
	stopChannel := make(chan bool)
	
	go func() {
		time.Sleep(2 * time.Second)
		messageChannel <- "Hello"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		stopChannel <- true
	}() 

	select {
		case message := <-messageChannel:
			println(message)
		case <-stopChannel:
			println("Stopped")
	}

}
package main

import (
	"time"
)

func processStop(stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			println("Stopped")
			return
		default:
			println("Processing")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	stop := make(chan struct{})
	go processStop(stop)

	time.Sleep(5 * time.Second)
	close(stop)
	time.Sleep(5 * time.Second)
}
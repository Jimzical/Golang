package main

import (
	"time"
)

func worker(stop <-chan string) {
	for {
		select {
		case <-stop:
			println("Worker stopped")
			return
		default:
			println("Worker processing while main function is busy")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	stop := make(chan string)
	go worker(stop)

	time.Sleep(5 * time.Second) // Let the worker work for 5 seconds
	stop <- "asd"                 // Send a signal to stop the worker
	time.Sleep(1 * time.Second)  // Give the worker some time to stop
}
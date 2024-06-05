package main

import "time"

func main() {
	stopChannel := make(chan bool)
	go myProcess(stopChannel)
	time.Sleep(1 * time.Second)
	stopChannel <- true
	time.Sleep(1 * time.Second)
	println("Done")

}

func myProcess(stopChannel chan bool) {
	for {
		select {
		case <-stopChannel:
			println("Stopped")
			return
		default:
			println("Processing")
			time.Sleep(1 * time.Second)
		}
	}

}
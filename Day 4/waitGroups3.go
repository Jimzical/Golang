package main

import (
	"sync"
	"time"
)

func main() {
	var stopIT bool
	var wg sync.WaitGroup

	wg.Add(1)
	go process(&wg, &stopIT)

	time.Sleep(3 * time.Second)
	stopIT = true
	wg.Wait()
	print("Done")
}

func process(wg *sync.WaitGroup, stopIT *bool) {
	defer wg.Done()

	for {
		println("Processing")
		time.Sleep(1 * time.Second)

		if *stopIT {
			println("Stopped Goroutine")
			return
		}
	}
}
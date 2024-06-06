package main

import (
	"sync/atomic"
	"time"
	"fmt"
)

func main() {
	var counter int32
	
	go func() {
		for i:=0; i<3; i++ {
			atomic.AddInt32(&counter, 1)	// Increment the counter
			fmt.Printf("Incrementing counter: %d\n", atomic.LoadInt32(&counter))
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		for i:=0; i<3; i++ {
			atomic.AddInt32(&counter, -1) 		// Decrement the counter
			fmt.Printf("Decrementing counter: %d\n", atomic.LoadInt32(&counter))
			time.Sleep(time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Printf("Final counter: %d\n", atomic.LoadInt32(&counter))
}
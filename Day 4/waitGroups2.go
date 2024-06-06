
package main

import "sync"

func main() {
	wg:=sync.WaitGroup{}

	for i:=0; i<3; i++ {
		wg.Add(1)			// Increment the counter by 1
		go doSomething(&wg,i)	// Run the goroutine and decrement the counter by 1
	}
	wg.Wait()			// Check if the counter is 0, then continue
	println("Done")
}

func doSomething(wg *sync.WaitGroup, i int) {
	println("Doing something", i)
	wg.Done()
}
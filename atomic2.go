package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	var ops atomic.Uint64	// import "sync/atomic"
	var wg sync.WaitGroup	// import "sync"

	for i:=0; i<50; i++ {
		wg.Add(1)
		go func() {
			for c:=0; c<100; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	println("ops: ", ops.Load())
}

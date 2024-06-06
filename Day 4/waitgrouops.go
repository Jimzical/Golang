package main

import "sync"

func printRam(wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		println("RAM")
	}
	wg.Done()
}

func printSita(wg *sync.WaitGroup) {
	for i := 0; i < 2; i++ {
		println("Sita")
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup		// can use wg:=sync.WaitGroup{} instead

	wg.Add(2)
	
	go printRam(&wg)
	go printSita(&wg)
	
	wg.Wait()
	println("Done")
}
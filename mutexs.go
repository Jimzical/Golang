package main

import (
	"sync"
	"time"
)

var mu sync.Mutex
var balance int

func Deposit(amount int) {
	mu.Lock()			// Lock the mutex
	balance += amount
	mu.Unlock()			// Unlock the mutex to avoid deadlock
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()		// We can use defer to unlock the mutex
	return balance
}

func main() {
	println("Initial balance: ", Balance())

	go Deposit(100)
	go Deposit(200)

	time.Sleep(1 * time.Second)
	println("Final balance: ", Balance())
}
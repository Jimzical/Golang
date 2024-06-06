package main

import "time"

var balance int

func main() {
	go func() {
		Deposit(200)                    // A1
		println("Balance =", Balance()) // A2	200
	}()

	go Deposit(100) // B
	time.Sleep(3 * time.Second)
	print("Main Goroutine finished")
}

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}

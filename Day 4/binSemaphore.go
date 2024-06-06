package main

import "time"

var sema = make(chan struct{}, 1)
var balance int

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance += amount  // deposit
	<-sema             // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance       // copy balance
	<-sema             // release token
	return b           // return balance
}

func main() {
	println("Balance =", Balance()) // 0

	go Deposit(200)
	go Deposit(100)

	time.Sleep(1 * time.Second)
	println("Balance =", Balance()) // 300
}

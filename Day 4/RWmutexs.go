package main

import "sync"

var mu sync.RWMutex
var balance int

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount
}

func Balance() int {
	mu.Lock



func main() {

}
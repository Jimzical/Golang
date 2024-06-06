package main

var deposits = make(chan int) // A
var balances = make(chan int) // B

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int // A
	for {				// Infinite loop ensures it will always be ready to receive a message
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}


func main() {
	go teller() // B
	go Deposit(200)
	println("Balance =", Balance()) // 200
	go Deposit(100)
	println("Balance =", Balance()) // 300

}
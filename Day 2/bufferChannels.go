package main

func main() {
	// ch := make(chan int, 2)

	// ch <- 1
	// ch <- 2

	// println(<-ch)
	// println(<-ch)

	ss := make(chan string, 2)
	ss <- "Hello"
	// ss <- "World"
	// ss <- "Extraa"     // Uncommenting this line will cause a deadlock error

	println(<-ss)
	println(<-ss)
}
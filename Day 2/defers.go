package main

func main() {
	// defer println("world")
	// println("hello")

	for i := 0; i < 3; i++ {
		defer println(i)		// OUTPUT: 2 1 0
	}
}
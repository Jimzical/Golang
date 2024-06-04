package main

// import "time"

func display(str string) {
	for w := 0; w < 5; w++ {
		// time.Sleep(1 * time.Second)
		println(str)
	}
}
func main() {
	println("Starting")
	go display("Go Routine")
	display("Normal function")

	println("Done!")
}
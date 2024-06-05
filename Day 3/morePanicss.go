package main

func handlePanic() {
	if r := recover(); r != nil {
		println("Recovered from panic: ", r)
	}
}

func main() {
	defer handlePanic()
	panic("Panic!")
	// Nothing after this line will be exceuted
}
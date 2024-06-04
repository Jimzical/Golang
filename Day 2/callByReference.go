package main

func main() {
	var x int = 5
	r1 := callByReference(&x)
	println("x:", x, "r1:", r1)
}

func callByReference(x *int) ( y int) {
	y = 10 * (*x)
	return 
}
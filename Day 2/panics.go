package main

func div(a,b int) int {
	if b == 0 {
		panic("Division by zero")
	}
	return a/b
}

func main() {
	println(div(10,2))
	println(div(10,0))
}
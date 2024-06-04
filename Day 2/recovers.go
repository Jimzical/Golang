package main

func div(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			println("Recovered in div", r)
			result = 0
		}
	}()

	if b == 0 {
		panic("Division by zero")
	}
	return a / b
}

func main() {
	println(div(10, 2))
	println(div(10, 0))
}
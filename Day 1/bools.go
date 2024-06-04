package main

func main() {
	var b1 bool = true
	var b2 bool = false

	println(b1)
	println(b2)
	println(b1 && b2)
	println(b1 || b2)
	println(!b1)
	println(!b2)
	println(b1 == b2)
	println(b1 != b2)


	if b1 {
		println("b1 is true")
	} else {
		println("b1 is false")
	}

	if !b2 {
		println("b2 is false")
	} else {
		println("b2 is true")
	}
}
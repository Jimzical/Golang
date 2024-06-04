package main

import "strconv"

func main() {
	// s:="helloworld"
	x := "123"
	y, err := strconv.Atoi(x)
	println(y,err)

	z := strconv.Itoa(y)
	println(z)

	x = "123.45"
	a , err := strconv.ParseFloat(x)
	println(a) 
}
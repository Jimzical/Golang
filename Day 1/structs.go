package main

import "fmt"

type Vertex struct {
	X      int
	Y      int
	ID     int
	Name   string
	Salary int
}

func main() {
	var v Vertex
	v.X = 12
	v.Y = 42
	fmt.Println(v)
	println(v.X)
	println(v.Y)
}
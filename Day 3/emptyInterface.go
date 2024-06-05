package main

import "fmt"

func PrintAny(value interface{}) {
	fmt.Println(value)

}


func main() {
	PrintAny(32)
	PrintAny("Hello")
	PrintAny(3.14)
}
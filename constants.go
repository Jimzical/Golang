package main

import "fmt"

const (
	sunday = iota
	monday = 2
	tuesday = 5
)

func main() {
	fmt.Println(sunday, monday, tuesday)
}
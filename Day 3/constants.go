package main

import "fmt"

const (
	sunday = iota
	monday 
	tuesday 
	wednesday  = iota
	thursday = iota
	friday
)

func main() {
	fmt.Println(sunday, monday, tuesday, wednesday, thursday, friday)
}
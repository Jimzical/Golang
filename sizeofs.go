package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Integers
	fmt.Println(unsafe.Sizeof(int(10)))	// 8
	fmt.Println(unsafe.Sizeof(uint(10)))	// 8
	fmt.Println(unsafe.Sizeof(int8(10)))	// 1
	fmt.Println(unsafe.Sizeof(int16(10)))	// 2
	fmt.Println(unsafe.Sizeof(int32(10)))	// 4
	fmt.Println(unsafe.Sizeof(int64(10)))	// 8

	// Bool
	fmt.Println(unsafe.Sizeof(bool(true)))	// 1


	var i = 10
	point := &i
    fmt.Println("Sizeof(j) :", unsafe.Sizeof(point)) // "8"
    fmt.Println("Sizeof(float32(0)) :", unsafe.Sizeof(float32(0))) // "4"
    fmt.Println("Sizeof(float64(0)) :", unsafe.Sizeof(float64(0))) // "8"

	// Empty struct
    fmt.Println("Sizeof(struct{}{}) :", unsafe.Sizeof(struct{}{})) // "0"

	// Interface
    var k interface{}
    fmt.Println("Sizeof(interface{}) :", unsafe.Sizeof(k)) // "16"

	// String
    s := "Hellosdfdssd"
    fmt.Println("Sizeof('Hello') :", unsafe.Sizeof(s)) // "0"

	fmt.Println("Sizeof([10]int{}) :", unsafe.Sizeof(arr)) // "80"	
}
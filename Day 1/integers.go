package main

import "fmt"

func main() {
	var i8 int8 = -1
	var i16 int16 = -1
	var i32 int32 = -1
	var i64 int64 = -1
	
	fmt.Printf("Signed Integers %d %d %d %d\n",i8,i16,i32,i64)

	// this gives overflow error
	// var ui8 uint8 = -1 
	// var ui16 uint16 = -1
	// var ui32 uint32 = -1
	// var ui64 uint64 = -1
	
	// fmt.Printf("Unsigned Integers %d %d %d %d\n",ui8,ui16,ui32,ui64)
	
	var ui8 uint8 = -1 * -1
	fmt.Println(ui8)   	// this gives 1

	var i int8
	var PI = 22.0/7.0
	i = int8(PI)
	fmt.Printf("%f %d\n", PI, i)

	// Convertion isses

	var apples int32 = 1
	var oranges int64 = 2
	// var ans int = apples + oranges		// Compile error
	var ans int = int(apples) + int(oranges)
	fmt.Println(ans)


}
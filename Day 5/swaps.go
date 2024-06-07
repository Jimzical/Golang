package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	swapper := reflect.Swapper(s)
	swapper(0, len(s)-1)
	fmt.Println(s) // 5 and 1 are swapped

	theList := []int{1, 2, 3, 4, 5}
	swap := reflect.Swapper(theList)
	fmt.Printf("Before Reverse :%v\n", theList)

	// Reversing a slice using Swapper() function
	for i := 0; i < len(theList)/2; i++ {
		pos := len(theList) - i - 1
		swap(i, pos)
	}
	fmt.Printf("After Reverse :%v\n", theList)

}
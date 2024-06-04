package main

import (
	"fmt"
	"strings"
)

func joinStr(strs ...string) string {
	return strings.Join(strs, "")
}

func main() {
	str1 := joinStr("Hello", " ", "World", "!")
	fmt.Println(str1)
	str2 := joinStr("Hello", " ", "World", "!", " ", "How", " ", "are", " ", "you", "?")
	fmt.Println(str2)

	println(sum(1,2,3,4,5))
	println(sum(1,2,3))
	// using an array
	nums := []int{1,2,3,4,5}
	println(sum(nums...))
}

func sum(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

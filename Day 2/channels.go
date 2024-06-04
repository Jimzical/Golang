package main

import "fmt"

func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	ch := make(chan int)
	go sum(s[0:3], ch)	 // 17
	go sum(s[3:6], ch)	// -5
	go sum(s[3:4], ch)	// -9

	
	fmt.Println(len(ch))
	x := <-ch		// -9
	y := <-ch		// -5
	z := <-ch		// 17
	
	fmt.Println(len(ch))
	println(x, y, z)	// -9 -5 17
}
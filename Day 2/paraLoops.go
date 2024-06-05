package main

import (
	// "context"
	// "sync"
	// "os"
	// "reflect"
)
func main() {
	a := []int{1, 2, 3, 4, 5}

	// normal
	for i := 0; i < len(a); i++ {
		println(a[i])
	}

	// parallel
	For(0, len(a), 1, func(i int) {
		println(a[i])
	})

}

func For(start, end, step int, f func(int)) {
	for i := start; i < end; i += step {
		go f(i)
	}
}
package main

import "runtime"
func main() {
	curVal := runtime.GOMAXPROCS(0)
	println(curVal)

	maxCore := runtime.NumCPU()
	println(maxCore)
	new := maxCore - 2
	runtime.GOMAXPROCS(new)
	curVal = runtime.GOMAXPROCS(0)
	println(curVal)
}
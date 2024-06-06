package main

func myChannelTest(i int,ch chan int) {
    ch <- i
}

func main() {
	ch := make(chan int)
	go myChannelTest(1,ch)
	x := <-ch		// -9
	y := <-ch		// -5
	go myChannelTest(3,ch)
	go myChannelTest(2,ch)
	z := <-ch		// 17
	
	// go func() {
	// 	ch <- 1
	// 	ch <- 2
	// 	ch <- 3
	// }()
	

	println(x, y, z)	// -9 -5 17
}
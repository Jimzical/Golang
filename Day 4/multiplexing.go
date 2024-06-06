package main

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case v := <-ch1:
				println("From ch1:", v)
			case v := <-ch2:
				println("From ch2:", v)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case ch1 <- i:
		case ch2 <- i:
		}
	}
}
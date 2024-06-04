package main

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
	printer(naturals)
	println()
	summer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		println(x)
	}
}

func summer( in <-chan int) {
	sum := 0
	for x := range in {
		sum += x
		println(x)
	}
	print(sum)
}
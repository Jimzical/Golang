package main

func main() {
	var data int
	go func() {
		data++
	}
	
	if data == 0 {
		println("Value is 0.")
	}
}
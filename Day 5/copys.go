package main

import (
	"fmt"
	"reflect"
)

func main() {
	destination := reflect.ValueOf([]string{"A", "B", "C"})
	source := reflect.ValueOf([]string{"D", "E", "F"})

	fmt.Println("destination: ", destination)
	
	counter := reflect.Copy(destination, source)
	fmt.Println("Count =", counter)
	fmt.Println("source: ", source)
	fmt.Println("destination: ", destination)
}

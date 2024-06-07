package main

import (
	"fmt"
	"reflect"
)

func main() {
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)
	fmt.Printf("type: %v\n", v.Type())
	fmt.Printf("kind: %v\n", v.Kind())

}
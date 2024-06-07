package main

import (
	"fmt"
	"reflect"
)

func main() {
	v1:= 3.4
	fmt.Println(reflect.ValueOf(v1))
	v2:= "Hello"
	fmt.Println(reflect.ValueOf(v2))
	v3:= true
	fmt.Println(reflect.ValueOf(v3))
	v4:= 10
	fmt.Println(reflect.ValueOf(v4))
	v5:= []int{1,2,3}
	fmt.Println(reflect.ValueOf(v5))
	v6:= [5]int{1,2,3,4,5}
	fmt.Println(reflect.ValueOf(v6))
	v7:= map[string]int{"a":1, "b":2}
	fmt.Println(reflect.ValueOf(v7))
}
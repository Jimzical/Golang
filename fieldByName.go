package main

import (
	"fmt"
	"reflect"
)

type struct_T struct {
	val1 int
	val2 string
}

func main() {
	s := struct_T{23, "hello"}
	v1 := reflect.ValueOf(s)
	v2 := reflect.ValueOf(&s).Elem()

	fmt.Println("Field 0:", v1.FieldByName("val1")) // 23
	fmt.Println("Field 1:", v2.FieldByName("val2")) // hello
}
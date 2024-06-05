package main


func main() {
	var i interface{} = 42
	s,ok:= i.(string)
	println(s,ok)

	// var i interface{} = "hello"
	// s,ok:= i.(string)
	// println(s,ok)

	// f,ok:= i.(float64)
	// println(f,ok)
}
package main

func main() {
	var str string = "helloworld"

	println(str[:5])  	// hello
	println(str[7:])	// rld
	println(str[:])		// helloworld
	println(str[2:4])	// ll

	s:="hello"
	s+=" world"

	println(s)

	// str[0] = "a"			// string immutable so cannot change
}
package main

func main() {
	type MyInt int
	var i int
	var j MyInt

	i = int(j)
	print(i, j)
}
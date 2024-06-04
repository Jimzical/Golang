package main
import ("fmt")


// outside the fucntion
var str1 string = "this is a string"	// type is explicit

func main() {

	// This can be done outsite the function as well and can be used to make global variables. 
	var str2 = "this is infered type"	// go infers the type
	fmt.Println(str1,str2)  // NOTE here it automatially adds a space btw the variables

	// This can only be used inside fucntiosn
	x := 2 // go will infer this as well
	println(x)

	// multiple assignments here of same type
	var a,b,c,d int = 1,2,5,7
	fmt.Println(a,b,c,d)

	var (
		q int // defalut 0
		w int = 1
		e string = "hi"
		r string // default is " "
		t bool = true
		y bool   // default false
		u float32 = 1.2
		i float32 // default 0
	)

	fmt.Println(q,w,e,r,t,y,u,i)

}


/*
OUTPUT
this is a string this is infered type
2
1 2 5 7
0 1 hi  true false 1.2 0
*/



// types of variables

// int
// float32
// string
// bool

// types of cases
// Pascal: Thisvariable
// Camel: ThisVariable
// Snake: this_variable


#GO Notes

# Basic
To create a basic program
```go
package main
import ("fmt")


func main() {
	fmt.Println("Hello, World!")
}

```
> OUTPUT: Hello, World!

## Function creation

- Ensure code does not have "{" at the start of the line. it will cause error
```go
func main() 
{			<--- ERROR
	fmt.Println("Hello, World!")
}

```
> OUTPUT: Error

# Print Statemnts
- Done using the fmt package

## Example

```go
package main
import ("fmt")

func main() {
	var string1 string = "this is from a variable"
	
	// Three types of prints
	fmt.Print("hi this is print")	// No space is added after this
	fmt.Println("println")			// adds newline at the end

	fmt.Printf("%v\n", string1)		// prints the value of the variable
	fmt.Printf("type is %T \n", string1)
	fmt.Printf("type is %T",2)
}	
```
> OUTPUT: hi this is printprintln
this is from a variable
type is string
type is int

# Variables
- Check the code for all info
```go
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
```

# IF comditions

- Same as C (almost)
- "{" cannot be at new line
- "else" cannot be on newline, needs to be "} else {" where } is from a prior if statemet  


## Example
```go
package main

func main() {
	// checking if number positive or negative
	var x int = -7

	if x>0 {
		println("positve")
	} else if x<0 {
		println("negative")
	} else {
		println("zero")
	}

}
```

# Constants
- Their values cant be changed. Use the `const` keyword

```go
package main
import ("fmt")


const PI = 3.14
func main() {
	fmt.Println(PI)
	// PI = 3
	// fmt.Println(PI)	// THIS GIVES ERROR as we can edit consts
	
}
```
> OUTPUT: 3.14

# Operations		// TODO check for missing items

## Arithemetic Operations

` * / % >> << & &^ + - ^ \`

- +,-,*,/: int, float and complex numbers
- %: `only applies to integers`
- sign of remainder always the same as dividend sign
- Follows BODMAS rules

Example 
`-5%3 and -5%-3 equal -2 `
>> does not matter what is on the RHS of `%`. only LHS sign taken

## Comparison Operations

` == != < <= > >= `

- only for basic types
- == and != can be used for strings
- strings are compared lexicographically
- strings are case sensitive
- Examples are in the Boolean topic

## Bitwise Binary Orperations 

- &     --> Bitwise And
- |     --> Bitwise Or
- ^     -->
- &^    --> Bitwise Clear
- <<    -->
- >>    -->




# Integers

Four types of ints --> int8, int16, int32, int 64

- Each int is a differnt from each other so int and int16 cannot be directly converted.
- float to int conversion removes the fractional part
- signed no are 2's complement
- higerorder bits are saved for the sign
- range of value is --> -2^n-1 to ((2^n-1) - 1)
- rnage of bits for non-negative is 0 to 2^n-1

## Signed Integers

```go
package main

import "fmt"

func main() {
	var i8 int8 = -1
	var i16 int16 = -1
	var i32 int32 = -1
	var i64 int64 = -1
	
	fmt.Printf("Signed Integers %d %d %d %d\n",i8,i16,i32,i64)
}

```
> OUTPUT -1,-1,-1,-1

## Unsigned Integers
```go
package main

import "fmt"

func main() {
	// this gives overflow error
	var ui8 uint8 = -1
	var ui16 uint16 = -1
	var ui32 uint32 = -1
	var ui64 uint64 = -1
	fmt.Printf("Unsigned Integers %d %d %d %d\n",ui8,ui16,ui32,ui64)
}

```
> OUTPUT: Overflow Error

```go
	var ui8 uint8 = -1 * -1
	fmt.Println(ui8)   	// this gives 1
```
> OUTPUT: 1

## Type casting
```go
	var i int8
	var PI = 22.0/7.0
	i = int8(PI)
	fmt.Printf("%f %d\n", PI, i)

```
>> OUTPUT 3.142857 3

```go
	var i int8
	var PI = 22/7    // removing the 0 from 22.0 --> This makes PI an INT
	i = int8(PI)
	fmt.Printf("%f %d\n", PI, i)

```
> OUTPUT %!f(int=3) 3

## Type Casting Problems
```go
	var apples int32 = 1
	var oranges int64 = 2
	var ans int = apples + oranges		// Compile error
	fmt.Println(ans)
```
>> OUTPUT: Error

```go
	var apples int32 = 1
	var oranges int64 = 2
	var ans int = int(apples) + int(oranges)
	fmt.Println(ans)
```
> OUTPUT: 3

# Floating Type Numbers

- Two types--> float32 and float64
- Limits found in Math Package.
- Limits for float32 are 1.4e-45 to 3.4e38
- Limits for float64 are 4.9e-324 to 1.8e308
- Float32 has 6 decimal digits of precision and Float64 has 15 decimal digits of precision
Eg:
```go
	var f1 float32 = 3.14e2
	var f2 float64 = 3.14e7
	fmt.Printf("%f %f\n", f1, f2)
```

## Max Values are in math package
```go
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
```
> OUTPUT: 340282346638528859811704183484516925440.000000
179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368.000000

## Examples
```go
	var f32 float32 = 16777216
	var f64 float64 = 16777216 * 2
	var f float64 = 8
	const Avagadro = 6.02214129e23

	fmt.Printf("%f %f %f %f\n",f32,f64,Avagadro,f)

```
> OUTPUT: 16777216.000000 33554432.000000 602214129000000000000000.000000 8.000000

## Defaults
```go
  var z float64
  fmt.Println(z,-z,1/z,-1/z,z/z)
```
> OUTPUT: 0 -0 +Inf -Inf NaN



# Complex Numbers

- Two types --> complex64 and complex128
- Complex numbers are represented as a pair of float32 or float64
- Built-in fucntions to get real and imaginary parts and to create the numbers
- Equality is checked by comparing the real and imaginary parts

## Example

```go
package main

import "fmt"

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)

	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}
```
> OUTPUT: </br>
> (-5+10i)</br>
> -5</br>
> 10</br>

# Boolean 
- It can be `true` or `false`
- conditions for if and else
- `&&` and `||` are used for and and or. These comparison operatros result in booleans
- no implicit conversion to 1 or 0
- if answer evaluated through LHS then RHS not evaluated


## Basic usage
```go
	var b1 bool = true
	var b2 bool = false

	println(b1) 		// true
	println(b2)			// false
	println(b1 && b2)	// flase
	println(b1 || b2)	// true
	println(!b1)		// false
	println(!b2)		// true
	println(b1 == b2)	// false
	println(b1 != b2)	// true

```
> OUTPUT:</br>
true</br>
false</br>
false</br>
true</br>
false</br>
true</br>
false</br>
true</br>


## Using in IF conditions
```go

	if b1 {
		println("b1 is true")
	} else {
		println("b1 is false")
	}

	if !b2 {
		println("b2 is false")
	} else {
		println("b2 is true")
	}
```
> OUTPUT:</br>
> b1 is true</br>
> b2 is false</br>

# Strings

- `Immutable string of bytes`. We cannot change it afterwordfs. `str[0] = "a" ` gives error
- May contain arbitrary data including bytes with value 0, but usually have human readable text
- Text strings are conventionally interpretted as UTF-8
- Built-in len fucntion returns number of bytes
- Immutability: safe for two copies to share same underlyimg memory making it cheap to copy (not sure what this is on about)
  

- s[i] slicing can be done to take data from a string

Example: s:="hi"

## Substring

- Substring operations `s[i:j]`  creats a new string from `i` till but not including `j`
- Either or both can be ommited
- Defaults are 0 and len(s)

Example
```go
package main

func main() {
	var str string = "helloworld"

	println(str[:5])  	// hello
	println(str[7:])	// rld
	println(str[:])		// helloworld
	println(str[2:4])	// ll

}

```
> OUTPUT:</br> hello </br>
rld </br>
helloworld </br>
ll

Example Assigning
```go
str:= "hello, world"
h = str[:5]		// {hello}	h points to first index of str
w = str[7:]		// {world}	w points to 7th index
```

Error
```go
str[0] = "a"		// gives error as string immutable
```
> OUTPUT: Error since strings are immutable
## Appending

- We can append using `+` operatior
Example: connecting
```go
println("good"+"bye")  // goodbye
```
Example: appending
```go
s:="hello"
s+=" world"

println(s)		// hello world
```


## String Integer Conversions
- Int to Str ==> strconv.Itoa()
- Str to Int ==> strconv.Atoi()
- Str to Float ==> strconv.ParseFloat

Example
```go
package main

import "strconv"

func main() {
	// s:="helloworld"
	x := "123"
	y, err := strconv.Atoi(x)
	println(y,err)
}
```
> OUTPUT: 123 (0x0,0x0)

# For Loops
To create for loops
- Ensure its `i:=0` and not `i=0` unless `i` has been explicitly stated somewhere prior
  
Example
```go
package main

func main() {
	sum := 0
	for i:=0; i<10; i++ {
		sum += 1
	}
	println(sum)
}

```
> OUTPUT: 10

Example where `i` is not required so it is `ommited`
```go
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	
	println(sum)
```
> OUTPUT: 1024


# Pointers
- Using pointers that point to a memory. It holds the memory address of a value
- `var p *int`
- `&` generates a pointer to its operand
- `*` denotes the pointers underlying value
- This is known as `dereferencing` or `indirecting`
Example
```go
package main

func main() {
	i := 42
	p := &i
	println(*p, p)	// *p is being read here
	*p = 21 		// *p is being set here
	// This is known as dereferencing or indirecting
}
```
> OUTPUT: 42 0xc000059f38


# Structs
- Collection of fields
- User-defined datatype
- accessed using dot
- Groups 0 or more values together. 
- Each value is called a field
- `If all fields of a struct are comparable. the struct is also comparable using == or !=`

```go
package main

type Vertex struct {
	X int
	Y int
}

func main() {
	var v Vertex
	v.X = 12
	v.Y = 42
	println(v.X)		// 12
	println(v.Y)		// 42
	// println(v)		// gives error, struct cant be printed
	fmt.Println(v)		// works; {12 42 0  0}


}
```
> OUTPUT: </br> {12 42 0  0} </br> 12  </br>42

# Arrays
- Colletion of similar items
- can be directly compared usign == or != to see if corresponding values are same
```go
var q [3]int = [3]int{1,2,3}
var r [3]int = [3]int{1,2}
println(r[2])	// 0 is the output
```

```go
q:= [...]int{1,2,3} // automatically sets that numebrs
printf("%T",q)	// [3]int

r:=[...]int{99:-1}		// sets a list with 99 0s and -1 at the end 
```

# Slice
- Varibale lenght sequence of elemetns
- has 3 componenss - pointer (to base location), lenght (no of eleemnts), capacity (lenght cannot exceed this)
- buitl in fucntions len and cap
- dynamically sized
- selects half-open range inlcues the 1st ekemetn and exclude the last
- does not store any data, just describe a section of array
- other slices with same underlying array can se
```go
low:=1
high:=4
a[low:high]
```

example
```go
prime := [...]int{1,2,3,4}
var s []int = prime[1:3]
println(s)
```

# Maps		// TODO HAVE TO ADD LATER
- Collection of key-value pairs


# Functions	// TODO HAVE TO ADD LATER

format
```go
func name(para_list) (return_list) {
	// code
}
```

# Return values		//TODO have to add missing items

- Multiple return values can be returned
- Named return values can be used to return values
- A return statement without arguments returns the named return values. This is known as a "naked"/"bare" return
Example
```go
package main

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	println(a, b)
}
```
> OUTPUT: world hello

## Bare Returns
- Used to return the named return values
Example
```go
package main

func split(sum int) (x, y int) {		// Also we have named return values here
	x = sum * 4 / 9
	y = sum - x
	return			// We can see there are no arguments here. This is a naked return
}

func main() {
	println(split(17))
}
```
> OUTPUT: 7 10
## Varidic Functions
- Functions that can take any number of arguments
- `...` is used to denote this
- The type of the varidic function is a slice
- Printf is an example of a varidic function
- Mainly used in situatuons where string formatting is needed
   
Example
```go
func sum(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	println(sum(1,2,3,4,5))
	println(sum(1,2,3))

	// using an array
	nums := []int{1,2,3,4}
	println(sum(nums...))
}
```
> OUTPUT: 15 6 10

# Deferred Fucntions
- Prefixed with `defer` keyword
- A deferred function is executed after the surrounding function returns
- The arguments are evaluated immediately but the function call is deferred
- Useful for situtions like closing a file after opening it or database connections
- Excuted in reverse order of which they were deferred
- They are pushed onto a stack and executed in LIFO order

Example
```go
package main

func main() {
	defer println("world")
	println("hello")
}
```
`OUTPUT: hello \n world

Example: LIFO
```go
	for i := 0; i < 5; i++ {
		defer println(i)
	}
```
> OUTPUT: 4 3 2 1 0

# Panic Function
- When Go program detects a problem during runtime and, it can panic
- Stops the normal execution of a function and all deferred functions are executed and the program crashes with a log message
- The log message has a error message for each gorooutine
- This can be used to debug

- This can be triggered using built-in `panic` function
- when panic occurs, it will start unwinding the stack of the goroutine, meaning it will exit all functions that were called till now
- When Top-level function is reached, the program will crash (i think)
Example
```go
package main

func div(a,b int) int {
	if b == 0 {
		panic("Division by zero")
	}
	return a/b
}

func main() {
	println(div(10,2))
	println(div(10,0))
}
```
> OUTPUT  
5 </br>
panic: Division by zero </br>
goroutine 1 [running]: </br>
main.div(...) </br>
        c:/Personal/Uni/CS/Golang/panics.go:5 </br>
main.main() </br>
        c:/Personal/Uni/CS/Golang/panics.go:12 +0x3e </br>
exit status 2 </br>

# Recover Function
- Used to regain control of a panicking goroutine
- Only useful inside deferred functions
- Panicking fucntion does not  continue where it left off, but returns normally
- If called at any other time, there is no effect
- If the goroutine is not panicking, it returns nil
- Recovering indescriminately is not a good idea as the state of the program variables is unknown after this
- Should not try to recover out of errors from a package fucntion

Example
```go
package main

func div(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			println("Recovered in div", r)
			result = 0
		}
	}()

	if b == 0 {
		panic("Division by zero")
	}
	return a / b
}

func main() {
	println(div(10, 2))
	println(div(10, 0))
}
```
> OUTPUT </br>
> 5 </br>
> Recovered in div (0x3e7040,0x40caa8) </br>
> 0 </br>

# Call by Reference

- When the memory address of a variable is passed to a function, it is known as call by reference
- This is done using pointers

Example
```go
package main

func main() {
	var x int = 5
	r1 := callByReference(&x)
	println("x:", x, "r1:", r1)		// both are 50 as x is passed by reference
}

func callByReference(x *int) (int) {
	*x = 10 * *x
	return *x
}
```
> OUTPUT: x: 50 r1: 50

# Method Declaration

- A method is a function with a special receiver argument
- The receiver appears in its own argument list between the func keyword and the method name
- Methods can be declared for any type that is declared in the same package

Example
- In this example, the Abs method has a receiver of type Vertex named v
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```
> OUTPUT: 5

Example: the class one
```go
package main

type rect struct {
	width, height int
}

func main() {
	r := rect{width: 10, height: 5}
	println("area: ", r.area())
	println("perim: ", r.perim())

	rp := &r
	println("area: ", rp.area())
	println("perim: ", rp.perim())
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

```
> OUTPUT: </br>
> area: 50 </br>
> perim: 30 </br>
> area: 50 </br>
> perim: 30 </br>
</br>
---------
Unit 1 End
---------

# Concurancy
- Making progress on more than one task, seemingly at the same time
- Context Switching is when OS saves the context of a process so it can pause it and come back to it later

# Parallisim
- Run mutiple tasks at the same time

## Goroutines

- Concurrency in Go is done through goroutines
- Similar to a thread
- It is an independent fucntion that runs on some seperate lightweight thread managed by the Go runtime
- When a program starts a goroutine calls the main function, which we call the main goroutine
- New goroutines are created using the `go` keyword followed by a function invocation
- This causes it to be called in a new goroutine



Example- We run this with the `go` keyword and without the `go` keyword, we can see that the go routine runs in the background and prints as soon as it can while Normal function in its own way making it interleaved with a pattern
```go
package main

import "time"

func display(str string) {
	for w := 0; w < 5; w++ {
		time.Sleep(1 * time.Second)
		println(str)
	}
}
func main() {
	go display("Go Routine")
	display("Normal function")

}
```
> OUTPUT: </br>
> Go Routine </br>
> Normal function </br>
> Normal function </br>
> Go Routine </br>
> Go Routine </br>
> Normal function </br>
> Normal function </br>
> Go Routine </br>
> Go Routine </br>
> Normal function </br>


Example: With both as goroutines we get random interleaved output
```go
func main() {
	go display("Go Routine")
	go display("Normal function")
}
```
> OUTPUT: </br>
Go Routine </br>
Normal function </br>
Normal function </br>
Go Routine </br>
Go Routine </br>
Normal function </br>
Normal function </br>
Go Routine </br>

- Sometimes the main function exits before the goroutines finish. This is because the main function does not wait for the goroutines to finish
- This is especially true if we remove the sleep function from the goroutine
- This is because goroutines return immediately and the main function exits before the goroutines can finish

Example: No sleep function
```go
package main

func display(str string) {
	for w := 0; w < 5; w++ {
		println(str)		// REMOVED the sleep fucntion
	}
}
func main() {
	println("Starting")
	go display("Go Routine")
	display("Normal function")

	println("Done!")
}

```
> OUTPUT: </br>
Starting </br>
Normal function </br>
Normal function </br>
Normal function </br>
Normal function </br>
Normal function </br>
Done! </br>


### Advantages of Goroutines
- Goroutines are cheap. They are only a few kb in stack size
- They are multiplexed onto multiple OS threads
- They are not threads. They are not managed by the OS, They are managed by the Go runtime

# Channels

- Channel of communication between goroutines
- Each channel is a conduit for values of a particular type, called the channel's element type
- Type of a channel is `chan` followed by the type of the elements. Eg `chan int`
- Channels are reference types, as in they are references to the data structure created by the make function
- They are built using the inbuilt `make` function
- Zero value of a channel is `nil`
- Channels can be buffered or unbuffered
- When we copy a channel or pass it as an argument to a function, we are copying a reference to the same channel, so the copy and original refer to the same data structure
- Two channels of the same type may be compared using `==`. The comparison is true if both are references to the same channel data structure. It may also be compared to 
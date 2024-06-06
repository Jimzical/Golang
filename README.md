# Go programming language

These are the notes im making for the `Go programming language` 


# Table of Contents
- [Go programming language](#go-programming-language)
- [Table of Contents](#table-of-contents)
- [Basic](#basic)
	- [Commenting](#commenting)
	- [Function creation](#function-creation)
- [Print Statemnts](#print-statemnts)
		- [Example](#example)
	- [Printf](#printf)
		- [Table](#table)
- [Declarations](#declarations)
	- [Predeclared Names](#predeclared-names)
- [Variables](#variables)
		- [Syntax](#syntax)
		- [Example](#example-1)
		- [Example: Using Global variables/ Package level variables](#example-using-global-variables-package-level-variables)
- [IF conditions](#if-conditions)
		- [Example](#example-2)
- [Constants](#constants)
		- [Syntax](#syntax-1)
		- [Example: using a constant](#example-using-a-constant)
- [Operations](#operations)
	- [Arithemetic Operations](#arithemetic-operations)
		- [Example](#example-3)
	- [Comparison Operations](#comparison-operations)
	- [Bitwise Binary Orperations](#bitwise-binary-orperations)
	- [Logical Operations](#logical-operations)
	- [Assignment Operations](#assignment-operations)
	- [Other Operations](#other-operations)
- [Integers](#integers)
		- [Syntax](#syntax-2)
	- [Signed Integers](#signed-integers)
	- [Unsigned Integers](#unsigned-integers)
	- [Type casting](#type-casting)
	- [Type Casting Problems](#type-casting-problems)
- [Floating Type Numbers](#floating-type-numbers)
	- [Max Values are in math package](#max-values-are-in-math-package)
		- [Examples](#examples)
	- [Defaults](#defaults)
- [Complex Numbers](#complex-numbers)
		- [Syntax](#syntax-3)
		- [Example](#example-4)
- [Boolean](#boolean)
	- [Basic usage](#basic-usage)
	- [Using in IF conditions](#using-in-if-conditions)
- [Strings](#strings)
	- [Substring](#substring)
		- [Example](#example-5)
		- [Example Assigning](#example-assigning)
	- [Appending](#appending)
		- [Example: connecting](#example-connecting)
		- [Example: appending](#example-appending)
	- [String Integer Conversions](#string-integer-conversions)
		- [Example](#example-6)
- [For Loops](#for-loops)
		- [Example](#example-7)
		- [Example where `i` is not required so it is `ommited`](#example-where-i-is-not-required-so-it-is-ommited)
		- [Example where `for` is used for lists](#example-where-for-is-used-for-lists)
		- [Example infinite loop](#example-infinite-loop)
- [Pointers](#pointers)
		- [Example](#example-8)
- [Structs](#structs)
- [Arrays](#arrays)
- [Slice](#slice)
		- [Example](#example-9)
- [Appends](#appends)
		- [Example](#example-10)
- [Maps](#maps)
		- [Syntax](#syntax-4)
		- [Example](#example-11)
- [Default Values](#default-values)
- [Functions](#functions)
		- [Syntax](#syntax-5)
		- [Example: Basic](#example-basic)
		- [Example: Multiple return values](#example-multiple-return-values)
		- [Example: Named return values](#example-named-return-values)
	- [Passing by value](#passing-by-value)
		- [Example: Passing by value](#example-passing-by-value)
	- [Passing by reference](#passing-by-reference)
		- [Example: Passing by reference](#example-passing-by-reference)
	- [Annonyms Functions](#annonyms-functions)
- [Return values](#return-values)
		- [Example](#example-12)
	- [Bare Returns](#bare-returns)
		- [Example](#example-13)
	- [Varidic Functions](#varidic-functions)
		- [Example](#example-14)
- [Deferred Fucntions](#deferred-fucntions)
		- [Syntax](#syntax-6)
		- [Example](#example-15)
		- [Example: LIFO](#example-lifo)
- [Panic Function](#panic-function)
		- [Syntax](#syntax-7)
		- [Example](#example-16)
- [Recover Function](#recover-function)
		- [Syntax](#syntax-8)
		- [Example](#example-17)
- [Call by Reference](#call-by-reference)
		- [Example](#example-18)
- [Call by Value](#call-by-value)
		- [Example](#example-19)
- [Method Declaration](#method-declaration)
		- [Syntax](#syntax-9)
		- [Example](#example-20)
		- [Example: the class one](#example-the-class-one)
- [Blank Identifier](#blank-identifier)
		- [Example](#example-21)
- [Concurancy](#concurancy)
- [Parallisim](#parallisim)
- [Goroutines](#goroutines)
		- [Syntax](#syntax-10)
		- [Example- We run this with the `go` keyword and without the `go` keyword, we can see that the go routine runs in the background and prints as soon as it can while Normal function in its own way making it interleaved with a pattern](#example--we-run-this-with-the-go-keyword-and-without-the-go-keyword-we-can-see-that-the-go-routine-runs-in-the-background-and-prints-as-soon-as-it-can-while-normal-function-in-its-own-way-making-it-interleaved-with-a-pattern)
		- [Example: With both as goroutines we get random interleaved output](#example-with-both-as-goroutines-we-get-random-interleaved-output)
		- [Example: No sleep function](#example-no-sleep-function)
		- [Advantages of Goroutines](#advantages-of-goroutines)
- [Channels](#channels)
	- [Operations](#operations-1)
		- [Syntax](#syntax-11)
		- [Example 1: Creating a channel](#example-1-creating-a-channel)
		- [Example 2: Sending and Receiving from Channels](#example-2-sending-and-receiving-from-channels)
		- [Example 3: Sending and Receiving from Channels](#example-3-sending-and-receiving-from-channels)
	- [Advantages of Channels](#advantages-of-channels)
	- [Properties of Channels](#properties-of-channels)
	- [Unbuffered Channels](#unbuffered-channels)
		- [Syntax](#syntax-12)
		- [Example: For a synchronous channel](#example-for-a-synchronous-channel)
	- [Buffered Channels](#buffered-channels)
		- [Syntax](#syntax-13)
		- [Example](#example-22)
		- [Example: Adding more values than the capacity](#example-adding-more-values-than-the-capacity)
		- [Example: Not sending any values](#example-not-sending-any-values)
	- [Unidirectional Channels](#unidirectional-channels)
		- [Example](#example-23)
	- [Looping in Parallel](#looping-in-parallel)
		- [For](#for)
		- [ForEach](#foreach)
- [Multiplexing with select](#multiplexing-with-select)
		- [Syntax](#syntax-14)
		- [Example](#example-24)
	- [Cancellation](#cancellation)
		- [Method 1: Using a Cancellation Channel](#method-1-using-a-cancellation-channel)
			- [Example](#example-25)
	- [Method 2: Using Channel Closure as a Termination Signal](#method-2-using-channel-closure-as-a-termination-signal)
			- [Example:](#example-26)
		- [Constants](#constants-1)
			- [Iota](#iota)
			- [Example: `iota`](#example-iota)
- [Interfaces](#interfaces)
		- [Syntax](#syntax-15)
		- [Example](#example-27)
	- [Empty Interface (IMPORTANT)](#empty-interface-important)
		- [Syntax](#syntax-16)
		- [Example: Simple usage](#example-simple-usage)
		- [Example: Proper an empty interface](#example-proper-an-empty-interface)
	- [Interface as a Contract](#interface-as-a-contract)
		- [Example: Using an interface as a contract](#example-using-an-interface-as-a-contract)
	- [Interface Types](#interface-types)
		- [Example](#example-28)
- [Type Assertion](#type-assertion)
		- [Syntax](#syntax-17)
		- [Example: Basic Assertion that is Correct](#example-basic-assertion-that-is-correct)
		- [Example: Basic Assertion that is Incorrect](#example-basic-assertion-that-is-incorrect)
- [Concurrancy with WaitGroup](#concurrancy-with-waitgroup)
		- [Methods](#methods)
		- [Syntax](#syntax-18)
		- [Example 1](#example-1)
		- [Example 2: Using a loop](#example-2-using-a-loop)
		- [Example 3: Terminating a goroutine using Shared Variable](#example-3-terminating-a-goroutine-using-shared-variable)
	- [Unit 2 End](#unit-2-end)
- [Race Condition and Mutual Exclusion](#race-condition-and-mutual-exclusion)
	- [Example: 1 for Race Condition](#example-1-for-race-condition)
		- [Example 2: For Race Condition](#example-2-for-race-condition)
	- [Race Condition Effects	//TODO: add later](#race-condition-effectstodo-add-later)
		- [Deadlocks](#deadlocks)
		- [Data Corruption](#data-corruption)
	- [Preventing Deadlocks		//TODO: more later](#preventing-deadlockstodo-more-later)
		- [Example: Using a Monitor](#example-using-a-monitor)
- [Mutual Exclusion](#mutual-exclusion)
	- [Binary Semaphores](#binary-semaphores)
		- [Example: Using a Binary Semaphore](#example-using-a-binary-semaphore)
- [Mutex](#mutex)
		- [Syntax: Creation](#syntax-creation)
		- [Syntax: Locking and Unlocking](#syntax-locking-and-unlocking)
		- [Example: Using a Mutex](#example-using-a-mutex)
	- [RW Mutex (Read-Write)](#rw-mutex-read-write)
		- [Syntax: Creation](#syntax-creation-1)
- [Atomic Operations](#atomic-operations)
	- [Common Atomic Operations](#common-atomic-operations)
		- [Load](#load)
		- [Store](#store)
		- [Add and Subtract](#add-and-subtract)
		- [Swap](#swap)
		- [Compare and Swap](#compare-and-swap)
		- [Example 1: Using Atomic Operations](#example-1-using-atomic-operations)
		- [Example 2: Using Atomic Operations with WaitGroup](#example-2-using-atomic-operations-with-waitgroup)
- [Race Detector](#race-detector)
	- [How it works](#how-it-works)
- [Semaphore](#semaphore)
	- [Bounded Channel Approach](#bounded-channel-approach)
		- [Example: Using a Bounded Channel](#example-using-a-bounded-channel)


# Basic
To create a basic program

- We start with the package name;  here it is `main`
- We then import the packages we need; using the `import` keyword
- We then create the main function; using the `func` keyword
- We then write the code inside the main function

- ; is not required at the end of the line
- new line is used to end the statement (basically works like python rather than C)
- Go is case sensitive
- We can add ; to have multiple statements on the same line. Eg: `fmt.Println("Hello, World!"); fmt.Println("Hello, World!");` is in the same line
```go
package main
import ("fmt")


func main() {
	fmt.Println("Hello, World!")
}

```
> OUTPUT: Hello, World!

## Commenting
- Single line comments are done using `//`
- Multi-line comments are done using `/* */`

```go
// This is a single line comment
/* This is a
multi-line comment */
```


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

### Example

```go
package main

import ("fmt")		// <----------- Imported a package

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
> OUTPUT: </br>
> hi this is printprintln </br>
> this is from a variable</br>
> type is string</br>
> type is int</br>

## Printf
- Used to format the output
- Has over a dozen verbs to format the output called `formating verbs`

### Table
| Verb | Description |
| --- | --- |
| %v | the value in a default format |
| %T | type of the value |
| %t | boolean |
| %d | decimal integer |
| %b | binary representation |
| %c | character (rune)|
| %x | hexadecimal |
| %f | float |
| %e | scientific notation |
| %s | string |
| %p | pointer |



# Declarations

- Go has 4 types of declarations 
  - `var` for variables
  - `const` for constants
  - `type` for type declarations
  - `func` for functions

## Predeclared Names
- there are about 36 predeclared names in Go for fucntions, types and constants
- **Constants** are `true`, `false`, `iota`, `nil`
- **Types** are `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`, `float32`, `float64`, `complex64`, `complex128`, `bool`, `byte`, `rune`, `string`, `error`
- **Functions** are `make`, `len`, `cap`, `new`, `append`, `copy`, `close`, `delete`, `complex`, `real`, `imag`, `panic`, `recover`

# Variables

- Variables are declared using the `var` keyword
- Variables can be declared outside the function as well
- Variables can be declared using the `:=` operator as well
- a var declaration can create a variable and assign a value to it. We can declare multiple variables at once using var (a,b,c int)
- Fucntion level variables are delcared within the function
- Package level variables are declared outside the function
	-Four Package level declarations are `var`, `const`, `type`, `func`
- Either the type or the value can be omitted but not both

### Syntax
```go
var variable_name type
var variable_name type = value
variable_name := value
var1, var2, var3 := 1, 2.4 , true
var var1, var2, var3 int
var var1, var2, var3 int = 1,2,3
var (
	var1 int
	var2 int = 2
	var3 int = 3
)
```



### Example

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

### Example: Using Global variables/ Package level variables
```go
package main

func main() {
	a := 10
	println(a)		// 10
	println(b)		// 20; This is a global variable, so it can be used here
}
var b= 20			// <---- b is a global variable
```
> OUTPUT: </br>
> 10 </br>
> 20 </br>

We can see here that the code outside the function runs first and then the code inside the function runs. This allows for global variables to be used inside the function

# IF conditions

- Same as C (almost)
- "{" cannot be at new line
- "else" cannot be on newline, needs to be "} else {" where } is from a prior if statemet  

### Example
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

### Syntax
```go
const variable_name
```


### Example: using a constant 
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

# Operations		

## Arithemetic Operations

` * / % >> << & &^ + - ^ \`

- +,-,*,/: int, float and complex numbers
- %: `only applies to integers`
- sign of remainder always the same as dividend sign
- Follows BODMAS rules

### Example 
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

` & | ^ &^ << >>`
- &     --> Bitwise And
- |     --> Bitwise Or
- ^     --> Bitwise XOR
- &^    --> Bitwise Clear	(AND NOT) {EG: 1010 &^ 0011 = 1000; ans is 0 if the bit is 1 in 0011 else it is the same as the 1010}
- <<    --> Bitwise Left Shift
- \>>   --> Bitwise Right Shift

## Logical Operations

` && || !`
- &&    --> Logical AND
- ||    --> Logical OR
- !     --> Logical NOT

## Assignment Operations

`= += -= *= /= %= &= |= ^= <<= >>= &^=`
- +=    --> a += b is same as a = a + b
- &=    --> a &= b is same as a = a & b
- <<=   --> a <<= b is same as a = a << b
- &^=   --> a &^= b is same as a = a &^ b
- %=    --> a %= b is same as a = a % b
- /=    --> a /= b is same as a = a / b
- |=    --> a |= b is same as a = a | b
- ^=    --> a ^= b is same as a = a ^ b
- >>=   --> a >>= b is same as a = a >> b
- -=    --> a -= b is same as a = a - b
- *=    --> a *= b is same as a = a * b
- =     --> a = b is same as a = b


## Other Operations

` & * <-`
- &     --> Address of
- *     --> Pointer to
- <-    --> Channel send/receive


# Integers

Four types of ints --> `int8`, `int16`, `int32`, `int64`

- Each `int` is a differnt from each other so `int` and `int16` cannot be directly converted.
- `int` is 32 or 64 bits depending on the system. 
- `int8` is same as `byte`
- `int32` is same as `rune`
- `float` to `int` conversion removes the fractional part
- signed no are 2's complement
- higerorder bits are saved for the sign
- range of value for int is --> `(-2^n)-1` to `[(2^n)-1] - 1`
- rnage of bits for non-negative is `0` to `(2^n)-1`
- range of value for int8 is --> `(-128)` to `127`
- range of values for uint8 is --> `0` to `255`


### Syntax
```go
var variable_name int
var variable_name int8
var variable_name int16
```

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
> OUTPUT 3.142857 3

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
> OUTPUT: Error

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

### Examples
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

### Syntax
```go
var variable_name1 complex128 = complex(int_x,int_y)
var variable_name2 complex64 = complex(int_x,int_y)
```


### Example
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

> ### Example: s:="hi"

## Substring

- Substring operations `s[i:j]`  creats a new string from `i` till but not including `j`
- Either or both can be ommited
- Defaults are 0 and len(s)

### Example
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

### Example Assigning
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
### Example: connecting
```go
println("good"+"bye")  // goodbye
```
### Example: appending
```go
s:="hello"
s+=" world"

println(s)		// hello world
```


## String Integer Conversions
- Int to Str ==> strconv.Itoa()
- Str to Int ==> strconv.Atoi()
- Str to Float ==> strconv.ParseFloat

### Example
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
  
### Example
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

### Example where `i` is not required so it is `ommited`
```go
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	
	println(sum)
```
> OUTPUT: 1024

### Example where `for` is used for lists	
```go
package main

func main() {
	var pow = []int{1,2,4,8,16,32,64,128}

	for i, v := range pow {
		println(i,v)
	}
}
```
> OUTPUT: </br>
> 0 1 </br>
> 1 2 </br>
> 2 4 </br>
> 3 8 </br>
> 4 16 </br>
> 5 32 </br>
> 6 64 </br>
> 7 128 </br>

Here we can see that i has the index and v has the value of the list

### Example infinite loop
```go
package main

func main() {
	for {
		// this is an infinite loop
		print("hi ")
	}
}
```
> OUTPUT: hi hi hi hi hi hi hi...

Here we can exit if there is a break statement

# Pointers
- Using pointers that point to a memory. It holds the memory address of a value
- `var p *int`
- `&` generates a pointer to its operand
- `*` denotes the pointers underlying value
- This is known as `dereferencing` or `indirecting`
### Example
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

### Example
```go
prime := [...]int{1,2,3,4}
var s []int = prime[1:3]
println(s)
```

# Appends

- Appends elements to the end of a slice
- If the underlying array is too small, a bigger array is allocated
- Using `append()` function

### Example
```go
func main() {
	var s []int
	s = append(s, 1)
	s = append(s, 2, 3, 4)
	fmt.Println(s)
}
```
> OUTPUT: [1 2 3 4]

Here we can see that the append function is used to add elements to the slice

# Maps		
- Unordered Collection of key-value pairs
- It is a reference to a Hash Table
- Map element is not a variable so we cant take its address; Eg `_ = &m["one"]` gives compiler error
- Maps can be created using the `make` function
- Maps can be compared using `==` and `!=` if all fields are comparable 
- Due to refernece type, if two maps are nil, they are equal
- Due to reference type, they are cheap to pass around
- Type of key and type of value must be same for all keys and all values respectively, if not it gives a compile error
  - Basically all keys together are same type. Eg `{1:"one", "two":2}` or `{1: "one", 2: 2} ` is not allowed, but `{"one":1, "two":2}` is allowed
- a Key Value pair can be of different types.
  
### Syntax
```go
var m map[key_type]value_type
m := map[key_type]value_type{key1:value1, key2:value2}

m := make(map[key_type]value_type)
m := make(map[key_type]value_type, capacity_hint)	// capacity_hint is just a hint for capacity, more elements can be added to the map
```

### Example
```go
var m map[string]int
m = make(map[string]int)
m := map[string]int{"one":1, "two":2}


m["one"] = 1
m["two"] = 2

println(m)
println(m["one"])
```
> OUTPUT: </br>
>  map[one:1 two:2] </br>
> 1

# Default Values

| Type | Default Value |
| --- | --- |
| int | 0 |
| float | 0.0 |
| bool | false |
| string | "" |
| pointer | nil |
| function | nil |
| interface | nil |
| slice | nil |
| channel | nil |
| map | nil |


# Functions	

- Functions are defined using the `func` keyword
- Return type is specified after the parameter list, They are also optional. If not specified, the function does not return anything
- Multiple return values can be returned
- Multiple parameters can be passed. They are also optional

- Parameter list are local to the function, with default value for each datatype
- If there is no return value, the return type is not specified
- There is no concept of default arguments in Go
- Fucntion with return-list needs a `return` statement
- Go supports recursion
  

### Syntax
```go
func name(para_list) (return_list) {
	// code
}
```

### Example: Basic
```go
func add(x int, y int) int {	// here int is used twice in para list
	return x + y
}
func sub(x, y int) int {		// here int is used only once in para list
	return x - y
}
```

### Example: Multiple return values
```go
func swap(x, y string) (string, string) {
	return y, x
}
```

### Example: Named return values
```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return					// Here we can see that there are no arguments here. This is a naked return/ bare return
}
```

## Passing by value
- This means the function gets a copy of the argument so the original value is not changed

### Example: Passing by value
```go
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	x := 10
	y := 20
	swap(x, y)
	println(x, y)
}
```
> OUTPUT: 10 20 </br>
> Here we can see that the values of x and y are not changed

## Passing by reference
- This means the function gets the memory address of the argument so the original value is changed 

### Example: Passing by reference
```go
func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x := 10
	y := 20
	swap(&x, &y)
	println(x, y)
}
```
> OUTPUT: 20 10 </br>
> here we can see that the values of x and y are changed

- If there is no body, it indicates that the function is implemented in some other language (kinda weird not sure what this is on about)
```go
func Sin(x float64) float64
```

## Annonyms Functions
- Named functions are only possible at the package level
- Anonymous functions are possible at the function level
- Functions that are declared without a name
- They can edit local variables for the function they are in

# Return values		

- Multiple return values can be returned
- Named return values can be used to return values
- A return statement without arguments returns the named return values. This is known as a "naked"/"bare" return
### Example
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
### Example
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
- Printf is an ### Example of a varidic function
- Mainly used in situatuons where string formatting is needed
   
### Example
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

### Syntax
```go
defer function_name()
```


### Example
```go
package main

func main() {
	defer println("world")
	println("hello")
}
```
`OUTPUT: hello \n world

### Example: LIFO
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

### Syntax
```go
func panic(/*String to be displayed*/)
```


### Example
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

### Syntax
```go

func main() {
	defer func() {
		if r := recover(); r != nil {
			// Code to handle the panic
		}
	}()
}

// OR

func handlePanic() {
	if r := recover()
	r != nil {
		// Code to handle the panic
	}
}

func main() {
	defer handlePanic()
	panic(/*String to be displayed*/)
}

```

### Example
```go
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
- The original value is changed
### Example
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

# Call by Value

- When the value of a variable is passed to a function, it is known as call by value
- This is done using normal variables
- The original value is not changed

### Example
```go
package main

func main() {
	var x int = 5
	r1 := callByValue(x)
	println("x:", x, "r1:", r1)		// x is 5 and r1 is 50
}

func callByValue(x int) (int) {
	x = 10 * x
	return x
}
```
> OUTPUT: x: 5 r1: 50

# Method Declaration		

- It is a OOP concept
  - An object is a value/variable that has methods
  - A method is a function with a receiver
  - With the help of the receiver parameter, the method can access the properties of the receiver
  - Go does not have classes
- Methods are  declared between the `func` keyword and the function name
- This parameter attaches the function to that type
- Methods can be declared for any type that is declared in the same package
- There is no inheritance in Go, but we can use composition to achieve similar results
- There is no self or this keyword in Go for the receiver. We can use any name for the receiver using the method.
  
### Syntax
```go
type variable struct {
	// fields
}

func (t variable) functionName(parameter list) {	// t is the method receiver
	// code
}
```


### Example
- In this ### Example, the Abs method has a receiver of type Vertex named v
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {		// Creating a method struct
	X, Y float64
}

func (v Vertex) Abs() float64 {				// The function Abs is attached to the Vertex type
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {			// Using the method
	v := Vertex{3, 4}
	fmt.Println(v.Abs())		// We can see that the v is passed to the Abs method
}
```
> OUTPUT: 5 </br>
Here we can see that the `v` variable/object has the ability to use the function/method `Abs` 

### Example: the class one
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


# Blank Identifier
- `_` is known as the blank identifier
- It is used to throw away(ignore) the values returned by a function
- Can be used multiple times 
- Helps avoid compiler errors when a variable is declared but not used

### Example
```go
func main() {
	var op1,op2 int = 10, 20
	_ = op1 + op2				// The result of the addition is ignored
}
```

---------
> Unit 1 End
---------

# Concurancy
- Making progress on more than one task, seemingly at the same time
- Context Switching is when OS saves the context of a process so it can pause it and come back to it later

# Parallisim
- Run mutiple tasks at the same time

# Goroutines

- Concurrency in Go is done through goroutines
- Similar to a thread
- It is an independent fucntion that runs on some seperate lightweight thread managed by the Go runtime
- When a program starts a goroutine calls the main function, which we call the main goroutine
- New goroutines are created using the `go` keyword followed by a function invocation
- This causes it to be called in a new goroutine

### Syntax
```go
go function_name()
```

### Example- We run this with the `go` keyword and without the `go` keyword, we can see that the go routine runs in the background and prints as soon as it can while Normal function in its own way making it interleaved with a pattern
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


### Example: With both as goroutines we get random interleaved output
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

### Example: No sleep function
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

## Operations
- Channel has 2 operations: `send` and `receive`, collectively known as `communication`
- a `send` statement sends a value from one goroutine through a channel to another goroutine
- a `receive` expression receives a value from the channel and assigns it to a variable
- They are used with the `<-` operator. 
- `send` is `ch <- v` (<- seperates the channel and the value)
- `receive` is `v = <-ch` (<- precedes the channel)
- `<-` is a receive operator where the result of the operation is discarded

### Syntax
```go
var ch chan int
ch = make(chan int)

// Using in a function
func goroutine(variable_name chan int) {
	// code
}
```

### Example 1: Creating a channel
```go
package main

import "fmt"

func main() {
	var char1 chan int		// creating a channel

	fmt.Println("Value of char1 is: ", char1)		// nil
	fmt.Printf("Type of char1 is: %T\n", char1)		// chan int

	char2 := make(chan int)
	fmt.Println("Value of char2 is: ", char2)		// 0xc000016120
	fmt.Printf("Type of char2 is: %T\n", char2)		// chan int
}
```
> OUTPUT: </br>
> Value of char1 is:  nil </br>
> Type of char1 is:  chan int </br>
> Value of char2 is:  0xc000016120 </br>
> Type of char2 is:  chan int </br>


### Example 2: Sending and Receiving from Channels
```go
package main

func myChannelTest(i int,ch chan int) {
    ch <- i
}

func main() {
	ch := make(chan int)
	
	go func() {
			ch <- 1
		ch <- 2
		ch <- 3
	}()
	
	x := <-ch		// 1
	y := <-ch		// 2
	z := <-ch		// 3

	println(x, y, z)	// 1 2 3
}
```
> OUTPUT: </br>
> 1 2 3

Here the channel acts like a queue as when the single goroutine runs, it sets the values in the order 1,2,3 and when the main function runs, it gets the values in the same order

### Example 3: Sending and Receiving from Channels
```go
package main

func myChannelTest(i int,ch chan int) {
    ch <- i
}

func main() {
	ch := make(chan int)
	go myChannelTest(1,ch)
	x := <-ch		// 1
	y := <-ch		// 3
	go myChannelTest(3,ch)
	go myChannelTest(2,ch)
	z := <-ch		// 2
	
	println(x, y, z)	// 1 3 2
}
```
> OUTPUT: </br>
> 1 3 2


Here what happens is that `x` gets 1 as it is right after the function and `z` gets 2 as it is right after the line where the function is called. `y` gets 3 as it is the last value in the channel (My Guess as thats what my experimentation shows me)

## Advantages of Channels
- Channels are typesafe. They only allow a particular type of data to be sent through them	
- They enable safe communication between goroutines and synchronisation
- They ensure the code runs predictably and efficiently with concurrent.
- **Blocking Send and Recive:** If a goroutine tries to send to a channel that is full, it will block until another goroutine receives from the channel. If a goroutine tries to receive from an empty channel, it will block until another goroutine sends to the channel
- **Zero Value Channel:** The zero value of a channel is `nil`. A nil channel is of no use and attempting to send or receive from it will cause a deadlock
- **For Loop Channel:** A loop can iterate over the sequntial values sent by a channel until it is closed. This is useful when the number of values is not known in advance

## Properties of Channels
- **Lenght of Channel:** This can be found using the `len` function. This is the number of elements currently buffered in the channel
- **Capacity of Channels:** This can be found using the `cap` function. This is the number of elements that can be buffered in the channel


## Unbuffered Channels
- A `send` operation on an unbuffered channel blocks the sending goroutine until another goroutine executes a corresponding `receive` on the same channel, at which point the value is transmitted and both goroutines may continue
- A `receive` operation on an unbuffered channel blocks the receiving goroutine until another goroutine executes a corresponding `send` on the same channel
- Communication for unbuffered channels is synchronous, Hence they are called `synchronous channels`

### Syntax
```go
ch := make(chan int)
```
- Here we have not set any limit for the channel


### Example: For a synchronous channel 
This is a pipeline as well as they are connected in a chain
```go
package main

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
}

func counter(out chan<- int) {		// Sends to `naturals` channel
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) { 		// Receives from `naturals` and sends to `squares`
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {		// Receives from `squares` channel
	for x := range in {
		println(x)
	}
}
```
> OUTPUT: </br>
> 0  </br>
> 1  </br>
> 4  </br>
> 9  </br>
> 16 </br>
> 25 </br>
> 36 </br>
> 49 </br>
> 64 </br>
> 81 </br>


## Buffered Channels
- This has a queue of elements
- The Max capacity of the queue is set when the channel is created	

### Syntax		
```go
ch := make(chan int, 3)
```
- Here we can send 3 values to the channel without blocking
- The fourth value will block.

### Example
```go
package main

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	println(<-ch)
	println(<-ch)

	ss := make(chan string, 2)
	ss <- "Hello"
	ss <- "World"

	println(<-ss)
	println(<-ss)
}
```
> OUTPUT: </br>
> 1 </br>
> 2 </br>
> Hello </br>
> World </br>

### Example: Adding more values than the capacity
```go
func main(){
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3		// <-- ERROR: This will block as it is the 3rd send when capacity is 2
}
```
> OUTPUT: fatal error: all goroutines are asleep - deadlock!

### Example: Not sending any values
```go
func main(){
	ch := make(chan int, 2)
	println(<-ch)		// <-- ERROR: This will block as there are no values to receive
}
```
> OUTPUT: fatal error: all goroutines are asleep - deadlock!

Hence we can see that we need to have the same number of sends and receives for the channel to work and this number should be less than or equal to the capacity of the channel
`n(sender) == n(receiver) <= capacity`  {Is another way to put it}

## Unidirectional Channels
- Channels can be unidirectional. This means that they can only send or receive data
- The type `chan <- int` is a send-only channel of integers
- The type `<- chan int` is a receive-only channel of integers

### Example
```go
func main() {
	send := make(chan<- int, 2)		// Send only channel
	receive := make(<-chan int, 2)	// Receive only channel


	send <- 1			// This will work
	println(<-send)		// ERROR: This will not work as it is a send only channel

	println(<-receive)	// This will work
	receive <- 2		// ERROR: This will not work as it is a receive only channel
}
```


## Looping in Parallel		
- We can loop over a channel in parallel
- Problems that are independent of each other are known as `embarrassingly parallel` problems
- These are the easiest to parallelise, we can do this through goroutines 
- They scale linearly with the amount of parallelism


NOT SURE WHAT THIS STUFF IS BELOW BUT ITS FROM SLIDES

### For
- Repeats function in parallel. Internally calls ForLoop
```go
func For(begin int, end int, f ForLoop) {
}
```

### ForEach
- Loops collections in parallel collection: slice, array, map, string, channel, etc. If multiple options, only first one is valid
```go
func ForEach(collection interface{}, f interface{}) {
}
```

- We can limit the amount of parallelism by using buffered channels

# Multiplexing with select

- Similar to a `switch` statement
- Used to choose from multiple send/receive channel operations
- There is no limit to the number of cases
- Values of each case must be operations on channels
- There are no `break` statements in the `Case`
- `Case` statements only run if the channel is possible to send or receive
  
- In case of multiple cases being ready, one is chosen at random
- When no `Case` is ready, the system goes into deadlock. (copilot says it goes to default though)


- The `select` statement lets a goroutine wait on multiple communication operations
- the `select` statement blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready
- A `select` with no cases blocks forever, ie waits forever

### Syntax
```go
select {
	case send_op:
		// code
	case receive_op:
		// code
	default:
		// code
}
```

### Example
```go
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case v := <-ch1:
				println("From ch1:", v)
			case v := <-ch2:
				println("From ch2:", v)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case ch1 <- i:
		case ch2 <- i:
		}
	}
}
```
> OUTPUT: </br>
> From ch1: 0 </br>
> From ch1: 1 </br>
> From ch1: 2 </br>
> From ch1: 3 </br>
> From ch2: 4 </br>
> From ch2: 5 </br>
> From ch2: 6 </br>
> From ch1: 7 </br>
> From ch1: 8 </br>
> From ch2:
## Cancellation
- There is no way to directly cancel a goroutine as it would leave the state of the program in an unknown state
- We can use a `cancellation` channel to signal the goroutine to stop
- We also defina a utility function to check if the channel is closed
- In general it is hard to know how many goroutines are running at any given time

### Method 1: Using a Cancellation Channel
- We create a channel on which no values are sent but whose closure signals the goroutine to stop

#### Example
```go
package main

func worker(stop <-chan bool) {
	for {
		select {
		case <-stop:
			println("Worker stopped")
			return
		default:
			println("Worker processing while main function is busy")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	stop := make(chan bool)
	go worker(stop)

	time.Sleep(5 * time.Second) // Let the worker work for 5 seconds
	stop <- true                 // Send a signal to stop the worker
	time.Sleep(1 * time.Second)  // Give the worker some time to stop
}
```
> OUTPUT: </br>
> Worker processing while main function is busy </br>
> Worker processing while main function is busy </br>
> Worker processing while main function is busy </br>
> Worker processing while main function is busy </br>
> Worker processing while main function is busy </br>
> Worker stopped

Here we can see that a signal is sent to the worker to stop after 5 seconds and the worker stops after 1 second
NOTE: The value sent to channel should match the channel type (eg chan bool, chan string), stop <- 1 will also work for bool though

## Method 2: Using Channel Closure as a Termination Signal
- Goroutine termination by utilizing the close fucntion on a channel

#### Example:
```go
func processStop(stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			println("Stopped")
			return
		default:
			println("Processing")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	stop := make(chan struct{})
	go processStop(stop)

	time.Sleep(5 * time.Second)
	close(stop)						// Using Closure
	time.Sleep(5 * time.Second)
}

```
> OUTPUT: </br>
> Processing </br>
> Processing </br>
> Processing </br>
> Processing </br>
> Processing </br>
> Stopped </br>

Here we can see that the goroutine stops after 5 Processing statements then stops after the stop channel is closed


### Constants

#### Iota
- This is a predeclared identifier
- Basically a counter for constants that increases by 1 for each line
- Always starts at 0

#### Example: `iota`
```go
package main

const (
	sunday = iota		// 0
	monday 				// 1
	tuesday				// 2
	wednesday 			// 3	// ALSO NOTE if we make wednesday = 5, all following values will be 5
	thursday = iota		// 4	// NOTE: adding iota here makes no difference even if it is not here	
	friday				// 5
)

func main() {
	println(sunday, monday, tuesday, wednesday, thursday, friday)
}
```
> OUTPUT: </br>
> 0 1 2 3 4 5 </br>

# Interfaces
- A set of method signatures that a type can implement
- flexible and powerful for writing reusable code
- A variable of an interface type can hold any value that implements those methods
- Interfaces are implemented implicitly. No explicit declaration as no `implements` keyword

### Syntax
```go
type interface_name interface {
	Method() signatures
}
```

### Example
```go
// interface
type Geometry interface {
	area() string
}

// rect object
type rect struct {
	width, height int
}

// circle object
type circle struct {
	radius float64
}

// method for rect
func (r rect) area() {
	fmt.Printf("Area of rectangle is: %d\n", r.width*r.height)
}

// method for circle
func (c circle) area() {
	fmt.Printf("Area of circle is: %.2f\n", 3.14*c.radius*c.radius)
}

func measure(g Geometry) {
	g.area()
}

func main() {
	r := rect{width: 10, height: 5}
	c := circle{radius: 5}

	var g Geometry
	g = r				// print(g) --> {10 5}
	measure(g)			// Area of rectangle is: 50
	g = c				// print(g) --> {5}
	measure(g)			// Area of circle is: 78.50
}
```
> OUTPUT: </br>
> Area of rectangle is: 50 </br>
> Area of circle is: 78.5 </br>

## Empty Interface (IMPORTANT)
- This is an interface with no methods
- Can hold any value regardless of its type
- Used when we want to accept values of any type


### Syntax
```go
var i interface{}
```

### Example: Simple usage
```go
var any interface{}
any = 42
any = "hi"
any = 3.14159
any = new(bytes.Buffer)
any = make(chan int, 3)
any = make(map[string]int)
```

### Example: Proper an empty interface
```go
func PrintAny (i interface{}) {
	fmt.Println(i)
}

func main() {
	PrintAny(42)		// prints 42
	PrintAny("Hello")	// prints Hello
	PrintAny(3.14159)	//	prints 3.14159
}
```
> OUTPUT: </br>
> 42 </br>
> Hello </br>
> 3.14159 </br>

Here we can see that the empty interface can take any type of value

## Interface as a Contract
- Interfaces are a contract that a method can implement


### Example: Using an interface as a contract
```go
func myPrintF(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
func main() {
	myPrintF("Hello %s\n", "World")
}
```
> OUTPUT: Hello World

Here we can see that the `myPrintF` function is a contract that the `fmt.Printf` function implements and `args` is an empty interface that can take any type of value


## Interface Types
- New interface types can be a combination of existing ones.
- This is called `embedding` or `composition` of an interface
  
### Example
```go
type Reader interface {					// Reader interface
	Read(p []byte) (n int, err error)
}	
type Writer interface {					// Writer interface
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {		// This is a combination of Reader and Writer
	Reader
	Writer
}
```
**Note(tho might not be important):** The `Write` method takes the data from the byte slice p writes into the underlying data stream and returns the number of bytes written and an error if any


# Type Assertion
- Basic asserstions are done using the `.` operator

### Syntax
```go

val, ok := interface_var(type)
// val is the value of the variable
// ok is a boolean value that is true if the assertion holds
```

### Example: Basic Assertion that is Correct
```go
var i interface{} = "hello"
s,ok := i.(string)
fmt.Println(s, ok)	// s = hello; ok = true
```
> OUTPUT: hello true

Here we can see that the assertion is correct so the value of `i` is stored in `s` and `ok` is true
### Example: Basic Assertion that is Incorrect
```go
var i interface{} = 42
s,ok := i.(string)
fmt.Println(s, ok)	// s = ""; ok = false
```
> OUTPUT:  false

Here we can see on wrong asserstion `s` gets the default value of the type it was checking

.....
 // TODO: Add this LATER
list/slice == comarision
creation of new struct using &
.....


# Concurrancy with WaitGroup
- To wait for multiple goroutines to finish, we can use a `WaitGroup`

### Methods
- `Add(int)` method is used to add the number of goroutines to wait for
- `Done()` method is used to signal that a goroutine is done. Decrements the counter by 1
- `Wait()` method is used to block until the counter is 0

> NOTE: You need to import the `sync` package to use the `WaitGroup`. The Counter starts at 0

### Syntax
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// code
}()

wg.Wait()
```

### Example 1
```go
package main

import "sync"

func printRam(wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		println("RAM")
	}
	wg.Done()				// Decrements the counter by 1
}

func printSita(wg *sync.WaitGroup) {
	for i := 0; i < 2; i++ {
		println("Sita")
	}
	wg.Done()				// Decrements the counter by 1
}

func main() {
	var wg sync.WaitGroup		// can use wg:=sync.WaitGroup{} instead

	wg.Add(2)			// adds 2 goroutines to wait for
	
	go printRam(&wg)		// counter is 1 after this
	go printSita(&wg)		// counter is 0 after this
	
	wg.Wait()			// waits until the counter for goroutines is 0
	println("Done")
}
```
> OUTPUT: </br>
> SITA </br>
> SITA </br>
> RAM </br>
> RAM </br>
> RAM </br>
> Done </br>

Here the order of SITA or RAM is not fixed as the goroutines run in parallel

### Example 2: Using a loop
```go
package main

import "sync"

func main() {
	wg:=sync.WaitGroup{}

	for i:=0; i<3; i++ {
		wg.Add(1)			// Increment the counter by 1
		go doSomething(&wg,i)	// Run the goroutine and decrement the counter by 1
	}
	wg.Wait()			// Check if the counter is 0, then continue
	println("Done")
}

func doSomething(wg *sync.WaitGroup, i int) {
	println("Doing something", i)
	wg.Done()
}
```
> OUTPUT: </br>
> Doing something 2</br>
> Doing something 1</br>
> Doing something 0</br>
> Done </br>

### Example 3: Terminating a goroutine using Shared Variable
```go
package main

import (
	"sync"
	"time"
)

func main() {
	var stopIT bool
	var wg sync.WaitGroup

	wg.Add(1)
	go process(&wg, &stopIT)

	time.Sleep(3 * time.Second)			// makes main fucntion wait for 3 seconds
	stopIT = true						// Shared variable to stop the goroutine
	wg.Wait()				
	print("Done")
}

func process(wg *sync.WaitGroup, stopIT *bool) {
	defer wg.Done()

	for {
		println("Processing")	// This will run until the stopIT is true every second in the background
		time.Sleep(1 * time.Second)

		if *stopIT {
			println("Stopped Goroutine")
			return
		}
	}
}
```
> OUTPUT: </br>
> Processing </br>
> Processing </br>
> Processing </br>
> Stopped Goroutine </br>
> Done </br>

- If we have number of `wg.Done()` > `wg.Add()` we get a panic as the counter goes below 0
- If we have number of `wg.Add()` > `wg.Done()` we get a deadlock as the counter never reaches 0

-----
Unit 2 End
-----

# Race Condition and Mutual Exclusion

- Race condition is when the output of a program depends on the order of execution of the goroutines
- Occurs when multiple goroutines access the same resource
- This can lead to unpredictable results

- Mutual Exclusion is a technique to ensure that only one goroutine can access a resource at a time
- This is done using a `mutex` or `lock`

## Example: 1 for Race Condition
```go
package main

import "time"

var balance int

func main() {
	go func() {
		Deposit(200)                    // A1
		println("Balance =", Balance()) // A2	200
	}()

	go Deposit(100) // B
	time.Sleep(3 * time.Second)
	print("Main Goroutine finished")
}

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}
```

> OUTPUT: </br>
> Balance = 200 </br>
> Main Goroutine finished </br>

Here we can see that the balance is updated by the goroutines so they are no longer guaranteed to giev the right answer
So Balance could be 300 if the goroutine B finishes first or 200 if the goroutine A finishes first

THIS IS CALLED DATA RACE


### Example 2: For Race Condition 
```go
package main

func main() {
	var data int		
	go func() {
		data++			// A1
	}
	
	if data == 0 {
		println(data)	// A2
	}
}
```
> OUTPUT: </br>
> 0

Here data may be 1 or 0 depending on if A1 or A2 runs first

## Race Condition Effects	//TODO: add later
### Deadlocks

### Data Corruption

## Preventing Deadlocks		//TODO: more later

- Avoiding direct access to shared variables
- Use a Channel to send a request a query/update a variable to goroutine 
- A goroutine that manages access called a `monitor` goroutine

### Example: Using a Monitor
```go
package main

var depoChannel = make(chan int) // A
var balanceChannel = make(chan int) // B

func Deposit(amount int) {
	depoChannel <- amount
}

func Balance() int {
	return <-balanceChannel
}

func teller() {
	var balance int // A
	for {				// Infinite loop ensures it will always be ready to receive a message
		select {
		case amount := <-depoChannel:	// Take the amount from the depoChannel channel
			balance += amount		// Add the amount to the balance
		case balanceChannel <- balance:	//	Send the balance to the balanceChannel channel
		}
	}
}


func main() {
	go teller() // B
	go Deposit(200)
	println("Balance =", Balance()) // 200
	go Deposit(100)
	println("Balance =", Balance()) // 300

}
```

> OUTPUT: </br>
> Balance = 200 </br>
> Balance = 300 </br>

Here we can see that the balance is updated correctly as the teller goroutine manages the balance


# Mutual Exclusion
>NOTE Empty Struct: `struct{}` does not have any values, so does not take up any memory

- Mutual Exclusion is a technique to ensure that only one goroutine can access a resource at a time
- Empty struct `struct{}` are used as placeholder
- Often used in situation where we want to signal with no additional data
- Usually when semaphores are created we see something like this 
  - ```go
	var sem = make(chan struct{}, numberOfSemaphores)
	sem <- struct{}{}
	```

## Binary Semaphores
- A binary semaphore is a semaphore that can have only two values: 0 and 1
- It is used to ensure only one goroutine can access a resource at a time

### Example: Using a Binary Semaphore
```go
package main

import "time"

var sema = make(chan struct{}, 1)
var balance int

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance += amount  // deposit
	<-sema             // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance       // copy balance
	<-sema             // release token
	return b           // return balance
}

func main() {
	println("Balance =", Balance()) // 0

	go Deposit(200)
	go Deposit(100)

	time.Sleep(1 * time.Second)
	println("Balance =", Balance()) // 300
}
```
> OUTPUT: </br>
> Balance = 0 </br>
> Balance = 300 </br>

# Mutex
- A mutex is a lock that allows only one goroutine to access a resource at a time

### Syntax: Creation
```go
var mu sync.Mutex
mu := sync.Mutex{}
```

### Syntax: Locking and Unlocking
```go
//func (m *Mutex) Lock()		// Locks m
//func (m *Mutex) Unlock()	// Unlocks m

mu.Lock()
// code
mu.Unlock()

```


### Example: Using a Mutex
```go
package main

import (
	"sync"
	"time"
)

var mu sync.Mutex
var balance int

func Deposit(amount int) {
	mu.Lock()			// Lock the mutex
	balance += amount
	mu.Unlock()			// Unlock the mutex to avoid deadlock
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()		// We can use defer to unlock the mutex
	return balance
}

func main() {
	println("Initial balance: ", Balance())

	go Deposit(100)
	go Deposit(200)

	time.Sleep(1 * time.Second)
	println("Final balance: ", Balance())
}
```
> OUTPUT: </br>
> Initial balance:  0 </br>
> Final balance:  300 </br>

## RW Mutex (Read-Write)
- It is a Read-Write Mutex
- The lock can be held by `multiple` of readers or a `single` writer
- The zero value of a `RWMutex` is an unlocked mutex

- Should not have recurive read locks as it makes it so only one reader can read at a time

### Syntax: Creation
```go
// func (rw *RWMutex) RLock()		// Locks rw for reading
// func (rw *RWMutex) RUnlock()	// Unlocks rw for reading

// Creation
rw := sync.RWMutex{}
var rw sync.RWMutex


// Usage
rw.RLock()
// code
rw.RUnlock()
```

# Atomic Operations
- Operations that are designed to be performed without interference from other operations
- They are used to ensure that a operations on shared variables are atomic

## Common Atomic Operations

### Load
- Atomically read the value of a variable
- Eg: `atomic.LoadInt32(&x)` loads the value of x (type is int32)
### Store
- Atomically set a value to a variable
- Eg: `atomic.StoreInt32(&x, 42)` sets the value of x to 42 (type is int32)
### Add and Subtract
- Atomically add or subtract a value from a variable
- Eg: `atomic.AddInt32(&x, 1)` adds 1 to the value of x (type is int32)
- Eg: `atomic.AddInt32(&x, -1)` subtracts 1 from the value of x (type is int32)
### Swap


### Compare and Swap


### Example 1: Using Atomic Operations
```go
package main

import (
	"sync/atomic"
	"time"
	"fmt"
)

func main() {
	var counter int32
	
	go func() {
		for i:=0; i<5; i++ {
			atomic.AddInt32(&counter, 1)	// Increment the counter
			fmt.Printf("Incrementing counter: %d\n", atomic.LoadInt32(&counter))
			time.Sleep(time.Millisecond)
		}
	}()

	go func() {
		for i:=0; i<5; i++ {
			atomic.AddInt32(&counter, -1) 		// Decrement the counter
			fmt.Printf("Decrementing counter: %d\n", atomic.LoadInt32(&counter))
			time.Sleep(time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Printf("Final counter: %d\n", atomic.LoadInt32(&counter))
}
```
> OUTPUT: </br>
> Decrementing counter: -1 </br>
> Incrementing counter: 0 </br>
> Incrementing counter: 1 </br>
> Decrementing counter: 0 </br>
> Decrementing counter: -1 </br>
> Incrementing counter: 0 </br>
> Final counter: 0 

### Example 2: Using Atomic Operations with WaitGroup
```go
package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	var ops atomic.Uint64	// import "sync/atomic"
	var wg sync.WaitGroup	// import "sync"

	for i:=0; i<50; i++ {
		wg.Add(1)
		go func() {
			for c:=0; c<100; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	println("ops: ", ops.Load())
}
```
> OUTPUT: </br>
> ops:  5000

# Race Detector
- A tool that detects race condition

## How it works
- It is integrated with go toolchain
- when the `-race flag` is used, the compiler instruments all memory accesses with code records when an how memory accessed, while the runtime detects unsynchronized access to the same memory location
- When such "racy" behaviour is detected, the runtime prints a warning.


# Semaphore
- A semaphore is a variable or abstract data type used to control access to a common resource by multiple processes in a concurrent system such as a multitasking operating system

## Bounded Channel Approach

### Example: Using a Bounded Channel
```go

func remoteDeleteEmployee() {
	// code
}

// sem is a channel that allows 10 goroutines to run at once
var sem = make(chan struct{}, 10)	// bufferred channel

func main() {
	for _, empl in range employees {
		sem <- 1		// acquire token
		go func() {
			remoteDeleteEmployee(empl.ID)
			<-sem			// release token
		}()
	}
}
```
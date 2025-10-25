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
	- [Examples](#examples)
		- [Example: Basic](#example-basic)
		- [Example: Using Global variables/ Package level variables](#example-using-global-variables-package-level-variables)
- [If conditions](#if-conditions)
	- [Syntax](#syntax-1)
	- [Example](#example-1)
- [Constants](#constants)
	- [Syntax](#syntax-2)
	- [Example: using a constant](#example-using-a-constant)
- [Operations](#operations)
	- [Arithemetic Operations](#arithemetic-operations)
		- [Example](#example-2)
	- [Comparison Operations](#comparison-operations)
	- [Bitwise Binary Orperations](#bitwise-binary-orperations)
	- [Logical Operations](#logical-operations)
	- [Assignment Operations](#assignment-operations)
	- [Other Operations](#other-operations)
- [Integers](#integers)
	- [Syntax](#syntax-3)
	- [Examples](#examples-1)
		- [Example: Signed Integers](#example-signed-integers)
		- [Example: Unsigned Integers](#example-unsigned-integers)
	- [Type Casting](#type-casting)
		- [Syntax](#syntax-4)
		- [Example: Type casting (Correct)](#example-type-casting-correct)
		- [Example: Type casting Problems (Incorrect Print)](#example-type-casting-problems-incorrect-print)
		- [Example: Type Casting Problems (int32 int64 addition)](#example-type-casting-problems-int32-int64-addition)
- [Floating Type Numbers](#floating-type-numbers)
	- [Max Values are in math package](#max-values-are-in-math-package)
		- [Examples](#examples-2)
	- [Defaults](#defaults)
- [Complex Numbers](#complex-numbers)
	- [Syntax](#syntax-5)
	- [Example](#example-3)
- [Boolean](#boolean)
	- [Basic usage](#basic-usage)
	- [Using in IF conditions](#using-in-if-conditions)
- [Strings](#strings)
	- [Substring](#substring)
		- [Example](#example-4)
		- [Example Assigning](#example-assigning)
	- [Appending](#appending)
		- [Example: connecting](#example-connecting)
		- [Example: appending](#example-appending)
	- [String Integer Conversions](#string-integer-conversions)
		- [Example](#example-5)
- [For Loops](#for-loops)
	- [Syntax](#syntax-6)
	- [Examples](#examples-3)
		- [Example: Basic](#example-basic-1)
		- [Example where `i` is not required so it is `ommited`](#example-where-i-is-not-required-so-it-is-ommited)
		- [Example where `for` is used for lists](#example-where-for-is-used-for-lists)
		- [Example infinite loop](#example-infinite-loop)
- [Pointers](#pointers)
	- [Syntax](#syntax-7)
	- [Example](#example-6)
- [Structs](#structs)
	- [Syntax](#syntax-8)
	- [Example](#example-7)
- [Arrays](#arrays)
	- [Syntax](#syntax-9)
	- [Examples](#examples-4)
		- [Example: Basic](#example-basic-2)
		- [Example: Ellipsis `...` `(Imp)`](#example-ellipsis--imp)
		- [Example: Comparisons `(Imp)`](#example-comparisons-imp)
- [Slice](#slice)
	- [Syntax](#syntax-10)
	- [Example](#example-8)
- [Appends](#appends)
	- [Syntax](#syntax-11)
	- [Example](#example-9)
- [Maps](#maps)
	- [Syntax](#syntax-12)
	- [Example](#example-10)
- [Default Values](#default-values)
- [Functions](#functions)
	- [Syntax](#syntax-13)
	- [Examples](#examples-5)
		- [Example: Basic](#example-basic-3)
		- [Example: Multiple return values](#example-multiple-return-values)
		- [Example: Named return values](#example-named-return-values)
	- [Passing by value](#passing-by-value)
		- [Example: Passing by value](#example-passing-by-value)
	- [Passing by reference](#passing-by-reference)
		- [Example: Passing by reference](#example-passing-by-reference)
	- [Annonyms Functions](#annonyms-functions)
		- [Syntax](#syntax-14)
		- [Example](#example-11)
- [Return values](#return-values)
		- [Example](#example-12)
	- [Bare Returns](#bare-returns)
		- [Example](#example-13)
	- [Varidic Functions](#varidic-functions)
		- [Example](#example-14)
- [Deferred Fucntions](#deferred-fucntions)
	- [Syntax](#syntax-15)
	- [Examples](#examples-6)
		- [Example: Basic](#example-basic-4)
		- [Example: Function call](#example-function-call)
		- [Example: LIFO](#example-lifo)
- [Panic Function](#panic-function)
	- [Syntax](#syntax-16)
	- [Example](#example-15)
- [Recover Function](#recover-function)
	- [Syntax](#syntax-17)
	- [Example](#example-16)
- [Call by Reference](#call-by-reference)
	- [Example](#example-17)
- [Call by Value](#call-by-value)
	- [Example](#example-18)
- [Method Declaration](#method-declaration)
	- [Syntax](#syntax-18)
	- [Examples](#examples-7)
		- [Example: Basic](#example-basic-5)
		- [Example: the class one](#example-the-class-one)
- [Blank Identifier](#blank-identifier)
	- [Example](#example-19)
	- [How to Start a Go Project](#how-to-start-a-go-project)
		- [1. Create a New Project Directory](#1-create-a-new-project-directory)
		- [2. Initialize a Go Module](#2-initialize-a-go-module)
		- [3. Create Your Main File](#3-create-your-main-file)
		- [4. Run Your Project](#4-run-your-project)
		- [5. Build Your Project](#5-build-your-project)
	- [Folder Structures](#folder-structures)
		- [General Principles](#general-principles)
		- [Example: Simple CLI Tool](#example-simple-cli-tool)
		- [Example: Microservice](#example-microservice)
		- [Example: Web Server](#example-web-server)
		- [Example: Library/Package](#example-librarypackage)
		- [Example: Monorepo (Multiple Services)](#example-monorepo-multiple-services)
		- [Example: Small Script](#example-small-script)
		- [Example: Large Enterprise App](#example-large-enterprise-app)
		- [Example: Test-Heavy Project](#example-test-heavy-project)
- [Pacakges](#pacakges)
	- [Importing Standard Library Packages](#importing-standard-library-packages)
		- [Syntax](#syntax-19)
		- [Example: Using `fmt` and `math`](#example-using-fmt-and-math)
		- [Notes](#notes)
	- [Installing and Importing Third-Party Packages](#installing-and-importing-third-party-packages)
		- [Installing a Package](#installing-a-package)
		- [Importing a Third-Party Package](#importing-a-third-party-package)
		- [Example: Using an External Package](#example-using-an-external-package)
		- [Notes](#notes-1)
	- [Creating Your Own Packages](#creating-your-own-packages)
		- [Basic Structure](#basic-structure)
		- [Notes](#notes-2)
	- [Project Structure and Multiple Files](#project-structure-and-multiple-files)
		- [Typical Project Layout](#typical-project-layout)
		- [How Multiple Files Work Together](#how-multiple-files-work-together)
	- [Importing and Using Local Packages](#importing-and-using-local-packages)
		- [Import Path](#import-path)
		- [Example: Using a Local Package](#example-using-a-local-package)
		- [Notes](#notes-3)
	- [Best Practices for Packages and Project Structure](#best-practices-for-packages-and-project-structure)
		- [1. Use Go Modules](#1-use-go-modules)
		- [2. Name Packages Clearly](#2-name-packages-clearly)
		- [3. Export Only What’s Needed](#3-export-only-whats-needed)
		- [4. Keep Package Scope Small](#4-keep-package-scope-small)
		- [5. Use Proper Import Paths](#5-use-proper-import-paths)
		- [6. Document Your Packages](#6-document-your-packages)
		- [7. Keep `main` Package Simple](#7-keep-main-package-simple)
		- [8. Use `go fmt` and `go vet`](#8-use-go-fmt-and-go-vet)
		- [9. Avoid Circular Imports](#9-avoid-circular-imports)
		- [10. Use Versioning for Dependencies](#10-use-versioning-for-dependencies)
	- [Useful Go Commands Reference](#useful-go-commands-reference)
		- [Module and Dependency Management](#module-and-dependency-management)
		- [Building and Running](#building-and-running)
		- [Formatting and Linting](#formatting-and-linting)
		- [Testing](#testing)
		- [Other Useful Commands](#other-useful-commands)
- [Concurancy](#concurancy)
- [Parallisim](#parallisim)
- [Goroutines](#goroutines)
	- [Syntax](#syntax-20)
	- [Examples](#examples-8)
		- [Example- Basic](#example--basic)
		- [Example: With both as goroutines we get random interleaved output](#example-with-both-as-goroutines-we-get-random-interleaved-output)
		- [Example: No sleep function](#example-no-sleep-function)
	- [Advantages of Goroutines](#advantages-of-goroutines)
- [Channels](#channels)
	- [Operations](#operations-1)
	- [Syntax](#syntax-21)
	- [Examples](#examples-9)
		- [Example 1: Creating a channel](#example-1-creating-a-channel)
		- [Example 2: Sending and Receiving from Channels](#example-2-sending-and-receiving-from-channels)
		- [Example 3: Sending and Receiving from Channels](#example-3-sending-and-receiving-from-channels)
	- [Advantages of Channels](#advantages-of-channels)
	- [Properties of Channels](#properties-of-channels)
	- [Unbuffered Channels](#unbuffered-channels)
		- [Syntax](#syntax-22)
		- [Example: For a synchronous channel](#example-for-a-synchronous-channel)
	- [Buffered Channels](#buffered-channels)
		- [Syntax](#syntax-23)
		- [Examples](#examples-10)
			- [Example: Basic](#example-basic-6)
			- [Example: Adding more values than the capacity](#example-adding-more-values-than-the-capacity)
			- [Example: Not sending any values](#example-not-sending-any-values)
	- [Unidirectional Channels](#unidirectional-channels)
		- [Example](#example-20)
	- [Looping in Parallel](#looping-in-parallel)
		- [For](#for)
		- [ForEach](#foreach)
- [Multiplexing with select](#multiplexing-with-select)
	- [Syntax](#syntax-24)
	- [Example](#example-21)
- [Cancellation](#cancellation)
	- [Method 1: Using a Cancellation Channel](#method-1-using-a-cancellation-channel)
		- [Example](#example-22)
	- [Method 2: Using Channel Closure as a Termination Signal](#method-2-using-channel-closure-as-a-termination-signal)
		- [Example:](#example-23)
- [Constants](#constants-1)
	- [Iota](#iota)
	- [Example: `iota`](#example-iota)
- [Interfaces](#interfaces)
	- [Syntax](#syntax-25)
	- [Example](#example-24)
	- [Empty Interface (IMPORTANT)](#empty-interface-important)
		- [Syntax](#syntax-26)
		- [Examples](#examples-11)
			- [Example: Simple usage](#example-simple-usage)
			- [Example: Proper an empty interface](#example-proper-an-empty-interface)
	- [Interface as a Contract (NOT COM)](#interface-as-a-contract-not-com)
		- [Example: Using an interface as a contract](#example-using-an-interface-as-a-contract)
	- [Interface Types](#interface-types)
		- [Example](#example-25)
- [Type Assertion](#type-assertion)
	- [Syntax](#syntax-27)
	- [Examples](#examples-12)
		- [Example: Basic Assertion that is Correct](#example-basic-assertion-that-is-correct)
		- [Example: Basic Assertion that is Incorrect](#example-basic-assertion-that-is-incorrect)
- [Concurrancy with WaitGroup](#concurrancy-with-waitgroup)
	- [Methods](#methods)
	- [Syntax](#syntax-28)
	- [Examples](#examples-13)
		- [Example 1](#example-1)
		- [Example 2: Using a loop](#example-2-using-a-loop)
		- [Example 3: Terminating a goroutine using Shared Variable](#example-3-terminating-a-goroutine-using-shared-variable)
- [Race Condition and Mutual Exclusion](#race-condition-and-mutual-exclusion)
	- [Examples](#examples-14)
		- [Example: 1 for Race Condition](#example-1-for-race-condition)
		- [Example 2: For Race Condition](#example-2-for-race-condition)
	- [Race Condition Effects](#race-condition-effects)
		- [Deadlocks](#deadlocks)
		- [Data Corruption](#data-corruption)
		- [Security Vulnerabilities](#security-vulnerabilities)
		- [Performance Degradation](#performance-degradation)
	- [Preventing Deadlocks](#preventing-deadlocks)
		- [Methods](#methods-1)
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
	- [Examples](#examples-15)
		- [Example 1: Using Atomic Operations](#example-1-using-atomic-operations)
		- [Example 2: Using Atomic Operations with WaitGroup](#example-2-using-atomic-operations-with-waitgroup)
- [Race Detector](#race-detector)
	- [How it works](#how-it-works)
- [Semaphore](#semaphore)
	- [Bounded Channel Approach](#bounded-channel-approach)
		- [Example: Using a Bounded Channel](#example-using-a-bounded-channel)
	- [GOMAXPROCS](#gomaxprocs)
		- [Syntax](#syntax-29)
	- [Context](#context)
		- [`Background` Context](#background-context)
			- [Syntax](#syntax-30)
		- [`TODO` Context](#todo-context)
			- [Syntax](#syntax-31)
	- [Weighted Semaphore](#weighted-semaphore)
		- [Weighted Semaphore Method](#weighted-semaphore-method)
			- [Syntax](#syntax-32)
		- [Aquire Semaphore Method](#aquire-semaphore-method)
			- [Syntax](#syntax-33)
		- [Release Semaphore Method](#release-semaphore-method)
			- [Syntax](#syntax-34)
		- [TryAcquire Semaphore Method](#tryacquire-semaphore-method)
			- [Syntax](#syntax-35)
		- [Example: Using a Weighted Semaphore](#example-using-a-weighted-semaphore)
- [Goroutines and Threads](#goroutines-and-threads)
	- [Difference between Goroutines and Threads](#difference-between-goroutines-and-threads)
	- [Goroutines Growable Stacks](#goroutines-growable-stacks)
	- [Goroutines Scheduling](#goroutines-scheduling)
- [Reflection](#reflection)
	- [Types and Interfaces](#types-and-interfaces)
- [Laws of Reflection](#laws-of-reflection)
	- [1. Reflection goes from interface value to reflection object](#1-reflection-goes-from-interface-value-to-reflection-object)
		- [Syntax](#syntax-36)
		- [Examples for `TypeOf`](#examples-for-typeof)
			- [Example 1: Getting the Type of a Variable](#example-1-getting-the-type-of-a-variable)
			- [Example 2: Getting the Value of a Variable](#example-2-getting-the-value-of-a-variable)
		- [Examples for `ValueOf`](#examples-for-valueof)
		- [Methods](#methods-2)
		- [Examples for `Kind`](#examples-for-kind)
			- [Example 1: Basic](#example-1-basic)
			- [Example 2: Difference btw Kind and Type](#example-2-difference-btw-kind-and-type)
	- [2. Reflection goes from reflection object to interface value](#2-reflection-goes-from-reflection-object-to-interface-value)
		- [Example](#example-26)
	- [3. To modify a reflection object, the value must be settable](#3-to-modify-a-reflection-object-the-value-must-be-settable)
		- [Settability](#settability)
			- [Example: Understanding Settablitity](#example-understanding-settablitity)
			- [Example: Settability of a Pointer](#example-settability-of-a-pointer)
			- [Syntax: Elem() Method](#syntax-elem-method)
			- [Example: Settability of a Struct Field](#example-settability-of-a-struct-field)
		- [Example: Using the CanSet() Method](#example-using-the-canset-method)
- [Why Reflection](#why-reflection)
- [Reflection in Go](#reflection-in-go)
	- [Methods](#methods-3)
		- [NumField Method](#numfield-method)
			- [Example: Using NumField](#example-using-numfield)
		- [Field Method](#field-method)
			- [Example: Using Field](#example-using-field)
		- [Copy Method](#copy-method)
			- [Example: Using Copy](#example-using-copy)
		- [DeepEqual Method](#deepequal-method)
			- [Examples](#examples-16)
				- [Example: Using DeepEqual for Structs](#example-using-deepequal-for-structs)
				- [Example: Using DeepEqual for Arrays](#example-using-deepequal-for-arrays)
				- [Example: Using DeepEqual for Functions](#example-using-deepequal-for-functions)
		- [Swapper	Method](#swappermethod)
			- [Examples](#examples-17)
				- [Example 1: Using Swapper](#example-1-using-swapper)
				- [Example 2: Reversing](#example-2-reversing)
		- [FieldByIndex Method](#fieldbyindex-method)
			- [Example: Using FieldByIndex](#example-using-fieldbyindex)
		- [FieldByName Method](#fieldbyname-method)
			- [Example: Using FieldByName](#example-using-fieldbyname)
- [Low Level Programming](#low-level-programming)
	- [Unsafe Package](#unsafe-package)
	- [Methods](#methods-4)
		- [Sizeof Method](#sizeof-method)
			- [Typical Sizes for Go](#typical-sizes-for-go)
			- [Example: Using Sizeof](#example-using-sizeof)
		- [Alignof Method](#alignof-method)
		- [Offsetof Method](#offsetof-method)
		- [Pointer Method](#pointer-method)
			- [Operations](#operations-2)
			- [Conversion](#conversion)
- [Debugging](#debugging)
	- [Syntax](#syntax-37)
	- [Common Commands](#common-commands)
	- [Delve Client](#delve-client)
		- [Starting Programs](#starting-programs)
		- [Manipulation of breakpoints](#manipulation-of-breakpoints)
		- [View program variables and memory](#view-program-variables-and-memory)
		- [List output and switch between threads and goroutines](#list-output-and-switch-between-threads-and-goroutines)
		- [View call stack and select frames](#view-call-stack-and-select-frames)
		- [Other commands:](#other-commands)
	- [Breakpoint](#breakpoint)
		- [Linespec locations](#linespec-locations)
		- [Commands](#commands)
	- [View](#view)
		- [Syntaxs](#syntaxs)
			- [Syntax: For Print](#syntax-for-print)
			- [Syntax: For Locals](#syntax-for-locals)
		- [Examples](#examples-18)
			- [Example: For Print](#example-for-print)
			- [Example: For Locals](#example-for-locals)
- [REPL](#repl)
- [Popular Third-Party Go Packages](#popular-third-party-go-packages)
	- [| **Go-Faker** | Fake data generation | `github.com/bxcodec/faker/v4` |](#-go-faker--fake-data-generation--githubcombxcodecfakerv4-)
		- [Niche or Lesser-Known Packages](#niche-or-lesser-known-packages)
	- [What Are Some Common Things Go Is Used For?](#what-are-some-common-things-go-is-used-for)
		- [Web Servers \& APIs](#web-servers--apis)
		- [Microservices](#microservices)
		- [CLI Tools](#cli-tools)
		- [Networking Tools](#networking-tools)
		- [Cloud Infrastructure](#cloud-infrastructure)
		- [DevOps \& Automation](#devops--automation)
		- [Database Clients \& Servers](#database-clients--servers)
		- [Distributed Systems](#distributed-systems)
		- [Data Processing \& ETL](#data-processing--etl)
		- [Game Development](#game-development)
		- [Web Scraping \& Automation](#web-scraping--automation)
		- [APIs for Mobile \& IoT](#apis-for-mobile--iot)
		- [Security Tools](#security-tools)
		- [Testing \& Validation](#testing--validation)

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

## Example

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

## Syntax
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



## Examples
### Example: Basic

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

# If conditions

- Same as C (almost)
- "{" cannot be at new line
- "else" cannot be on newline, needs to be "} else {" where } is from a prior if statemet  


## Syntax
```go
if condition {
	// code
} else if condition {
	// code
} else {
	// code
}
```


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

## Syntax
```go
const variable_name
```


## Example: using a constant 
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

- `+`,`-`,`*`,`/`: int, float and complex numbers
- `%`: `only applies to integers`
- Sign of remainder always the same as dividend sign
- Follows BODMAS rules

### Example 
`-5%3 and -5%-3 equal -2 `
> does not matter what is on the RHS of `%`. only LHS sign taken

## Comparison Operations

` == != < <= > >= `

- only for basic types
- `==` and `!=` can be used for strings
- strings are compared lexicographically
- strings are case sensitive
- Examples are in the Boolean topic

## Bitwise Binary Orperations 

` & | ^ &^ << >>`
| Operator | Description 				|
|----------|----------------------------|
| `&`      | Bitwise And 				|
| `\|`      | Bitwise Or  				|
| `^`      | Bitwise XOR 				|
| `&^`     | Bitwise Clear (AND NOT)	|
| `<<`     | Bitwise Left Shift 		|
| `>>`     | Bitwise Right Shift 		|

> Bitwise Clear: `a &^ b` is same as `a & (^b)`

## Logical Operations

` && || !`
| Operator | Description |
| -------- | ----------- |
| `&&`     | Logical AND |
| `\|\|`   | Logical OR  |
| `!`      | Logical NOT |

## Assignment Operations

`= += -= *= /= %= &= |= ^= <<= >>= &^=`
| Operator | Equivalent                    |
| -------- | ----------------------------- |
| `+=`     | a += b is same as a = a + b   |
| `&=`     | a &= b is same as a = a & b   |
| `<<=`    | a <<= b is same as a = a << b |
| `&^=`    | a &^= b is same as a = a &^ b |
| `%=`     | a %= b is same as a = a % b   |
| `/=`     | a /= b is same as a = a / b   |
| `\|= `   | a \|= b is same as a = a \| b |
| `^=`     | a ^= b is same as a = a ^ b   |
| `>>=`    | a >>= b is same as a = a >> b |
| `-=`     | a -= b is same as a = a - b   |
| `*=`     | a *= b is same as a = a * b   |
| `=`      | a = b is same as a = b        |


## Other Operations

` & * <-`
| Symbol | Description          |
| ------ | -------------------- |
| `&`    | Address of           |
| `*`    | Pointer to           |
| `<-`   | Channel send/receive |


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


## Syntax
```go
var variable_name int
var variable_name int8
var variable_name int16
```

## Examples
### Example: Signed Integers
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

### Example: Unsigned Integers
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

## Type Casting
- To convert one type to another
  
### Syntax
```go
var := type(variable_name)		
```

### Example: Type casting (Correct)
```go
	var i int8
	var PI = 22.0/7.0
	i = int8(PI)
	fmt.Printf("%f %d\n", PI, i)

```
> OUTPUT 3.142857 3

### Example: Type casting Problems (Incorrect Print)
```go
	var i int8
	var PI = 22/7    // PI is int now since its not 22.0/7.0
	i = int8(PI)
	fmt.Printf("%f %d\n", PI, i)

```
> OUTPUT %!f(int=3) 3

### Example: Type Casting Problems (int32 int64 addition)
```go
	var apples int32 = 1
	var oranges int64 = 2
	var ans int = apples + oranges		// Compile error
	fmt.Println(ans)
```
> OUTPUT: Error

------- 
<h3> Corrected Code </h3>

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

## Syntax
```go
var variable_name1 complex128 = complex(int_x,int_y)
var variable_name2 complex64 = complex(int_x,int_y)
```


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

## Syntax
```go
for i:=0; i<10; i++ {
	// code here
}
```

## Examples
### Example: Basic
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


## Syntax
```go
var p *int
```

## Example
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

## Syntax
```go
// Creating a struct
type StructName struct {
	field1 type1
	field2 type2
}

// Creating a variable of the struct
var variable_name StructName

// Assigning values to the struct
variable_name.field1 = value1

// Creating a struct with values
variable_name := StructName{field1: value1, field2: value2}

// Creating a struct with values without field names
variable_name := StructName{value1, value2}

// IMP
// Creating a new instance of the struct
variable_name  = &StructName{value1, value2}
```

## Example
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

## Syntax
```go
var variable_name [size]type
```

## Examples
### Example: Basic
```go
var q [3]int = [3]int{1,2,3}
var r [3]int = [3]int{1,2}
println(r[2])	// 0 is the output
```


### Example: Ellipsis `...` `(Imp)`
```go
q:= [...]int{1,2,3} // automatically sets that numebrs
printf("%T",q)	// [3]int

r:=[...]int{99:-1}		// sets a list with 99 0s and -1 at the end 
```


### Example: Comparisons `(Imp)`
```go
a := [4]int{1,2,3,4}
b := [...]int{1,2,3,4}
c := [4]int{1,2,3,5}
println(a==b)	// true
println(a==c)	// false
println(b==c)	// false
```
> OUTPUT: </br>
> true </br>
> false </br>
> false </br>

# Slice
- Varibale lenght sequence of elemetns
- has 3 componenss - pointer (to base location), lenght (no of eleemnts), capacity (lenght cannot exceed this)
- buitl in fucntions len and cap
- dynamically sized
- selects half-open range inlcues the 1st ekemetn and exclude the last
- does not store any data, just describe a section of array
- other slices with same underlying array can se

## Syntax
```go
low:= int
high:= int
a[low:high]
```

## Example
```go
prime := [...]int{1,2,3,4}
var s []int = prime[1:3]
println(s)
```

# Appends

- Appends elements to the end of a slice
- If the underlying array is too small, a bigger array is allocated
- Using `append()` function

## Syntax
```go
list_var = append(list_var, value1, value2...)
```

## Example
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
  
## Syntax
```go
var m map[key_type]value_type
m := map[key_type]value_type{key1:value1, key2:value2}

m := make(map[key_type]value_type)
m := make(map[key_type]value_type, capacity_hint)	// capacity_hint is just a hint for capacity, more elements can be added to the map
```

## Example
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
  

## Syntax
```go
func name(para_list) (return_list) {
	// code
}
```

## Examples

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

### Syntax
```go
func(para_list) (return_list) {		// notice there is not funcName here
	// code
}
```

### Example
```go
package main

func main() {
	add := func(x, y int) int {		// here we can see that there is no function name
		return x + y
	}

	println(add(1, 3))		// 4

	func(x, y int) {
		println(x + y)		// 5
	}(2, 3)		// here we can see that the function is called immediately

	func () {
		println("hello")		// here we can see that there are no parameters
	}()
}
```
> OUTPUT: </br>
> 4 </br>
> 5 </br>
> hello </br>


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
- `Printf` is an example of a varidic function
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

## Syntax
```go
defer function_name()
```

## Examples
### Example: Basic
```go
package main

func main() {
	defer println("world")
	println("hello")
}
```
`OUTPUT: hello \n world

### Example: Function call
```go
package main

func OpenDB() {
	println("Opening DB")
}

func CloseDB() {
	println("Closing DB")
}

func main() {
	defer CloseDB()
	OpenDB()
	printn("Doing something")
}
```
> OUTPUT: </br>
> Opening DB </br>
> Doing something </br>
> Closing DB </br>

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

## Syntax
```go
func panic(/*String to be displayed*/)
```


## Example
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

## Syntax
```go

func main() {
	defer func() {
		if r := recover(); r != nil {
			// Code to handle the panic
		}
	}()
}

// ----------- OR ---------------------------------------------

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

## Example
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

## Example
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

## Example
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
  
## Syntax
```go
type variable struct {
	// fields
}

func (t variable) functionName(parameter list) {	// t is the method receiver
	// code
}
```

## Examples

### Example: Basic
- In this example, the Abs method has a receiver of type Vertex named v
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

## Example
```go
func main() {
	var op1,op2 int = 10, 20
	_ = op1 + op2				// The result of the addition is ignored
}
```

<!--  **Unit 1 End** -->
<!-- This is 2025 addtion, not from the summer class era -->
## How to Start a Go Project

Starting a Go project is straightforward and follows a few key steps. Here’s a practical guide:

### 1. Create a New Project Directory

Pick a folder for your project and navigate to it:

```sh
mkdir myproject
cd myproject
```

### 2. Initialize a Go Module

This sets up your project for dependency management and proper imports.  
Use your desired module name (often a repo URL for real projects, but can be anything for local work):
> If we wish to change the module name later, we can directly edit the `go.mod` file.

```sh
go mod init myproject
```
- This creates a `go.mod` file, which tracks your module name and dependencies.

### 3. Create Your Main File

Start with a `main.go` file:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go project!")
}
```

### 4. Run Your Project

Use the Go toolchain to run your code:

```sh
go run main.go
```
### 5. Build Your Project

To build a binary:

```sh
go build
```

## Folder Structures

A well-organized folder structure makes your Go project easier to maintain, scale, and understand. Here are some general principles and examples:

### General Principles

- Keep related code together in packages (folders).
- Use clear, short, lowercase names for folders and packages.
- Separate your application entry point (`main.go`) from reusable code (in packages).
- Place tests in the same folder as the code they test, using `_test.go` suffix.
- For larger projects, consider grouping by feature or layer (e.g., `api/`, `internal/`, `pkg/`).

### Example: Simple CLI Tool

```
mycli/
├── go.mod
├── main.go
├── cmd/
│   └── root.go
├── utils/
│   └── helpers.go
```

- `main.go`: Entry point.
- `cmd/`: Command logic.
- `utils/`: Helper functions.

### Example: Microservice

```
userservice/
├── go.mod
├── main.go
├── api/
│   └── handlers.go
├── model/
│   └── user.go
├── db/
│   └── repository.go
├── internal/
│   └── auth.go
```

- `api/`: HTTP handlers.
- `model/`: Data models.
- `db/`: Database logic.
- `internal/`: Private/internal code.

### Example: Web Server

```
webapp/
├── go.mod
├── main.go
├── routes/
│   └── router.go
├── templates/
│   └── index.html
├── static/
│   └── style.css
```

- `routes/`: Routing logic.
- `templates/`: HTML templates.
- `static/`: Static assets.

### Example: Library/Package

```
mathlib/
├── go.mod
├── mathlib.go
├── utils.go
├── mathlib_test.go
```

- All code in package folder.
- Tests in same folder.

### Example: Monorepo (Multiple Services)

```
company/
├── go.mod
├── serviceA/
│   └── main.go
├── serviceB/
│   └── main.go
├── shared/
│   └── logger.go
```

- Multiple services in one repo.
- Shared code in `shared/`.

### Example: Small Script

```
quicktool/
├── go.mod
├── main.go
```

- Just a single file.

### Example: Large Enterprise App

```
bigapp/
├── go.mod
├── main.go
├── api/
├── domain/
├── service/
├── repository/
├── internal/
├── pkg/
```

- Grouped by layer or feature.

### Example: Test-Heavy Project

```
testproj/
├── go.mod
├── main.go
├── core/
│   ├── core.go
│   └── core_test.go
├── mocks/
│   └── mock_db.go
```

- Tests and mocks close to code.





# Pacakges
- A package is a collection of source files in the same directory that are compiled together
  
## Importing Standard Library Packages

Go comes with a rich standard library. To use a package, import it at the top of your  file using the `import` keyword.

### Syntax

```go
import "package_name"
```

You can import multiple packages using parentheses:

```go
import (
    "fmt"
    "math"
)
```

### Example: Using `fmt` and `math`

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Square root of 16 is", math.Sqrt(16))
}
```
> OUTPUT: Square root of 16 is 4

### Notes

- Only imported packages that are used in your code are allowed; unused imports cause a compile error.
- Standard library packages do not require installation—just import and use.
 
## Installing and Importing Third-Party Packages

Go makes it easy to use external libraries from sources like GitHub. You install packages using the command line, and then import them in your code.

### Installing a Package

Use `go get` to download and install a package:

```sh
go get github.com/gorilla/mux
```

This will:
- Download the package and its dependencies.
- Add it to your project's `go.mod` file (if you are using Go modules).

### Importing a Third-Party Package

After installation, import the package at the top of your `.go` file:

```go
import "github.com/gorilla/mux"
```

### Example: Using an External Package

```go
package main

import (
    "fmt"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    fmt.Println("Router created:", r)
}
```
> OUTPUT: Router created: &{...}

### Notes

- Always run `go mod init <module-name>` in your project folder before installing packages (if not already done).
- Use `go get -u <package>` to update a package.
- You can find packages at [pkg.go.dev](https://pkg.go.dev/) or GitHub.

## Creating Your Own Packages

We can create our own packages, letting you organize related functions, types, and variables into packages. A package is simply a directory containing one or more files with the same `package` name.

### Basic Structure

- Each `.go` file starts with a `package` declaration.
- The directory name is usually the package name.

**Example: Creating a simple package**

Suppose you want a package called `mathutils`:

```
project/
│
├── main.go
└── mathutils/
    └── mathutils.go
```

**mathutils/mathutils.go**
```go
package mathutils

func Add(a, b int) int {
    return a + b
}

func Sub(a, b int) int {
    return a - b
}
```

**main.go**
```go
package main

import (
    "fmt"
    "project/mathutils"
)

func main() {
    fmt.Println(mathutils.Add(2, 3)) // 5
    fmt.Println(mathutils.Sub(5, 2)) // 3
}
```

### Notes

- The package name in the import path is the folder name.
- All files in a folder must use the same `package` name.
- Only exported identifiers (starting with a capital letter) are accessible from other packages.
- Use Go modules (`go mod init <module-name>`) for proper import paths.

## Project Structure and Multiple Files

We can organize our code into multiple files and folders for better maintainability.

### Typical Project Layout

```
project/
│
├── go.mod
├── main.go
├── mathutils/
│   └── mathutils.go
├── helpers/
│   └── helpers.go
```

- `go.mod`: Defines your module and manages dependencies.
- `main.go`: Entry point of your application.
- Subfolders (like `mathutils`, `helpers`): Contain your own packages.

### How Multiple Files Work Together

- All `.go` files in the same folder and with the same `package` name are compiled together.
- You can split code logically (e.g., types in one file, functions in another).

**Example: Splitting code in a package**

```
mathutils/
├── add.go
├── sub.go
```

**add.go**
```go
package mathutils

func Add(a, b int) int {
    return a + b
}
```

**sub.go**
```go
package mathutils

func Sub(a, b int) int {
    return a - b
}
```

Both files are part of the `mathutils` package and can access each other's functions.

## Importing and Using Local Packages

Once you've created your own package (e.g., `mathutils`), you can import and use it in other files within your project.

### Import Path

- The import path is based on your module name (from `go.mod`) and the folder structure.
- If your module name is `project`, and your package is in the `mathutils` folder, import it as:
  ```go
  import "project/mathutils"
  ```

### Example: Using a Local Package

**main.go**
```go
package main

import (
    "fmt"
    "project/mathutils"
)

func main() {
    fmt.Println(mathutils.Add(2, 3)) // 5
    fmt.Println(mathutils.Sub(5, 2)) // 3
}
```

**mathutils/mathutils.go**
```go
package mathutils

func Add(a, b int) int {
    return a + b
}

func Sub(a, b int) int {
    return a - b
}
```

### Notes

- The package name in the import path matches the folder name.
- Only exported functions (starting with a capital letter) are accessible outside the package.
- If you change your module name in `go.mod`, update your import paths accordingly.
- Use `go build` or `go run main.go` to compile and run your project.

- Use folders to organize related code.
- All files in a folder must use the same package name.
- Only exported identifiers (starting with a capital letter) are accessible outside the package. (Public)
- Ones that are with lowercase are private to the package. (Private)
- Examples:
 ```go
package mathutils
func Add(a, b int) int { // Public function
	return a + b
}

func Subtract(a, b int) int { // Public function
	return a - b
}	
```
## Best Practices for Packages and Project Structure

Following best practices helps keep your Go code clean, maintainable, and idiomatic.

### 1. Use Go Modules

- Always initialize your project with Go modules:
  ```sh
  go mod init <module-name>
  ```
- This makes dependency management and imports reliable.

### 2. Name Packages Clearly

- Use short, meaningful, lowercase names for packages (e.g., `utils`, `mathutils`).
- Avoid underscores and hyphens.

### 3. Export Only What’s Needed

- Only capitalize functions, types, or variables that should be accessible outside the package.
- Keep internal details unexported (lowercase).

### 4. Keep Package Scope Small

- Group related code together.
- Avoid huge packages; split logically (e.g., `mathutils`, `helpers`).

### 5. Use Proper Import Paths

- Import local packages using your module name (from `go.mod`).
- Update imports if you rename your module.

### 6. Document Your Packages

- Add comments above exported functions/types.
- Use `go doc` to view documentation.

### 7. Keep `main` Package Simple

- The `main` package should only contain the entry point and high-level orchestration.
- Move logic to other packages.

### 8. Use `go fmt` and `go vet`

- Format your code:
  ```sh
  go fmt ./...
  ```
- Check for common mistakes:
  ```sh
  go vet ./...
  ```

### 9. Avoid Circular Imports

- Packages should not import each other in a cycle.

### 10. Use Versioning for Dependencies

- Specify versions for third-party packages:
  ```sh
  go get github.com/gorilla/mux@v1.8.0
  ```

## Useful Go Commands Reference

Here are some essential Go commands for working with packages, modules, and projects from the command line:

### Module and Dependency Management

- **Initialize a new module:**
  ```sh
  go mod init <module-name>
  ```
- **Download and install a package:**
  ```sh
  go get <package-path>
  ```
- **Update a package:**
  ```sh
  go get -u <package-path>
  ```
- **Tidy up dependencies (remove unused, add missing):**
  ```sh
  go mod tidy
  ```
- **List all dependencies:**
  ```sh
  go list -m all
  ```

### Building and Running

- **Build your project:**
  ```sh
  go build
  ```
- **Run your main file:**
  ```sh
  go run main.go
  ```
- **Run all files in the current folder:**
  ```sh
  go run .
  ```

### Formatting and Linting

- **Format your code:**
  ```sh
  go fmt ./...
  ```
- **Check for common mistakes:**
  ```sh
  go vet ./...
  ```

### Testing

- **Run tests:**
  ```sh
  go test
  ```
- **Run tests with verbose output:**
  ```sh
  go test -v
  ```

### Other Useful Commands

- **Show documentation for a package:**
  ```sh
  go doc <package>
  ```
- **Show environment info:**
  ```sh
  go env
  ```
- **Show Go version:**
  ```sh
  go version
  ```

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
- Goroutines are nonpreemptive 

## Syntax
```go
go function_name()
```

## Examples
### Example- Basic
We run this with the `go` keyword and without the `go` keyword, we can see that the go routine runs in the background and prints as soon as it can while Normal function in its own way making it interleaved with a pattern
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


## Advantages of Goroutines
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
| Term      | Example    | Description                                                             |
| --------- | ---------- | ----------------------------------------------------------------------- |
| `send`    | `ch <- v`  | Sends a value from one goroutine through a channel to another goroutine |
| `receive` | `v = <-ch` | Receives a value from the channel and assigns it to a variable          |
- `<-` is a receive operator where the result of the operation is discarded

## Syntax
```go
var ch chan int
ch = make(chan int)

// Using in a function
func goroutine(variable_name chan int) {
	// code
}
```

## Examples 
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

This is because we first put x into the channel using the first goroutine. Then we try to get y from the channel, but there is no value in the channel, so it blocks until one of the other goroutines puts a value into the channel. The order of execution of the goroutines is not guaranteed, so we can get either 2 or 3 as y. Finally, we get z from the channel, which will be the remaining value. This is because channels are FIFO (First In First Out) data structures.
  
> OLD EXPLANATION: (probably wrong) </br>
> Here what happens is that `x` gets 1 as it is right after the function and `z` gets 2 as it is right after the line where the function is called. `y` gets 3 as it is the last value in the channel (My Guess as thats what my experimentation shows me)

## Advantages of Channels
- Channels are typesafe. They only allow a particular type of data to be sent through them	
- They enable safe communication between goroutines and synchronisation
- They ensure the code runs predictably and efficiently with concurrent.
- **Blocking Send and Recive:** If a goroutine tries to send to a channel that is full, it will block until another goroutine receives from the channel. If a goroutine tries to receive from an empty channel, it will block until another goroutine sends to the channel
- **Zero Value Channel:** The zero value of a channel is `nil`. A nil channel is of no use and attempting to send or receive from it will cause a deadlock
- **For Loop Channel:** A loop can iterate over the sequntial values sent by a channel until it is closed. This is useful when the number of values is not known in advance

## Properties of Channels
- **Lenght of Channel:** This can be found using the `len` function. This is the number of elements **currently** buffered in the channel
- **Capacity of Channels:** This can be found using the `cap` function. This is the number of elements that **can be** buffered in the channel


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

### Examples
#### Example: Basic
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

#### Example: Adding more values than the capacity
```go
func main(){
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3		// <-- ERROR: This will block as it is the 3rd send when capacity is 2
}
```
> OUTPUT: fatal error: all goroutines are asleep - deadlock!

#### Example: Not sending any values
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

## Syntax
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

## Example
```go
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for { // since this select is in an infinite for loop in a go func, it will keep listening forever
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
> From ch2: 9 </br>

Here the ch1 or ch2 is selected at random as the select statement chooses one at random if multiple are ready
The order of the output will vary

# Cancellation
- There is no way to directly cancel a goroutine as it would leave the state of the program in an unknown state
- We can use a `cancellation` channel to signal the goroutine to stop
- We also defina a utility function to check if the channel is closed
- In general it is hard to know how many goroutines are running at any given time

## Method 1: Using a Cancellation Channel
- We create a channel on which no values are sent but whose closure signals the goroutine to stop

### Example
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

### Example:
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
	time.Sleep(1 * time.Second)
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


# Constants

## Iota
- This is a predeclared identifier
- Basically a counter for constants that increases by 1 for each line
- Always starts at 0

## Example: `iota`
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

## Syntax
```go
type interface_name interface {
	Method() signatures
}
```

## Example
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


func main() {
	r := rect{width: 10, height: 5}
	c := circle{radius: 5}

	var g Geometry
	g = r				// print(g) --> {10 5}
	g.area()			// Area of rectangle is: 50
	g = c				// print(g) --> {5}
	g.area()			// Area of circle is: 78.50
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

### Examples
#### Example: Simple usage
```go
var any interface{}
any = 42
any = "hi"
any = 3.14159
any = new(bytes.Buffer)
any = make(chan int, 3)
any = make(map[string]int)
```

#### Example: Proper an empty interface
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

## Interface as a Contract (NOT COM)
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

## Syntax
```go

val, ok := interface_var(type)
// val is the value of the variable
// ok is a boolean value that is true if the assertion holds
```

## Examples
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



# Concurrancy with WaitGroup
- To wait for multiple goroutines to finish, we can use a `WaitGroup`

## Methods
- `Add(int)` method is used to add the number of goroutines to wait for
- `Done()` method is used to signal that a goroutine is done. Decrements the counter by 1
- `Wait()` method is used to block until the counter is 0

> NOTE: You need to import the `sync` package to use the `WaitGroup`. The Counter starts at 0

## Syntax
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
	defer wg.Done()
	// code
}()

wg.Wait()
```

## Examples
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

Note: the order here is 2 , 1 , 0 as the goroutines run in parallel and the order is not fixed so its basically going to be random

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
		time.Sleep(1 * time.Second)		// So it checks once every second

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

<!-- **Unit 2 End** -->

# Race Condition and Mutual Exclusion

- Race condition is when the output of a program depends on the order of execution of the goroutines
- Occurs when multiple goroutines access the same resource
- This can lead to unpredictable results

- Mutual Exclusion is a technique to ensure that only one goroutine can access a resource at a time
- This is done using a `mutex` or `lock`

## Examples
### Example: 1 for Race Condition
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

Here we can see that the balance is updated by the goroutines so they are no longer guaranteed to give the right answer
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
	
	println(data)	// A2
}
```
> OUTPUT: </br>
> 0

Here data may be 1 or 0 depending on if A1 or A2 runs first

## Race Condition Effects	
### Deadlocks
- Can lead to a deadlock
- Happens when two or more goroutines are waiting for each other to release a resource
### Data Corruption
- Can lead to data corruption when multiple goroutines write to the same resource

### Security Vulnerabilities
- Can lead to security vulnerabilities where an attack gains unauthorized access to a resource

### Performance Degradation
- Can lead to performance degradation when multiple goroutines are waiting for a resource


## Preventing Deadlocks		

- Avoiding direct access to shared variables
- "Do not communicate by sharing memory; instead, share memory by communicating" - Rob Pike
- Use a Channel to send a request a query/update a variable to goroutine 
- A goroutine that manages access called a `monitor` goroutine

### Methods
- Using a channel to send a request to a goroutine
- Use locks or semaphores to manage access to shared resources 
- Atomic operations to ensure that operations on shared variables are atomic
- Avoid global variables as they can be accessed by multiple goroutines


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

## Example: Using a Binary Semaphore
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

## Syntax: Creation
```go
var mu sync.Mutex
mu := sync.Mutex{}
```

## Syntax: Locking and Unlocking
```go
//func (m *Mutex) Lock()		// Locks m
//func (m *Mutex) Unlock()	// Unlocks m

mu.Lock()
// code
mu.Unlock()

```


## Example: Using a Mutex
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
- Needs to import the `sync/atomic` package
  
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
- Atomically swap the value of a variable
- Eg: `atomic.SwapInt32(&x, 42)` swaps the value of x with 42 (type is int32)

### Compare and Swap
- Atomically compare and swap the value of a variable
- If the current value of the variable is equal to the old value, the new value is set
- Used to implement lock-free algorithms
  
## Examples
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

## GOMAXPROCS

- The `GOMAXPROCS` environment variable limits the number of OS threads that can execute Go code simultaneously
- The default value is the number of CPUs on the machine. If n < 1, it does not change the current setting
- Using this we can optimize the concurancy and parallelism.
- requires the `runtime` package to be imported

### Syntax
```go
var n int = runtime.GOMAXPROCS(n)
```	

### Example
```go
package main

import "runtime"
func main() {
	curVal := runtime.GOMAXPROCS(0)
	println(curVal)

	maxCore := runtime.NumCPU()
	println(maxCore)
	new := maxCore - 2
	runtime.GOMAXPROCS(new)
	curVal = runtime.GOMAXPROCS(0)
	println(curVal)
}
```
> OUTPUT: </br>
> 12 </br>
> 12 </br>
> 10 </br>

## Context
- The `context` package provides a way to pass cancellation signals, deadlines, and other request-scoped values across API boundaries and between processes
- There are two functions

### `Background` Context
- This is the root of the context tree
- Used when there is no parent context
- Retruns a non-nil, empty Context. 
- It is never cancelled, has no values, and has no deadline
- Typically used for main, init, and tests, and as the top-level context for incoming requests

#### Syntax
```go
ctx := context.Background()
```

### `TODO` Context
- This is used when a context is needed but there isn't one available
- It is used when a function requires a context but the caller doesn't have one
- It is used to avoid nil pointers

#### Syntax
```go
ctx := context.TODO()
```

## Weighted Semaphore
- A weighted semaphore is a semaphore that allows multiple units to be acquired or released at once
- requires the `golang.org/x/sync/semaphore` package to be imported

### Weighted Semaphore Method
- `func NewWeighted(n int) *Weighted`
- This creates a new weighted semaphore with the given initial value

#### Syntax
```go
sem := semaphore.NewWeighted(1)	// creates a new weighted semaphore with an initial value of 1
```

### Aquire Semaphore Method
- `func (s *Weighted) Aquire(ctx context.Context, n int) error`
- This acquires n units from the semaphore, blocking until they are available or ctx is done
- On success, returns nil.
- On failure, returns ctx.Err()
- If `ctx` already done, Acquire may still succeed without blocking

#### Syntax
```go
sem := semaphore.NewWeighted(1)
if err := sem.Acquire(ctx, 1); err != nil {
	// handle error
}
```


### Release Semaphore Method
- `func (s *Weighted) Release(n int)`
- This releases n units back to the semaphore
- It panics if n is negative or if n exceeds the number of units held by the semaphore

#### Syntax
```go
sem := semaphore.NewWeighted(1)
sem.Release(1)
```


### TryAcquire Semaphore Method 
- `func (s *Weighted) TryAcquire(n int) bool`
- This tries to acquire n units from the semaphore without blocking
- On success, returns true
- On failure, returns false

#### Syntax
```go
sem := semaphore.NewWeighted(1)
if sem.TryAcquire(1) {
	// success
} else {
	// failure
}
```


### Example: Using a Weighted Semaphore
```go
package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	pool := semaphore.NewWeighted(2)
	go swim("A", pool)
	go swim("B", pool)
	go swim("C", pool)
	go swim("D", pool)

	// For brevity, better use sync.WaitGroup
	time.Sleep(5 * time.Second)
	log.Println("Main: Done, shutting down")
}

func swim(name string, pool *semaphore.Weighted) {
	log.Printf("%v: I want to swim\n", name)

	// In real applications, pass in your context such as HTTP request context
	ctx := context.Background()

	// We can also Acquire/Release more than 1 when the workloads consume different amount of resources
	if err := pool.Acquire(ctx, 1); err != nil {
		log.Printf("%v: Ops, something went wrong! I cannot acquire a lane\n", name)
		return
	}

	log.Printf("%v: I got a lane, I'm swimming\n", name)
	time.Sleep(time.Second)
	log.Printf("%v: I'm done. Releasing my lane\n", name)
	pool.Release(1)
}
```
> 2024/06/06 15:27:45 D: I want to swim </br>
> 2024/06/06 15:27:45 D: I got a lane, I'm swimming </br>
> 2024/06/06 15:27:45 A: I want to swim </br>
> 2024/06/06 15:27:45 A: I got a lane, I'm swimming </br>
> 2024/06/06 15:27:45 B: I want to swim </br>
> 2024/06/06 15:27:45 C: I want to swim </br>
> 2024/06/06 15:27:46 A: I'm done. Releasing my lane </br>
> 2024/06/06 15:27:46 B: I got a lane, I'm swimming </br>
> 2024/06/06 15:27:46 D: I'm done. Releasing my lane </br>
> 2024/06/06 15:27:46 C: I got a lane, I'm swimming </br>
> 2024/06/06 15:27:47 B: I'm done. Releasing my lane </br>
> 2024/06/06 15:27:47 C: I'm done. Releasing my lane </br>
> 2024/06/06 15:27:50 Main: Done, shutting down

# Goroutines and Threads

## Difference between Goroutines and Threads
| Category | Goroutines | Threads |
|----------|------------|---------|
| Stack Management | <ul><li>Can grow and shrink their stack as needed</li></ul> | <ul><li>Have a fixed stack size</li></ul> |
| Scheduling | <ul><li>Scheduled by the Go runtime</li><li>Multiplexed onto a small number of OS threads</li><li>Lightweight</li></ul> | <ul><li>Scheduled by the OS</li><li>Scheduled by the OS kernel</li><li>Heavyweight comparatively</li></ul> |
| Concurrency Model | <ul><li>Use a `communicating sequential processes` (CSP) model</li></ul> | <ul><li>Use a `shared memory` model</li></ul> |
| Performance | <ul><li>Faster to create and manage</li><li>More efficient</li><li>Enable high concurrency and scalability</li></ul> | <ul><li>Slower to create and manage</li><li>Less efficient</li><li>Enable low concurrency and scalability comparatively</li></ul> |

## Goroutines Growable Stacks
- Goroutines have growable stacks
- The stack starts at 2KB and can grow to 1GB
- It is like an OS thread but more lightweight as it is not always 2MB

## Goroutines Scheduling
- Go scheduler does not periodically interrupt by a hardware timer like OS threads. Instead done implicitly by certain Go constructs
- Because it does not need to switch to kernel context, it is faster than OS threads

# Reflection
- Ability of a program to inspect variables and values at runtime for their type
  - Can create, modify, and delete variables, functions, and structs
- Reflection in computing is the ability to examine its own structure  at runtime
- `reflect` package is used for reflection in Go
- Three main concepts
  - *Type*
  - *Value*
  - *Kind*

## Types and Interfaces
- Every Variables is statically typed and fixed at compile time

- `Type` is the representation of a Go type at runtime
```go
type MyInt int
var i int
var j MyInt
```
Here `i` and `j` are of different types so they cannot be assigned to each other without conversion

- An interface variable can store any concrete (non-interface) value as long as that value implements the interface’s methods.

- A well-known pair of examples is `io.Reader` and `io.Writer`. Any type that implements a Read (or Write) method with this signature is said to implement `io.Reader` (or `io.Writer`).

- An empty interface `interface{}` can store any value of any type. **But Go’s interfaces are still statically typed as even if they hold a value of any type, The value will always satisfy the interface**

# Laws of Reflection
## 1. Reflection goes from interface value to reflection object
- At basic level, reflection is just a way to get the type and value of an interface variable
- We need to know two concepts
  - `Type` --> `reflect.TypeOf`
    - Returns a `reflect.Type` object that represents the type of the interface variable
  - `Value` --> `reflect.ValueOf` (also it's easy to get the type from here too)
    - Returns a `reflect.Value` object that holds the value of the interface variable

### Syntax
```go
// func TypeOf(interfaceVar interface{}) Type

import "reflect"
i:= reflect.TypeOf(interface_var
j:= reflect.ValueOf(interface_var)


```

- `TypeOf` retruns reflection `Type` that returns dynamic type of the interface variable. If interface variable is `nil`, it returns `nil`
- `ValueOf` returns reflection `Value` that holds the value of the interface variable. If the interface variable is `nil`, it returns a `zero value`

### Examples for `TypeOf`

#### Example 1: Getting the Type of a Variable
```go
	import "reflect"
	func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	}
```
> OUTPUT: type: float64

Here it may look like there is no interface, but the signature of `reflect.TypeOf` includes an empty interface, so it can take any type of value. Basically when `x` is passed to `reflect.TypeOf`, it is stored in an empty interface then passed as an argument. The empty interface is later unpacked to get the type of `x`


#### Example 2: Getting the Value of a Variable
```go
v1:= 3.4
println(reflect.TypeOf(v1)) // float64
v2:= "Hello"
println(reflect.TypeOf(v2)) // string
v3:= true
println(reflect.TypeOf(v3)) // bool
v4:= 10
println(reflect.TypeOf(v4)) // int
v5:= []int{1,2,3}
println(reflect.TypeOf(v5)) // []int
v6:= [5]int{1,2,3,4,5}
println(reflect.TypeOf(v6)) // [5]int
v7:= map[string]int{"a":1, "b":2}
println(reflect.TypeOf(v7)) // map[string]int
```
> OUTPUT: </br>
> 3.4 </br>
> string </br>
> bool </br> 
> int </br>
> []int </br>
> [5]int </br>
> map[string]int </br>

### Examples for `ValueOf`
```go
	v1:= 3.4
	fmt.Println(reflect.ValueOf(v1)) // 3.4
	v2:= "Hello"
	fmt.Println(reflect.ValueOf(v2)) // Hello
	v3:= true
	fmt.Println(reflect.ValueOf(v3)) // true
	v4:= 10
	fmt.Println(reflect.ValueOf(v4)) // 10
	v5:= []int{1,2,3}
	fmt.Println(reflect.ValueOf(v5)) // [1 2 3]
	v6:= [5]int{1,2,3,4,5}
	fmt.Println(reflect.ValueOf(v6)) // [1 2 3 4 5]
	v7:= map[string]int{"a":1, "b":2}
	fmt.Println(reflect.ValueOf(v7)) // map[a:1 b:2]
```
> OUTPUT: </br>
> 3.4 </br>
> Hello </br>
> true </br>
> 10 </br>
> [1 2 3] </br>
> [1 2 3 4 5] </br>
> map[a:1 b:2] </br>

### Methods
- `Value` also has a `Type` method that returns the `Type` of the value 
- Both `Type` and `Value` have a `Kind` method that returns a constant indication what sort of item is stored. Eg: `Uint`, `Float64`, `Slice`, `Map`, etc.
 ```go
	func (v Value) Kind() Kind
```
- Kind returns the v's Kind
- If v is the zero Value (IsValid returns false), Kind returns Invalid

### Examples for `Kind`

#### Example 1: Basic
```go
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("value:", v)			// value: 3.4
fmt.Println("type:", v.Type())		// type: float64
fmt.Println("kind is float64:", v.Kind() == reflect.Float64)	// kind is float64: true
fmt.Println("value:", v.Float())	// value: 3.4
//fmt.Println("value:", v.Int()) 	// panic: reflect as v is float64 not int
```
> OUTPUT: </br>
> value: 3.4 </br>
> type: float64 </br>
> kind is float64: true </br>
> value: 3.4 </br>

#### Example 2: Difference btw Kind and Type
```go
type MyInt int
var x MyInt = 7
v := reflect.ValueOf(x)
fmt.Println("value:", v)			// value: 7
fmt.Println("type:", v.Type())		// type: main.MyInt
fmt.Printf("kind: %v\n", v.Kind())	// kind: int
```
> OUTPUT: </br>
> value: 7 </br>
> type: main.MyInt </br>

Here the `Kind` of `v` is still `int` as `MyInt` is an alias for `int`.
This means `Kind` cannot differentiate between `int` and `MyInt` while `Type` can.

## 2. Reflection goes from reflection object to interface value
- Reflection in Go can generate its own inverse
- Inverse of `reflect.ValueOf` is `Value.Interface`
	- Given `reflect.Value` object, we can recover the interface value it represents
- Inverse of `reflect.TypeOf` is `reflect.New` and `reflect.Zero` {copilot answer}

### Example
```go
	var x float64 = 3.4
	v := reflect.ValueOf(x)			// v is a reflect.Value
	fmt.Println("value:", v)		// "value: 3.4"
	// y will have type float64.
	y := v.Interface().(float64)	// y is an interface{}.
	fmt.Println("y:", y)			// "y: 3.4"
```
> OUTPUT: </br>
> value: 3.4 </br>
> y: 3.4 </br>

This prints the `float64` value represented by reflection object `v`

- In short  the Interface method is the inverse of the ValueOf function, except that its result is always of static type `interface{}`.
> Reiterating: Reflection goes from interface values to reflection objects and back
again.


## 3. To modify a reflection object, the value must be settable
- A reflection object is settable if it is addressable and was not obtained by the use of unexported struct fields
```go
var x float64 = 3.4
v := reflect.ValueOf(x)
v.SetFloat(7.1) // Error: will panic.
```
> OUTPUT: panic: reflect.Value.SetFloat using unaddressable value
The Problem here isnt that `7.1` is not addressable; But that `v` is not settable

### Settability
- **Settability** is a property of a reflection value, and **not all** of them have this.
- It is like addressability but stricter. It is the Property of a reflection value that it can be modified
- It is determined by whether the reflection value holds the original value
#### Example: Understanding Settablitity
```go
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("settability of v:", v.CanSet())	// false
```
> OUTPUT: settability of v: false

Here we pass a copy of x to `reflect.ValueOf` so the interface value created is not original.
Hence if 
```go
v.SetFloat(7.1)
```
was allowed to succeed, it would change the copy of `x` and not the original `x` which is confusing and illegal, so *settability* protects against this

#### Example: Settability of a Pointer
- To Modify a reflection object, we must give a pointer to the value
```go
var x float64 = 3.4
p := reflect.ValueOf(&x) // Note: take the address of x.
fmt.Println("type of p:", p.Type())
fmt.Println("settability of p:", p.CanSet())	// false
```
> OUTPUT: </br>
> type of p: *float64 </br>
> settability of p: false </br>

The *settability* of `p` is false, but that is not the value we want to update. To get to what `p` points to we need to use the `Elem()` method of `Value` (`v` in this case)
Therefore
```go
v := p.Elem()
fmt.Println("settability of v:",v.CanSet())	// now true
```
> OUTPUT: settability of v: true

#### Syntax: Elem() Method
```go
// func (v Value) Elem() Value
p := reflect.ValueOf(&x)
v := p.Elem()
```
- Returns the value that the interface v contains or that the pointer v points to.
- If `v` is neither an interface nor a pointer it panics
- If `v` nil, it returns a zero Value

Now we can update `x` like
```go
v.SetFloat(7.1)
fmt.Println(v.Interface())	// 7.1
fmt.Println(x)				// 7.1
```
> OUTPUT: </br>
> 7.1 </br>
> 7.1 </br>

> <h2>Thus, This is how we can modify a reflection object</h2>

#### Example: Settability of a Struct Field
```go
type struct_T struct {
	val1 int
	val2 string
}

func main() {
	s := struct_T{23, "hello"}
	// v := reflect.ValueOf(s)
	// fmt.Println("settability of v:", v.CanSet())	// false

	v := reflect.ValueOf(&s).Elem()
	fmt.Println("settability of v:", v.CanSet())	// true

	f := v.Field(0)
	fmt.PrintF("%s %s = %v\n", s.Type().Field(0).Name, f.Type(), f.interface())	// val1 int = 23
}
```
> OUTPUT: </br>
> settability of v: true </br>
> val1 int = 23 </br>


### Example: Using the CanSet() Method
- `CanSet` method is used to check if a reflection value is settable
- Using `Set` method on a non-settable value will panic
```go
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("settability of v:", v.CanSet())	// false
```
> OUTPUT: settability of v: false

# Why Reflection
- When we need to create a function to deal with values of different types uniformily. This can happen when
  - Values dont have a common interface
  - Dont have known representation
  - Dont exist at time of design
  - or all three
- Example: `fmt.Println` can print any type of value, even user-defined types
- Often times data in empty interface are structs and reflection is used to extract the fields
- Such problems are solved during **runtime using reflection**



# Reflection in Go
- Done using the `reflect` package
- Has two types
  - `TypeOf` accept any interface{} and returns a `Type`
  - `ValueOf` accepts any interface{} and returns a `Value`. `reflect.Value` can hold any value

## Methods
### NumField Method
- Returns the number of fields in a struct
- If argument not a struct, it panics
`func (v Value) NumField() int`

#### Example: Using NumField
```go
type struct_T struct {
	val1 int
	val2 string
}

func main() {
	s := struct_T{23, "hello"}
	v := reflect.ValueOf(s)
	fmt.Println("Number of fields:", v.NumField())	// 2
}
```
> OUTPUT: Number of fields: 2

### Field Method
- Returns the `i`th field of a struct
`func (v Value) Field(i int) Value`

#### Example: Using Field
```go
type struct_T struct {
	val1 int
	val2 string
}

func main() {
	s := struct_T{23, "hello"}
	v := reflect.ValueOf(s)
	f := v.Field(0)
	fmt.Println("Field 0:", f)	// 23
}
```
> OUTPUT: 23

### Copy Method
- Copies the value from source to destination until 
  - source exhausts 
  - destination is full
`func Copy(dst, src Value) int`

#### Example: Using Copy
```go
	destination := reflect.ValueOf([]string{"A", "B", "C"})
	source := reflect.ValueOf([]string{"D", "E", "F"})

	fmt.Println("destination: ", destination)		// [A B C]
	
	counter := reflect.Copy(destination, source)
	fmt.Println("Count =", counter)					// 3
	fmt.Println("source: ", source)					// [D E F]
	fmt.Println("destination: ", destination)		// [D E F]
```
> OUTPUT: </br>
> destination: [A B C] </br>
> Count = 3 </br>
> source: [D E F] </br>
> destination: [D E F] </br>


### DeepEqual Method
- Returns `true` or `false` if two values are deeply equal
- **Struct** values are deeply equal if their corresponding fields are deeply equal, both exported and unexported
  - Exported fields are compared by name
  - Unexported fields are compared by type
- **Array** values are deeply equal if their corresponding elements are deeply equal
- **Function** values are deeply equal if both are nil; otherwise they are not deeply equal
- **Interface** values are deeply equal if they hold deeply equal concrete values
- Values like `int`, `float64`, `bool`, etc are deeply equal if they are equal using the `==` operator

#### Examples
##### Example: Using DeepEqual for Structs
```go
type struct_T struct {
	val1 int
	val2 string
}

func main() {
	s1 := struct_T{23, "hello"}
	s2 := struct_T{23, "hello"}
	fmt.Println("DeepEqual:", reflect.DeepEqual(s1, s2))	// true
}
```
> OUTPUT: DeepEqual: true

##### Example: Using DeepEqual for Arrays
```go
s1 := [3]int{1, 2, 3}
s2 := [3]int{1, 2, 3}
s3 := [3]int{1, 2, 4}
fmt.Println("DeepEqual:", reflect.DeepEqual(s1, s2))	// true
fmt.Println("DeepEqual:", reflect.DeepEqual(s1, s3))	// false
```
> OUTPUT: </br>
> DeepEqual: true </br>
> DeepEqual: false </br>

##### Example: Using DeepEqual for Functions
```go
f1 := func() {}
f2 := func() {}
fmt.Println("DeepEqual:", reflect.DeepEqual(f1, f2))	// false

// for true
f1 = nil
f2 = nil
fmt.Println("DeepEqual:", reflect.DeepEqual(f1, f2))	// true
```
> OUTPUT: </br>
> DeepEqual: false
> DeepEqual: true

### Swapper	Method
- To swap values in a slice
- Can be used to sort or reverse a slice

#### Examples
##### Example 1: Using Swapper
```go
	s := []int{1, 2, 3, 4, 5}
	swapper := reflect.Swapper(s)
	swapper(0, len(s)-1)
	fmt.Println(s)	// 5 and 1 are swapped
```
> OUTPUT: [5 2 3 4 1]

##### Example 2: Reversing
```go
theList := []int{1, 2, 3, 4, 5}
swap := reflect.Swapper(theList)
fmt.Printf("Before Reverse :%v\n", theList)

// Reversing a slice using Swapper() function
for i := 0; i < len(theList)/2; i++ {
	pos := len(theList) - i - 1
	swap(i, pos)
}
fmt.Printf("After Reverse :%v\n", theList)
```
> OUTPUT: </br>
> Before Reverse :[1 2 3 4 5] </br>
> After Reverse :[5 4 3 2 1]


### FieldByIndex Method
- To get a nested field of a struct
`func (reflect.Type).FieldByIndex(index []int) reflect.StructField`

#### Example: Using FieldByIndex
```go
type FirstStct struct {
	A int
	B string
	C float64
}
type SecondStct struct {
	FirstStct
	D bool
}
func main() {
	s := SecondStct{FirstStct: FirstStct{10, "ABCD", 15.20}, D: true}
	t := reflect.TypeOf(s)
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0}))		// FirstStct (SecondStct.child_1)
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))	// A (FirstStct.child_1)
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))	// B (FirstStct.child_2)
	fmt.Printf("%#v\n", t.FieldByIndex([]int{1}))		// D (SecondStct.child_2)

}

```
> OUTPUT </br>
> reflect.StructField{Name:"FirstStct", PkgPath:"", Type:(*reflect.rtype)(0x61d4a0), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true} </br>
> reflect.StructField{Name:"A", PkgPath:"", Type:(*reflect.rtype)(0x614040), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:false} </br>
> reflect.StructField{Name:"B", PkgPath:"", Type:(*reflect.rtype)(0x613e00), Tag:"", Offset:0x8, Index:[]int{1}, Anonymous:false} </br>
> reflect.StructField{Name:"D", PkgPath:"", Type:(*reflect.rtype)(0x614200), Tag:"", Offset:0x20, Index:[]int{1}, Anonymous:false} </br>


### FieldByName Method
- To get a field of a struct by name

#### Example: Using FieldByName
```go
type struct_T struct {
	val1 int
	val2 string
}

func main() {
	s := struct_T{23, "hello"}
	v1 := reflect.ValueOf(s)
	v2 := reflect.ValueOf(&s).Elem()

	fmt.Println("Field 0:", v1.FieldByName("val1"))	// 23
	fmt.Println("Field 1:", v2.FieldByName("val2"))	// hello
}
```
> OUTPUT: </br>
> Field 0: 23 </br>
> Field 1: hello </br>


# Low Level Programming
- Design of Go has certain safety policies
  - During compilation, Go checks for type safety, array bounds, etc
  - Strict rules for type conversion prevent direct memory access
  - Automatic garbage collection prevents memory leaks
- In Go
  - the Go scheduler freely moves goroutines from one thread to another.
  - A pointer identifies a variable without revealing the variable’s numeric address.
  - Addresses may change as the garbage collector moves variables; pointers are transparently updated

- All these features make Go more predictable than C
- Hiding these low level details makes Go very portable

## Unsafe Package
- It provides access to low-level programming features that expose the underlying memory model of Go
- It is implemented by the compiler
- Some environments may not allow the use of the `unsafe` package
- Package unsafe contains operations that step around the type safety of Go programs.
- This is used extensively in packages like `os`, `runtime`, `syscall`, and `net`
- Packages that import this may be non-portable

## Methods
### Sizeof Method
- `func Sizeof(v ArbitraryType) uintptr`
- Returns the size of a type in bytes of ` hypothetical variable `h` if `var h = x`
- Does not include any memory referenced by `x`, only direct memory used by the fixed size of the data structure
  - Size of a **Slice** is the size of the slice descriptor, not the size of the underlying array
  - For **Struct**, the size includes any padding added by field alignment
- The Return value is a `Go constant` if `v` does not have variable lenght (Eg arrays have variable length)
- A call to Sizeof is a constant expression of type uintptr

#### Typical Sizes for Go

| Type      | Size        |
|-----------|-------------|
| bool      | 1 byte      |
| intN      | N/8 bytes   |
| uintN     | N/8 bytes   |
| floatN    | N/8 bytes   |
| complexN  | N/8 bytes   |
| int       | 1 word      |
| uint      | 1 word      |
| uintptr   | 1 word      |
| *T        | 1 word      |
| string    | 2 words     |
| []T       | 3 words     |
| map       | 1 word      |
| func      | 1 word      |
| chan      | 1 word      |
| interface | 2 words     |

A word is the size of a pointer, either 4 or 8 bytes depending on the architecture (32 or 64 bit)

Values should be aligned
- memory addresses for 2-byte values should be divisible by 2 (Eg: int16)
- memory addresses for 4-byte values should be divisible by 4 (Eg: int32)
- memory addresses for 8-byte values should be divisible by 8 (Eg: int64)

For this reason, the size of a value of an aggregate type (a struct or array) is at least the sum of
the sizes of its fields or elements but may be greater due to the presence of ‘‘holes.’’ 

Holes are unused spaces added by the compiler to ensure that the following field or element is
properly aligned relative to the start of the struct or array.
Hence it may be more space effiecent to order them in such a way that they are packed as tightly as possible

For example
```go
type StructA struct {
	a bool    // 1 byte
	b float64 // 8 bytes
	c int32   // 4 bytes
}

type StructB struct {
	b float64 // 8 bytes
	c int32   // 4 bytes
	a bool    // 1 byte
}

// StructA: 24 bytes
// StructB: 16 bytes
```

This is because
```
StructA:
		| a (1 byte)| <7 bytes padding> |					// 8 bytes
		| b (8 bytes) | c (4 bytes) | <4 bytes padding> |	// 16 bytes (c would be split here usually)

StructB: 
		| b (8 bytes) | 								//	8 bytes
		|c (4 bytes) | a  (1 byte)| <3 bytes padding> | // 8 bytes
```
#### Example: Using Sizeof
```go
import (
	"fmt"
	"unsafe"
)

func main() {
	// Integers
	fmt.Println(unsafe.Sizeof(int(10)))	// 8
	fmt.Println(unsafe.Sizeof(uint(10)))	// 8
	fmt.Println(unsafe.Sizeof(int8(10)))	// 1
	fmt.Println(unsafe.Sizeof(int16(10)))	// 2
	fmt.Println(unsafe.Sizeof(int32(10)))	// 4
	fmt.Println(unsafe.Sizeof(int64(10)))	// 8

	// Bool
	fmt.Println(unsafe.Sizeof(bool(true)))	// 1


	var i = 10
	point := &i
    fmt.Println("Sizeof(j) :", unsafe.Sizeof(point)) // "8"
    fmt.Println("Sizeof(float32(0)) :", unsafe.Sizeof(float32(0))) // "4"
    fmt.Println("Sizeof(float64(0)) :", unsafe.Sizeof(float64(0))) // "8"

	// Empty struct
    fmt.Println("Sizeof(struct{}{}) :", unsafe.Sizeof(struct{}{})) // "0"

	// Interface
    var k interface{}
    fmt.Println("Sizeof(interface{}) :", unsafe.Sizeof(k)) // "16"

	// String
    s := "Hello"
    fmt.Println("Sizeof('Hello') :", unsafe.Sizeof(s)) // "16"
	
}
```

### Alignof Method
- `func Alignof(v ArbitraryType) uintptr`
- Returns the required alignment of a type in bytes
- It is the same as the value returned by reflect.TypeOf(x).Align(). 
- For struct `s` with field `f`, Alignof(s.f) retruns the required alignment of the field within the struct
- The return value of Alignof is a Go constant if the type of the argument does not have variable size.

Offset for a struct
```go
type struct_T struct {
	val1 bool
	val2 int
	val3 []int
}
```

| Field | Size | Align | Offset |
|-------|------|-------|--------|
| val1  | 1    | 1     | 0      |
| val2  | 8    | 8     | 8      |
| val3  | 24   | 8     | 16     |

### Offsetof Method
- `func Offsetof(x ArbitraryType) uintptr`
- Returns the offset within the struct of the field represented by x, which must be of the form structValue.field
- Simply put, it returns the number of bytes between the start of the struct and the start of the field.
- The return value of Offsetof is a Go constant if the type of the argument does not have variable size

### Pointer Method
- `type Pointer *ArbitraryType`
- An unsafe pointer can be used to share data between different data types or to manipulate data at the byte level, providing the ability to optimize certain kinds of operations.
- Most times its written as `*T`
- Special type of pointer that can point to any type of data
- They are compatible and can be compared with `nil` which is the zero value of a pointer

#### Operations
- There are 4 special operations that can be done on pointers
  - A pointer value of any type can be converted to a Pointer.
  - A Pointer can be converted to a pointer value of any type.
  - A uintptr can be converted to a Pointer.
  - A Pointer can be converted to a uintptr.
   
#### Conversion
- Normal pointers can be converted to unsafe pointers and vice versa
- By converting a float64 pointer to a uintptr, we can get the address of the float64 value
```go
func Float64bits(f float64) uint64 {
return *(*uint64)(unsafe.Pointer(&f))
}
fmt.Printf("%#016x\n", Float64bits(1.0)) // "0x3ff0000000000000"
```

# Debugging
- Done using `Delve` debugger
- Delve is a debugger for the Go programming language

## Syntax
```sh
dlv cmd
```


```sh
dlv debug file.go	//  compiles the binary from source
dlv exec file.go	//   runs the binary

```

## Common Commands
| Command | Alias                    | Description               |
| ------- | ------------------------ | ------------------------- |
| `func`  | dlv funcs func_name      | Print a list of functions |
| `exit`  | dlv exit                 | Exit the debugger         |
| `list`  | dlv list file.go:line_no | Show source code          |

## Delve Client
### Starting Programs
- Delve has a client that can be used to interact with the debugger
| Command            | Alias | Description                                                                                                        |
| ------------------ | ----- | ------------------------------------------------------------------------------------------------------------------ |
| `call`             | -     | Resumes the process by embedding a function call                                                                   |
| `continue`         | `c`   | Executes until a breakpoint or program termination                                                                 |
| `next`             | `n`   | Goes to the next line of source code                                                                               |
| `rebuild`          | -     | Rebuilds the target executable and restarts it. Does not work if the executable was not built by the delve program |
| `restart`          | `r`   | Restarts the process                                                                                               |
| `step`             | `s`   | Single step through the program                                                                                    |
| `step-instruction` | `si`  | A single step on a single processor instruction                                                                    |
| `stepout`          | `so`  | Exit the current function                                                                                          |

### Manipulation of breakpoints
| Command       | Alias  | Description                                       |
| ------------- | ------ | ------------------------------------------------- |
| `break`       | `b`    | Sets a breakpoint                                 |
| `breakpoints` | `bp`   | Outputs information about active breakpoints      |
| `clear`       | -      | Deletes a breakpoint                              |
| `clearall`    | -      | Removes multiple breakpoints                      |
| `condition`   | `cond` | Sets the breakpoint condition                     |
| `on`          | -      | Executes the command when a breakpoint is reached |
| `trace`       | `t`    | Sets the trace point                              |


### View program variables and memory
| Command      | Alias | Description                                                          |
| ------------ | ----- | -------------------------------------------------------------------- |
| `args`       | -     | Print function arguments                                             |
| `display`    | -     | Display the value of the expression each time the program is stopped |
| `examinemem` | `x`   | Examine memory                                                       |
| `locals`     | -     | Print local variables                                                |
| `print`      | `p`   | Evaluate the expression                                              |
| `regs`       | -     | Output the contents of the processor registers                       |
| `set`        | -     | Change the value of a variable                                       |
| `vars`       | -     | Output package variables                                             |
| `whatis`     | -     | Output the type of an expression                                     |

### List output and switch between threads and goroutines
| Command      | Alias | Description                                        |
| ------------ | ----- | -------------------------------------------------- |
| `goroutine`  | `gr`  | Displays or changes the current goroutine          |
| `goroutines` | `grs` | List of program goroutines                         |
| `thread`     | `tr`  | Navigates to the specified thread                  |
| `threads`    | -     | Output information for each thread being monitored |

### View call stack and select frames

| Command    | Alias | Description                                                 |
| ---------- | ----- | ----------------------------------------------------------- |
| `deferred` | -     | Execute a command in the context of a deferred call         |
| `down`     | -     | Move the current frame down                                 |
| `frame`    | -     | Set the current frame or execute a command on another frame |
| `stack`    | `bt`  | Output a stack trace                                        |
| `up`       | -     | Move the current frame up                                   |
### Other commands:
| Command     | Alias       | Description                                         |
| ----------- | ----------- | --------------------------------------------------- |
| `config`    | -           | Changes the configuration settings                  |
| `edit`      | -           | Opens where you are in $DELVE_EDITOR or $EDITOR     |
| `exit`      | `quit`, `q` | Exit the debugger                                   |
| `funcs`     | -           | Print a list of functions                           |
| `help`      | `h`         | Print a help message                                |
| `libraries` | -           | List of loaded dynamic libraries                    |
| `list`      | `ls`, `l`   | Show source code                                    |
| `source`    | -           | Executes a file containing a list of delve commands |
| `sources`   | -           | Execute a list of source files                      |
| `types`     | -           | Execute a list of types                             |


## Breakpoint
- A breakpoint instructs the debugger to stop at a particular code location in the user's program, returning control of the debugger to them
- Most common way to interact with the debugger

### Linespec locations
- A linespec location is a colon-separated list that includes a source file name, source line number, and function and/or label names.
- Example: `file.go:LineNo`

### Commands
| Name        | Command                    | Description                |
| ----------- | -------------------------- | -------------------------- |
| break       | `dlv break file.go:LineNo` | Add a breakpoint           |
| breakpoints | `dlv breakpoints`          | List all breakpoints       |
| clear       | `dlv clear line_no`        | Remove a breakpoint        |
| clearall    | `dlv clearall`             | Remove all breakpoints     |
| continue    | `dlv continue`             | Continue execution         |
| next        | `dlv next`                 | Move to next line          |
| step        | `dlv step`                 | Move inside function       |
| stepout     | `dlv stepout`              | Return to calling function |
| restart     | `dlv restart`              | Restart the program        |

## View

- **Print** allows us to print the value of a variable and evaluate an expression
- **Locals** used to examine the content of all local variables
   
### Syntaxs
#### Syntax: For Print
```sh
dlv print var_name
```

#### Syntax: For Locals
```sh
dlv list						// list local var
```

### Examples
#### Example: For Print
```go
package main

func main() {
	var x int = 10
	var y int = 20
	var z int = x + y
}
```

```sh
dlv print x		// 10
dlv print y		// 20
dlv print z		// 30
```
> OUTPUT: </br>
> 10 </br>
> 20 </br>
> 30 </br>

#### Example: For Locals
```go
package main

func main() {
	var x int = 10
	var y int = 20
	var z int = x + y
}
```

```sh
dlv list
```
> OUTPUT: </br>
> 1: var x int = 10 </br>
> 2: var y int = 20 </br>
> 3: var z int = 30 </br>

# REPL
- A read–eval–print loop (REPL), also termed an interactive toplevel or language shell
- Simply put, it is a simple, interactive computer programming environment that takes single user inputs, executes them, and returns the result to the user
- A Program written in a REPL environment is executed piecewise.
  
# Popular Third-Party Go Packages

Here are more popular, useful, and fun Go packages to spark ideas for your projects:

| Package | Use Case | Import Path |
| --- | --- | --- |
| **Gin** | HTTP web framework | `github.com/gin-gonic/gin` |
| **Echo** | HTTP web framework | `github.com/labstack/echo/v4` |
| **Gorilla Mux** | HTTP router | `github.com/gorilla/mux` |
| **Colly** | Web scraping | `github.com/gocolly/colly` |
| **GORM** | ORM for SQL databases | `gorm.io/gorm` |
| **sqlx** | SQL toolkit | `github.com/jmoiron/sqlx` |
| **MongoDB Driver** | MongoDB client | `go.mongodb.org/mongo-driver/mongo` |
| **Badger** | Embedded key-value DB | `github.com/dgraph-io/badger/v4` |
| **grpc-go** | gRPC server/client | `google.golang.org/grpc` |
| **Cobra** | CLI framework | `github.com/spf13/cobra` |
| **Lipgloss** | Build beautiful terminal UIs (styling, colors, layouts) | `github.com/charmbracelet/lipgloss` |
| **Viper** | Config management | `github.com/spf13/viper` |
| **Zap** | Fast logging | `go.uber.org/zap` |
| **Logrus** | Structured logger | `github.com/sirupsen/logrus` |
| **Go-SDL2** | Game dev (SDL2 bindings) | `github.com/veandco/go-sdl2/sdl` |
| **Testify** | Unit testing/assertions | `github.com/stretchr/testify` |
| **Godotenv** | Load `.env` files | `github.com/joho/godotenv` |
| **JWT-Go** | JSON Web Tokens | `github.com/golang-jwt/jwt/v5` |
| **Fiber** | Express-style web framework | `github.com/gofiber/fiber/v2` |
| **GoQuery** | jQuery-style HTML parsing | `github.com/PuerkitoBio/goquery` |
| **GoCV** | Computer vision (OpenCV bindings) | `gocv.io/x/gocv` |
| **Go-Redis** | Redis client | `github.com/redis/go-redis/v9` |
| **Minio** | S3-compatible object storage | `github.com/minio/minio` |
| **Go-Mail** | Email sending | `github.com/go-mail/mail` |
| **Go-Chart** | Chart/graph generation | `github.com/wcharczuk/go-chart/v2` |
| **Go-Humble** | Terminal UI | `github.com/gizak/termui/v3` |
| **Go-Playground/Validator** | Struct validation | `github.com/go-playground/validator/v10` |
| **Go-Socket.IO** | Real-time websockets | `github.com/googollee/go-socket.io` |
| **Go-TGBot** | Telegram bot API | `github.com/go-telegram-bot-api/telegram-bot-api/v5` |
| **Go-Discord** | Discord bot API | `github.com/bwmarrin/discordgo` |
| **Go-Excelize** | Excel file manipulation | `github.com/xuri/excelize/v2` |
| **Go-Colorable** | Cross-platform color output | `github.com/mattn/go-colorable` |
| **Go-Geo** | Geospatial calculations | `github.com/paulmach/go.geo` |
 **Go-Image** | Image manipulation | `github.com/disintegration/imaging` |
| **Go-Faker** | Fake data generation | `github.com/bxcodec/faker/v4` |
---

### Niche or Lesser-Known Packages

| Package | Use Case | Import Path |
| --- | --- | --- |
| **Go-Rogue** | Roguelike game engine | `github.com/JoelOtter/termloop` |
| **Go-Noise** | Procedural noise generation | `github.com/ojrac/opensimplex-go` |
| **Go-Barcode** | Barcode generation | `github.com/boombuler/barcode` |
| **Go-Sound** | Audio playback/processing | `github.com/faiface/beep` |
| **Go-QR** | QR code generation | `github.com/skip2/go-qrcode` |

---

> **Explore more at [pkg.go.dev](https://pkg.go.dev/) and GitHub!**

## What Are Some Common Things Go Is Used For?

Go is a versatile language, popular for its simplicity, speed, and built-in concurrency. Here are some of the most common use cases, along with popular libraries for each:

---

### Web Servers & APIs
Go is widely used to build fast, scalable web servers and RESTful APIs.

**Popular Libraries:**
- [Gin](https://github.com/gin-gonic/gin)
- [Echo](https://github.com/labstack/echo)
- [Fiber](https://github.com/gofiber/fiber)
- [Gorilla Mux](https://github.com/gorilla/mux)

---

### Microservices
Its concurrency model and performance make Go ideal for microservice architectures.

**Popular Libraries:**
- [grpc-go](https://google.golang.org/grpc) (gRPC)
- [Go-Kit](https://github.com/go-kit/kit)
- [Go-Micro](https://github.com/go-micro/go-micro)

---

### CLI Tools
Many command-line tools are written in Go due to its easy cross-compilation and static binaries.

**Popular Libraries:**
- [Cobra](https://github.com/spf13/cobra)
- [urfave/cli](https://github.com/urfave/cli)
- [Lipgloss](https://github.com/charmbracelet/lipgloss) (for beautiful terminal UIs)

---

### Networking Tools
Go’s standard library makes it easy to build network utilities, proxies, and custom protocols.

**Popular Libraries:**
- [gnet](https://github.com/panjf2000/gnet)
- [Go-Socket.IO](https://github.com/googollee/go-socket.io)

---

### Cloud Infrastructure
Major cloud projects (Kubernetes, Terraform, Docker) are written in Go.

**Popular Libraries:**
- [client-go](https://github.com/kubernetes/client-go) (Kubernetes API)
- [Terraform SDK](https://github.com/hashicorp/terraform-plugin-sdk)

---

### DevOps & Automation
Go is used for CI/CD tools, deployment scripts, and infrastructure automation.

**Popular Libraries:**
- [GoCDK](https://github.com/google/go-cloud)
- [Sentry-Go](https://github.com/getsentry/sentry-go)

---

### Database Clients & Servers
Go is used for building database drivers, clients, and even lightweight database servers.

**Popular Libraries:**
- [GORM](https://gorm.io/gorm) (ORM for SQL)
- [sqlx](https://github.com/jmoiron/sqlx)
- [Go-Redis](https://github.com/redis/go-redis)
- [MongoDB Driver](https://go.mongodb.org/mongo-driver/mongo)
- [Badger](https://github.com/dgraph-io/badger)

---

### Distributed Systems
Its concurrency and networking strengths make Go a top choice for distributed systems and event-driven architectures.

**Popular Libraries:**
- [NATS](https://github.com/nats-io/nats.go)
- [etcd](https://github.com/etcd-io/etcd)

---

### Data Processing & ETL
Go is used for high-performance data pipelines and ETL jobs.

**Popular Libraries:**
- [Go-CSV](https://github.com/gocarina/gocsv)
- [Go-Excelize](https://github.com/xuri/excelize)

---

### Game Development
While less common, Go can be used for simple games and game servers.

**Popular Libraries:**
- [Go-SDL2](https://github.com/veandco/go-sdl2/sdl)
- [Termloop](https://github.com/JoelOtter/termloop)
- [Ebiten](https://github.com/hajimehoshi/ebiten)

---

### Web Scraping & Automation
Libraries like Colly make web scraping straightforward.

**Popular Libraries:**
- [Colly](https://github.com/gocolly/colly)
- [GoQuery](https://github.com/PuerkitoBio/goquery)

---

### APIs for Mobile & IoT
Go is used to build backend APIs for mobile apps and IoT devices.

**Popular Libraries:**
- [Gin](https://github.com/gin-gonic/gin)
- [Echo](https://github.com/labstack/echo)

---

### Security Tools
Go’s performance and static binaries are great for building security scanners and analysis tools.

**Popular Libraries:**
- [ZAP](https://go.uber.org/zap) (logging)
- [JWT-Go](https://github.com/golang-jwt/jwt)

---

### Testing & Validation
Go is used for robust testing and validation.

**Popular Libraries:**
- [Testify](https://github.com/stretchr/testify)
- [Go-Playground/Validator](https://github.com/go-playground/validator)

---

Go’s ecosystem and package library make it suitable for almost any backend or infrastructure task.
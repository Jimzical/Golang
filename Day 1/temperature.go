package main

func main() {
	const freezingF = 32.0
	const boilingF = 212.0

	println(fToc(freezingF))

}

func fToc(f float64) float64 {
	return 5/9 * (f - 32)
}


// TYPES of INTS
// int8, int16, int32, int64

// .Each int is a differnt from each other so int and int16 cannot be directly converted.
// float to int conversion removes the fractional part


// signed no are 2's complement
// higerorder bits are saved for the sign
// range of value is --> -2^n-1 to ((2^n-1) - 1)
// rnage of bits for non-negative is 0 to 2^n-1


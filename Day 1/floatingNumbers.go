package main

import (
	"fmt"
	"math"
)

func main() {
	// exmaple to show the usage of %g is printf
	var f32 float32 = 16777216
	var f64 float64 = 16777216 * 2
	var f float64 = 8
	const Avagadro = 6.02214129e23

	fmt.Printf("%f\n",math.MaxFloat32)
	fmt.Printf("%f\n",math.MaxFloat64)

	fmt.Printf("%f %f %f %f\n",f32,f64,Avagadro,f)

	var z float64
	fmt.Println(z,-z,1/z,-1/z,z/z)
}
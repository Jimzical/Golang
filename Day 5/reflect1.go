package main
import "reflect"
import "fmt"

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)			// v is a reflect.Value
	fmt.Println("value:", v)		// "value: 3.4"
	// y will have type float64.
	y := v.Interface().(float64)	// y is an interface{}.
	fmt.Println("y:", y)			// "y: 3.4"
}

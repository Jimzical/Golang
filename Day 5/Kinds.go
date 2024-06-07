package main
import 
(
	"reflect"
	"fmt"
)
func main() {
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("value:", v)
fmt.Println("type:", v.Type())
fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
fmt.Println("value:", v.Float())
fmt.Println("value:", v.Int()) // panic: reflect
}
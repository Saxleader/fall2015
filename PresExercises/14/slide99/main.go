package main
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Please enter a float64:")
	var x float64
	fmt.Scan(&x)
	fmt.Println("The least integer value greater than",x,"is",math.Ceil(x))
}

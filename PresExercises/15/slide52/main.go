package main
import "fmt"

func main() {
	ints := make([]int,5,10)
	fmt.Println(ints)
	//False, make did not return a pointer
}

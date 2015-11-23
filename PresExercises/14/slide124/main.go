package main
import "fmt"

func main() {
	var variable int32 = 120
	fmt.Println(string(variable))
	fmt.Println(string('a'))
	fmt.Println(string([]byte{'h','e','l','l','o'}))
	fmt.Println([]byte{'h','e','l','l','o'})
	fmt.Println(float64(12))
//	fmt.Println(int(12.1230123))
}

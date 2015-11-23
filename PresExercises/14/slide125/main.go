package main
import "fmt"

func main() {
	var input1 interface{} = 123
	var input2 interface{} = "123"
	i, ok := input1.(int)
	if !ok {
		fmt.Println("Not an int")
	} else {
		fmt.Println("Is an int")
	}
	fmt.Println(i)
	i, ok = input2.(int)
	if !ok {
		fmt.Println("Not an int")
	} else {
		fmt.Println("Is an int")
	}
	fmt.Println(i)
}

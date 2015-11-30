package main
import "fmt"

type test struct {
	a string
}

func main() {
	p := new(test)
	p.a = "hi"
	fmt.Println(p)
	fmt.Println(p.a)
	//Yes, you can use new to create a variable of a struct
}

package main
import "fmt"

func main() {
	var myint *int = new(int)
	fmt.Println(myint)
	fmt.Println(*myint)
	//True, new did return a pointer
}

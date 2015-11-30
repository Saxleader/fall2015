package main
import "fmt"

func main() {
	var mystr *string = new(string)
	fmt.Println(mystr)
	fmt.Println(*mystr)
	//True, new did return a pointer
}

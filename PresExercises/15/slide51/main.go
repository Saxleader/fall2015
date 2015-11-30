package main
import "fmt"

func main() {
	var mybool *bool = new(bool)
	fmt.Println(mybool)
	fmt.Println(*mybool)
	//True, new did return a pointer
}

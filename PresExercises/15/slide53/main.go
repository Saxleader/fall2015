package main
import "fmt"

func main() {
	mymap := make(map[int]string)
	fmt.Println(mymap)
	//False, make did not return a pointer
}

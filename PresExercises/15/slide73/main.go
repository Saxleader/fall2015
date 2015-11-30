package main
import "fmt"

type test struct {
	a string
}

func main() {
	p := make(test)
	p.a = "hi"
	fmt.Println(p)
	fmt.Println(p.a)
	//No, you cannot use make for anything except slices, maps, and channels
}

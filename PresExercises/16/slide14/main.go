package main
import "fmt"

func even(i int) bool {
	if i % 2 == 0 {
		return true
	}else {
		return false
	}
}

func myFunc(i int) (int, bool) {
	return i/2,even(i)
}

func main() {
	a,b := myFunc(1)
	c,d := myFunc(2)
	fmt.Println("1:",a,b)
	fmt.Println("2:",c,d)
}

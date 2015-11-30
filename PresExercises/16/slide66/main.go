package main
import "fmt"

func fibonacci() func() int {
	first, second := 0,1
	return func() int {
		current := first
		first, second = second, first+second
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i <= 10; i++ {
		fmt.Println(f())
	}
}

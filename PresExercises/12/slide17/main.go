package main
import "fmt"

func main() {
	fmt.Println("Printing all even numbers from 1 - 1000:")
	for i := 1; i <= 1000; i++ {
		if (i%2)==0 {
			fmt.Println(i)
		}
	}
}

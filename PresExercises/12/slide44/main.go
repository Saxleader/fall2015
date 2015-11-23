package main
import "fmt"

func main() {
	fmt.Println("Counting to 4:")
	for i := 0; i < 5; i++ {
		fmt.Println("Count:",i)
	}
	fmt.Println("Counting to 4 again:")
	i:=0
	for i < 5{
		fmt.Println("Count:",i)
		i++
	}
	fmt.Println("Counting to 4 last time:")
	i=0
	for{
		fmt.Println("Count:",i)
		if i>4 {
			break
		}
		i++
	}
}

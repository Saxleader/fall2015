package main
import "fmt"

func Max(i ...int) int {
	max := i[0]
	for _,val := range i{
		if val > max {
			max = val
		}
	}
	return max
}

func main() {
	ints := []int{4,65,12,5,18,2,99}
	fmt.Println("List:", ints)
	fmt.Println("Max:",Max(ints...))
}

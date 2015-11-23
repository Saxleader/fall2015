package main
import "fmt"

func MySum(i []int) int{
	n:=0
	for _, val := range i {
		n+=val
	}
	return n
}

func main() {
	myInts:= make([]int,5)
	myInts[0]= 1
	myInts[1]= 2
	myInts[2]= 3
	myInts[3]= 4
	myInts[4]= 5
	total := MySum(myInts)
	fmt.Println(total)
}

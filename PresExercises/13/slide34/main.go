package main
import "fmt"

func main() {
	fmt.Println("Please input some numbers:")
	myInts := make([]int,10)
	temp := make([]interface{},10)
	for i := range myInts {
		temp[i] = &myInts[i]
	}
	n, _ := fmt.Scanln(temp...)
//	if err != nil {
//		fmt.Println("Scanln err:",err)
//		return
//	}
	myInts = myInts[:n]
	total := 0
	for _,val := range myInts {
		total += val
	}
	fmt.Println("The sum of the numbers is",total)
}

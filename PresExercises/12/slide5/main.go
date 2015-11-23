package main
import "fmt"

func main() {
	var a,b int
	fmt.Println("Please enter two numbers:")
	fmt.Scan(&a,&b)
	fmt.Println("The remainder is:",a%b)
}

package main
import "fmt"

func main() {
	var A int = 4
	fmt.Println("Memory Address for A:",&A)
	fmt.Println("Value of A:", A)
	var B *int = &A
	fmt.Println("Memory Address for B:",&B)
	fmt.Println("Value of B:",B)
	fmt.Println("Value of *B:",*B)
	fmt.Println("WOW, (&A == B) and (*B == A)")
}

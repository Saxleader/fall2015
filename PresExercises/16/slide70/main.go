package main
import "fmt"

func main() {
	fmt.Println("Hello, Corey")
	defer func() { fmt.Println("Goodbye, Corey") }()
	fmt.Println("Do some things")
	fmt.Println("Execute some code")
	fmt.Println("Finish some things up")
}

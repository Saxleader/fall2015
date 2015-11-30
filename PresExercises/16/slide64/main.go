package main
import "fmt"

func hello() func() {
	return func() { fmt.Println("Hello") }
}

func main() {
	sayHello := hello()
	sayHello()
}

package main
import "fmt"

func sayHello() func(string) {
	return func(name string) {fmt.Println("Hello,",name)}
}

func main() {
	greeting := sayHello()
	greeting("Corey")
}

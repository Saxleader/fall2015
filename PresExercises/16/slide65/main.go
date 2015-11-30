package main
import (
	"fmt"
	"reflect"
)

func hello() func() {
	return func() { fmt.Println("Hello") }
}

func main() {
	sayHello := hello()
	sayHello()
	fmt.Println(reflect.TypeOf(sayHello))
}

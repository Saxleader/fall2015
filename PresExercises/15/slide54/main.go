package main
import (
	"fmt"
	"reflect"
)

func main() {
	var a *int = new(int)
	var b *string = new(string)
	var c *bool = new(bool)
	fmt.Println(reflect.TypeOf(*a),":",*a)
	fmt.Println(reflect.TypeOf(*b),":",*b)
	fmt.Println(reflect.TypeOf(*c),":",*c)
}

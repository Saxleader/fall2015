package main
import (
	"fmt"
	"reflect"
)

func main() {
	var a *[]int = new([]int)
	var b *map[string]string = new(map[string]string)
	fmt.Println(reflect.TypeOf(*a),":",*a)
	fmt.Println(reflect.TypeOf(*b),":",*b)
}

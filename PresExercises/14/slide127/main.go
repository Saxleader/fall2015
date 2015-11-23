package main
import (
	"fmt"
	"reflect"
)

func CheckType(a interface{}){
	fmt.Println(reflect.TypeOf(a))
	fmt.Printf("%T\n\n",a)
}

func main() {
	var a int
	var b string
	var c bool
	var d map[string]int
	CheckType(a)
	CheckType(b)
	CheckType(c)
	CheckType(d)
}

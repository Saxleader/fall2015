package main
import (
	"fmt"
	"strconv"
)

func sentence(name string, age int) string {
	return name+" is "+strconv.Itoa(age)+" years old."
}

func main() {
	fmt.Println(sentence("John",27))
}

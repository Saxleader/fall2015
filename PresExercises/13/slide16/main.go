package main
import "fmt"

func printStrings(a ...string){
	for _,val := range a{
		fmt.Println(val)
	}
}

func main() {
	a:="Corey"
	b:="Brandy"
	c:= []string{"Alex","Aaron","Taylor"}
	printStrings(a,b)
	printStrings(c...)
}

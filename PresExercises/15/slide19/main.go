package main
import "fmt"

func main() {
	name1 := []string{"Corey","Brandy"}
	name2 := []string{"Alex","Taylor"}
	name := append(name1,name2...)
	fmt.Println(name)
}

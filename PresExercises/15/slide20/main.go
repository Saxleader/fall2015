package main
import "fmt"

func main() {
	name := []string{"Corey","Brandy","Alex","Taylor"}
	fmt.Println(name)
	name = append(name[:2],name[3:]...)
	fmt.Println(name)
}

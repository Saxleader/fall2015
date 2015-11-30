package main
import "fmt"

func main() {
	name := "Corey Dihel"
	fmt.Println("Name: ",name)
	fmt.Println("Last Name:",string(name[6:]))
}
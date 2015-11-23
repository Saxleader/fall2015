package main
import "fmt"

func main() {
	name := "My name is Corey Dihel"
	fmt.Println("Initials:",string(name[11]),string(name[17]))
	fmt.Println("First Name:",name[11:16])
}

package main
import "fmt"

const (
	A int = iota
	B int = iota
	C int = iota
)

func main() {
	fmt.Println(B,"is less than",C,", but larger than",A)
}

package main
import (
	"fmt"
	"strconv"
)

func main() {
	num := 3
	char := "56"
	numstr := strconv.Itoa(num)
	charint, _ := strconv.Atoi(char)
	total := num + charint
	totalstr := strconv.Itoa(total)
	fmt.Println(char+" + "+numstr+" = "+totalstr)
}

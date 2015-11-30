package main
import "fmt"

func dogYears(age int) (dogYears int) {
	dogYears = age * 7
	return
}

func main() {
	fmt.Println("7 human years =",dogYears(7),"dog years")
}

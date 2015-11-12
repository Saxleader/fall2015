package main

import "fmt"

func main() {
	fmt.Println("What is your name?")
	var input string
	fmt.Scan(&input)
	fmt.Println("Hello",input)
}

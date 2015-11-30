package main
import "fmt"

func main() {
	ages := map[string]int{
		"Corey":26,
		"Brandy":27,
		"Taylor":24,
		"Bleu":1,
		"Alex":23,
	}
	ages["Mark"]=14
	ages["Brandy"]=26
	if _, exists := ages["Alex"]; exists {
		delete(ages, "Alex")
	}
	for key,val := range ages{
		fmt.Println(key,"is",val,"years old.")
	}
	fmt.Println("Length of map is",len(ages))
}

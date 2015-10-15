package main

import(
"fmt"
uuid "github.com/nu7hatch/gouuid"
)

func main(){
	myUUID, _ := uuid.NewV4()
	myString := myUUID.String()
	fmt.Println(myString)
}
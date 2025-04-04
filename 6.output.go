package main

import (
	"fmt"
)

func main(){
	var name, age = "John", 23
	fmt.Print(name, age)
	fmt.Print("\n")
	fmt.Print(name, " ", age)
	fmt.Print("\n")
	fmt.Println(name, age)
	fmt.Println(name, " ", age)
	fmt.Printf("name has value: %v and type %T\n", name, name)
	fmt.Printf("age has value: %v and type %T\n", age, age)

	var queue, number = 7, 44
	fmt.Print(queue, number) // give automatic space if neither string type.

}
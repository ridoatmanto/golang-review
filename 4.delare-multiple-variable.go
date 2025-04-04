package main

import (
	"fmt"
)

func main(){
	// same type
	var a, b, c int = 1,2,3
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// different type but inferred
	var name, age= "John", 35.5
	fmt.Println(name)
	fmt.Println(age)

	// direct assign but multiple type
	address, number := "pennsylvania", 64
	fmt.Println(address)
	fmt.Println(number)

	// assigment in a block for better readibility
	var (
		customerName string
		customerAge float32 = 32.5
		customerAddress string = "Ohio"
	)
	fmt.Println(customerName)
	fmt.Println(customerAge)
	fmt.Println(customerAddress)

	// Naming convention could be:
	// Camel Case: myVariableName
	// Pascal Case: MyVariableName
	// Snake Case: my_variable_name
}
package main

import (
	"fmt"
)

func main(){
	fmt.Println("Declare variable without assignment:")
	var name string
	var age float32
	var limit int
	var certified bool
	
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(limit)
	age = 3.5
	fmt.Println(certified)
	fmt.Println("limit", age, "years")


	fmt.Println("Assignment after declare:")
	var student string
	student = "John"
	fmt.Println(student)
	// IMPORTANT: Declare with := must have assigned value.

	//var -> Can be inside and outside function
	// := -> Only inside function

	// var -> Declaration and assignment can be separated.
	// := -> Declaration and Assignment can't be done separately. 
	// (must be done in the same line)
}
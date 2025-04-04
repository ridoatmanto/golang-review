package main

import (
	"fmt"
)
const A int = 1 // with type
const PHI = 3.14 // inferred without type(based on value).

func main(){
	// Constant rules.
	// Usually in UPPERCASE letter for easy identify and differentiation form variables.
	// Can be declared both inside and outside of function.
	// Unchangeable and Read Only.

	
	// assigment in a block for better readibility
	var (
		STARTFROM int = 0
		KEY = "ABC"
		DEFAULTAGE = 13
	)
	fmt.Println(STARTFROM)
	fmt.Println(KEY)
	fmt.Println(DEFAULTAGE)
}
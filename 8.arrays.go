package main

import (
	"fmt"
)

func main(){
	// In Go, Arrays has fixed length.
	var array_age = [3]int{1,2,3}
	var array_name = [3]string{"John", "Travolta", "Doe"}

	fmt.Println("Age:", array_age)
	fmt.Println("Name:", array_name)

	// If not declared, compiler will state by its content number values.
	array_address := [...]string{"Ohio", "Hawaii", "Texas"}
	fmt.Println(array_address)
	fmt.Println("Access Array:", array_address[1])
	array_address[1] = "Alaska"
	fmt.Println("Change Array:", array_address[1])

	// How to initialized array
	arr1 := [5]int{} // empty
	arr2 := [5]int{1,2} // partial assign
	arr3 := [5]int{1,2,4,5,3} // fully assign
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)
	fmt.Println("arr3:", arr3)
	fmt.Println("arr3 length:", len(arr3))

	arr4 := [5]int{1:30,4:60} // assign spesific index
	fmt.Println("arr4:", arr4)
	fmt.Println("arr4 length:", len(arr4))
}
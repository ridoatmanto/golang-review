package main

import (
	"fmt"
)

func main(){
	// Slices similar with array but more powerful and flexible.
	// Can grow and shrink fit.

	//There's several ways to create a slice.
	// :=
	// slice from another array
	// with make.

	myslice := []int{} //empty slice
	myslice2 := []int{1,2,3} //initialize during declaration

	fmt.Println("myslice", myslice)
	fmt.Println("len(myslice)", len(myslice)) // check length
	fmt.Println("cap(myslice)", cap(myslice)) // check capacity

	fmt.Println("myslice2", myslice2)
	fmt.Println("len(myslice2)", len(myslice2))
	fmt.Println("cap(myslice2)", cap(myslice2))

	// myslice := myArray[start:capacity] // create slice from existing array
	// capacity comes FromStart array.
	allNumber := [6]int{1,2,3,4,5}
	mySliceArray := allNumber[2:4]
	fmt.Println("allNumber", allNumber)
	fmt.Println("len(mySliceArray)", len(mySliceArray))
	fmt.Println("cap(mySliceArray)", cap(mySliceArray))

	// using make
	// slice_name := make([]type, length, capacity)

	sliceName :=make([]int, 5, 10)
	fmt.Println("sliceName", sliceName)
	fmt.Println("len(sliceName)", len(sliceName))
	fmt.Println("cap(sliceName)", cap(sliceName))

	sliceName2 := make([]int, 5) // omitted/removed capacity
	fmt.Println("sliceName2", sliceName2)
	fmt.Println("len(sliceName2)", len(sliceName2))
	fmt.Println("cap(sliceName2)", cap(sliceName2))
	// capacity will be 5.

}
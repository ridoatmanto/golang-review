package main

import (
	"fmt"
)

func main(){
	
	myslice := []int{1,2,3} // initialize during declaration

	fmt.Println("myslice", myslice[0])
	fmt.Println("myslice[1]", myslice[1]) // access element.
	myslice[1] = 5; // Change element.
	
	// append(sliceSource, element1, element2)
	newSlice := append(myslice, 7, 8)
	fmt.Println("newSlice", newSlice)

	slice1 := []int{5,3,1}
	slice2 := []int{2,4,6}

	slice3 := append(slice1, slice2...) // need three dots after second slice.
	fmt.Println("slice3", slice3)

	arr1 := [6]int{9,10,11,12,13,14}
	sliceA := arr1[1:5] // get from index 1, with 5 capacity from start.
	fmt.Println(sliceA)
	fmt.Println("len(sliceA)", len(sliceA)) // 4
	fmt.Println("cap(sliceA)", cap(sliceA)) // 5

	sliceB := append(sliceA, 77, 88) // must be a slice. Not and Array.
	fmt.Println("sliceB", sliceB)

	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	// neededNumbers := numbers[:len(numbers)-10] // 5
	neededNumbers := numbers[:5] // 5
	numbersCopy := make([]int, len(neededNumbers)+1) // 6
	// numbersCopy := make([]int, len(neededNumbers)) // 5
	// We can made more capacity than neededNumbers
	fmt.Println("neededNumbers", neededNumbers)
	fmt.Println("numbersCopy", numbersCopy)
	
	copy(numbersCopy, neededNumbers) // 
	fmt.Println("numbersCopy", numbersCopy)
}
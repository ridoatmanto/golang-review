package main

import (
	"fmt"
)

func main(){
	var total = 15 + 17
	fmt.Println("total", total)

	var (
		sum1 = 100 + 50
		sum2 = sum1 + 250
		sum3 = sum2 + sum2
	)
	fmt.Println("sum3", sum3)

	var add = 15 + 17
	var minus = 15 - 17
	var multiply = 15 * 17
	var div = 15 * 17
	var mod = 17 % 15
	var mod2 = 15 % 17
	var increment = 15
	increment++
	var decrement = 15
	decrement--
	fmt.Println("add", add)
	fmt.Println("minus", minus)
	fmt.Println("multiply", multiply)
	fmt.Println("div", div)
	fmt.Println("mod", mod)
	fmt.Println("mod2", mod2)
	fmt.Println("increment", increment)
	fmt.Println("decrement", decrement)
	decrement += 10
	fmt.Println("decrement add 10", decrement)
	decrement %= 7
	fmt.Println("decrement mod 7", decrement)
	

}
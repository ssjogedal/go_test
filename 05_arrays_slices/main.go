package main

import "fmt"

func main() {
	// Arrays
	//var fruitArr [2]string

	// Assign values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	// Declare and assign
	//fruitArr := [2]string{"apple", "orange"}

	//fmt.Println(fruitArr)
	//fmt.Println(fruitArr[1])

	fruitSlice := []string{"apple", "orange", "grape", "banana"}

	// Number of items in slice
	fmt.Println(len(fruitSlice))
	// Range between [Start:Stop]
	fmt.Println(fruitSlice[1:3])

}

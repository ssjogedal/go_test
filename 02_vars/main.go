package main

import "fmt"

func main() {

	// Using var "string" and "int" unnecessary
	// var name = "Simon"
	var age int = 28

	// const -> Constant value. Non changeble
	const isCool = true

	// Shorthand 1.1 & 1.2
	// name := "Simon"
	size := 1.3

	firstName, lastName := "Simon", "Sjogedal"

	fmt.Println(firstName, lastName, age, isCool)

	// Printf formats according to format specifier
	fmt.Printf("%T\n", size)
}

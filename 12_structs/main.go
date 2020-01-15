package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	// Alternative
	// firstName, lastName, city, gender string
	// age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old"
}

// hasBirthday method (pointer reciever * ) since we are going to change something
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever * )
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "F" {
		p.lastName = spouseLastName
	} else {
		return
	}

}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Simon", lastName: "Sjogedal", city: "Vancouver", gender: "M", age: 28}

	// Alterantive (Better if not hardcoded)
	person2 := Person{"Ellinor", "Svalberg", "Stockholm", "F", 25}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person1.lastName)

	person2.age++

	fmt.Println(person2)

	person1.hasBirthday()
	person2.getMarried(person1.lastName)
	fmt.Println(person2.greet())

}

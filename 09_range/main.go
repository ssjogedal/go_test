package main

import "fmt"

func main() {

	// Slice of IDs
	ids := []int{33, 45, 67, 2, 88, 11, 53, 14}

	// Loop through IDs using range
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index(i)
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum:", sum)

	// Range with maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}

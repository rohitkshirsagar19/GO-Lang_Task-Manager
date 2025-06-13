package main

import "fmt"

func main() {

	// Variable declaration
	var name string = "Task-Manager"

	// Print Statement
	fmt.Println("Welcome to", name)

	// Function call
	sum := add(5, 10)
	fmt.Println("Sum of 5 and 10 is:", sum)

	// Loop to print numbers from 1 to 5
	for i := 0; i < 5; i++ {
		fmt.Println("Number: ", i+1)
	}
}

// Function to add two integers
func add(a int, b int) int {
	return a + b
}

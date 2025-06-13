package main

import "fmt"

// GO Data Structure
type Task struct {
	ID     int
	Name   string
	Status bool
}

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

	// Create a Task instance
	tasks := []Task{
		{ID: 1, Name: "Write Code", Status: true},
		{ID: 2, Name: "Review Code", Status: false},
		{ID: 3, Name: "Test Code", Status: true},
	}

	// Apend task
	tasks = append(tasks, Task{ID: 4, Name: "Deploy Code", Status: false})

	// Loop over the tasks
	for _, task := range tasks {
		fmt.Printf("ID: %d , Name: %s , Status: %v \n", task.ID, task.Name, task.Status)
	}

	// Mapping
	taskStatus := map[int]bool{
		1: true,
		0: false,
	}
	fmt.Println("Task 1 done?", taskStatus[1])

}

// Function to add two integers
func add(a int, b int) int {
	return a + b
}

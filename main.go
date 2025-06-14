package main

import (
	"GO-Lang_Task-Manager/tasks"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := "tasks.json"
	taskList, err := tasks.LoadTasks(filename)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: add <task_name>")
			return
		}
		addTask(&taskList, os.Args[2])
	case "list":
		listTasks(taskList)
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: done <task_id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID:", err)
			return
		}
		markDone(&taskList, id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: delete <task_id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID:", err)
			return
		}
		deleteTask(&taskList, id)
	default:
		printUsage()
		return
	}

	if err := tasks.SaveTasks(filename, taskList); err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func addTask(taskList *[]tasks.Task, name string) {
	id := 1
	if len(*taskList) > 0 {
		id = (*taskList)[len(*taskList)-1].ID + 1
	}
	*taskList = append(*taskList, tasks.Task{ID: id, Name: name, Done: false})
	fmt.Println("Added task:", name)
}

func listTasks(taskList []tasks.Task) {
	if len(taskList) == 0 {
		fmt.Println("No tasks")
		return
	}
	for _, task := range taskList {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("ID: %d, Name: %s, Status: %s\n", task.ID, task.Name, status)
	}
}

func markDone(taskList *[]tasks.Task, id int) {
	for i, task := range *taskList {
		if task.ID == id {
			(*taskList)[i].Done = true
			fmt.Println("Marked task", id, "as done")
			return
		}
	}
	fmt.Println("Task", id, "not found")
}

func deleteTask(taskList *[]tasks.Task, id int) {
	for i, task := range *taskList {
		if task.ID == id {
			*taskList = append((*taskList)[:i], (*taskList)[i+1:]...)
			fmt.Println("Deleted task", id)
			return
		}
	}
	fmt.Println("Task", id, "not found")
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add <task_name>    Add a new task")
	fmt.Println("  list               List all tasks")
	fmt.Println("  done <task_id>     Mark task as done")
	fmt.Println("  delete <task_id>   Delete a task")
}

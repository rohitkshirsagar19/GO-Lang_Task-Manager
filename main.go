package main

import (
	"GO-Lang_Task-Manager/tasks"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	filename := "tasks.json"
	taskList, err := tasks.LoadTasks(filename)
	if err != nil && !os.IsNotExist(err) {
		color.Red("✗ Error loading tasks: %v", err)
		return
	}

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			color.Yellow("⚠ Usage: add <task_name>")
			return
		}
		color.Cyan("⏳ Adding task...")
		addTask(&taskList, os.Args[2])
	case "list":
		listTasks(taskList)
	case "done":
		if len(os.Args) < 3 {
			color.Yellow("⚠ Usage: done <task_id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			color.Red("✗ Invalid ID: %v", err)
			return
		}
		color.Cyan("⏳ Marking task as done...")
		markDone(&taskList, id)
	case "delete":
		if len(os.Args) < 3 {
			color.Yellow("⚠ Usage: delete <task_id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			color.Red("✗ Invalid ID: %v", err)
			return
		}
		color.Cyan("⏳ Deleting task...")
		deleteTask(&taskList, id)
	default:
		printUsage()
		return
	}

	color.Cyan("⏳ Saving tasks...")
	if err := tasks.SaveTasks(filename, taskList); err != nil {
		color.Red("✗ Error saving tasks: %v", err)
	}
	color.Green("✓ Tasks saved successfully")
}

func addTask(taskList *[]tasks.Task, name string) {
	id := 1
	if len(*taskList) > 0 {
		id = (*taskList)[len(*taskList)-1].ID + 1
	}
	*taskList = append(*taskList, tasks.Task{ID: id, Name: name, Done: false})
	color.Green("✓ Added task: %s", name)
}

func listTasks(taskList []tasks.Task) {
	if len(taskList) == 0 {
		color.Yellow("⚠ No tasks found")
		return
	}

	// Table header
	fmt.Println()
	color.Cyan("┌───────┬───────────────────────────┬──────────┐")
	color.Cyan("│ ID    │ Task Name                 │ Status   │")
	color.Cyan("├───────┬───────────────────────────┬──────────┤")

	// Table rows
	for _, task := range taskList {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		name := task.Name
		if len(name) > 25 {
			name = name[:22] + "..."
		}
		if task.Done {
			color.Green("│ %-5d │ %-25s │ %-8s │", task.ID, name, status)
		} else {
			color.Yellow("│ %-5d │ %-25s │ %-8s │", task.ID, name, status)
		}
	}

	// Table footer
	color.Cyan("└───────┴───────────────────────────┴──────────┘")
	fmt.Println()
}

func markDone(taskList *[]tasks.Task, id int) {
	for i, task := range *taskList {
		if task.ID == id {
			(*taskList)[i].Done = true
			color.Green("✓ Marked task %d as done", id)
			return
		}
	}
	color.Red("✗ Task %d not found", id)
}

func deleteTask(taskList *[]tasks.Task, id int) {
	for i, task := range *taskList {
		if task.ID == id {
			*taskList = append((*taskList)[:i], (*taskList)[i+1:]...)
			color.Green("✓ Deleted task %d", id)
			return
		}
	}
	color.Red("✗ Task %d not found", id)
}

func printUsage() {
	fmt.Println()
	color.Cyan("=== Task Manager CLI ===")
	fmt.Println()
	fmt.Println("Usage:")
	color.Green("  add <task_name>    Add a new task")
	color.Green("  list               List all tasks")
	color.Green("  done <task_id>     Mark task as done")
	color.Green("  delete <task_id>   Delete a task")
	fmt.Println()
}

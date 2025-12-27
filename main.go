package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task represents a single todo item
type Task struct {
	ID        int
	Title     string
	Completed bool
}

var tasks []Task
var nextID = 1

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n==== TODO LIST ====")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			addTask(reader)
		case "2":
			listTasks()
		case "3":
			completeTask(reader)
		case "4":
			deleteTask(reader)
		case "5":
			fmt.Println("Exiting... Have a productive day 🚀")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func addTask(reader *bufio.Reader) {
	fmt.Print("Enter task title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	task := Task{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}

	tasks = append(tasks, task)
	nextID++

	fmt.Println("✅ Task added successfully.")
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("\n--- Tasks ---")
	for _, task := range tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("[%d] %s %s\n", task.ID, status, task.Title)
	}
}

func completeTask(reader *bufio.Reader) {
	fmt.Print("Enter task ID to complete: ")
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			fmt.Println("🎯 Task marked as completed.")
			return
		}
	}

	fmt.Println("Task not found.")
}

func deleteTask(reader *bufio.Reader) {
	fmt.Print("Enter task ID to delete: ")
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("🗑️ Task deleted.")
			return
		}
	}

	fmt.Println("Task not found.")
}

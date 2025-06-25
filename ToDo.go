package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// getInput reads a line of input from console and trims whitespace
func getInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input!", err)
	}
	return strings.TrimSpace(input)
}

func main() {
	myTasks := map[int]string{}

	reader := bufio.NewReader(os.Stdin)

	for {
		// Display the menu
		fmt.Println("\nEnter 1 to add new task")
		fmt.Println("Enter 2 to view your tasks")
		fmt.Println("Enter 3 to delete a task")
		fmt.Println("Enter 4 to exit")

		fmt.Print("Choose an option: ")

		option := getInput(reader)

		if option == "1" {
			fmt.Print("/nEnter the task: ")
			task := getInput(reader)

			id := len(myTasks) + 1
			myTasks[id] = task
			fmt.Println("Task added successfully")
		} else if option == "2" {
			if len(myTasks) == 0 {
				fmt.Println("No tasks found. Please add a task first.")
				continue
			}

			fmt.Println("/nYour ToDo tasks:")
			for id, task := range myTasks {
				fmt.Printf("%d. %s\n", id, task)
			}
		} else if option == "3" {
			if len(myTasks) == 0 {
				fmt.Println("No tasks found. Please add a task first.")
				continue
			}

			fmt.Println("Your ToDo tasks:")
			for id, task := range myTasks {
				fmt.Printf("%d. %s\n", id, task)
			}
			fmt.Print("Enter the task number which needs to be deleted: ")

			input := getInput(reader)
			id, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid number.", err)
				continue
			}

			if _, exists := myTasks[id]; exists {
				delete(myTasks, id)
				fmt.Println("/nTask deleted successfully.")
			} else {
				fmt.Println("/nTask not found.")
			}
		} else if option == "4" {
			fmt.Println("----------")
			fmt.Println("Goodbye.")
			fmt.Println("----------")
			break
		} else {
			fmt.Println("Invalid option.")
		}
	}
}

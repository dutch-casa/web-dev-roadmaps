package main

import (
	"fmt"

	"boundary-drawing/display"
	"boundary-drawing/task"
)

// main is the composition — the imperative shell.
// It creates data, orchestrates operations, and prints output.
// The logic lives in the task and display packages.

func main() {
	tasks := []task.Task{
		task.New("Set up development environment", 1),
		task.New("Learn Git basics", 1),
		task.New("Complete naming exercises", 2),
		task.New("Read about control flow", 2),
		task.New("Build terminal Wordle", 3),
	}

	task.Start(&tasks[0])
	task.Complete(&tasks[0])
	task.Start(&tasks[1])

	fmt.Println("=== Task Report ===")
	fmt.Println()
	for _, t := range tasks {
		fmt.Println("  " + display.FormatTask(t))
	}
	fmt.Println()
	fmt.Println(display.FormatSummary(tasks))
}

package main

import (
	"fmt"
	"strings"
	"time"
)

// Boundary drawing exercise
//
// This program is a tiny task manager. Everything is in one file.
// Your job: split it into packages at the natural boundaries.
//
// Suggested packages:
//   task/     — the Task type, status constants, and task logic
//   display/  — formatting and display functions
//   main.go   — just the composition (create tasks, print report)
//
// Rules:
//   - Each package should be understandable without reading the others
//   - The task package should know nothing about printing or formatting
//   - The display package should know about tasks but not about main()
//   - Only export what callers actually need
//   - "go run ." must produce the same output
//
// Think about: where are the natural seams? What knowledge belongs
// where? What should be hidden?

type TaskStatus int

const (
	StatusTodo TaskStatus = iota
	StatusInProgress
	StatusDone
)

func (s TaskStatus) String() string {
	return [...]string{"todo", "in-progress", "done"}[s]
}

type Task struct {
	Title     string
	Status    TaskStatus
	Priority  int // 1 = highest
	CreatedAt time.Time
	DoneAt    time.Time
}

func NewTask(title string, priority int) Task {
	return Task{
		Title:     title,
		Status:    StatusTodo,
		Priority:  priority,
		CreatedAt: time.Now(),
	}
}

func StartTask(t *Task) {
	t.Status = StatusInProgress
}

func CompleteTask(t *Task) {
	t.Status = StatusDone
	t.DoneAt = time.Now()
}

func IsOverdue(t Task, deadline time.Time) bool {
	return t.Status != StatusDone && time.Now().After(deadline)
}

func FilterByStatus(tasks []Task, status TaskStatus) []Task {
	var result []Task
	for _, t := range tasks {
		if t.Status == status {
			result = append(result, t)
		}
	}
	return result
}

func CountByStatus(tasks []Task) map[TaskStatus]int {
	counts := make(map[TaskStatus]int)
	for _, t := range tasks {
		counts[t.Status]++
	}
	return counts
}

func FormatTask(t Task) string {
	priority := strings.Repeat("!", t.Priority)
	return fmt.Sprintf("[%s] %-30s %s", t.Status, t.Title, priority)
}

func FormatSummary(tasks []Task) string {
	counts := CountByStatus(tasks)
	return fmt.Sprintf("Total: %d | Todo: %d | In Progress: %d | Done: %d",
		len(tasks),
		counts[StatusTodo],
		counts[StatusInProgress],
		counts[StatusDone])
}

func PrintReport(tasks []Task) {
	fmt.Println("=== Task Report ===")
	fmt.Println()
	for _, t := range tasks {
		fmt.Println("  " + FormatTask(t))
	}
	fmt.Println()
	fmt.Println(FormatSummary(tasks))
}

func main() {
	tasks := []Task{
		NewTask("Set up development environment", 1),
		NewTask("Learn Git basics", 1),
		NewTask("Complete naming exercises", 2),
		NewTask("Read about control flow", 2),
		NewTask("Build terminal Wordle", 3),
	}

	StartTask(&tasks[0])
	CompleteTask(&tasks[0])
	StartTask(&tasks[1])

	PrintReport(tasks)
}

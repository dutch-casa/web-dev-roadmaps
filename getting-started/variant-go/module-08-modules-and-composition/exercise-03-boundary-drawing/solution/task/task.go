// Package task defines the Task type, status values, and pure
// operations on tasks. It knows nothing about display or I/O.
package task

import "time"

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

func (s Status) String() string {
	return [...]string{"todo", "in-progress", "done"}[s]
}

type Task struct {
	Title     string
	Status    Status
	Priority  int
	CreatedAt time.Time
	DoneAt    time.Time
}

func New(title string, priority int) Task {
	return Task{
		Title:     title,
		Status:    Todo,
		Priority:  priority,
		CreatedAt: time.Now(),
	}
}

func Start(t *Task) {
	t.Status = InProgress
}

func Complete(t *Task) {
	t.Status = Done
	t.DoneAt = time.Now()
}

func FilterByStatus(tasks []Task, status Status) []Task {
	var result []Task
	for _, t := range tasks {
		if t.Status == status {
			result = append(result, t)
		}
	}
	return result
}

func CountByStatus(tasks []Task) map[Status]int {
	counts := make(map[Status]int)
	for _, t := range tasks {
		counts[t.Status]++
	}
	return counts
}

// Package display handles formatting tasks for terminal output.
// It knows about the task package but nothing about main() or I/O.
// Functions return strings — the caller decides where to print them.
package display

import (
	"fmt"
	"strings"

	"boundary-drawing/task"
)

func FormatTask(t task.Task) string {
	priority := strings.Repeat("!", t.Priority)
	return fmt.Sprintf("[%s] %-30s %s", t.Status, t.Title, priority)
}

func FormatSummary(tasks []task.Task) string {
	counts := task.CountByStatus(tasks)
	return fmt.Sprintf("Total: %d | Todo: %d | In Progress: %d | Done: %d",
		len(tasks),
		counts[task.Todo],
		counts[task.InProgress],
		counts[task.Done])
}

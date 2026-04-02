package main

import (
	"fmt"
	"strings"
)

// File organization exercise
//
// All the code below lives in one file and it's a mess.
// Your job: split it into multiple files based on what each piece does.
//
// Suggested split:
//   main.go      — just the main function (composition)
//   student.go   — the Student type and related functions
//   format.go    — formatting and display functions
//   validate.go  — validation logic
//
// Rules:
//   - All files stay in package main (same directory)
//   - "go run ." must produce the same output
//   - Each file should be understandable without reading the others
//   - Put types at the top of their file, functions below

func formatCurrency(cents int) string {
	dollars := cents / 100
	remainder := cents % 100
	return fmt.Sprintf("$%d.%02d", dollars, remainder)
}

func validateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

type Student struct {
	Name    string
	Email   string
	Scores  []int
	Tuition int // in cents
}

func validateStudent(s Student) []string {
	var errors []string
	if s.Name == "" {
		errors = append(errors, "name is required")
	}
	if !validateEmail(s.Email) {
		errors = append(errors, "invalid email")
	}
	if len(s.Scores) == 0 {
		errors = append(errors, "at least one score required")
	}
	for _, score := range s.Scores {
		if score < 0 || score > 100 {
			errors = append(errors, fmt.Sprintf("score %d out of range", score))
		}
	}
	if s.Tuition < 0 {
		errors = append(errors, "tuition cannot be negative")
	}
	return errors
}

func averageScore(scores []int) float64 {
	total := 0
	for _, s := range scores {
		total += s
	}
	return float64(total) / float64(len(scores))
}

func formatStudentSummary(s Student) string {
	avg := averageScore(s.Scores)
	return fmt.Sprintf("%-10s %-25s avg=%.1f tuition=%s",
		s.Name, s.Email, avg, formatCurrency(s.Tuition))
}

func formatValidationErrors(name string, errs []string) string {
	return fmt.Sprintf("INVALID %-10s: %s", name, strings.Join(errs, "; "))
}

func main() {
	students := []Student{
		{Name: "Alice", Email: "alice@auburn.edu", Scores: []int{92, 88, 95}, Tuition: 1250000},
		{Name: "Bob", Email: "bob@auburn.edu", Scores: []int{71, 65, 78}, Tuition: 1250000},
		{Name: "", Email: "nobody", Scores: []int{}, Tuition: -100},
		{Name: "Clara", Email: "clara@auburn.edu", Scores: []int{86, 91, 83}, Tuition: 1100000},
	}

	fmt.Println("=== Student Report ===")
	fmt.Println()

	for _, s := range students {
		errs := validateStudent(s)
		if len(errs) > 0 {
			fmt.Println(formatValidationErrors(s.Name, errs))
			continue
		}
		fmt.Println(formatStudentSummary(s))
	}
}

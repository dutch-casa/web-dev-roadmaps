package main

import "fmt"

const appName = "Grade Calculator"
const version = "1.0.0"
const maxScore = 100.0

func main() {
	// studentName doesn't change in this program, but Go's const only works
	// with basic types. A string constant is fine here.
	const studentName = "Jordan"

	// scores is a slice — Go doesn't allow const slices. It must be a var.
	// We don't modify it, but the language requires it to be a variable.
	scores := []float64{88, 92, 75, 96, 84}

	// total is accumulated in a loop — it changes on every iteration.
	total := 0.0
	for _, s := range scores {
		total += s
	}

	// average is computed once and never changes after this line.
	// But it depends on runtime values (total, len), so it can't be a const.
	average := total / float64(len(scores))

	// passing is derived from average. Same situation — runtime value.
	passing := average >= 70.0

	// grade is reassigned inside the if/else chain. Must be a variable.
	grade := "F"
	if average >= 90 {
		grade = "A"
	} else if average >= 80 {
		grade = "B"
	} else if average >= 70 {
		grade = "C"
	} else if average >= 60 {
		grade = "D"
	}

	fmt.Println(appName, version)
	fmt.Printf("Student: %s\n", studentName)
	fmt.Printf("Scores: %v\n", scores)
	fmt.Printf("Average: %.1f / %.0f\n", average, maxScore)
	fmt.Printf("Passing: %v\n", passing)
	fmt.Printf("Grade: %s\n", grade)
}

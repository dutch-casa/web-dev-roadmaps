package main

import "fmt"

// Every value below is declared as a variable.
// Your job: convert every one that COULD be a const into a const.
// For each one you leave as a variable, add a comment explaining
// why it needs to stay mutable.
//
// Run with: go run .

func main() {
	var appName = "Grade Calculator"
	var version = "1.0.0"
	var maxScore = 100.0

	var studentName = "Jordan"
	var scores = []float64{88, 92, 75, 96, 84}

	var total = 0.0
	for _, s := range scores {
		total += s
	}

	var average = total / float64(len(scores))

	var passing = average >= 70.0

	var grade = "F"
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

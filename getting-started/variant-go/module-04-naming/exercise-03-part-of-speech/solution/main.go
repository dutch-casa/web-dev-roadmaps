package main

import "fmt"

// Types are nouns
type Applicant struct {
	FirstName string
	LastName  string
	Age       int
	IsActive  bool // boolean is a predicate
	Scores    []float64 // collection is a plural noun
}

// Booleans are predicates — reads as "is eligible"
func isEligible(applicant Applicant) bool {
	return applicant.Age >= 18 && applicant.IsActive
}

// Functions are verbs — "calculate average score"
func calculateAverageScore(scores []float64) float64 {
	total := 0.0
	for _, s := range scores {
		total += s
	}
	if len(scores) == 0 {
		return 0
	}
	return total / float64(len(scores))
}

// Function is a verb — "format decision"
func formatDecision(applicant Applicant) string {
	if !isEligible(applicant) {
		return "DENIED"
	}
	average := calculateAverageScore(applicant.Scores)

	standing := "regular"
	if average >= 90 {
		standing = "honors"
	}

	return fmt.Sprintf("%s %s (approved) - avg: %.1f - %s",
		applicant.FirstName, applicant.LastName, average, standing)
}

func main() {
	// Collection is a plural noun
	applicants := []Applicant{
		{FirstName: "Jordan", LastName: "Lee", Age: 20, IsActive: true, Scores: []float64{92, 88, 95}},
		{FirstName: "Sam", LastName: "Park", Age: 17, IsActive: true, Scores: []float64{85, 79, 91}},
		{FirstName: "Alex", LastName: "Chen", Age: 22, IsActive: false, Scores: []float64{96, 94, 98}},
		{FirstName: "Casey", LastName: "Jones", Age: 19, IsActive: true, Scores: []float64{72, 68, 75}},
	}

	for _, applicant := range applicants {
		fmt.Println(formatDecision(applicant))
	}
}

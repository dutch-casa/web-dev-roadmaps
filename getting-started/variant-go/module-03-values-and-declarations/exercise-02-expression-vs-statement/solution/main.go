package main

import (
	"fmt"
	"strings"
)

// Problem 1: The result is a single expression built from string operations.
// If there's a middle initial, include it. Otherwise, skip it.
func formatName(first, middle, last string) string {
	base := strings.ToUpper(last) + ", " + first
	if middle != "" {
		return base + " " + string(middle[0]) + "."
	}
	return base
}

// Problem 2: The boolean expression IS the answer.
// No temporary variables. No if/else. The condition is the return value.
func isEligible(age int, hasPermission bool) bool {
	return age >= 18 && hasPermission
}

// Problem 3: Guard clauses turn a nested ladder into a linear scan.
// Each threshold gets one line. The reader sees the mapping instantly.
func letterGrade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

// Problem 4: A helper function expresses the decision.
// The formatting and the decision are separate concerns.
func describeTemp(celsius float64) string {
	fahrenheit := celsius*9/5 + 32
	label := tempLabel(fahrenheit)
	return fmt.Sprintf("%.0f°F — %s", fahrenheit, label)
}

func tempLabel(f float64) string {
	switch {
	case f > 100:
		return "scorching"
	case f > 80:
		return "hot"
	case f > 60:
		return "pleasant"
	case f > 40:
		return "cold"
	default:
		return "freezing"
	}
}

func main() {
	fmt.Println(formatName("Rosa", "Louise", "Parks"))
	fmt.Println(formatName("Guido", "", "van Rossum"))
	fmt.Println()

	fmt.Println("Eligible (21, true):", isEligible(21, true))
	fmt.Println("Eligible (16, true):", isEligible(16, true))
	fmt.Println()

	for _, score := range []int{95, 82, 74, 65, 48} {
		fmt.Printf("Score %d → %s\n", score, letterGrade(score))
	}
	fmt.Println()

	for _, temp := range []float64{42, 18, -5, 30, 55} {
		fmt.Println(describeTemp(temp))
	}
}

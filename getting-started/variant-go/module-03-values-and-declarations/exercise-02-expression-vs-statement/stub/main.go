package main

import (
	"fmt"
	"strings"
)

// Expression vs. statement
//
// Expressions produce values. Statements perform actions.
// The practical difference: expressions compose (you can nest them),
// statements are sequential steps.
//
// This exercise is about recognizing when statement-heavy code
// can be simplified by thinking in terms of expressions.
// Each function below does something clunky with temporary variables
// and reassignment. Rewrite each one so the logic flows through
// expressions instead of being stapled together with statements.
//
// Run "go run ." before and after — output must be identical.

// Problem 1: String building through reassignment.
// This builds a formatted name by mutating a variable five times.
// Rewrite it so the result is computed in one or two expressions.
func formatName(first, middle, last string) string {
	var result string

	result = strings.ToUpper(last)

	result = result + ", "

	result = result + first

	if middle != "" {
		result = result + " " + string(middle[0]) + "."
	}

	return result
}

// Problem 2: Boolean logic buried under if/else.
// The if/else statements assign true or false to a variable.
// That's just... the boolean expression itself. Simplify.
func isEligible(age int, hasPermission bool) bool {
	var ageCheck bool
	if age >= 18 {
		ageCheck = true
	} else {
		ageCheck = false
	}

	var result bool
	if ageCheck && hasPermission {
		result = true
	} else {
		result = false
	}

	return result
}

// Problem 3: Computing a letter grade through a chain of reassignment.
// Each branch sets the same variable. There's a simpler way to express
// "map a number to a category."
func letterGrade(score int) string {
	var grade string
	if score >= 90 {
		grade = "A"
	} else {
		if score >= 80 {
			grade = "B"
		} else {
			if score >= 70 {
				grade = "C"
			} else {
				if score >= 60 {
					grade = "D"
				} else {
					grade = "F"
				}
			}
		}
	}
	return grade
}

// Problem 4: Describing a temperature by mutating through branches.
// This function reassigns `desc` in every branch, then uses it once.
// Express the decision as a value instead.
func describeTemp(celsius float64) string {
	fahrenheit := celsius*9/5 + 32

	var desc string
	if fahrenheit > 100 {
		desc = "scorching"
	} else if fahrenheit > 80 {
		desc = "hot"
	} else if fahrenheit > 60 {
		desc = "pleasant"
	} else if fahrenheit > 40 {
		desc = "cold"
	} else {
		desc = "freezing"
	}

	return fmt.Sprintf("%.0f°F — %s", fahrenheit, desc)
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

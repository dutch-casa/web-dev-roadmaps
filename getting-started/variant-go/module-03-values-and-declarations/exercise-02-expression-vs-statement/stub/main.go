package main

import (
	"fmt"
	"strings"
)

// Part 1: Classify each line in the comments below as
// "expression" or "statement". Write your answer next to each one.
//
// An expression produces a value.
// A statement performs an action.
//
//   3 + 4                          →
//   x := 10                        →
//   len("hello")                   →
//   fmt.Println("hi")              →
//   x > 0 && y < 10               →
//   for i := range 5 { }          →
//   strings.ToUpper("go")         →
//   if x > 0 { }                  →
//   x * 2 + 1                     →
//   return total                  →

// Part 2: The function below is written in a statement-heavy style.
// It builds a result string by repeatedly reassigning a variable.
// Rewrite it to use expressions — build the result by combining
// string operations instead of accumulating into a variable.

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

// Part 3: This function uses a temporary variable where an expression
// would be clearer. Simplify it.

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

func main() {
	fmt.Println(formatName("Rosa", "Louise", "Parks"))
	fmt.Println(formatName("Guido", "", "van Rossum"))
	fmt.Println(isEligible(21, true))
	fmt.Println(isEligible(16, true))
}

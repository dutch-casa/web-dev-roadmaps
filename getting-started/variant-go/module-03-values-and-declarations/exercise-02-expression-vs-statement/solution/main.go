package main

import (
	"fmt"
	"strings"
)

// Part 1 answers:
//
//   3 + 4                          → expression (produces 7)
//   x := 10                        → statement (declares and assigns)
//   len("hello")                   → expression (produces 5)
//   fmt.Println("hi")              → statement (performs I/O; technically returns values, but used for its side effect)
//   x > 0 && y < 10               → expression (produces a bool)
//   for i := range 5 { }          → statement (controls flow)
//   strings.ToUpper("go")         → expression (produces "GO")
//   if x > 0 { }                  → statement (controls flow)
//   x * 2 + 1                     → expression (produces a number)
//   return total                  → statement (exits the function)

// Part 2: Expression-oriented rewrite.
// Instead of accumulating into a variable, we build the string in one expression.
func formatName(first, middle, last string) string {
	base := strings.ToUpper(last) + ", " + first
	if middle != "" {
		return base + " " + string(middle[0]) + "."
	}
	return base
}

// Part 3: Simplified.
// The boolean expression IS the result — no temporary variables needed.
func isEligible(age int, hasPermission bool) bool {
	return age >= 18 && hasPermission
}

func main() {
	fmt.Println(formatName("Rosa", "Louise", "Parks"))
	fmt.Println(formatName("Guido", "", "van Rossum"))
	fmt.Println(isEligible(21, true))
	fmt.Println(isEligible(16, true))
}

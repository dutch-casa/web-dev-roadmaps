package main

import (
	"fmt"
	"strings"
)

// Fixed: Short names in a tight loop. The scope is 4 lines.
// 'total' and 'n' are fine here — 'n' is used immediately.
func sumOfSquares(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n * n
	}
	return total
}

// Fixed: Longer names for a function where variables are used far
// from their declaration. 'rating' is used 10 lines below where
// it's assigned — the reader needs the name to remind them what it is.
func printGradeReport(gradesByStudent map[string][]int) {
	for name, grades := range gradesByStudent {
		total := 0
		for _, g := range grades {
			total += g
		}
		average := float64(total) / float64(len(grades))

		rating := "fail"
		if average >= 90 {
			rating = "excellent"
		} else if average >= 80 {
			rating = "good"
		} else if average >= 70 {
			rating = "satisfactory"
		} else if average >= 60 {
			rating = "needs improvement"
		}

		fmt.Printf("%-10s avg=%.1f  %s\n", name, average, rating)
	}
}

// Fixed: The function name describes the transformation.
// Parameters tell the caller what to pass.
func repeatString(text string, count int) string {
	return strings.Repeat(text, count)
}

func main() {
	fmt.Println("Sum of squares:", sumOfSquares([]int{1, 2, 3, 4, 5}))

	fmt.Println()
	printGradeReport(map[string][]int{
		"Alice": {92, 88, 95},
		"Bob":   {71, 65, 78},
		"Clara": {86, 91, 83},
	})

	fmt.Println()
	fmt.Println(repeatString("Go! ", 3))
}

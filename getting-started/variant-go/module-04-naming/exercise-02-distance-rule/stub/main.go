package main

import "fmt"

// Distance rule exercise
//
// Each function below has naming problems related to scope distance.
// Some names are too long for their scope. Some are too short.
// Fix each one and add a comment explaining your reasoning.

// Problem 1: Names are too verbose for a tight loop.
func sumOfSquares(inputNumbers []int) int {
	totalSumOfAllSquaredValues := 0
	for currentIndex := range inputNumbers {
		currentSquaredValue := inputNumbers[currentIndex] * inputNumbers[currentIndex]
		totalSumOfAllSquaredValues += currentSquaredValue
	}
	return totalSumOfAllSquaredValues
}

// Problem 2: Names are too short for a function this long.
// The variable 'r' is used 15 lines after it's assigned.
func processStudentGrades(grades map[string][]int) {
	for n, g := range grades {
		t := 0
		for _, v := range g {
			t += v
		}
		a := float64(t) / float64(len(g))

		r := "fail"
		if a >= 90 {
			r = "excellent"
		} else if a >= 80 {
			r = "good"
		} else if a >= 70 {
			r = "satisfactory"
		} else if a >= 60 {
			r = "needs improvement"
		}

		fmt.Printf("%-10s avg=%.1f  %s\n", n, a, r)
	}
}

// Problem 3: This exported function has a name that doesn't describe
// what it actually does. It also has parameter names that don't help.
func DoIt(s string, n int) string {
	result := ""
	for range n {
		result += s
	}
	return result
}

func main() {
	fmt.Println("Sum of squares:", sumOfSquares([]int{1, 2, 3, 4, 5}))

	fmt.Println()
	processStudentGrades(map[string][]int{
		"Alice": {92, 88, 95},
		"Bob":   {71, 65, 78},
		"Clara": {86, 91, 83},
	})

	fmt.Println()
	fmt.Println(DoIt("Go! ", 3))
}

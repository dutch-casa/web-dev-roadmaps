package main

import "fmt"

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

package main

import (
	"fmt"
	"strings"
)

// Name audit
//
// This program works correctly. The problem is that every name is
// meaningless. Your job: rename every variable, function, and parameter
// so that someone can understand the code without reading the implementation.
//
// Rules:
//   - Don't change the behavior. Only change names.
//   - The code should read almost like English when you're done.
//   - Run "go run ." before and after to verify it still works.

type thing struct {
	x string
	y int
	z float64
	w bool
}

func do1(data []thing) []thing {
	var result []thing
	for _, d := range data {
		if d.w {
			result = append(result, d)
		}
	}
	return result
}

func do2(data []thing) float64 {
	temp := 0.0
	n := 0
	for _, d := range data {
		temp += d.z
		n++
	}
	if n == 0 {
		return 0
	}
	return temp / float64(n)
}

func do3(data []thing) map[int][]thing {
	m := make(map[int][]thing)
	for _, d := range data {
		m[d.y] = append(m[d.y], d)
	}
	return m
}

func do4(d thing) string {
	var parts []string
	parts = append(parts, d.x)
	parts = append(parts, fmt.Sprintf("(year %d)", d.y))
	parts = append(parts, fmt.Sprintf("$%.2f", d.z))
	if d.w {
		parts = append(parts, "[available]")
	} else {
		parts = append(parts, "[sold]")
	}
	return strings.Join(parts, " ")
}

func main() {
	stuff := []thing{
		{x: "Go Programming", y: 2023, z: 49.99, w: true},
		{x: "The Art of SQL", y: 2021, z: 39.95, w: false},
		{x: "Systems Design", y: 2024, z: 54.99, w: true},
		{x: "Network Protocols", y: 2022, z: 44.50, w: true},
		{x: "Data Structures", y: 2023, z: 42.00, w: false},
	}

	fmt.Println("All:")
	for _, d := range stuff {
		fmt.Println(" ", do4(d))
	}

	temp := do1(stuff)
	fmt.Printf("\nAvailable: %d\n", len(temp))
	for _, d := range temp {
		fmt.Println(" ", do4(d))
	}

	fmt.Printf("\nAverage price: $%.2f\n", do2(stuff))

	m := do3(stuff)
	fmt.Println("\nBy year:")
	for k, v := range m {
		fmt.Printf("  %d: %d items\n", k, len(v))
	}
}

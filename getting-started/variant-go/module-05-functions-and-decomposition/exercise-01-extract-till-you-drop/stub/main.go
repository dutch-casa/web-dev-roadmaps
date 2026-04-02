package main

import (
	"fmt"
	"strings"
)

// Extract till you drop
//
// This function does everything in one place: parsing, validating,
// computing, formatting, and printing. Your job is to break it into
// small functions where each one does one thing.
//
// Guidelines:
//   - No function should exceed ~25 lines
//   - Each extracted function gets a clear verb name
//   - The main composition function should read like a summary
//   - Keep pure logic separate from printing (side effects)
//   - Run "go run ." before and after — output must be identical

func generateReport(rawData string) {
	lines := strings.Split(rawData, "\n")
	var names []string
	var scores [][]int

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Printf("WARNING: skipping malformed line: %q\n", line)
			continue
		}
		name := strings.TrimSpace(parts[0])
		if name == "" {
			fmt.Printf("WARNING: skipping line with empty name\n")
			continue
		}
		scoreStrs := strings.Split(strings.TrimSpace(parts[1]), ",")
		var studentScores []int
		valid := true
		for _, s := range scoreStrs {
			s = strings.TrimSpace(s)
			n := 0
			for _, ch := range s {
				if ch < '0' || ch > '9' {
					fmt.Printf("WARNING: invalid score %q for %s\n", s, name)
					valid = false
					break
				}
				n = n*10 + int(ch-'0')
			}
			if !valid {
				break
			}
			if n < 0 || n > 100 {
				fmt.Printf("WARNING: score %d out of range for %s\n", n, name)
				valid = false
				break
			}
			studentScores = append(studentScores, n)
		}
		if !valid {
			continue
		}
		names = append(names, name)
		scores = append(scores, studentScores)
	}

	if len(names) == 0 {
		fmt.Println("No valid student data found.")
		return
	}

	fmt.Println("=== Grade Report ===")
	fmt.Println()

	classTotal := 0.0
	for i, name := range names {
		total := 0
		for _, s := range scores[i] {
			total += s
		}
		avg := float64(total) / float64(len(scores[i]))
		classTotal += avg

		grade := "F"
		if avg >= 90 {
			grade = "A"
		} else if avg >= 80 {
			grade = "B"
		} else if avg >= 70 {
			grade = "C"
		} else if avg >= 60 {
			grade = "D"
		}

		bar := strings.Repeat("█", int(avg/5))
		fmt.Printf("%-12s avg=%5.1f  grade=%s  %s\n", name, avg, grade, bar)
	}

	classAvg := classTotal / float64(len(names))
	fmt.Println()
	fmt.Printf("Class average: %.1f\n", classAvg)
	fmt.Printf("Students: %d\n", len(names))
}

func main() {
	data := `
		Alice: 92, 88, 95, 91
		Bob: 71, 65, 78, 72
		Clara: 86, 91, 83, 89
		Dave: 55, 62, 48, 51
		Eve: 98, 95, 100, 97
	`
	generateReport(data)
}

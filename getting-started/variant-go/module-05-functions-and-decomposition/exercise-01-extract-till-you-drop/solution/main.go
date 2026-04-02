package main

import (
	"fmt"
	"strings"
)

type StudentRecord struct {
	Name   string
	Scores []int
}

// Grade thresholds as a data table instead of branching.
var gradeThresholds = []struct {
	MinAverage float64
	Letter     string
}{
	{90, "A"},
	{80, "B"},
	{70, "C"},
	{60, "D"},
}

// --- Parsing (boundary layer: handles raw input) ---

func parseScore(raw string) (int, bool) {
	s := strings.TrimSpace(raw)
	n := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0, false
		}
		n = n*10 + int(ch-'0')
	}
	if n < 0 || n > 100 {
		return 0, false
	}
	return n, true
}

func parseLine(line string) (StudentRecord, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return StudentRecord{}, fmt.Errorf("malformed line: %q", line)
	}
	name := strings.TrimSpace(parts[0])
	if name == "" {
		return StudentRecord{}, fmt.Errorf("empty name in line: %q", line)
	}
	scoreStrs := strings.Split(strings.TrimSpace(parts[1]), ",")
	scores := make([]int, 0, len(scoreStrs))
	for _, raw := range scoreStrs {
		score, ok := parseScore(raw)
		if !ok {
			return StudentRecord{}, fmt.Errorf("invalid score %q for %s", raw, name)
		}
		scores = append(scores, score)
	}
	return StudentRecord{Name: name, Scores: scores}, nil
}

func parseStudentData(rawData string) []StudentRecord {
	var records []StudentRecord
	for _, line := range strings.Split(rawData, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		record, err := parseLine(line)
		if err != nil {
			fmt.Printf("WARNING: %v\n", err)
			continue
		}
		records = append(records, record)
	}
	return records
}

// --- Pure computation (no I/O) ---

func averageScore(scores []int) float64 {
	total := 0
	for _, s := range scores {
		total += s
	}
	return float64(total) / float64(len(scores))
}

func letterGrade(avg float64) string {
	for _, threshold := range gradeThresholds {
		if avg >= threshold.MinAverage {
			return threshold.Letter
		}
	}
	return "F"
}

func classAverage(records []StudentRecord) float64 {
	total := 0.0
	for _, r := range records {
		total += averageScore(r.Scores)
	}
	return total / float64(len(records))
}

// --- Formatting (pure: produces strings, no side effects) ---

func formatStudentLine(record StudentRecord) string {
	avg := averageScore(record.Scores)
	grade := letterGrade(avg)
	bar := strings.Repeat("█", int(avg/5))
	return fmt.Sprintf("%-12s avg=%5.1f  grade=%s  %s", record.Name, avg, grade, bar)
}

// --- Composition (the only place with side effects) ---

func printReport(records []StudentRecord) {
	if len(records) == 0 {
		fmt.Println("No valid student data found.")
		return
	}

	fmt.Println("=== Grade Report ===")
	fmt.Println()
	for _, r := range records {
		fmt.Println(formatStudentLine(r))
	}
	fmt.Println()
	fmt.Printf("Class average: %.1f\n", classAverage(records))
	fmt.Printf("Students: %d\n", len(records))
}

func main() {
	data := `
		Alice: 92, 88, 95, 91
		Bob: 71, 65, 78, 72
		Clara: 86, 91, 83, 89
		Dave: 55, 62, 48, 51
		Eve: 98, 95, 100, 97
	`
	records := parseStudentData(data)
	printReport(records)
}

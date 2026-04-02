package main

import (
	"fmt"
	"strings"
)

func formatCurrency(cents int) string {
	dollars := cents / 100
	remainder := cents % 100
	return fmt.Sprintf("$%d.%02d", dollars, remainder)
}

func formatStudentSummary(s Student) string {
	avg := averageScore(s.Scores)
	return fmt.Sprintf("%-10s %-25s avg=%.1f tuition=%s",
		s.Name, s.Email, avg, formatCurrency(s.Tuition))
}

func formatValidationErrors(name string, errs []string) string {
	return fmt.Sprintf("INVALID %-10s: %s", name, strings.Join(errs, "; "))
}

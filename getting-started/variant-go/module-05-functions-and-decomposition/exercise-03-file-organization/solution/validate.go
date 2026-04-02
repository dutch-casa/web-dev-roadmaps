package main

import (
	"fmt"
	"strings"
)

func validateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func validateStudent(s Student) []string {
	var errors []string
	if s.Name == "" {
		errors = append(errors, "name is required")
	}
	if !validateEmail(s.Email) {
		errors = append(errors, "invalid email")
	}
	if len(s.Scores) == 0 {
		errors = append(errors, "at least one score required")
	}
	for _, score := range s.Scores {
		if score < 0 || score > 100 {
			errors = append(errors, fmt.Sprintf("score %d out of range", score))
		}
	}
	if s.Tuition < 0 {
		errors = append(errors, "tuition cannot be negative")
	}
	return errors
}

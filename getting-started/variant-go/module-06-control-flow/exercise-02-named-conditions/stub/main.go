package main

import (
	"fmt"
	"time"
)

// Named conditions
//
// Each function below has complex boolean expressions used inline.
// Extract each one into a named boolean variable that explains
// the intent. The code should read like English when you're done.
//
// Example:
//   Before: if x > 0 && x < 100 && y != 0
//   After:  inRange := x > 0 && x < 100 && y != 0
//           if inRange { ... }

type Employee struct {
	Name           string
	Department     string
	YearsEmployed  int
	IsFullTime     bool
	IsManager      bool
	AnnualSalary   int
	LastReviewDate time.Time
	OnProbation    bool
}

// Problem 1: Who qualifies for a bonus?
func qualifiesForBonus(emp Employee) bool {
	return emp.IsFullTime && emp.YearsEmployed >= 2 && !emp.OnProbation && emp.AnnualSalary < 150000 && time.Since(emp.LastReviewDate).Hours() < 365*24
}

// Problem 2: Who can approve a purchase?
func canApprovePurchase(emp Employee, amount int) bool {
	if emp.IsManager && emp.IsFullTime && !emp.OnProbation {
		if amount <= 5000 || (amount <= 50000 && emp.YearsEmployed >= 5) || (emp.Department == "Finance" && emp.YearsEmployed >= 10) {
			return true
		}
	}
	return false
}

// Problem 3: Should we send a reminder email?
func shouldSendReminder(emp Employee, lastEmailDays int, hasUnreadNotifications bool, emailOptIn bool) bool {
	return emailOptIn && !emp.OnProbation && emp.IsFullTime && lastEmailDays > 7 && hasUnreadNotifications && emp.Department != "Legal"
}

func main() {
	emp := Employee{
		Name:           "Jordan",
		Department:     "Engineering",
		YearsEmployed:  6,
		IsFullTime:     true,
		IsManager:      true,
		AnnualSalary:   95000,
		LastReviewDate: time.Now().AddDate(0, -6, 0),
		OnProbation:    false,
	}

	fmt.Printf("%s qualifies for bonus: %v\n", emp.Name, qualifiesForBonus(emp))
	fmt.Printf("%s can approve $3000: %v\n", emp.Name, canApprovePurchase(emp, 3000))
	fmt.Printf("%s can approve $25000: %v\n", emp.Name, canApprovePurchase(emp, 25000))
	fmt.Printf("Send reminder to %s: %v\n", emp.Name, shouldSendReminder(emp, 10, true, true))
}

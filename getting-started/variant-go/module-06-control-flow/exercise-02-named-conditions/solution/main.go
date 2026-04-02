package main

import (
	"fmt"
	"time"
)

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

func qualifiesForBonus(emp Employee) bool {
	isEligibleStatus := emp.IsFullTime && !emp.OnProbation
	hasEnoughTenure := emp.YearsEmployed >= 2
	isBelowSalaryCap := emp.AnnualSalary < 150000
	hasRecentReview := time.Since(emp.LastReviewDate).Hours() < 365*24

	return isEligibleStatus && hasEnoughTenure && isBelowSalaryCap && hasRecentReview
}

func canApprovePurchase(emp Employee, amount int) bool {
	isActiveManager := emp.IsManager && emp.IsFullTime && !emp.OnProbation
	if !isActiveManager {
		return false
	}

	isSmallPurchase := amount <= 5000
	isMediumWithSeniority := amount <= 50000 && emp.YearsEmployed >= 5
	isFinanceVeteran := emp.Department == "Finance" && emp.YearsEmployed >= 10

	return isSmallPurchase || isMediumWithSeniority || isFinanceVeteran
}

func shouldSendReminder(emp Employee, lastEmailDays int, hasUnreadNotifications bool, emailOptIn bool) bool {
	wantsEmail := emailOptIn && emp.Department != "Legal"
	isActiveEmployee := emp.IsFullTime && !emp.OnProbation
	isDueForReminder := lastEmailDays > 7 && hasUnreadNotifications

	return wantsEmail && isActiveEmployee && isDueForReminder
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

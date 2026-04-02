package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// Pure vs. impure
//
// Part 1: Label each function below as PURE or IMPURE.
// Write your answer as a comment above each function.
//
// A pure function:
//   - Always returns the same output for the same input
//   - Does not modify anything outside itself
//   - Does not read from anything that could change (time, random, I/O)
//
// Part 2: Refactor the three impure functions at the bottom
// into pure cores with thin impure wrappers.

// Label: ___
func double(x int) int {
	return x * 2
}

// Label: ___
func greetUser(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Label: ___
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Label: ___
func randomGreeting(name string) string {
	greetings := []string{"Hey", "Hello", "Hi", "Howdy"}
	i := rand.IntN(len(greetings))
	return fmt.Sprintf("%s, %s!", greetings[i], name)
}

// Label: ___
func currentYear() int {
	return time.Now().Year()
}

// Label: ___
func contains(items []string, target string) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}

// --- Part 2: Refactor these ---

// This function mixes computation with printing.
// Extract the pure computation into its own function.
func printDiscount(price float64, discountPercent float64) {
	discounted := price - (price * discountPercent / 100)
	if discounted < 0 {
		discounted = 0
	}
	fmt.Printf("Original: $%.2f → Discounted: $%.2f (%.0f%% off)\n",
		price, discounted, discountPercent)
}

// This function uses the current time to determine a greeting.
// Extract the pure decision logic so it can be tested with any hour.
func timeBasedGreeting(name string) string {
	hour := time.Now().Hour()
	if hour < 12 {
		return fmt.Sprintf("Good morning, %s!", name)
	} else if hour < 17 {
		return fmt.Sprintf("Good afternoon, %s!", name)
	}
	return fmt.Sprintf("Good evening, %s!", name)
}

// This function generates a random password and prints it.
// Split the generation from the printing.
func generateAndPrintPassword(length int) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, length)
	for i := range password {
		password[i] = chars[rand.IntN(len(chars))]
	}
	fmt.Printf("Your new password: %s\n", string(password))
}

func main() {
	fmt.Println(double(21))
	greetUser("Auburn")
	fmt.Println(maxValue(10, 20))
	fmt.Println(randomGreeting("Jordan"))
	fmt.Println(currentYear())
	fmt.Println(contains([]string{"go", "rust", "python"}, "go"))

	fmt.Println()
	printDiscount(99.99, 20)
	fmt.Println(timeBasedGreeting("Sam"))
	generateAndPrintPassword(16)
}

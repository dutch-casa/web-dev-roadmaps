package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"time"
)

// Part 1 answers:

// PURE — same input always gives same output, no side effects.
func double(x int) int {
	return x * 2
}

// IMPURE — prints to stdout (side effect).
func greetUser(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// PURE — deterministic, no side effects.
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// IMPURE — uses random number generator (different result each call).
func randomGreeting(name string) string {
	greetings := []string{"Hey", "Hello", "Hi", "Howdy"}
	i := rand.IntN(len(greetings))
	return fmt.Sprintf("%s, %s!", greetings[i], name)
}

// IMPURE — reads the system clock (result changes over time).
func currentYear() int {
	return time.Now().Year()
}

// PURE — deterministic, no side effects.
func contains(items []string, target string) bool {
	return slices.Contains(items, target)
}

// --- Part 2: Refactored ---

// Pure core: computes the discounted price.
func applyDiscount(price float64, discountPercent float64) float64 {
	discounted := price - (price * discountPercent / 100)
	return max(discounted, 0)
}

// Impure wrapper: calls the pure function and prints the result.
func printDiscount(price float64, discountPercent float64) {
	discounted := applyDiscount(price, discountPercent)
	fmt.Printf("Original: $%.2f → Discounted: $%.2f (%.0f%% off)\n",
		price, discounted, discountPercent)
}

// Pure core: decides the greeting based on an hour value.
// Can be tested with any hour — no dependency on the clock.
func greetingForHour(name string, hour int) string {
	if hour < 12 {
		return fmt.Sprintf("Good morning, %s!", name)
	} else if hour < 17 {
		return fmt.Sprintf("Good afternoon, %s!", name)
	}
	return fmt.Sprintf("Good evening, %s!", name)
}

// Impure wrapper: reads the clock and passes the hour to the pure function.
func timeBasedGreeting(name string) string {
	return greetingForHour(name, time.Now().Hour())
}

// Pure core: generates a password from a given random source.
// In real code you'd pass a rand source; here we keep it simple.
func generatePassword(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, length)
	for i := range password {
		password[i] = chars[rand.IntN(len(chars))]
	}
	return string(password)
}

// Impure wrapper: generates and prints.
func generateAndPrintPassword(length int) {
	fmt.Printf("Your new password: %s\n", generatePassword(length))
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

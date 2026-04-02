package main

import "fmt"

// FP vs OOP decision
//
// Two problems. Each one fits one paradigm better than the other.
// Implement both, then write a paragraph explaining your choice.
//
// Problem A: Sales report pipeline
//
// You have a list of sales records. You need to:
//   1. Filter out refunded sales
//   2. Convert prices from cents to dollars
//   3. Group sales by category
//   4. Calculate the total revenue per category
//
// This is a data pipeline. Data goes in, transformed data comes out.
// No state changes. No identity. Pure transformations.
//
// Implement this using standalone functions that take data in and
// return data out. No methods on types. No mutation of the input.

type Sale struct {
	ID       string
	Category string
	Amount   int  // cents
	Refunded bool
}

type CategoryTotal struct {
	Category string
	Revenue  float64 // dollars
}

// TODO: Implement the pipeline as standalone functions:
//   filterActive(sales []Sale) []Sale
//   centsToDollars(cents int) float64
//   groupByCategory(sales []Sale) map[string][]Sale
//   categoryTotals(grouped map[string][]Sale) []CategoryTotal
//
// Then compose them in main().

// Problem B: A thermostat
//
// A thermostat has:
//   - A target temperature (set by the user)
//   - A current temperature (read from a sensor)
//   - A mode: heating, cooling, or idle
//   - Rules: if current < target - 1, heat. If current > target + 1, cool.
//     Otherwise, idle.
//
// The thermostat is a *thing* with identity and state that changes over
// time. You call methods on it: SetTarget, ReadTemperature, UpdateMode.
// The mode depends on the relationship between current and target.
//
// Implement this as a struct with methods.

// TODO: Define the Thermostat type and its methods:
//   NewThermostat(target float64) *Thermostat
//   (t *Thermostat) SetTarget(temp float64)
//   (t *Thermostat) ReadTemperature(current float64)
//   (t *Thermostat) UpdateMode()
//   (t *Thermostat) Status() string

func main() {
	// --- Problem A: Pipeline ---
	fmt.Println("=== Sales Report ===")
	sales := []Sale{
		{ID: "001", Category: "Electronics", Amount: 29999, Refunded: false},
		{ID: "002", Category: "Books", Amount: 1499, Refunded: false},
		{ID: "003", Category: "Electronics", Amount: 5999, Refunded: true},
		{ID: "004", Category: "Books", Amount: 2499, Refunded: false},
		{ID: "005", Category: "Clothing", Amount: 4999, Refunded: false},
		{ID: "006", Category: "Electronics", Amount: 14999, Refunded: false},
		{ID: "007", Category: "Clothing", Amount: 7999, Refunded: true},
		{ID: "008", Category: "Books", Amount: 999, Refunded: false},
	}
	_ = sales
	// TODO: Run the pipeline and print totals per category

	// --- Problem B: Thermostat ---
	fmt.Println("\n=== Thermostat ===")
	// TODO: Create a thermostat, set targets, read temperatures,
	// update mode, and print status at each step

	fmt.Println("TODO: Implement both problems")
}

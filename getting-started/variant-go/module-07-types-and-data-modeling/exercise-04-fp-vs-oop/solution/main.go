package main

import (
	"fmt"
	"sort"
)

// ============================================================
// Problem A: Sales report pipeline (functional style)
//
// Each function takes data in and returns data out. No mutation.
// No methods. No state. Pure transformations you can compose.
// ============================================================

type Sale struct {
	ID       string
	Category string
	Amount   int
	Refunded bool
}

type CategoryTotal struct {
	Category string
	Revenue  float64
}

func filterActive(sales []Sale) []Sale {
	var active []Sale
	for _, s := range sales {
		if !s.Refunded {
			active = append(active, s)
		}
	}
	return active
}

func centsToDollars(cents int) float64 {
	return float64(cents) / 100.0
}

func groupByCategory(sales []Sale) map[string][]Sale {
	grouped := make(map[string][]Sale)
	for _, s := range sales {
		grouped[s.Category] = append(grouped[s.Category], s)
	}
	return grouped
}

func categoryTotals(grouped map[string][]Sale) []CategoryTotal {
	var totals []CategoryTotal
	for category, sales := range grouped {
		revenue := 0.0
		for _, s := range sales {
			revenue += centsToDollars(s.Amount)
		}
		totals = append(totals, CategoryTotal{Category: category, Revenue: revenue})
	}
	sort.Slice(totals, func(i, j int) bool {
		return totals[i].Revenue > totals[j].Revenue
	})
	return totals
}

// ============================================================
// Problem B: Thermostat (object-oriented style)
//
// The thermostat is a thing with identity. It has state that
// changes over time. Methods read and modify that state.
// The mode depends on the relationship between internal fields.
// ============================================================

type ThermostatMode int

const (
	Idle ThermostatMode = iota
	Heating
	Cooling
)

func (m ThermostatMode) String() string {
	return [...]string{"idle", "heating", "cooling"}[m]
}

type Thermostat struct {
	target  float64
	current float64
	mode    ThermostatMode
}

func NewThermostat(target float64) *Thermostat {
	return &Thermostat{
		target: target,
		mode:   Idle,
	}
}

func (t *Thermostat) SetTarget(temp float64) {
	t.target = temp
}

func (t *Thermostat) ReadTemperature(current float64) {
	t.current = current
}

// UpdateMode determines heating/cooling/idle based on the gap
// between current and target temperature. The 1-degree deadband
// prevents the system from oscillating rapidly when the temperature
// is close to the target.
func (t *Thermostat) UpdateMode() {
	switch {
	case t.current < t.target-1:
		t.mode = Heating
	case t.current > t.target+1:
		t.mode = Cooling
	default:
		t.mode = Idle
	}
}

func (t *Thermostat) Status() string {
	return fmt.Sprintf("target=%.1f°  current=%.1f°  mode=%s",
		t.target, t.current, t.mode)
}

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

	// Compose the pipeline: filter → group → total
	active := filterActive(sales)
	grouped := groupByCategory(active)
	totals := categoryTotals(grouped)

	for _, t := range totals {
		fmt.Printf("  %-15s $%.2f\n", t.Category, t.Revenue)
	}

	// --- Problem B: Thermostat ---
	fmt.Println("\n=== Thermostat ===")
	thermo := NewThermostat(72.0)

	// Simulate sensor readings over time
	readings := []float64{65.0, 68.0, 71.5, 72.0, 74.0, 76.0, 73.0, 72.5}
	for _, temp := range readings {
		thermo.ReadTemperature(temp)
		thermo.UpdateMode()
		fmt.Printf("  %s\n", thermo.Status())
	}

	thermo.SetTarget(68.0)
	fmt.Println("  (target changed to 68°)")
	thermo.UpdateMode()
	fmt.Printf("  %s\n", thermo.Status())
}

// ============================================================
// Why each paradigm fits its problem:
//
// Problem A (pipeline) is functional because:
//   - There's no identity. A sale is a value, not a thing.
//   - No state changes over time. You transform data, you don't
//     mutate it.
//   - Each step is independently testable: pass in data, assert
//     on the return value.
//   - The pipeline composes naturally: filter → group → total.
//
// Problem B (thermostat) is object-oriented because:
//   - The thermostat IS a thing. It has identity — there's one
//     specific thermostat in the room, not a value you copy.
//   - Its state changes over time: the temperature changes, the
//     mode changes, the target changes.
//   - The mode depends on the internal relationship between
//     target and current — it's behavior tied to state.
//   - Construction matters: NewThermostat guarantees a valid
//     starting state.
//
// If you implemented the sales report as a class with methods
// like addSale() and getTotal(), you'd be hiding the pipeline
// behind unnecessary state management. If you implemented the
// thermostat as pure functions passing structs around, you'd be
// fighting the fact that it's fundamentally a stateful thing.
//
// The right question is always: am I computing a value, or
// managing a thing?
// ============================================================

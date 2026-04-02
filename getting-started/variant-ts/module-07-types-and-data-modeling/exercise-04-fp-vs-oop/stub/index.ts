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
// Implement this using standalone functions and array methods
// (.filter(), .reduce(), .map()). No classes. No mutation of the input.

type Sale = {
  id: string;
  category: string;
  amount: number; // cents
  refunded: boolean;
};

type CategoryTotal = {
  category: string;
  revenue: number; // dollars
};

// TODO: Implement the pipeline as standalone functions:
//   filterActive(sales: Sale[]): Sale[]
//   centsToDollars(cents: number): number
//   groupByCategory(sales: Sale[]): Map<string, Sale[]>
//   categoryTotals(grouped: Map<string, Sale[]>): CategoryTotal[]
//
// Then compose them below.

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
// time. You call methods on it: setTarget, readTemperature, updateMode.
// The mode depends on the relationship between current and target.
//
// Implement this as a class with private fields.

// TODO: Define the ThermostatMode type as a string literal union
// TODO: Define the Thermostat class with:
//   constructor(target: number)
//   setTarget(temp: number): void
//   readTemperature(current: number): void
//   updateMode(): void
//   status(): string

const sales: Sale[] = [
  { id: "001", category: "Electronics", amount: 29999, refunded: false },
  { id: "002", category: "Books", amount: 1499, refunded: false },
  { id: "003", category: "Electronics", amount: 5999, refunded: true },
  { id: "004", category: "Books", amount: 2499, refunded: false },
  { id: "005", category: "Clothing", amount: 4999, refunded: false },
  { id: "006", category: "Electronics", amount: 14999, refunded: false },
  { id: "007", category: "Clothing", amount: 7999, refunded: true },
  { id: "008", category: "Books", amount: 999, refunded: false },
];

// --- Problem A: Pipeline ---
console.log("=== Sales Report ===");
// TODO: Run the pipeline and print totals per category

// --- Problem B: Thermostat ---
console.log("\n=== Thermostat ===");
// TODO: Create a thermostat, set targets, read temperatures,
// update mode, and print status at each step

console.log("TODO: Implement both problems");

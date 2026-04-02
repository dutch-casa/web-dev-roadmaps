// ============================================================
// Problem A: Sales report pipeline (functional style)
//
// Each function takes data in and returns data out. No mutation.
// No classes. No state. Pure transformations you can compose.
// TS/JS array methods make this natural.
// ============================================================

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

const filterActive = (sales: Sale[]): Sale[] =>
  sales.filter((s) => !s.refunded);

const centsToDollars = (cents: number): number =>
  cents / 100;

const groupByCategory = (sales: Sale[]): Map<string, Sale[]> =>
  sales.reduce((grouped, sale) => {
    const existing = grouped.get(sale.category) ?? [];
    grouped.set(sale.category, [...existing, sale]);
    return grouped;
  }, new Map<string, Sale[]>());

const categoryTotals = (grouped: Map<string, Sale[]>): CategoryTotal[] =>
  Array.from(grouped.entries())
    .map(([category, sales]) => ({
      category,
      revenue: sales.reduce((sum, s) => sum + centsToDollars(s.amount), 0),
    }))
    .sort((a, b) => b.revenue - a.revenue);

// ============================================================
// Problem B: Thermostat (object-oriented style)
//
// The thermostat is a thing with identity. It has state that
// changes over time. Methods read and modify that state.
// The mode depends on the relationship between internal fields.
// ============================================================

type ThermostatMode = "idle" | "heating" | "cooling";

class Thermostat {
  #target: number;
  #current: number;
  #mode: ThermostatMode;

  constructor(target: number) {
    this.#target = target;
    this.#current = 0;
    this.#mode = "idle";
  }

  setTarget(temp: number): void {
    this.#target = temp;
  }

  readTemperature(current: number): void {
    this.#current = current;
  }

  // UpdateMode determines heating/cooling/idle based on the gap
  // between current and target temperature. The 1-degree deadband
  // prevents the system from oscillating rapidly when the temperature
  // is close to the target.
  updateMode(): void {
    if (this.#current < this.#target - 1) {
      this.#mode = "heating";
    } else if (this.#current > this.#target + 1) {
      this.#mode = "cooling";
    } else {
      this.#mode = "idle";
    }
  }

  status(): string {
    return `target=${this.#target.toFixed(1)}\u00B0  current=${this.#current.toFixed(1)}\u00B0  mode=${this.#mode}`;
  }
}

// --- Problem A: Pipeline ---
console.log("=== Sales Report ===");
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

// Compose the pipeline: filter -> group -> total
const active = filterActive(sales);
const grouped = groupByCategory(active);
const totals = categoryTotals(grouped);

for (const t of totals) {
  console.log(`  ${t.category.padEnd(15)} $${t.revenue.toFixed(2)}`);
}

// --- Problem B: Thermostat ---
console.log("\n=== Thermostat ===");
const thermo = new Thermostat(72.0);

// Simulate sensor readings over time
const readings = [65.0, 68.0, 71.5, 72.0, 74.0, 76.0, 73.0, 72.5];
for (const temp of readings) {
  thermo.readTemperature(temp);
  thermo.updateMode();
  console.log(`  ${thermo.status()}`);
}

thermo.setTarget(68.0);
console.log("  (target changed to 68\u00B0)");
thermo.updateMode();
console.log(`  ${thermo.status()}`);

// ============================================================
// Why each paradigm fits its problem:
//
// Problem A (pipeline) is functional because:
//   - There's no identity. A sale is a value, not a thing.
//   - No state changes over time. You transform data, you don't
//     mutate it.
//   - Each step is independently testable: pass in data, assert
//     on the return value.
//   - The pipeline composes naturally: filter -> group -> total.
//   - JS/TS array methods (.filter, .reduce, .map) make this
//     the native idiom. You don't fight the language.
//
// Problem B (thermostat) is object-oriented because:
//   - The thermostat IS a thing. It has identity — there's one
//     specific thermostat in the room, not a value you copy.
//   - Its state changes over time: the temperature changes, the
//     mode changes, the target changes.
//   - The mode depends on the internal relationship between
//     target and current — it's behavior tied to state.
//   - Private fields (#target, #current, #mode) protect the
//     invariants from outside interference.
//
// If you implemented the sales report as a class with methods
// like addSale() and getTotal(), you'd be hiding the pipeline
// behind unnecessary state management. If you implemented the
// thermostat as pure functions passing objects around, you'd be
// fighting the fact that it's fundamentally a stateful thing.
//
// The right question is always: am I computing a value, or
// managing a thing?
// ============================================================

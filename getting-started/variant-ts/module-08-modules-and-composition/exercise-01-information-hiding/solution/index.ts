import { Store } from "./inventory/index";

const store = new Store("Auburn Supply Co.");

const r1 = store.addItem("Notebook", 499, 50, "supplies");
const r2 = store.addItem("Pen", 149, 200, "supplies");
store.addItem("Backpack", 3999, 25, "bags");

if ("error" in r1) throw new Error(r1.error);
if ("error" in r2) throw new Error(r2.error);

const sku1 = r1.sku;
const sku2 = r2.sku;

console.log(`Store: ${store.name}`);
console.log(`Items: ${store.itemCount}`);
console.log(`Total inventory value: $${(store.totalValue() / 100).toFixed(2)}`);

console.log();

// Sell some items — through the interface, which enforces invariants.
const err1 = store.sell(sku1, 5);
if (err1) console.log("Error:", err1);

const err2 = store.sell(sku2, 10);
if (err2) console.log("Error:", err2);

// This would have been possible with the old design:
//   store.items[0].stock = -999
//
// Now it's impossible. The fields are private (#field). The only way
// to reduce stock is through sell(), which checks for negatives.

// Try to oversell — the error is informative.
const err3 = store.sell(sku1, 9999);
if (err3) console.log("Oversell blocked:", err3);

console.log(`After sales: $${(store.totalValue() / 100).toFixed(2)}`);

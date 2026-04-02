// This code reaches into the module's internals.
// After you seal the module, update this to use only the public interface.

import {
  type Store,
  addItem,
  removeStock,
  totalValue,
} from "./inventory/index";

const store: Store = {
  name: "Auburn Supply Co.",
  nextSku: 0,
  items: [],
};

// Add some items
const sku1 = addItem(store, "Notebook", 499, 50, "supplies");
const sku2 = addItem(store, "Pen", 149, 200, "supplies");
const sku3 = addItem(store, "Backpack", 3999, 25, "bags");

console.log(`Store: ${store.name}`);
console.log(`Items: ${store.items.length}`);
console.log(`Total inventory value: $${(totalValue(store) / 100).toFixed(2)}`);

console.log();

// Sell some items
const err1 = removeStock(store, sku1, 5);
if (err1) console.log("Error:", err1);

const err2 = removeStock(store, sku2, 10);
if (err2) console.log("Error:", err2);

// This is the kind of thing callers shouldn't be able to do:
// directly reaching in and breaking invariants
store.items[0].stock = -999; // whoops

console.log(`After sales: $${(totalValue(store) / 100).toFixed(2)}`);

void sku3;

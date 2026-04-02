// Information hiding exercise
//
// This module exposes everything. Every field is public. Every
// helper function is exported. The caller can reach into the guts
// of the inventory and break invariants.
//
// Your job: seal it. Make the interface as small as possible while
// keeping the same functionality. The rules:
//
//   - Only export what callers actually need
//   - Enforce invariants through the interface (e.g., stock can't go negative)
//   - Provide a constructor/factory that ensures valid initial state
//   - Every internal detail should be hidden (unexported or private)
//
// After you're done, go to index.ts and update the usage to work
// with the sealed interface.

export type Item = {
  name: string;
  priceCents: number;
  stock: number;
  category: string;
  sku: string;
};

export type Store = {
  items: Item[];
  name: string;
  nextSku: number;
};

export const generateSku = (store: Store): string => {
  store.nextSku++;
  return `SKU-${String(store.nextSku).padStart(4, "0")}`;
};

export const findItemBySku = (store: Store, sku: string): Item | undefined =>
  store.items.find((item) => item.sku === sku);

export const addItem = (
  store: Store,
  name: string,
  priceCents: number,
  stock: number,
  category: string,
): string => {
  const sku = generateSku(store);
  store.items.push({ name, priceCents, stock, category, sku });
  return sku;
};

export const removeStock = (store: Store, sku: string, quantity: number): string | null => {
  const item = findItemBySku(store, sku);
  if (!item) {
    return `item ${sku} not found`;
  }
  item.stock -= quantity; // BUG: no check for negative stock!
  return null;
};

export const totalValue = (store: Store): number =>
  store.items.reduce((total, item) => total + item.priceCents * item.stock, 0);

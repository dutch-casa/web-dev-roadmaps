// The sealed interface. Callers see: createStore, and the Store class
// methods: addItem, sell, totalValue, itemCount, name.
// Nothing else. They can't touch stock directly, can't generate
// SKUs, can't access the array of items.

type Item = {
  name: string;
  priceCents: number;
  stock: number;
  category: string;
  sku: string;
};

export class Store {
  #name: string;
  #items: Item[] = [];
  #nextSku = 0;

  constructor(name: string) {
    this.#name = name;
  }

  get name(): string {
    return this.#name;
  }

  get itemCount(): number {
    return this.#items.length;
  }

  #generateSku(): string {
    this.#nextSku++;
    return `SKU-${String(this.#nextSku).padStart(4, "0")}`;
  }

  #findBySku(sku: string): Item | undefined {
    return this.#items.find((item) => item.sku === sku);
  }

  // addItem creates a new inventory item and returns its SKU.
  // Price must be positive. Stock must be non-negative.
  addItem(name: string, priceCents: number, stock: number, category: string): string {
    if (priceCents <= 0) {
      throw new Error(`price must be positive, got ${priceCents}`);
    }
    if (stock < 0) {
      throw new Error(`stock must be non-negative, got ${stock}`);
    }
    const sku = this.#generateSku();
    this.#items.push({ name, priceCents, stock, category, sku });
    return sku;
  }

  // sell reduces stock for an item. Quantity must be positive and
  // cannot exceed available stock.
  sell(sku: string, quantity: number): string | null {
    if (quantity <= 0) {
      return `quantity must be positive, got ${quantity}`;
    }
    const item = this.#findBySku(sku);
    if (!item) {
      return `item ${sku} not found`;
    }
    if (item.stock < quantity) {
      return `insufficient stock for ${item.name}: have ${item.stock}, need ${quantity}`;
    }
    item.stock -= quantity;
    return null;
  }

  // totalValue returns the total inventory value in cents.
  totalValue(): number {
    return this.#items.reduce(
      (total, item) => total + item.priceCents * item.stock,
      0,
    );
  }
}

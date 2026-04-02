// Option B: The deep approach
//
// Three methods. Expiration is automatic. The caller never sees
// entries, timestamps, or cleanup logic.

type Entry = {
  value: string;
  expiresAt: number;
};

export class Cache {
  #entries: Map<string, Entry> = new Map();

  constructor(cleanupIntervalMs: number) {
    setInterval(() => this.#evictExpired(), cleanupIntervalMs);
  }

  set(key: string, value: string, ttlMs: number): void {
    this.#entries.set(key, { value, expiresAt: Date.now() + ttlMs });
  }

  get(key: string): string | undefined {
    const entry = this.#entries.get(key);
    if (!entry) return undefined;
    if (Date.now() > entry.expiresAt) {
      this.#entries.delete(key);
      return undefined;
    }
    return entry.value;
  }

  delete(key: string): void {
    this.#entries.delete(key);
  }

  #evictExpired(): void {
    const now = Date.now();
    for (const [key, entry] of this.#entries) {
      if (now > entry.expiresAt) {
        this.#entries.delete(key);
      }
    }
  }
}

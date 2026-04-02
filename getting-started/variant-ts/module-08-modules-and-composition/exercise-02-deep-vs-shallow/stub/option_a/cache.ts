// Option A: The shallow approach
//
// Every internal detail is exposed. The caller manages expiration,
// creates entries, checks timestamps, and drives the cleanup.

export type CacheEntry = {
  key: string;
  value: string;
  createdAt: number;  // Date.now() timestamp
  expiresAt: number;  // Date.now() timestamp
  isExpired: boolean;
};

export class Cache {
  entries: Map<string, CacheEntry> = new Map();

  putEntry(entry: CacheEntry): void {
    this.entries.set(entry.key, entry);
  }

  getEntry(key: string): CacheEntry | undefined {
    return this.entries.get(key);
  }

  deleteEntry(key: string): void {
    this.entries.delete(key);
  }

  cleanupExpired(): number {
    let removed = 0;
    for (const [key, entry] of this.entries) {
      if (checkExpired(entry)) {
        this.entries.delete(key);
        removed++;
      }
    }
    return removed;
  }

  size(): number {
    return this.entries.size;
  }
}

export const createEntry = (key: string, value: string, ttlMs: number): CacheEntry => {
  const now = Date.now();
  return {
    key,
    value,
    createdAt: now,
    expiresAt: now + ttlMs,
    isExpired: false,
  };
};

export const checkExpired = (entry: CacheEntry): boolean =>
  Date.now() > entry.expiresAt;

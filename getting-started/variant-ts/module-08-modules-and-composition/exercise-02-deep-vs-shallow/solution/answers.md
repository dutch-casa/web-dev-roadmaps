# Deep vs. shallow -- answers

## Counting exported names

**Option A (shallow):** 8 exported names
- Types: `CacheEntry` (1)
- Classes: `Cache` (1)
- Functions: `createEntry`, `checkExpired` (2)
- Methods on Cache: `putEntry`, `getEntry`, `deleteEntry`, `cleanupExpired`, `size` (5)
- Public fields on Cache: `entries` (1)
- Fields on CacheEntry: `key`, `value`, `createdAt`, `expiresAt`, `isExpired` (5)

A caller must understand 14+ concepts to use this module correctly.

**Option B (deep):** 2 exported names
- Classes: `Cache` (1)
- Methods: `set`, `get`, `delete` (3)
- Public fields: none

A caller must understand 4 concepts.

## Caller code for a basic workflow

**Option A:**
```typescript
import { Cache, createEntry, checkExpired } from "./option_a/cache.ts";

const cache = new Cache();
const entry = createEntry("user:123", "Jordan", 5 * 60 * 1000);
cache.putEntry(entry);
// ... later ...
const got = cache.getEntry("user:123");
if (got && !checkExpired(got)) {
  console.log(got.value);
}
// caller must remember to clean up
cache.cleanupExpired();
```

7 lines. The caller manages entry creation, expiration checking, and cleanup.

**Option B:**
```typescript
import { Cache } from "./option_b/cache.ts";

const cache = new Cache(60_000);
cache.set("user:123", "Jordan", 5 * 60 * 1000);
// ... later ...
const val = cache.get("user:123");
if (val !== undefined) {
  console.log(val);
}
```

4 lines. Expiration is handled automatically. No cleanup to remember.

## Which is deeper?

**Option B is deeper.** Simple interface (4 names), significant hidden complexity (automatic cleanup via setInterval, transparent expiration on read, private fields). The caller doesn't need to think about timestamps, entry objects, or cleanup schedules.

**Option A is shallower.** The interface is almost as complex as the implementation. The caller must understand and manage expiration manually. The `CacheEntry` type leaks internal details. The `checkExpired` function exists because the design didn't handle expiration automatically.

## What if you change the storage to a database?

**Option A:** Every caller that touches `CacheEntry`, `entries`, `checkExpired`, or `cleanupExpired` breaks. The internal data structure is part of the interface.

**Option B:** Only the internals change. The `Entry` type is private. The `#entries` map is private. Callers use `set`, `get`, `delete` -- those signatures don't change. Zero caller code changes.

## Which one to maintain?

Option B. It's easier to change the internals without breaking callers, easier for callers to use correctly, and harder for callers to misuse. The shallow version invites bugs (forgetting to check expiration, forgetting to run cleanup).

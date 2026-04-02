# Deep vs. shallow — answers

## Counting exported names

**Option A (shallow):** 11 exported names
- Types: `Cache`, `CacheEntry` (2)
- Functions: `NewCache`, `CreateEntry`, `CheckExpired` (3)
- Methods: `PutEntry`, `GetEntry`, `DeleteEntry`, `CleanupExpired`, `Size` (5)
- Exported fields on `CacheEntry`: `Key`, `Value`, `CreatedAt`, `ExpiresAt`, `IsExpired` (5)
- Exported fields on `Cache`: `Entries` (1)

A caller must understand 17+ concepts to use this package correctly.

**Option B (deep):** 4 exported names
- Types: `Cache` (1)
- Functions: `New` (1)
- Methods: `Set`, `Get`, `Delete` (3)
- Exported fields: none

A caller must understand 4 concepts.

## Caller code for a basic workflow

**Option A:**
```go
cache := option_a.NewCache()
entry := option_a.CreateEntry("user:123", "Jordan", 5*time.Minute)
cache.PutEntry(entry)
// ... later ...
got, ok := cache.GetEntry("user:123")
if ok && !option_a.CheckExpired(got) {
    fmt.Println(got.Value)
}
// caller must remember to clean up
cache.CleanupExpired()
```

7 lines. The caller manages entry creation, expiration checking, and cleanup.

**Option B:**
```go
cache := option_b.New(time.Minute)
cache.Set("user:123", "Jordan", 5*time.Minute)
// ... later ...
if val, ok := cache.Get("user:123"); ok {
    fmt.Println(val)
}
```

4 lines. Expiration is handled automatically. No cleanup to remember.

## Which is deeper?

**Option B is deeper.** Simple interface (4 names), significant hidden complexity (automatic cleanup, transparent expiration on read, background goroutine). The caller doesn't need to think about timestamps, entry structs, or cleanup schedules.

**Option A is shallower.** The interface is almost as complex as the implementation. The caller must understand and manage expiration manually. The `CacheEntry` struct leaks internal details. The `CheckExpired` function exists because the design didn't handle expiration automatically.

## What if you change the storage to a database?

**Option A:** Every caller that touches `CacheEntry`, `Entries`, `CheckExpired`, or `CleanupExpired` breaks. The internal data structure is part of the interface.

**Option B:** Only the internals change. The `entry` type is unexported. The `entries` map is unexported. Callers use `Set`, `Get`, `Delete` — those signatures don't change. Zero caller code changes.

## Which one to maintain?

Option B. It's easier to change the internals without breaking callers, easier for callers to use correctly, and harder for callers to misuse. The shallow version invites bugs (forgetting to check expiration, forgetting to run cleanup).

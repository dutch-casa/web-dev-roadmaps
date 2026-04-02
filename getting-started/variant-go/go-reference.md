# Go Language Reference

Keep this open while you work through the exercises. Not a tutorial — a lookup tool.

---

## Variables and constants

```go
const maxRetries = 3                // constant — cannot change
const (                             // iota: auto-incrementing constants
    Red   = iota                    // 0
    Green                           // 1
    Blue                            // 2
)

var name string                     // explicit type, zero value ""
var count = 10                      // type inferred
age := 25                           // short declaration (functions only)
```

**When to use which:**

| Form | When |
|------|------|
| `const` | Value known at compile time, never changes |
| `var` | Need zero value, or package-level variable |
| `:=` | Inside a function, type is obvious from the right side |

**Zero values — every type has one:**

| Type | Zero value |
|------|-----------|
| `int`, `float64` | `0` |
| `string` | `""` |
| `bool` | `false` |
| pointer, slice, map, func, interface, channel | `nil` |

> **Gotcha:** `:=` only works inside functions. At package level you must use `var`.

---

## Types

| Type | What it is | Literal |
|------|-----------|---------|
| `int` | Whole numbers | `42` |
| `float64` | Decimals | `3.14` |
| `string` | Text | `"hello"` |
| `bool` | Logic | `true` |
| `byte` | Alias for `uint8` | `'A'` |
| `rune` | Alias for `int32` (Unicode codepoint) | `'界'` |

**Conversions are explicit** -- Go never converts for you: `int(3.14)` truncates to 3. `strconv.Itoa(42)` gives `"42"`. `strconv.Atoi("42")` gives `42, nil`.

**Type definition vs alias:** `type Celsius float64` creates a new type (needs conversion). `type F = float64` is an alias (identical to float64).

---

## Scope

| Scope | Declared where | Visible where |
|-------|---------------|--------------|
| Package | Outside any function | Entire package |
| Function | Inside a function | That function |
| Block | Inside `if`/`for`/`switch` | That block only |

**Exported vs unexported:** Uppercase `ProcessOrder` = visible outside the package. Lowercase `processOrder` = package-private.

> **Gotcha:** Shadowing. An inner `:=` silently creates a new variable that hides an outer one. The compiler won't warn you.

```go
err := doFirst()
if true {
    err := doSecond()   // new variable — shadows the outer err
    _ = err
}
// outer err still from doFirst()
```

---

## Strings

```go
s := "hello\nworld"      // interpreted — escape sequences work
r := `hello\nworld`      // raw — backslash is literal, great for regex/paths
msg := fmt.Sprintf("score: %d, name: %s, ratio: %.2f", 85, "Gus", 0.75)
```

**strings package essentials:**

| Function | What it does |
|----------|-------------|
| `strings.Contains(s, "go")` | Does s contain substring? |
| `strings.HasPrefix(s, "go")` | Starts with? |
| `strings.HasSuffix(s, ".go")` | Ends with? |
| `strings.Join([]string{"a","b"}, ",")` | `"a,b"` |
| `strings.Split("a,b", ",")` | `["a", "b"]` |
| `strings.TrimSpace(" hi ")` | `"hi"` |
| `strings.ToUpper(s)` / `ToLower(s)` | Case conversion |
| `strings.Replace(s, "old", "new", -1)` | Replace all (-1 = no limit) |
| `strings.Index(s, "sub")` | Position of first match, -1 if absent |

**Runes vs bytes:**

`len("Go界")` is 5 (bytes). `utf8.RuneCountInString("Go界")` is 3 (characters). `for _, r := range s` gives runes. `s[i]` gives a byte.

> **Gotcha:** `s[i]` gives you a byte, not a character. For Unicode-safe iteration, always use `range`.

---

## Slices

```go
nums := []int{1, 2, 3}          // literal
zeros := make([]int, 5)         // [0 0 0 0 0] — length 5
buf := make([]int, 0, 10)       // length 0, capacity 10
```

```go
nums = append(nums, 4, 5)       // always reassign the result
len(nums)                        // 5
nums[0]                          // 1
sub := nums[1:3]                 // [2, 3] — start inclusive, end exclusive
```

**Iteration:** `for i, v := range nums` (index + value), `for _, v := range nums` (value only), `for i := range 10` (0..9, Go 1.22+).

**slices package (Go 1.21+):**

| Function | Does |
|----------|------|
| `slices.Contains(s, 3)` | Is 3 in the slice? |
| `slices.Index(s, 3)` | Position of 3, or -1 |
| `slices.Sort(s)` | Sort in place |
| `slices.SortFunc(s, cmp)` | Custom sort (`cmp` returns neg/zero/pos) |

**Nil vs empty:** `var s []int` is nil. `s := []int{}` is empty. Both have `len == 0` and both work with `append`. Only nil equals `nil`.

> **Gotcha:** `append` returns a new slice. If you write `append(s, v)` without assigning back, the value is lost.

---

## Maps

```go
m := map[string]int{"a": 1, "b": 2}   // literal
m2 := make(map[string]int)             // empty, ready to use
v := m["a"]                      // returns zero value if missing
v, ok := m["z"]                  // comma-ok: 0, false
m["c"] = 3                       // set
delete(m, "a")                   // remove
len(m)                           // count
for k, v := range m { }         // iteration — order is random
```

> **Gotcha:** A nil map reads fine (returns zero values) but panics on write. Always initialize with `make` or a literal before writing.

---

## Structs

```go
type Point struct {
    X, Y float64
}

p := Point{X: 1, Y: 2}         // always use field names
fmt.Println(p.X)                // 1
```

**Methods — pointer vs value receiver:**

```go
func (p Point) Distance() float64  { return math.Sqrt(p.X*p.X + p.Y*p.Y) }  // value — read only
func (p *Point) Scale(f float64)   { p.X *= f; p.Y *= f }                    // pointer — mutates
```

| Receiver | Can modify? | Call on value? | Call on pointer? |
|----------|------------|---------------|-----------------|
| Value `(p Point)` | No | Yes | Yes |
| Pointer `(p *Point)` | Yes | Yes (auto &) | Yes |

**Embedding (composition, not inheritance):**

```go
type Circle struct {
    Point                        // embedded — Circle gets X, Y, Distance()
    Radius float64
}
c := Circle{Point: Point{X: 1, Y: 2}, Radius: 5}
c.X                              // promoted from Point
```

**JSON tags:** `` `json:"name"` ``, `` `json:"name,omitempty"` ``. Unexported fields are invisible to `encoding/json`.

---

## Functions

```go
func add(a, b int) int { return a + b }                         // basic
func divide(a, b float64) (float64, error) { ... }              // multiple returns
func sum(nums ...int) int { ... }                                // variadic
```

**Multiple returns (the idiomatic error pattern):**

```go
result, err := divide(10, 0)
if err != nil {
    return err
}
```

**First-class functions and closures:**

```go
double := func(n int) int { return n * 2 }   // functions are values
func counter() func() int {                   // closure captures n
    n := 0
    return func() int { n++; return n }
}
```

**defer — runs when the function returns, LIFO order:**

```go
f, err := os.Open(path)
if err != nil { return err }
defer f.Close()                               // guaranteed cleanup
```

> **Gotcha:** `defer` captures arguments at the time of the defer statement, not when it runs. `defer fmt.Println(x)` captures the current value of x.

---

## Interfaces

**Implicit satisfaction — no `implements` keyword:**

```go
type Stringer interface { String() string }

type Dog struct{ Name string }
func (d Dog) String() string { return d.Name }   // Dog satisfies Stringer
```

**Type assertions and type switches:**

```go
var x any                        // any == interface{} — holds anything
x = 42
n, ok := x.(int)                // 42, true
s, ok := x.(string)             // "", false

switch v := x.(type) {          // type switch
case int:    fmt.Println(v)
case string: fmt.Println(v)
}
```

**Key interfaces:** `io.Reader` (Read), `io.Writer` (Write), `fmt.Stringer` (String), `error` (Error). Keep your own interfaces small -- one or two methods. Accept interfaces, return structs.

---

## Error handling

```go
err := errors.New("not found")               // create
err = fmt.Errorf("user %d: %w", id, err)    // wrap with context
if err != nil { return err }                  // check immediately

if errors.Is(err, os.ErrNotExist) { }        // match sentinel
var pe *os.PathError
if errors.As(err, &pe) { }                   // match type
```

**Custom error type:**

```go
type ValidationError struct { Field, Message string }
func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

**Rules:** Lowercase, no punctuation (`"connection refused"`). Wrap with context (`fmt.Errorf("open config: %w", err)`). Check immediately.

> **Gotcha:** `panic`/`recover` exist. Almost never use them. They are not exceptions.

---

## Control flow

**if** (with optional init statement):

```go
if err := doThing(); err != nil { return err }   // err scoped to if/else
```

**for — the only loop keyword:**

| Form | Use |
|------|-----|
| `for { }` | Infinite |
| `for cond { }` | While |
| `for i := 0; i < n; i++ { }` | Classic |
| `for i, v := range slice { }` | Slice |
| `for k, v := range m { }` | Map (random order) |
| `for i, r := range "Go界" { }` | String (runes) |
| `for i := range 10 { }` | Integer (Go 1.22+) |

**switch — no fallthrough by default:**

```go
switch status {
case "ok":    handle()
case "error": fail()
default:      unknown()
}

switch {                             // tagless — replaces if/else chains
case n < 0:  return "negative"
case n == 0: return "zero"
default:     return "positive"
}
```

**select** — multiplexes channel operations. Covered in the Backend track.

---

## Packages and imports

```go
package main                         // entry point — must have func main()

import (                             // group imports
    "fmt"
    "strings"
)
import _ "net/http/pprof"            // blank import — side effects only
```

| Command | Does |
|---------|------|
| `go mod init github.com/you/project` | Create module |
| `go mod tidy` | Sync dependencies |
| `go run .` | Run current package |
| `go build .` | Compile |
| `go test ./...` | Test everything |

Code under `internal/` is only importable by the parent module. Use it to hide implementation details.

---

## Concurrency (brief intro)

```go
go doWork()                          // launch a goroutine

ch := make(chan int)                  // unbuffered channel
ch <- 42                             // send (blocks until received)
v := <-ch                            // receive (blocks until sent)
bch := make(chan int, 5)             // buffered — up to 5 before blocking
```

**sync.WaitGroup** — call `Add` before launching, `Done` inside (via defer), `Wait` to block.

**sync.Mutex** — `Lock()` / `Unlock()` around shared state. Keep the critical section small.

---

## Common patterns

**Comma-ok idiom** -- used in map lookups (`v, ok := m[key]`), type assertions (`v, ok := x.(T)`), and channel receives (`v, ok := <-ch`).

**Constructor pattern:** `func NewThing(args) *Thing { return &Thing{...} }`

**Table-driven tests:**

```go
tests := []struct{ a, b, want int }{
    {1, 2, 3}, {0, 0, 0}, {-1, 1, 0},
}
for _, tt := range tests {
    if got := add(tt.a, tt.b); got != tt.want {
        t.Errorf("add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
    }
}
```

**Functional options pattern** -- look this up when you need configurable constructors with clean APIs.

---

## Style rules

| Rule | Example |
|------|---------|
| `const` by default | `const maxRetries = 3` |
| Guard clauses, early return | `if err != nil { return err }` |
| No `else` after `return` | Return early, then continue with the happy path |
| `gofmt` is non-negotiable | Run it. Always. `gofmt -w .` |
| Uppercase = exported | `ProcessOrder` visible outside package |
| Lowercase = unexported | `processOrder` package-private |
| Error messages: lowercase, no punctuation | `"connection refused"` not `"Connection refused."` |
| Receiver names: short, not `this`/`self` | `func (s *Server)` not `func (this *Server)` |
| Accept interfaces, return structs | Functions take narrow inputs, return concrete types |

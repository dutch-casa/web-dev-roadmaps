# 08 · Modules & Composition

```mermaid
---
config:
  layout: elk
---
flowchart LR
  A["01 · Environment"] --> B["02 · Git"]
  B --> C["03 · Values & Declarations"]
  C --> D["04 · Naming"]
  D --> E["05 · Functions & Decomposition"]
  E --> F["06 · Control Flow"]
  F --> G["07 · Types & Data Modeling"]
  G --> H["08 · Modules & Composition"]
  H --> I["09 · Capstone: Terminal Wordle"]
  style H fill:#f90,color:#000
```

*In Module 07, you learned to model data so invalid states can't exist. Now you'll learn to draw boundaries around that data — hiding complexity so that each piece of your program can be understood on its own.*

This is the architectural module. Everything before this taught you how to write good code inside a function, inside a type. This module teaches you where to draw the lines *between* those things. Get the boundaries right, and each part of a system becomes understandable in isolation. Get them wrong, and every change ripples everywhere.

The ideas here — information hiding, deep modules, temporal decomposition, pulling complexity downward — are foundational to software engineering. They originate in a 1972 paper on decomposing systems into modules and were refined into a practical design philosophy over the decades since. The Resources section at the bottom has the primary sources.

MIT's 6.033 (Computer System Engineering) teaches these same principles at the systems level: modularity reduces fate-sharing, abstraction separates interface from implementation, and enforced boundaries are what make large systems possible. We're applying that same thinking to Go packages.

---

## 1. What a module actually is

A module is any unit of code with two faces: an **interface** and an **implementation**.

The interface is everything a developer working in a *different* module must know in order to use this one. The implementation is the code that carries out the promises made by the interface.

The value of a module is a ratio:

> **The best modules are those that provide powerful functionality yet have simple interfaces.** The benefit of a module is its functionality. The cost of a module (in terms of system complexity) is its interface.

A module that hides a lot behind a narrow opening is valuable. A module whose interface is nearly as complex as its implementation saves you nothing — you still have to understand the whole thing.

In Go, the concrete unit is the **package**. A package's interface is its exported names: types, functions, constants, and variables that start with an uppercase letter. Its implementation is everything unexported. The compiler enforces the boundary. You don't need a convention or a code review to maintain it — lowercase means hidden, period.

---

## 2. Deep modules vs. shallow modules

The most useful mental model for modules. Picture a module as a rectangle. The top edge is the interface — its width represents how much a caller needs to know. The body is the implementation — its height represents how much work the module does.

```
Deep module:              Shallow module:

┌───────┐                 ┌────────────────────────────┐
│       │                 │                            │
│       │                 └────────────────────────────┘
│       │
│       │                 Wide interface, thin implementation.
│       │                 The caller learns almost as much
│       │                 as the module contains.
│       │
│       │
└───────┘

Narrow interface,
deep implementation.
Enormous value.
```

**Deep: Unix file I/O.** Five functions — `open`, `read`, `write`, `lseek`, `close`. Behind that interface: file systems (ext4, NTFS, ZFS), disk drivers, block caching, permissions, journaling, network file systems, RAID controllers. The caller knows none of this. The caller calls `open` and gets a file descriptor. The depth is staggering; the interface is five functions.

**Shallow: Java's classic I/O.** Want to read lines from a file with buffering?

```java
BufferedReader br = new BufferedReader(
    new InputStreamReader(
        new FileInputStream("data.txt")));
```

Three classes layered on each other, each with its own interface to learn. The complexity that should be hidden is pushed onto the caller. Every layer is shallow — it adds a thin wrapper and a new set of concepts.

Go's `os.Open` + `bufio.NewScanner` is not as deep as Unix's raw I/O, but it's deeper than Java's stack:

```go
f, err := os.Open("data.txt")
if err != nil {
    return err
}
defer f.Close()

scanner := bufio.NewScanner(f)
for scanner.Scan() {
    line := scanner.Text()
    // use line
}
return scanner.Err()
```

Two objects. Two concepts. That's the interface. Behind it: buffered reads, newline splitting, memory management, error propagation.

**The measure**: how much does a caller need to know to use this module? Less is better.

A Go proverb says the same thing from the interface side: **"The bigger the interface, the weaker the abstraction."** An interface with one method (`io.Reader`) has hundreds of implementations. An interface with twenty methods has maybe one. Small interfaces make deep modules possible.

---

## 3. Information hiding

This is the oldest and most important idea in software engineering. The 1972 paper that introduced it states:

> Every module is characterized by its knowledge of a design decision which it hides from all others. Its interface or definition is chosen to reveal as little as possible about its inner workings.

A design decision — the file format, the scoring algorithm, the storage layout — lives in exactly one module. Change that decision, and only that module changes. The rest of the system is insulated.

Go makes this structural, not just a convention:

```go
package game

// Feedback is exported — callers see it.
type Feedback struct {
    Guess   string
    Letters [5]LetterResult
}

// LetterResult is exported — callers need it to render output.
type LetterResult int

const (
    Absent LetterResult = iota
    Misplaced
    Correct
)

// score is unexported — the scoring algorithm is hidden.
// Callers can't call it, depend on it, or break if it changes.
func score(guess, answer string) Feedback {
    var fb Feedback
    fb.Guess = guess
    // ... scoring logic ...
    return fb
}

// Check is exported — this is the interface.
// It takes a guess and the answer, returns feedback.
// The caller doesn't know how scoring works. Doesn't need to.
func Check(guess, answer string) Feedback {
    return score(guess, answer)
}
```

The scoring algorithm is a design decision. It lives inside `score`, which is unexported. If you change from "check exact matches first, then misplaced" to some other strategy, nothing outside the `game` package notices. That's information hiding at work. The caller sees `Check` — one function, two strings in, one struct out. The implementation can be as intricate as it needs to be.

**Information leakage** is the failure mode. It happens when a design decision shows up in multiple modules. If the scoring logic leaks into the UI layer — say the terminal code knows to check exact matches before misplaced ones — then changing the scoring algorithm requires changing two packages. The boundary failed. The modules are coupled by shared knowledge that should have been private.

---

## 4. Temporal decomposition (anti-pattern)

This is the most common way beginners draw wrong boundaries. It is a major red flag.

**Temporal decomposition** means splitting code into modules based on *when* things happen rather than *what knowledge* each module needs. "First we read the file, then we process it, then we write it" — so we make three modules: `reader`, `processor`, `writer`.

The problem: the reader and writer both need to know the file format. That knowledge now lives in two places. Change the format, and two modules change. The boundary was drawn along the timeline instead of along lines of knowledge.

```go
// BAD: temporal decomposition.
// reader and writer both know the file format.

package reader

func ReadConfig(path string) (map[string]string, error) {
    // knows: file path, file format (key=value lines),
    // parsing rules, error handling
}

package writer

func WriteConfig(path string, data map[string]string) error {
    // knows: file path, file format (key=value lines),
    // serialization rules, error handling
}
```

The file format is a design decision. It should live in one place:

```go
// GOOD: one module owns the format.

package config

// Load reads a config file. The format is hidden.
func Load(path string) (map[string]string, error) { ... }

// Save writes a config file. The format is hidden.
func Save(path string, data map[string]string) error { ... }
```

Now the format is a private detail of the `config` package. Change from `key=value` to JSON, and only `config` changes. The reader and writer are in the same module because they share the same knowledge — not because they happen at different times.

Since, in most cases, design decisions transcend time of execution, modules will not correspond to steps in the processing.

---

## 5. Pass-through methods (anti-pattern)

Another red flag. A pass-through method does little except call another method with the same or nearly the same signature. It adds interface surface without adding depth.

```go
// BAD: pass-through. GameService.SubmitGuess does nothing
// except call Game.SubmitGuess with the same arguments.

package service

type GameService struct {
    game *game.Game
}

func (s *GameService) SubmitGuess(guess string) (game.Feedback, error) {
    return s.game.SubmitGuess(guess)
}
```

This is a shallow module by definition. Its interface is the same size as its implementation. It adds a name the caller must learn, a type the caller must construct, and a layer the caller must navigate — all to do nothing. The existence of `GameService` means someone drew a boundary where there was no design decision to hide.

The fix is usually to eliminate the pass-through layer entirely and let callers use the underlying module directly. If you find yourself writing a method that just delegates, ask: *what does this layer know that the caller doesn't?* If the answer is nothing, the layer shouldn't exist.

Each layer in a system should provide a different abstraction. If two adjacent layers have the same interface, at least one of them is not useful. This is the "different layer, different abstraction" rule.

---

## 6. Pull complexity downward

The most direct piece of architectural advice:

> It is more important for a module to have a simple interface than a simple implementation.

When you can handle complexity inside a module rather than exposing it through the interface, do it. Suffer as the implementer so your callers don't have to.

**Configuration parameters** are the canonical failure to pull complexity down. Instead of computing a reasonable default, the developer punts the decision to the caller:

```go
// BAD: pushing complexity upward.
// The caller now has to know what a good timeout is.
func NewClient(host string, port int, timeout time.Duration,
    retries int, backoffFactor float64) *Client { ... }
```

```go
// GOOD: pulling complexity downward.
// Reasonable defaults computed inside.
// The caller provides only what they actually know: the host.
func NewClient(host string) *Client {
    return &Client{
        host:          host,
        port:          443,
        timeout:       10 * time.Second,
        retries:       3,
        backoffFactor: 2.0,
    }
}
```

Ask yourself: *does the caller really know better, or are you just too lazy to figure out a good default?* Most of the time, the module author is the one with the expertise to make the call. Pull that decision down into the implementation.

This doesn't mean you can never offer configuration. It means the default path should require zero configuration. Callers who need to override can use an option pattern or a separate config struct. But the simple case — which is most cases — should be one line.

---

## 7. Functional core / imperative shell

Everything above converges on one architecture: separate pure logic from I/O.

The **functional core** is the deepest module in your system. Business rules, data transformations, scoring algorithms. It takes values in and returns values out. No printing, no file reading, no network calls. It's pure, testable, and portable.

The **imperative shell** is thin. It reads from the outside world (stdin, files, HTTP), calls the core, and writes results back. It's impure but simple. It has almost no logic of its own.

```
┌──────────────────────────────────────────┐
│ Imperative shell (I/O, effects)          │
│                                          │
│  Read input → ┌──────────────────┐       │
│               │ Functional core  │       │
│               │ (pure logic)     │       │
│  Write output ← └──────────────────┘      │
│                                          │
└──────────────────────────────────────────┘
```

In the Wordle capstone you'll build in Module 09, this looks like:

- **`game/`** — the functional core. It knows the word list, the scoring algorithm, the rules about six guesses. It knows nothing about terminals, colors, or user input. Given a guess string, it returns a `Feedback` struct. Pure in, pure out.
- **`ui/`** — the imperative shell. It reads from `stdin`, calls `game.Check`, renders colored output. It knows nothing about how scoring works.
- **`main.go`** — composition root. It wires `game` and `ui` together. Three lines of real code.

The `game` package is a deep module: narrow interface (`Check`), deep implementation (word matching, letter frequency handling, state tracking). The `ui` package is deliberately shallow — it exists to touch the impure world. `main.go` has no logic at all; it's pure wiring.

This is information hiding applied end-to-end. Change the scoring algorithm, and `ui/` doesn't change. Change the terminal library, and `game/` doesn't change. Swap the terminal for a web frontend, and `game/` still doesn't change — it never knew about terminals in the first place.

---

## 8. When to create a Go package

A package represents a **domain concept**, not an architectural layer.

**Symptoms of wrong boundaries:**

| Package name | Problem |
|---|---|
| `utils` | Not a concept. A drawer of unrelated things. Every function here belongs somewhere else. |
| `helpers` | Same as `utils`. |
| `models` | An architectural layer, not a domain concept. The types in `models` are actually about accounts, orders, games — put them there. |
| `common` | Shared knowledge often means shared coupling. Ask what each piece is *about* and move it there. |

**Signs of a good package:**

- You can explain what it does in one sentence **without the word "and"**. If you need "and," you have two packages.
- The name is a noun that describes a domain concept: `account`, `game`, `auth`, `inventory`.
- The exported surface is smaller than the implementation. If everything is exported, the package isn't hiding anything.
- Callers don't need to read the source code to use it. The exported names and their types tell the whole story.

From [Effective Go](https://go.dev/doc/effective_go#package-names): package names should be short, concise, lowercase, single-word. No underscores, no mixedCaps. The package name is part of every qualified reference — `bufio.Reader`, not `bufio.BufReader` — so avoid stuttering.

From [How to Write Go Code](https://go.dev/doc/code): a package is all the `.go` files in a single directory, compiled together. A module is a collection of packages with a shared `go.mod`. One repository typically contains one module.

Create a new package when you have a design decision to hide — a file format, a scoring algorithm, a protocol, an external service. If there's no decision to hide, there's no boundary to draw.

---

## Putting it together: how this connects to what you've learned

| Module | What it taught | How Module 08 uses it |
|---|---|---|
| 03 Values & Declarations | Constants, variables, zero values | A module's interface is its exported declarations |
| 04 Naming | Names carry meaning | Package names *are* the module's first impression |
| 05 Functions & Decomposition | Break work into functions | A module's depth comes from well-decomposed internals |
| 06 Control Flow | Guard clauses, clear paths | Implementation complexity stays inside the module |
| 07 Types & Data Modeling | Make invalid states unrepresentable | Types enforce the module's invariants at the boundary |

Module 08 doesn't add new syntax. It adds the discipline of deciding *where things go* — which is harder than any syntax.

---

## Exercises

1. **[Information hiding](exercise-01-information-hiding/)** — seal a module by reducing its exported surface to only what callers need
2. **[Deep vs. shallow](exercise-02-deep-vs-shallow/)** — compare two implementations and measure their interface complexity
3. **[Boundary drawing](exercise-03-boundary-drawing/)** — split a program into packages at the natural seams

---

## Resources

- [Go — How to Write Go Code](https://go.dev/doc/code) — official guide to packages and modules
- [Go — Effective Go: Package names](https://go.dev/doc/effective_go#package-names) — naming packages well
- [Parnas — "On the Criteria to Be Used in Decomposing Systems into Modules" (1972)](https://dl.acm.org/doi/10.1145/361598.361623) — the origin of information hiding
- [Ousterhout — *A Philosophy of Software Design*](https://web.stanford.edu/~ouster/cgi-bin/aposd.php) — chapters 4-8 are directly relevant
- [MIT 6.033 — Computer System Engineering](https://ocw.mit.edu/courses/6-033-computer-system-engineering-spring-2018/) — modularity and abstraction at the systems level
- [Rob Pike — Go Proverbs](https://go-proverbs.github.io/) — "The bigger the interface, the weaker the abstraction"

---

## What's next

Module 09 is the capstone. You'll build a complete terminal Wordle game — and every boundary decision you make will be a test of what you learned here. The `game` package will be your deep module. The `ui` package will be your thin shell. `main.go` will wire them together. If you drew the boundaries right, changing the scoring algorithm won't touch the UI, and changing the terminal library won't touch the game logic. That's the whole point.

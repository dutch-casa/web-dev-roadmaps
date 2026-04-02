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

*In Module 07, you learned to model data so invalid states can't exist. Now you'll learn to draw boundaries around that data — hiding complexity so each piece of your program can be understood on its own.*

Two APIs for a cache. Which one would you rather learn?

```
API A (11 exported names):               API B (4 exported names):
  Cache                                    Cache
  CacheConfig                              create
  CacheOption                              get
  createCache                              set
  withTTL
  withMaxSize
  withEvictionPolicy
  get
  set
  remove
  flush
```

API B hides more behind a narrower opening. You learn four things and you're productive. API A demands you understand configuration types, option patterns, and eviction policies before you store a single value. The second cache might do just as much work — it just doesn't make *you* do the work of understanding its internals.

That ratio — power of implementation divided by complexity of interface — is the measure of a good module.

## Deep modules vs. shallow modules

Picture a module as a rectangle. The top edge is the interface (what callers learn). The body is the implementation (what the module does).

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
└───────┘

Narrow interface,
deep implementation.
```

**Deep: Unix file I/O.** Five functions — `open`, `read`, `write`, `lseek`, `close`. Behind that: file systems, disk drivers, block caching, permissions, journaling, network file systems. The caller calls `open` and gets a file descriptor. The depth is staggering; the interface is five functions.

**Shallow: a pass-through wrapper.** A `GameService` whose `submitGuess` just calls `Game.submitGuess` with the same arguments. It adds a name to learn, a type to construct, and a layer to navigate — all to do nothing. If a layer doesn't know something the caller doesn't, the layer shouldn't exist.

## ES modules: `export` and `import`

In TypeScript (and modern JavaScript), the file is the module boundary. Everything in a file is private by default. You choose what to expose with `export`:

```ts
// game/evaluate.ts

// Exported — this is the interface
export const evaluateGuess = (target: string, guess: string): LetterResult[] => {
  return score(target, guess);
};

// Not exported — the algorithm is hidden
// Callers can't import it, depend on it, or break if it changes
const score = (target: string, guess: string): LetterResult[] => {
  // ... scoring logic ...
};
```

Callers import what they need:

```ts
// main.ts
import { evaluateGuess } from "./game/evaluate";
```

### Named exports vs. default exports

```ts
// Named export — explicit, greppable, refactor-friendly
export const parse = (input: string): Result => { ... };
export type Config = { host: string; port: number };

// Default export — one per file, name is up to the importer
export default class Game { ... }
```

Default to named exports. They're explicit — the name is the same everywhere. Default exports let importers pick any name, which fragments search results and makes renames invisible.

### Barrel files — and why they're usually wrong

A barrel file re-exports from multiple files:

```ts
// game/index.ts
export { evaluateGuess } from "./evaluate";
export { isValidWord, randomWord } from "./words";
export type { LetterResult, GameState } from "./types";
```

The appeal: `import { evaluateGuess, isValidWord } from "./game"` instead of two separate imports.

The cost: bundlers and tree-shakers struggle with barrels. Adding an unused export to the barrel can pull in code the consumer never wanted. Circular dependencies sneak in. Barrel files hide what depends on what. Use them sparingly — for genuinely public API boundaries, not for convenience.

## Information hiding

A design decision — the file format, the scoring algorithm, the storage layout — lives in exactly one module. Change that decision, and only that module changes. The rest of the system is insulated.

If the scoring algorithm leaks into the UI layer — say the terminal code knows to check exact matches before misplaced ones — then changing the algorithm requires changing two files. The boundary failed. The modules are coupled by shared knowledge that should have been private.

**Temporal decomposition** is the most common way beginners draw wrong boundaries. "First we read, then we process, then we write" — so three modules: `reader.ts`, `processor.ts`, `writer.ts`. But the reader and writer both know the file format. That knowledge now lives in two places. The boundary was drawn along the timeline instead of along lines of knowledge. Put them in one module: `config.ts` with `load` and `save`.

## Functional core / imperative shell

Everything above converges on one architecture: separate pure logic from I/O. Module 05's pure vs. impure distinction, applied at the module level.

The **functional core** is the deepest module. Business rules, data transforms, scoring algorithms. Values in, values out. No `console.log`, no file reading. Pure, testable, portable.

The **imperative shell** is thin. It reads from the outside world, calls the core, writes results back. Impure but simple. Almost no logic of its own.

In the Wordle capstone (Module 09), this looks like:

- **`game/`** — the functional core. Given a guess string, returns a result. Knows nothing about terminals. Module 07's types (illegal states unrepresentable) enforce the rules.
- **`ui/`** — the imperative shell. Reads stdin, calls `game.evaluateGuess`, renders output. Knows nothing about scoring.
- **`index.ts`** — wires them together. Three lines of real code.

Change the scoring algorithm, `ui/` doesn't change. Swap the terminal for a web frontend, `game/` doesn't change. That's information hiding applied end-to-end.

## When to create a module

A module (file or directory) represents a **domain concept**, not an architectural layer.

**Bad module names** — these aren't concepts, they're junk drawers: `utils.ts`, `helpers.ts`, `models.ts`, `common.ts`. The types in `models` are actually about accounts, orders, games — put them there.

**Good module signs:**
- You can explain it in one sentence **without "and"** (Module 04's naming lesson — if the name is hard, the design is muddled)
- The name is a noun: `account`, `game`, `auth`, `inventory`
- More is unexported than exported
- Callers don't need to read the source to use it — the exported names tell the story

Create a module when you have a design decision to hide. If there's no decision to hide, there's no boundary to draw.

## Exercises

1. **[Information hiding](exercise-01-information-hiding/)** — seal a module by reducing its exported surface to only what callers need
2. **[Deep vs. shallow](exercise-02-deep-vs-shallow/)** — compare two implementations and measure their interface complexity
3. **[Boundary drawing](exercise-03-boundary-drawing/)** — split a program into modules at the natural seams

## Resources

- [MDN — JavaScript modules](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Modules) — ES modules reference
- [TypeScript Handbook — Modules](https://www.typescriptlang.org/docs/handbook/2/modules.html) — TypeScript module system
- [Bun — Module resolution](https://bun.sh/docs/runtime/modules) — how Bun resolves imports
- [Parnas — "On the Criteria to Be Used in Decomposing Systems into Modules" (1972)](https://dl.acm.org/doi/10.1145/361598.361623) — the origin of information hiding
- Ousterhout, John. *A Philosophy of Software Design* — chapters 4-8 on deep modules and pulling complexity downward
- [MIT 6.033 — Computer System Engineering](https://ocw.mit.edu/courses/6-033-computer-system-engineering-spring-2018/) — modularity and abstraction at the systems level

*Next: [Module 09 · Capstone: Terminal Wordle](../module-09-capstone-wordle/) — build a complete game. Every boundary decision you make will test what you learned here.*

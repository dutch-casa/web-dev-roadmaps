# 03 · Values & Declarations

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
  style C fill:#f90,color:#000
```

*In Module 02, you learned to track the history of your code. Now you need to understand what code actually does — and it starts with data.*

Strip away the syntax, the frameworks, the tooling. A program takes data in, transforms it, and produces data out. Keyboard input, file contents, network request — doesn't matter where it comes from. Text on screen, database row, signal to another process — doesn't matter where it goes. The core activity is always: take values, transform them, produce new values.

Once you see programs as data pipelines instead of instruction sequences, you stop following tutorials and start solving problems.

## Values and types

A value is a piece of data. `42`. `"hello"`. `true`. Values don't change — `42` is always `42`.

Values have types. The type tells you what operations make sense. You can add two numbers. You can concatenate two strings. You can't add a number to a string in any meaningful way — and JavaScript will silently coerce it instead of telling you, which is why TypeScript exists (Module 07).

| Type | What it is | Example |
|------|-----------|---------|
| `number` | All numbers (integer and decimal) | `42`, `3.14`, `-7` |
| `string` | Text | `"hello"`, `'world'`, `` `template` `` |
| `boolean` | True or false | `true`, `false` |
| `null` | Intentional absence | `null` |
| `undefined` | Uninitialized or missing | `undefined` |
| `bigint` | Arbitrary-precision integers | `9007199254740993n` |
| `symbol` | Unique identifiers | `Symbol("id")` |

JavaScript has one `number` type for both integers and floats. No `int` vs `float64` distinction. This keeps things simple until you hit floating-point precision — `0.1 + 0.2 !== 0.3`. That's IEEE 754, not a JavaScript bug.

## Binding names to values

A declaration binds a name to a value:

```js
const name = "Auburn";      // immutable binding — this is the default
let count = 42;             // mutable binding — use when you need to reassign
```

`const` means the binding cannot be reassigned. `let` means it can. There's also `var` — never use it. `var` has broken scoping rules left over from 1995. Pretend it doesn't exist.

```js
const x = 10;
// x = 20;    // TypeError — good, the binding is immutable

let score = 85;
score = 92;   // fine — score is a mutable binding
```

## Const by default

When you read a `const`, you know its value without tracing any code. It's the same at line 1 as at line 500. `let` creates a burden — you have to track every reassignment to know the current value.

**Default to `const`.** Reach for `let` only when the value genuinely needs to change — loop counters, accumulators, state that evolves. If you find yourself writing `let` more than `const`, something is off.

One subtlety: `const` prevents *reassignment*, not *mutation*. A `const` object can still have its properties changed:

```js
const user = { name: "Dutch", age: 30 };
user.age = 31;       // allowed — mutating the object, not the binding
// user = {};        // TypeError — can't reassign the binding
```

This is a footgun. Module 07 covers how TypeScript's `as const` and `Readonly` types close this gap.

## Expressions compose, statements don't

An **expression** produces a value: `2 + 3`, `"hello".length`, `age >= 18`.

A **statement** performs an action: `const x = 10`, `console.log(x)`, `if (x > 0) { ... }`.

Expressions nest. Statements don't. JavaScript leans more expression-oriented than many languages — arrow functions, ternaries, and template literals are all expressions:

```js
// Statement-heavy
const input = "  Hello World  ";
const trimmed = input.trim();
const lowered = trimmed.toLowerCase();
const length = lowered.length;

// Expression-oriented: one pipeline
const length = "  Hello World  ".trim().toLowerCase().length;
```

Template literals turn string building from statement gymnastics into a single expression:

```js
// Concatenation (statement-flavored)
const greeting = "Hello, " + name + "! You are " + age + " years old.";

// Template literal (expression-flavored)
const greeting = `Hello, ${name}! You are ${age} years old.`;
```

Neither is always better. Short pipeline? Expression form. Intermediate values need names for clarity? Statements. You'll develop the judgment.

## Destructuring — pull data apart

Destructuring lets you extract values from objects and arrays in a single declaration. This is one of JavaScript's most useful features.

```js
// Object destructuring
const user = { name: "Dutch", age: 30, city: "Auburn" };
const { name, age } = user;  // name = "Dutch", age = 30

// Array destructuring
const rgb = [255, 128, 0];
const [red, green, blue] = rgb;  // red = 255, green = 128, blue = 0

// With rest syntax
const [first, ...rest] = [1, 2, 3, 4];  // first = 1, rest = [2, 3, 4]
```

Destructuring is not just convenient — it's a declaration of intent. `const { name, age } = user` says "I need these two things and nothing else." It narrows the scope.

## Data has shape

The idea that connects values to everything else: **choosing the right shape for your data is the most important decision in programming.**

```js
const scores = [98, 85, 92, 77];                    // ordered collection, same type

const capitals = {                                    // key-value pairs, fast lookup
  Alabama: "Montgomery",
  Georgia: "Atlanta",
};

const student = {                                     // named fields, different types
  name: "Dutch",
  major: "CS",
  gpa: 3.8,
};
```

When you model a domain, the first question isn't "what should the code do?" It's "what does the data look like?" Get the shape right and the code writes itself. Get it wrong and you fight the representation at every step. Module 07 goes deep on this. For now, sketch the shapes before you write any logic.

## Exercises

1. **[Const by default](exercise-01-const-by-default/)** — convert `let` declarations to `const` where possible and explain why each remaining `let` must stay mutable
2. **[Expression vs. statement](exercise-02-expression-vs-statement/)** — rewrite statement-heavy code using expressions, simplify boolean logic buried under temps
3. **[Data shape sketcher](exercise-03-data-shape-sketcher/)** — model real-world domains as JavaScript objects before writing any logic

## Resources

- [MDN — JavaScript data types and data structures](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Data_structures) — the reference for all JS types
- [MDN — Destructuring assignment](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment) — full syntax reference
- [Will Sentance — JavaScript: The Hard Parts v3 (FEM)](https://frontendmasters.com/courses/javascript-hard-parts-v2/) — what actually happens when JS executes

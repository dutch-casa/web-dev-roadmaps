# JavaScript + TypeScript Quick Reference

Keep this open while working through exercises. Modules 01-06 use plain JS. Modules 07-09 use TypeScript.

---

## Variables and Constants

```js
const name = "Jordan";       // default — use this
let count = 0;               // only when you need to reassign
// never use var
```

`const` means the *binding* can't be reassigned. The value can still be mutated:

```js
const arr = [1, 2, 3];
arr.push(4);                 // fine — same binding, mutated value
// arr = [5, 6];             // error — can't reassign
```

**Destructuring**

```js
const { name, age } = user;
const [first, ...rest] = [1, 2, 3];    // first = 1, rest = [2, 3]
const { name: userName } = user;        // rename during destructuring
```

---

## Types (JavaScript)

### Primitives

| Type        | Example              |
|-------------|----------------------|
| `string`    | `"hello"`, `` `hi` ``|
| `number`    | `42`, `3.14`, `NaN`  |
| `boolean`   | `true`, `false`      |
| `null`      | `null`               |
| `undefined` | `undefined`          |
| `bigint`    | `9007199254740993n`  |
| `symbol`    | `Symbol("id")`       |

```js
typeof "hello"    // "string"
typeof 42         // "number"
typeof null       // "object"  <-- historical bug, just memorize it
```

### Truthy and Falsy

All values are truthy **except** these six:

| Falsy value   | Notes                          |
|---------------|--------------------------------|
| `false`       |                                |
| `0`           | also `-0`                      |
| `""`          | empty string                   |
| `null`        | intentional absence            |
| `undefined`   | uninitialized or missing       |
| `NaN`         | failed number parse            |

> **Gotcha**: `[]` and `{}` are truthy. An empty array is *not* falsy.

### Equality

Always use `===` and `!==`. Never `==` or `!=`.

```js
0 == ""       // true  (type coercion — nonsense)
0 === ""      // false (no coercion — correct)
```

### null vs undefined

- `undefined` — JS gives you this. Uninitialized variable, missing property, missing argument.
- `null` — you give this. Explicit "nothing here."

---

## Types (TypeScript) **TypeScript**

> Applies from Module 07 onward.

### Annotations and inference

```ts
const name: string = "Jordan";         // explicit — usually unnecessary
const age = 20;                        // inferred as number
function greet(name: string): string { // annotate at boundaries
  return `Hello, ${name}`;
}
```

Annotate: function parameters, return types, and module boundaries. Let inference handle local variables.

### type vs interface

```ts
// type — unions, intersections, primitives, tuples
type Status = "active" | "inactive";
type Point = { x: number; y: number };

// interface — object shapes that might be extended
interface User {
  name: string;
  age: number;
}
interface Admin extends User {
  role: "admin";
}
```

### Union types

```ts
type Status = "loading" | "success" | "error";
type StringOrNumber = string | number;
```

### Discriminated unions

```ts
type Result<T> =
  | { tag: "ok"; value: T }
  | { tag: "error"; message: string };

function handle(r: Result<number>) {
  if (r.tag === "ok") {
    console.log(r.value);    // TS knows value exists here
  } else {
    console.log(r.message);  // TS knows message exists here
  }
}
```

### as const and satisfies

```ts
const DIRECTIONS = ["north", "south", "east", "west"] as const;
// type: readonly ["north", "south", "east", "west"]

const config = {
  port: 3000,
  host: "localhost",
} satisfies Record<string, string | number>;
// type-checked but not widened
```

### Generics

```ts
function identity<T>(value: T): T {
  return value;
}
function first<T>(arr: T[]): T | undefined {
  return arr[0];
}
```

### Utility types

| Utility            | What it does                              |
|--------------------|-------------------------------------------|
| `Partial<T>`       | All properties optional                   |
| `Required<T>`      | All properties required                   |
| `Pick<T, K>`       | Keep only keys `K`                        |
| `Omit<T, K>`       | Remove keys `K`                           |
| `Record<K, V>`     | Object with keys `K` and values `V`       |
| `Readonly<T>`      | All properties `readonly`                 |

### Type narrowing

```ts
function process(value: unknown) {
  if (typeof value === "string") {
    value.toUpperCase();              // narrowed to string
  }
  if (value !== null && typeof value === "object" && "tag" in value) {
    // narrowed to object with 'tag' property
  }
}
```

### unknown vs any

- `unknown` — safe. Must narrow before use. Use at boundaries.
- `any` — disables type checking. Never use it.

### No enum

```ts
// bad — avoid enum
// enum Direction { North, South }

// good — string literal union
type Direction = "north" | "south" | "east" | "west";

// good — const object when you need runtime values
const Direction = {
  North: "north",
  South: "south",
} as const;
type Direction = (typeof Direction)[keyof typeof Direction];
```

---

## Scope

| Scope    | Mechanism                                  |
|----------|--------------------------------------------|
| Block    | `const`/`let` live inside `{ }`            |
| Function | Parameters + local vars                    |
| Module   | Each file is its own scope (ES modules)    |
| Closure  | A function captures its enclosing scope    |

```js
function makeCounter() {
  let count = 0;                   // captured by closure
  return () => ++count;
}
const counter = makeCounter();
counter(); // 1
counter(); // 2
```

> **Gotcha**: `var` and `function` declarations are hoisted to the top of their scope. One more reason to use `const`/`let` only.

---

## Strings

Template literals — always prefer over concatenation:

```js
const greeting = `Hello, ${name}!`;
const multiline = `line one
line two`;
```

| Method              | Returns                            |
|---------------------|------------------------------------|
| `.length`           | character count (property)         |
| `.trim()`           | removes leading/trailing whitespace|
| `.toLowerCase()`    | lowercase copy                     |
| `.toUpperCase()`    | uppercase copy                     |
| `.includes(str)`    | `boolean`                          |
| `.startsWith(str)`  | `boolean`                          |
| `.endsWith(str)`    | `boolean`                          |
| `.split(sep)`       | `string[]`                         |
| `.replace(a, b)`    | replaces first match               |
| `.replaceAll(a, b)` | replaces all matches               |
| `.slice(start, end)`| substring                          |
| `.indexOf(str)`     | index or `-1`                      |
| `.padStart(n, ch)`  | pads from left                     |
| `.padEnd(n, ch)`    | pads from right                    |

`String.raw` — raw template string (no escape processing):
```js
String.raw`\n is not a newline here` // "\\n is not a newline here"
```

> Strings are immutable. Every method returns a new string.

---

## Numbers

All numbers are `number` (64-bit IEEE 754 float). No separate integer type.

```js
parseInt("42", 10)          // 42  — always pass radix
parseFloat("3.14")          // 3.14
(3.14159).toFixed(2)        // "3.14" — returns string
```

| Function                | Purpose                       |
|-------------------------|-------------------------------|
| `Number.isNaN(x)`       | true if `NaN` (not `isNaN()`) |
| `Number.isFinite(x)`    | excludes `Infinity` and `NaN` |
| `Number.isInteger(x)`   | whole number check            |
| `Math.floor(x)`         | round down                    |
| `Math.ceil(x)`          | round up                      |
| `Math.round(x)`         | nearest integer               |
| `Math.max(a, b, ...)`   | largest                       |
| `Math.min(a, b, ...)`   | smallest                      |
| `Math.random()`         | `[0, 1)` — not crypto-safe   |

> **Gotcha**: `0.1 + 0.2 !== 0.3`. Floating-point math. Compare with a tolerance or use integers (cents, not dollars).

---

## Arrays

### Creating

```js
const nums = [1, 2, 3];
const five = Array.from({ length: 5 }, (_, i) => i);  // [0,1,2,3,4]
const zeros = new Array(5).fill(0);                    // [0,0,0,0,0]
```

### Mutating methods

| Method        | Effect                 |
|---------------|------------------------|
| `.push(x)`    | add to end             |
| `.pop()`      | remove from end        |
| `.unshift(x)` | add to start           |
| `.shift()`    | remove from start      |
| `.sort(fn)`   | sort in place          |

### Non-mutating transforms

| Method          | Returns                                |
|-----------------|----------------------------------------|
| `.map(fn)`      | new array, same length                 |
| `.filter(fn)`   | new array, matching items              |
| `.flatMap(fn)`  | map + flatten one level                |
| `.toSorted(fn)` | sorted copy (ES2023)                   |
| `.reduce(fn, init)` | single accumulated value          |

### Search

| Method           | Returns                            |
|------------------|------------------------------------|
| `.find(fn)`      | first match or `undefined`         |
| `.findIndex(fn)` | first match index or `-1`          |
| `.includes(x)`   | `boolean`                          |
| `.indexOf(x)`    | index or `-1`                      |
| `.some(fn)`      | true if any match                  |
| `.every(fn)`     | true if all match                  |

### Iteration

```js
for (const item of arr) { }              // preferred
for (const [i, item] of arr.entries()) { } // when you need index
arr.forEach(item => { });                 // no break, no await — avoid
```

### Spread and destructuring

```js
const merged = [...arr1, ...arr2];
const copy = [...arr];                    // shallow copy
const [first, second, ...rest] = arr;
```

---

## Objects

```js
const user = { name: "Jordan", age: 20 };

user.name               // "Jordan"
user["name"]            // "Jordan" — use when key is dynamic

const updated = { ...user, name: "Sam" };  // shallow copy + override
const { name, age } = user;                // destructuring
const key = "name";
const obj = { [key]: "Jordan" };           // computed key
```

| Function               | Returns                          |
|------------------------|----------------------------------|
| `Object.keys(obj)`    | `string[]` of keys               |
| `Object.values(obj)`  | array of values                  |
| `Object.entries(obj)` | `[key, value][]` pairs           |
| `Object.freeze(obj)`  | shallow immutability             |

**Optional chaining and nullish coalescing**

```js
const avatar = user?.profile?.avatar;       // undefined if any link is null/undefined
const port = config.port ?? 3000;           // default only for null/undefined, not 0 or ""
```

> **Gotcha**: `||` treats `0`, `""`, and `false` as falsy. Use `??` when those are valid values.

---

## Maps and Sets

### Map

```js
const m = new Map();
m.set("key", "value");
m.get("key");           // "value"
m.has("key");           // true
m.delete("key");
m.size;                 // 0

for (const [k, v] of m) { }
```

Use **Map** when keys are dynamic or non-string. Use plain objects for fixed-shape data.

### Set

```js
const s = new Set([1, 2, 3, 3]);    // {1, 2, 3}
s.add(4);
s.has(2);               // true
s.delete(2);
s.size;                 // 3

const unique = [...new Set(array)]; // deduplicate an array
for (const item of s) { }
```

---

## Functions

### Arrow functions (preferred)

```js
const add = (a, b) => a + b;               // implicit return
const greet = (name) => {                   // block body
  const msg = `Hello, ${name}`;
  return msg;
};
```

### Regular functions

```js
function add(a, b) {
  return a + b;
}
```

### Parameters

```js
function greet(name = "world") { }          // default
function sum(...nums) { }                   // rest — nums is an array
```

### First-class functions

```js
const ops = [Math.floor, Math.ceil];
const apply = (fn, x) => fn(x);            // pass function as argument
```

### this

Arrow functions inherit `this` from their enclosing scope. Regular functions get their own `this` based on how they're called. Prefer arrow functions to avoid surprises.

### TypeScript function signatures **TypeScript**

```ts
function add(a: number, b: number): number {
  return a + b;
}
const identity = <T>(value: T): T => value;
```

---

## Classes **TypeScript**

```ts
class Counter {
  #count: number;                  // # = real private (runtime-enforced)
  readonly name: string;           // immutable after construction

  constructor(name: string, start = 0) {
    this.name = name;
    this.#count = start;
  }

  get value(): number {
    return this.#count;
  }

  increment(): void {
    this.#count++;
  }
}
```

```ts
interface Printable {
  print(): void;
}
class Report implements Printable {
  print() { console.log("report"); }
}
```

Use classes for entities with identity and lifecycle. Not everything needs to be a class.

---

## Error Handling

### try/catch

```js
try {
  const data = JSON.parse(input);
} catch (err) {
  console.error("Invalid JSON:", err.message);
} finally {
  // always runs
}
```

### Custom errors

```js
class NotFoundError extends Error {
  constructor(id) {
    super(`Not found: ${id}`);
    this.name = "NotFoundError";
  }
}
```

### Result pattern (preferred in TypeScript) **TypeScript**

```ts
type Result<T> =
  | { tag: "ok"; value: T }
  | { tag: "error"; message: string };

function parse(input: string): Result<number> {
  const n = Number(input);
  if (Number.isNaN(n)) return { tag: "error", message: "not a number" };
  return { tag: "ok", value: n };
}
```

Prefer returning error values over throwing. Reserve `throw` for truly exceptional cases.

---

## Modules (ES Modules)

```js
// math.js — named exports (preferred)
export const add = (a, b) => a + b;
export const sub = (a, b) => a - b;

// app.js — named imports
import { add, sub } from "./math";

// re-export
export { add } from "./math";

// side-effect import (runs the file, imports nothing)
import "./setup";
```

> Avoid `export default`. Named exports give better autocomplete and refactoring.

In Bun, no file extension needed in import paths.

---

## Promises and Async

```js
async function fetchUser(id) {
  const res = await fetch(`/api/users/${id}`);
  if (!res.ok) throw new Error(`HTTP ${res.status}`);
  return res.json();
}

// parallel
const [user, posts] = await Promise.all([
  fetchUser(1),
  fetchPosts(1),
]);

// error handling
try {
  const data = await fetchUser(1);
} catch (err) {
  console.error(err.message);
}
```

---

## Iteration Patterns

| Pattern                            | Use when                              |
|------------------------------------|---------------------------------------|
| `for (const x of arr)`            | default — arrays, strings, maps, sets |
| `for (const [i, x] of arr.entries())` | need the index                   |
| `for (let i = 0; i < n; i++)`     | numeric range or index math           |
| `.forEach(fn)`                     | avoid — no `break`, no `await`        |
| `for (const k in obj)`            | almost never — includes prototype keys|

---

## Common Patterns

```js
// optional chaining
const name = user?.profile?.name;

// nullish coalescing
const port = config.port ?? 3000;

// logical assignment
options.timeout ??= 5000;            // assign only if null/undefined

// immutable object update
const next = { ...state, count: state.count + 1 };

// immutable array append
const updated = [...items, newItem];

// deep copy (no lodash needed)
const clone = structuredClone(original);

// tagged template literals (used by libraries like sql`...`)
const result = sql`SELECT * FROM ${table}`;
```

---

## Bun-Specific

| Command             | Purpose                                   |
|---------------------|-------------------------------------------|
| `bun run file.ts`   | run TypeScript directly, no build step    |
| `bun run file.js`   | run JavaScript                            |
| `bun init`           | initialize a project                      |
| `bun add package`    | install a dependency                      |
| `bun test`           | built-in test runner                      |

### Bun APIs

```js
// synchronous terminal input (Bun only, not in browsers)
const name = prompt("What is your name?");

// read a file
const file = Bun.file("data.txt");
const text = await file.text();

// HTTP server (Backend track)
Bun.serve({
  fetch(req) {
    return new Response("Hello");
  },
});
```

# 06 · Control Flow

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
  style F fill:#f90,color:#000
```

*In Module 05, you learned to decompose problems into functions. Now you'll learn what happens inside those functions — how to shape control flow so the reader follows a straight line, not a maze.*

What does this function do?

```js
const processOrder = (order) => {
  if (order.items !== null) {
    if (order.items.length > 0) {
      if (order.customer.isActive) {
        if (order.total() > 0) {
          return buildReceipt(order);
        } else {
          throw new Error("empty total");
        }
      } else {
        throw new Error("inactive customer");
      }
    } else {
      throw new Error("no items");
    }
  } else {
    throw new Error("null items");
  }
};
```

Four levels deep. You have to hold the entire nesting stack in your head to match each error to its condition. Now read this:

```js
const processOrder = (order) => {
  if (order.items == null) throw new Error("null items");
  if (order.items.length === 0) throw new Error("no items");
  if (!order.customer.isActive) throw new Error("inactive customer");
  if (order.total() <= 0) throw new Error("empty total");

  return buildReceipt(order);
};
```

Same logic. Same number of paths. But each check is self-contained: condition, consequence, done. The happy path flows straight down the page.

## Guard clauses

A guard clause checks for a condition that prevents the function from doing its work, then returns (or throws) immediately. After the guards, the remaining code runs at the shallowest indentation.

No else after return. The successful flow continues unindented.

**Helps when** a function has multiple preconditions — input validation, null checks. **Hurts when** the "error" and "success" paths are symmetric and equally complex.

## Optional chaining and nullish coalescing

JavaScript has two operators that eliminate entire categories of null-checking branches:

```js
// Without optional chaining
let city;
if (user && user.address && user.address.city) {
  city = user.address.city;
} else {
  city = "Unknown";
}

// With optional chaining + nullish coalescing
const city = user?.address?.city ?? "Unknown";
```

`?.` short-circuits to `undefined` if anything in the chain is `null` or `undefined`. `??` provides a fallback for `null` or `undefined` only (not `0` or `""`). Use `??` over `||` for defaults — `||` treats `0`, `""`, and `false` as falsy, which is almost never what you want.

These are guard clauses at the expression level. Same principle: handle the bad case, let the good case flow.

## Named conditions

When a conditional has multiple clauses, extract it into a named boolean. The name documents intent — Module 04's naming lesson applied directly.

```js
// Before: reader parses boolean logic inline
if (user.age >= 18 && user.emailVerified && !user.suspended &&
    user.subscriptionEnd > Date.now()) {
  grantAccess(user);
}

// After: the name tells you what the check means
const canAccess = user.age >= 18
  && user.emailVerified
  && !user.suspended
  && user.subscriptionEnd > Date.now();

if (canAccess) {
  grantAccess(user);
}
```

One extra line saves every future reader from parsing four sub-expressions. When the condition is wrong, the name tells you what was *intended*.

**Helps when** the expression has more than two clauses. **Hurts when** the condition is trivial — `const isPositive = x > 0` before `if (isPositive)` adds noise.

## Data tables replacing branches

When a switch maps inputs to outputs and every case has the same shape, an object is clearer. You saw this in Module 05's data table technique.

```js
const httpStatusTexts = {
  200: "OK",
  201: "Created",
  400: "Bad Request",
  404: "Not Found",
  500: "Internal Server Error",
};

const httpStatusText = (code) => httpStatusTexts[code] ?? "Unknown";
```

Data separated from mechanism. Adding a status is one line. **Does not work when** the branches have different shapes — if case A computes, case B calls an API, and case C validates, an object can't represent that.

## Switch

JavaScript's `switch` falls through by default (unlike Go). Always use `break` or `return`:

```js
const describeDay = (day) => {
  switch (day) {
    case "Monday":
    case "Tuesday":
    case "Wednesday":
    case "Thursday":
    case "Friday":
      return "weekday";
    case "Saturday":
    case "Sunday":
      return "weekend";
    default:
      return `unknown day: ${day}`;
  }
};
```

Always include a `default` case. In Module 07, you'll learn how TypeScript's discriminated unions and exhaustive checking can guarantee at compile time that every case is handled.

## The tension

There's a real tension here, and it doesn't have a clean resolution. When you split a function into three helpers that are only called from one place, you haven't reduced complexity — you've spread it across more locations. A function you have to cross-reference is sometimes worse than 30 lines you can read in place.

Guard clauses can create long preambles that push the interesting code off-screen. Named conditions can over-abstract simple checks. Data tables hide different-shaped logic behind a uniform facade.

No formula resolves this. If the extracted piece has a clear name, a clear contract, and could be tested independently, extract it. If it's just "the middle part of the other function," leave it inline.

## Exercises

1. **[Flatten the pyramid](exercise-01-flatten-the-pyramid/)** — refactor deeply nested code using guard clauses
2. **[Named conditions](exercise-02-named-conditions/)** — extract complex boolean expressions into named variables
3. **[Linear flow](exercise-03-linear-flow/)** — rewrite tangled control flow into a clean top-to-bottom structure

## Resources

- [MDN — Control flow and error handling](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Control_flow_and_error_handling) — JavaScript control structures reference
- [MDN — Optional chaining](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Optional_chaining) — `?.` operator
- [MDN — Nullish coalescing](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Nullish_coalescing) — `??` operator
- McConnell, Steve. *Code Complete*, 2nd ed. — Chapter 19: General Control Issues
- Ousterhout, John. *A Philosophy of Software Design* — deep vs. shallow modules, complexity as the central problem
- Muratori, Casey. ["Clean Code, Horrible Performance"](https://www.computerenhance.com/p/clean-code-horrible-performance) — the cost of scattering logic across abstractions

*Next: [Module 07 · Types & Data Modeling](../module-07-types-and-data-modeling/) — constrain the value space so invalid states can't exist. This is where TypeScript enters.*

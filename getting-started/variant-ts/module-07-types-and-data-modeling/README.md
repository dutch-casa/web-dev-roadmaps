# 07 · Types & Data Modeling

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
  style G fill:#f90,color:#000
```

*In Module 06, you learned to write code that flows in a straight line — guard clauses, named conditions, flat paths. Now you'll learn the tool that makes most of those checks unnecessary.*

**From this module forward, everything is TypeScript.** Plain JavaScript served you well — you learned values, functions, control flow without a type system in the way. Now you'll see what you've been missing. TypeScript adds compile-time types to JavaScript: same language, same runtime, but the compiler catches bugs before your code runs. You write `.ts` files. Bun runs them directly — no build step.

This is the most important module in the track. Everything else teaches you to write code that humans can read. Types teach you to write code where entire categories of bugs *cannot exist*.

## Why TypeScript

JavaScript is dynamically typed. A variable can be a string, then a number, then `undefined`, and nobody complains until production:

```js
const greet = (name) => `Hello, ${name.toUpperCase()}!`;
greet(42);  // runtime crash: name.toUpperCase is not a function
```

TypeScript catches this before the code runs:

```ts
const greet = (name: string): string => `Hello, ${name.toUpperCase()}!`;
greet(42);  // compile error: Argument of type 'number' is not assignable to parameter of type 'string'
```

The type annotation is the contract. The compiler enforces it. The runtime never sees a violation because the violation never ships.

## Make illegal states unrepresentable

Look at this object:

```ts
type Order = {
  id: string;
  status: string;       // "draft", "confirmed", "shipped", "delivered"
  address: string;      // empty in draft
  trackingNum: string;  // empty until shipped
  shippedAt: Date;      // undefined until shipped
  deliveredAt: Date;    // undefined until delivered
};
```

How many invalid states can you construct? A "draft" with a tracking number. A "delivered" with no ship date. A status of `"banana"`. Dozens. Every function that touches this type must check for impossible combinations. You're writing validators, defensive `if` statements, and tests for states that should not exist. The type created work instead of preventing it.

Now look at this:

```ts
type DraftOrder = {
  tag: "draft";
  id: string;
};

type ConfirmedOrder = {
  tag: "confirmed";
  id: string;
  address: string;
};

type ShippedOrder = {
  tag: "shipped";
  id: string;
  address: string;
  trackingNum: string;
  shippedAt: Date;
};

type DeliveredOrder = {
  tag: "delivered";
  id: string;
  address: string;
  trackingNum: string;
  shippedAt: Date;
  deliveredAt: Date;
};

type Order = DraftOrder | ConfirmedOrder | ShippedOrder | DeliveredOrder;
```

A `DraftOrder` cannot have a tracking number — the field doesn't exist. A `DeliveredOrder` always has a ship date — the compiler requires it at construction. Invalid states: zero. The guard clauses from Module 06? Gone. Not refactored, not hidden. Eliminated by the type system.

The state transitions become function signatures:

```ts
const confirm = (d: DraftOrder, address: string): ConfirmedOrder => ({ ...d, tag: "confirmed", address });
const ship = (c: ConfirmedOrder, trackingNum: string): ShippedOrder => ({ ...c, tag: "shipped", trackingNum, shippedAt: new Date() });
const deliver = (s: ShippedOrder): DeliveredOrder => ({ ...s, tag: "delivered", deliveredAt: new Date() });
```

You cannot ship a draft. You cannot deliver a confirmation. The compiler enforces the business rules.

## Discriminated unions

The `tag` field above is a **discriminant** — a literal type that tells TypeScript which variant you have. When you `switch` on it, TypeScript narrows the type automatically:

```ts
const describeOrder = (order: Order): string => {
  switch (order.tag) {
    case "draft":     return `Draft #${order.id}`;
    case "confirmed": return `Confirmed #${order.id} to ${order.address}`;
    case "shipped":   return `Shipped #${order.id} via ${order.trackingNum}`;
    case "delivered":  return `Delivered #${order.id} at ${order.deliveredAt.toISOString()}`;
  }
};
```

Inside each case, TypeScript knows the exact shape. `order.trackingNum` is valid in `"shipped"` and an error in `"draft"`. No casts, no assertions, no runtime checks.

For exhaustive checking, use a `never` helper: `const exhaustive = (value: never): never => { throw new Error(`Unhandled: ${value}`); }`. Add it as the `default` case and any unhandled variant becomes a compile error. Module 06's "always write a default case" meets compile-time enforcement.

## `type` vs `interface`

Both define object shapes. Use `type` for unions, intersections, and mapped types. Use `interface` when you need declaration merging (rare) or want to name an object shape that may be extended.

```ts
// type: unions, computed types, primitives
type Result = { tag: "ok"; value: string } | { tag: "err"; message: string };
type ID = string & { readonly __brand: "ID" };

// interface: object shapes, extensible contracts
interface Logger {
  log(message: string): void;
  error(message: string): void;
}
```

For this track, default to `type`. It's more general. Reach for `interface` when you're defining a contract that multiple implementations will satisfy.

## `as const` and `satisfies`

`as const` makes a value deeply readonly and narrows its type to the literal:

```ts
const DIRECTIONS = ["north", "east", "south", "west"] as const;
type Direction = (typeof DIRECTIONS)[number];  // "north" | "east" | "south" | "west"
```

No runtime enum. The array is the source of truth, and the type is derived from it.

`satisfies` checks that a value matches a type *without widening it*:

```ts
type Config = { host: string; port: number };
const config = { host: "localhost", port: 3000 } satisfies Config;
// config.host is type "localhost", not string — the literal is preserved
```

Use `satisfies` when you want type checking *and* inference. Use `:` annotation when you want to widen to the declared type.

## Strict null checks

With `strictNullChecks` enabled (always enable it), `null` and `undefined` are not assignable to other types. When a value might be absent, say so: `const findUser = (id: string): User | undefined`. The caller must narrow before accessing properties — TypeScript enforces this.

Three kinds of "not the happy path" — don't conflate them:

| Situation | Meaning | TypeScript idiom |
|-----------|---------|-----------------|
| **Absence** | A value legitimately doesn't exist | `T \| undefined`, optional properties |
| **Failure** | An operation went wrong | Result type or thrown Error |
| **Invalidity** | Input violates a domain rule | Reject at construction (branded types) |

A user not found is absence. The database being unreachable is failure. A negative user ID is invalidity. Each demands a different response.

## FP vs OOP — when to use each

This is not a religious debate. Different problems have different shapes.

**Use values and pure functions when** you're computing a result from inputs. `const finalPrice = applyDiscount(applyTax(basePrice, 0.08), 15)` — no state, same inputs same output. Functions compose (Module 03's expressions lesson again).

**Use classes with methods when** the data has identity — it represents a thing that exists over time. A game board, a connection pool, a session. Construction invariants matter. The state changes *are* the point. Mark internals `private` and expose only the operations that maintain invariants.

The judgment: is this a **value** you're computing, or a **thing** you're managing? Values get functions. Things get classes. Most real programs use both. The mistake is reaching for one reflexively.

## Domain types over primitives

A common temptation: use `string` for everything. Email addresses, user IDs, statuses. This is primitive obsession — a generic type where a domain type would prevent misuse.

```ts
// Dangerous: nothing stops you from swapping arguments
const sendConfirmation = (email: string, username: string): void => { ... };

// Better: branded types catch the mixup
type Email = string & { readonly __brand: "Email" };
type Username = string & { readonly __brand: "Username" };
const sendConfirmation = (email: Email, username: Username): void => { ... };
```

Zero runtime cost. Zero readability cost. `sendConfirmation(username, email)` with arguments swapped is now a compile error instead of a production bug.

## Exercises

1. **[Illegal states](exercise-01-illegal-states/)** — model a domain where bad states are impossible to construct
2. **[Errors as values](exercise-02-errors-as-values/)** — handle three kinds of failure using discriminated unions
3. **[Type vs. interface](exercise-03-type-vs-interface/)** — implement a problem both ways and discuss tradeoffs
4. **[FP vs. OOP decision](exercise-04-fp-vs-oop/)** — two problems, two paradigms, explain your choices

## Resources

- [TypeScript Handbook](https://www.typescriptlang.org/docs/handbook/) — the official guide to TypeScript
- [Anjana Vakil — TypeScript First Steps](https://anjana.dev/typescript-first-steps/) — approachable introduction to TypeScript
- [Mike North — TypeScript Fundamentals v4](https://www.typescript-training.com/course/fundamentals-v4) — deep-dive on the type system
- [Mike North — Domain Modeling with TypeScript](https://www.typescript-training.com/course/domain-modeling-with-ts) — making illegal states unrepresentable
- [Rich Hickey — "Simple Made Easy" (Strange Loop 2011)](https://www.infoq.com/presentations/Simple-Made-Easy/) — the distinction between simplicity and familiarity
- Ousterhout, John. *A Philosophy of Software Design* — deep modules, information hiding, defining errors out of existence

*Next: [Module 08 · Modules & Composition](../module-08-modules-and-composition/) — organize types into modules, hide decisions behind boundaries, compose them into programs.*

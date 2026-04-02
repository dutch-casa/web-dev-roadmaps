# 05 · Functions & Decomposition

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
  style E fill:#f90,color:#000
```

*In Module 04, you learned that naming is design — and that `validateAndSaveUser` is a warning sign. Now you'll learn where to make the cut, and when to leave the knife in the drawer.*

A function is a named transformation. Data in, data out. The name describes what happens (Module 04). The body does it.

```js
const celsiusToFahrenheit = (c) => c * 9 / 5 + 32;
```

No printing. No file reads. No mutation. Given the same input, always the same output. Not all functions have these properties. Learning which do and which don't is what this module is about.

## Pure vs. impure

A **side effect** is anything a function does besides returning a value.

```js
// Pure: returns a value, nothing else happens
const totalPrice = (items) => {
  let total = 0;
  for (const item of items) {
    total += item.price * item.quantity;
  }
  return total;
};

// Impure: writes to stdout — that's a side effect
const printReceipt = (items) => {
  for (const item of items) {
    const dollars = (item.price / 100).toFixed(2);
    console.log(`  ${item.name}: $${dollars}`);
  }
};
```

You can't avoid side effects entirely — a program that never prints or writes is useless. But you can isolate them. Pure functions do the thinking. Impure functions do the doing. If you can return a value instead of printing it, do that. The caller can always print later.

## Four refactoring techniques

Each has a cost (more indirection) and a benefit (easier to read, test, change). Don't apply them everywhere. Recognize which fits the mess you're looking at.

### 1. Extract function

```js
// Before: parsing buried in a loop
for (const line of lines) {
  const parts = line.split(",");
  if (parts.length !== 3) continue;
  const name = parts[0].trim();
  const score = parseInt(parts[1].trim(), 10);
  if (isNaN(score)) continue;
  const grade = parts[2].trim();
  records.push({ name, score, grade });
}

// After: the loop reads like a summary
for (const line of lines) {
  const record = parseLine(line);
  if (record) records.push(record);
}
```

**Helps when** the block does a coherent thing you can name clearly. **Hurts when** the block is 2-3 lines, or the extracted function needs 5+ parameters — you drew the boundary wrong.

### 2. Data table

```js
// Before: structurally identical cases
switch (status) {
  case "todo":        return "○";
  case "in-progress": return "◐";
  case "done":        return "●";
  default:            return "?";
}

// After: the mapping is visible at a glance
const statusIcons = {
  todo: "○",
  "in-progress": "◐",
  done: "●",
};
```

**Helps when** every case has the same shape. **Hurts when** cases differ — some compute, some validate, some call APIs. An object can't represent that honestly.

### 3. Separate traversal from decision

```js
// Before: selection and action tangled
const notifyOverdue = (loans) => {
  for (const loan of loans) {
    if (loan.dueDate < Date.now() && !loan.returned) {
      sendReminder(loan.borrower, loan.book);
    }
  }
};

// After: overdueLoans is pure and testable
const overdueLoans = (loans) =>
  loans.filter((loan) => loan.dueDate < Date.now() && !loan.returned);
```

JavaScript's array methods — `.filter()`, `.map()`, `.reduce()` — make this pattern natural. The filter is a pure function. The action happens separately. **Helps when** the filter is complex or reused. **Hurts when** the filter is one trivial line.

### 4. Push I/O up

```js
// Before: decision and printing mixed
const reportWeather = (tempC) => {
  const tempF = tempC * 9 / 5 + 32;
  if (tempF > 100) {
    console.log("Heat warning!");
  }
};

// After: pure function returns a value, caller prints
const weatherStatus = (tempC) => {
  const tempF = tempC * 9 / 5 + 32;
  if (tempF > 100) return "Heat warning!";
  return `${Math.round(tempF)}°F — normal`;
};
```

**Helps when** you want to test logic without capturing stdout. **Hurts when** both computation and I/O are one line each.

## When NOT to refactor

Every technique above adds indirection. The cost is real.

- **The code is short and clear.** A 10-line function with one if/else doesn't need to be three functions.
- **You'd create a shallow wrapper.** If the new function just calls another with the same arguments, it adds a layer with zero value.
- **It's a one-shot script.** Refactoring is an investment in the future. No future, no payoff.
- **You can't name it.** If the name doesn't come easily, the boundary isn't real. That's Module 04's lesson at work.

The goal is not "small functions." The goal is code where each unit is understandable on its own.

## Exercises

1. **[Extract till you drop](exercise-01-extract-till-you-drop/)** — break a 100-line function into named pieces using the techniques above
2. **[Pure vs. impure](exercise-02-pure-vs-impure/)** — identify side effects and refactor impure functions into pure cores with thin wrappers
3. **[File organization](exercise-03-file-organization/)** — reorganize a single file so it reads top to bottom

## Resources

- [MDN — Functions](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Functions) — JavaScript function reference
- [MDN — Array methods](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array) — `.map()`, `.filter()`, `.reduce()`, `.find()`
- [Casey Muratori — Semantic Compression](https://caseymuratori.com/blog_0015) — code compression that isn't just "make functions smaller"
- Ousterhout, John. *A Philosophy of Software Design*, Ch. 4-5 — deep modules and information hiding at function scale

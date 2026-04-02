// Pure vs. impure
//
// Part 1: Label each function below as PURE or IMPURE.
// Write your answer as a comment above each function.
//
// A pure function:
//   - Always returns the same output for the same input
//   - Does not modify anything outside itself
//   - Does not read from anything that could change (time, random, I/O)
//
// Part 2: Refactor the three impure functions at the bottom
// into pure cores with thin impure wrappers.
//
// Run with: bun run index.js

// Label: ___
const double = (x) => x * 2;

// Label: ___
const greetUser = (name) => {
  console.log(`Hello, ${name}!`);
};

// Label: ___
const maxValue = (a, b) => {
  if (a > b) {
    return a;
  }
  return b;
};

// Label: ___
const randomGreeting = (name) => {
  const greetings = ["Hey", "Hello", "Hi", "Howdy"];
  const i = Math.floor(Math.random() * greetings.length);
  return `${greetings[i]}, ${name}!`;
};

// Label: ___
const currentYear = () => new Date().getFullYear();

// Label: ___
const contains = (items, target) => {
  for (const item of items) {
    if (item === target) {
      return true;
    }
  }
  return false;
};

// --- Part 2: Refactor these ---

// This function mixes computation with printing.
// Extract the pure computation into its own function.
const printDiscount = (price, discountPercent) => {
  let discounted = price - (price * discountPercent / 100);
  if (discounted < 0) {
    discounted = 0;
  }
  console.log(`Original: $${price.toFixed(2)} \u2192 Discounted: $${discounted.toFixed(2)} (${discountPercent}% off)`);
};

// This function uses the current time to determine a greeting.
// Extract the pure decision logic so it can be tested with any hour.
const timeBasedGreeting = (name) => {
  const hour = new Date().getHours();
  if (hour < 12) {
    return `Good morning, ${name}!`;
  } else if (hour < 17) {
    return `Good afternoon, ${name}!`;
  }
  return `Good evening, ${name}!`;
};

// This function generates a random password and prints it.
// Split the generation from the printing.
const generateAndPrintPassword = (length) => {
  const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
  let password = "";
  for (let i = 0; i < length; i++) {
    password += chars[Math.floor(Math.random() * chars.length)];
  }
  console.log(`Your new password: ${password}`);
};

console.log(double(21));
greetUser("Auburn");
console.log(maxValue(10, 20));
console.log(randomGreeting("Jordan"));
console.log(currentYear());
console.log(contains(["js", "rust", "python"], "js"));

console.log();
printDiscount(99.99, 20);
console.log(timeBasedGreeting("Sam"));
generateAndPrintPassword(16);

// Part 1 answers:

// PURE — same input always gives same output, no side effects.
const double = (x) => x * 2;

// IMPURE — prints to stdout (side effect).
const greetUser = (name) => {
  console.log(`Hello, ${name}!`);
};

// PURE — deterministic, no side effects.
const maxValue = (a, b) => {
  if (a > b) {
    return a;
  }
  return b;
};

// IMPURE — uses Math.random() (different result each call).
const randomGreeting = (name) => {
  const greetings = ["Hey", "Hello", "Hi", "Howdy"];
  const i = Math.floor(Math.random() * greetings.length);
  return `${greetings[i]}, ${name}!`;
};

// IMPURE — reads the system clock (result changes over time).
const currentYear = () => new Date().getFullYear();

// PURE — deterministic, no side effects.
const contains = (items, target) => items.includes(target);

// --- Part 2: Refactored ---

// Pure core: computes the discounted price.
const applyDiscount = (price, discountPercent) => {
  const discounted = price - (price * discountPercent / 100);
  return Math.max(discounted, 0);
};

// Impure wrapper: calls the pure function and prints the result.
const printDiscount = (price, discountPercent) => {
  const discounted = applyDiscount(price, discountPercent);
  console.log(`Original: $${price.toFixed(2)} \u2192 Discounted: $${discounted.toFixed(2)} (${discountPercent}% off)`);
};

// Pure core: decides the greeting based on an hour value.
// Can be tested with any hour — no dependency on the clock.
const greetingForHour = (name, hour) => {
  if (hour < 12) {
    return `Good morning, ${name}!`;
  } else if (hour < 17) {
    return `Good afternoon, ${name}!`;
  }
  return `Good evening, ${name}!`;
};

// Impure wrapper: reads the clock and passes the hour to the pure function.
const timeBasedGreeting = (name) => {
  return greetingForHour(name, new Date().getHours());
};

// Pure core: generates a password from random values.
// In real code you'd pass a random source; here we keep it simple.
const generatePassword = (length) => {
  const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
  let password = "";
  for (let i = 0; i < length; i++) {
    password += chars[Math.floor(Math.random() * chars.length)];
  }
  return password;
};

// Impure wrapper: generates and prints.
const generateAndPrintPassword = (length) => {
  console.log(`Your new password: ${generatePassword(length)}`);
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

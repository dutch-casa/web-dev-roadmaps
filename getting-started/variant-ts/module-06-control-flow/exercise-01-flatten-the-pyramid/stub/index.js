// Flatten the pyramid
//
// Each function below has deeply nested control flow.
// Refactor each one using guard clauses so the maximum
// nesting depth is 2 (one level inside the function body,
// one more for a loop if needed).
//
// Rules:
//   - Don't change the behavior. Output must be identical.
//   - Use guard clauses (early returns) to handle edge cases at the top.
//   - The "happy path" logic should be at the shallowest indentation level.
//
// Run with: bun run index.js

// Problem 1: Nested validation
const registerUser = (user) => {
  if (user.name !== "") {
    if (user.email !== "") {
      if (user.email.includes("@")) {
        if (user.age >= 13) {
          if (!user.isBanned) {
            return `Welcome, ${user.name}! Registration complete.`;
          } else {
            return "Error: user is banned";
          }
        } else {
          return "Error: must be at least 13 years old";
        }
      } else {
        return "Error: invalid email format";
      }
    } else {
      return "Error: email is required";
    }
  } else {
    return "Error: name is required";
  }
};

// Problem 2: Nested loop with conditions
const calculateCartTotal = (items, couponCode) => {
  if (items !== null && items !== undefined) {
    if (items.length > 0) {
      let total = 0;
      for (const item of items) {
        if (item.quantity > 0) {
          if (item.product.stock >= item.quantity) {
            if (item.product.price > 0) {
              total += item.product.price * item.quantity;
            } else {
              return `Error: invalid price for ${item.product.name}`;
            }
          } else {
            return `Error: not enough stock for ${item.product.name}`;
          }
        } else {
          return `Error: invalid quantity for ${item.product.name}`;
        }
      }
      if (couponCode !== "") {
        if (couponCode === "SAVE10") {
          total = total - Math.floor(total / 10);
        } else if (couponCode === "SAVE20") {
          total = total - Math.floor(total / 5);
        } else {
          return "Error: invalid coupon code";
        }
      }
      return `Total: $${Math.floor(total / 100)}.${String(total % 100).padStart(2, "0")}`;
    } else {
      return "Error: cart is empty";
    }
  } else {
    return "Error: cart is null";
  }
};

console.log(registerUser({ name: "Alice", email: "alice@email.com", age: 20, isBanned: false }));
console.log(registerUser({ name: "", email: "a@b.com", age: 20, isBanned: false }));
console.log(registerUser({ name: "Bob", email: "bob@email.com", age: 10, isBanned: false }));
console.log(registerUser({ name: "Eve", email: "eve@email.com", age: 25, isBanned: true }));

console.log();

const items = [
  { product: { name: "Widget", price: 999, stock: 50 }, quantity: 2 },
  { product: { name: "Gadget", price: 2499, stock: 10 }, quantity: 1 },
];
console.log(calculateCartTotal(items, ""));
console.log(calculateCartTotal(items, "SAVE10"));
console.log(calculateCartTotal(null, ""));
console.log(calculateCartTotal([], ""));

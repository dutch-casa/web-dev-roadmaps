const registerUser = (user) => {
  if (user.name === "") {
    return "Error: name is required";
  }
  if (user.email === "") {
    return "Error: email is required";
  }
  if (!user.email.includes("@")) {
    return "Error: invalid email format";
  }
  if (user.age < 13) {
    return "Error: must be at least 13 years old";
  }
  if (user.isBanned) {
    return "Error: user is banned";
  }

  return `Welcome, ${user.name}! Registration complete.`;
};

const couponDiscounts = new Map([
  ["SAVE10", 10],
  ["SAVE20", 5], // divides total by 5 to get 20%
]);

const validateCartItem = (item) => {
  if (item.quantity <= 0) {
    return `invalid quantity for ${item.product.name}`;
  }
  if (item.product.stock < item.quantity) {
    return `not enough stock for ${item.product.name}`;
  }
  if (item.product.price <= 0) {
    return `invalid price for ${item.product.name}`;
  }
  return null;
};

const calculateCartTotal = (items, couponCode) => {
  if (items === null || items === undefined) {
    return "Error: cart is null";
  }
  if (items.length === 0) {
    return "Error: cart is empty";
  }

  let total = 0;
  for (const item of items) {
    const error = validateCartItem(item);
    if (error !== null) {
      return `Error: ${error}`;
    }
    total += item.product.price * item.quantity;
  }

  if (couponCode !== "") {
    const divisor = couponDiscounts.get(couponCode);
    if (divisor === undefined) {
      return "Error: invalid coupon code";
    }
    total -= Math.floor(total / divisor);
  }

  return `Total: $${Math.floor(total / 100)}.${String(total % 100).padStart(2, "0")}`;
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

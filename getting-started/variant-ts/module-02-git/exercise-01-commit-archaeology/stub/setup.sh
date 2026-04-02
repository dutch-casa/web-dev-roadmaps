#!/bin/bash

# This script creates a sample repo with history for you to explore.
# Run it once, then answer the questions in questions.md.

set -e

REPO_DIR="archaeology-repo"

rm -rf "$REPO_DIR"
mkdir "$REPO_DIR"
cd "$REPO_DIR"
git init

# Commit 1
cat > index.js << 'EOF'
console.log("Welcome to the store");
EOF
cat > package.json << 'EOF'
{
  "name": "store",
  "version": "1.0.0",
  "module": "index.js"
}
EOF
git add .
git commit -m "initial commit: basic store greeting"

# Commit 2
cat > index.js << 'EOF'
const greet = (name) => {
  return `Welcome to the store, ${name}!`;
};

console.log(greet("customer"));
EOF
git add .
git commit -m "add personalized greeting function"

# Commit 3
cat > inventory.js << 'EOF'
const defaultInventory = () => {
  return [
    { name: "Notebook", price: 4.99, stock: 50 },
    { name: "Pen", price: 1.49, stock: 200 },
    { name: "Eraser", price: 0.99, stock: 100 },
  ];
};

module.exports = { defaultInventory };
EOF
git add .
git commit -m "add inventory with three items"

# Commit 4
cat > index.js << 'EOF'
const { defaultInventory } = require("./inventory");

const greet = (name) => {
  return `Welcome to the store, ${name}!`;
};

const listItems = (items) => {
  for (const item of items) {
    const name = item.name.padEnd(10);
    console.log(`  ${name} $${item.price.toFixed(2)} (${item.stock} in stock)`);
  }
};

console.log(greet("customer"));
console.log("\nToday's inventory:");
listItems(defaultInventory());
EOF
git add .
git commit -m "display inventory in main output"

# Commit 5
cat > inventory.js << 'EOF'
const defaultInventory = () => {
  return [
    { name: "Notebook", price: 4.99, stock: 50 },
    { name: "Pen", price: 1.49, stock: 200 },
    { name: "Eraser", price: 0.99, stock: 100 },
    { name: "Ruler", price: 2.49, stock: 75 },
  ];
};

module.exports = { defaultInventory };
EOF
git add .
git commit -m "add ruler to inventory"

echo ""
echo "Done. The repo is in ./$REPO_DIR"
echo "cd into it and start answering questions.md"

#!/bin/bash

# This script creates a repo where a bug was introduced in one of the commits.
# Your job: use git bisect to find which commit broke it, then revert it.

set -e

REPO_DIR="bisect-repo"

rm -rf "$REPO_DIR"
mkdir "$REPO_DIR"
cd "$REPO_DIR"
git init

cat > package.json << 'EOF'
{
  "name": "calculator",
  "version": "1.0.0",
  "module": "index.js"
}
EOF

# Commit 1 — working
cat > calc.js << 'EOF'
const add = (a, b) => a + b;
const subtract = (a, b) => a - b;
const multiply = (a, b) => a * b;

module.exports = { add, subtract, multiply };
EOF
cat > index.js << 'EOF'
const { add, subtract, multiply } = require("./calc");

console.log("2 + 3 =", add(2, 3));
console.log("10 - 4 =", subtract(10, 4));
console.log("5 * 6 =", multiply(5, 6));
EOF
git add .
git commit -m "initial calculator with add, subtract, multiply"

# Commit 2 — working
cat > calc.js << 'EOF'
const add = (a, b) => a + b;
const subtract = (a, b) => a - b;
const multiply = (a, b) => a * b;

const divide = (a, b) => {
  if (b === 0) {
    return 0;
  }
  return Math.floor(a / b);
};

module.exports = { add, subtract, multiply, divide };
EOF
cat > index.js << 'EOF'
const { add, subtract, multiply, divide } = require("./calc");

console.log("2 + 3 =", add(2, 3));
console.log("10 - 4 =", subtract(10, 4));
console.log("5 * 6 =", multiply(5, 6));
console.log("15 / 3 =", divide(15, 3));
EOF
git add .
git commit -m "add divide function with zero check"

# Commit 3 — BUG INTRODUCED: subtract is broken
cat > calc.js << 'EOF'
const add = (a, b) => a + b;
const subtract = (a, b) => a + b;
const multiply = (a, b) => a * b;

const divide = (a, b) => {
  if (b === 0) {
    return 0;
  }
  return Math.floor(a / b);
};

module.exports = { add, subtract, multiply, divide };
EOF
git add .
git commit -m "refactor: clean up calc functions"

# Commit 4 — working (but bug still present)
cat > index.js << 'EOF'
const { add, subtract, multiply, divide } = require("./calc");

console.log("=== Calculator ===");
console.log("2 + 3 =", add(2, 3));
console.log("10 - 4 =", subtract(10, 4));
console.log("5 * 6 =", multiply(5, 6));
console.log("15 / 3 =", divide(15, 3));
EOF
git add .
git commit -m "add header to calculator output"

# Commit 5 — working (but bug still present)
cat > calc.js << 'EOF'
const add = (a, b) => a + b;
const subtract = (a, b) => a + b;
const multiply = (a, b) => a * b;

const divide = (a, b) => {
  if (b === 0) {
    return 0;
  }
  return Math.floor(a / b);
};

const modulo = (a, b) => {
  if (b === 0) {
    return 0;
  }
  return a % b;
};

module.exports = { add, subtract, multiply, divide, modulo };
EOF
cat > index.js << 'EOF'
const { add, subtract, multiply, divide, modulo } = require("./calc");

console.log("=== Calculator ===");
console.log("2 + 3 =", add(2, 3));
console.log("10 - 4 =", subtract(10, 4));
console.log("5 * 6 =", multiply(5, 6));
console.log("15 / 3 =", divide(15, 3));
console.log("17 % 5 =", modulo(17, 5));
EOF
git add .
git commit -m "add modulo function"

echo ""
echo "Done. The repo is in ./$REPO_DIR"
echo ""
echo "Something is broken: 10 - 4 should be 6, but it's not."
echo "Use git bisect to find which commit introduced the bug."
echo "Then use git revert to undo it without destroying history."

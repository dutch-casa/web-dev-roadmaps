#!/bin/bash

# This script creates a repo where a bug was introduced in one of the commits.
# Your job: use git bisect to find which commit broke it, then revert it.

set -e

REPO_DIR="bisect-repo"

rm -rf "$REPO_DIR"
mkdir "$REPO_DIR"
cd "$REPO_DIR"
git init

cat > go.mod << 'EOF'
module calculator

go 1.26
EOF

# Commit 1 — working
cat > calc.go << 'EOF'
package main

func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a - b }
func multiply(a, b int) int { return a * b }
EOF
cat > main.go << 'EOF'
package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", add(2, 3))
	fmt.Println("10 - 4 =", subtract(10, 4))
	fmt.Println("5 * 6 =", multiply(5, 6))
}
EOF
git add .
git commit -m "initial calculator with add, subtract, multiply"

# Commit 2 — working
cat >> calc.go << 'EOF'

func divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
EOF
cat > main.go << 'EOF'
package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", add(2, 3))
	fmt.Println("10 - 4 =", subtract(10, 4))
	fmt.Println("5 * 6 =", multiply(5, 6))
	fmt.Println("15 / 3 =", divide(15, 3))
}
EOF
git add .
git commit -m "add divide function with zero check"

# Commit 3 — BUG INTRODUCED: subtract is broken
cat > calc.go << 'EOF'
package main

func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a + b }
func multiply(a, b int) int { return a * b }

func divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
EOF
git add .
git commit -m "refactor: clean up calc functions"

# Commit 4 — working (but bug still present)
cat > main.go << 'EOF'
package main

import "fmt"

func main() {
	fmt.Println("=== Calculator ===")
	fmt.Println("2 + 3 =", add(2, 3))
	fmt.Println("10 - 4 =", subtract(10, 4))
	fmt.Println("5 * 6 =", multiply(5, 6))
	fmt.Println("15 / 3 =", divide(15, 3))
}
EOF
git add .
git commit -m "add header to calculator output"

# Commit 5 — working (but bug still present)
cat > calc.go << 'EOF'
package main

func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a + b }
func multiply(a, b int) int { return a * b }

func divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func modulo(a, b int) int {
	if b == 0 {
		return 0
	}
	return a % b
}
EOF
cat > main.go << 'EOF'
package main

import "fmt"

func main() {
	fmt.Println("=== Calculator ===")
	fmt.Println("2 + 3 =", add(2, 3))
	fmt.Println("10 - 4 =", subtract(10, 4))
	fmt.Println("5 * 6 =", multiply(5, 6))
	fmt.Println("15 / 3 =", divide(15, 3))
	fmt.Println("17 % 5 =", modulo(17, 5))
}
EOF
git add .
git commit -m "add modulo function"

echo ""
echo "Done. The repo is in ./$REPO_DIR"
echo ""
echo "Something is broken: 10 - 4 should be 6, but it's not."
echo "Use git bisect to find which commit introduced the bug."
echo "Then use git revert to undo it without destroying history."

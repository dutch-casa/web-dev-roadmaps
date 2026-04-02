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
cat > main.go << 'EOF'
package main

import "fmt"

func main() {
	fmt.Println("Welcome to the store")
}
EOF
cat > go.mod << 'EOF'
module store

go 1.26
EOF
git add .
git commit -m "initial commit: basic store greeting"

# Commit 2
cat > main.go << 'EOF'
package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Welcome to the store, %s!", name)
}

func main() {
	fmt.Println(greet("customer"))
}
EOF
git add .
git commit -m "add personalized greeting function"

# Commit 3
cat > inventory.go << 'EOF'
package main

type Item struct {
	Name  string
	Price float64
	Stock int
}

func defaultInventory() []Item {
	return []Item{
		{Name: "Notebook", Price: 4.99, Stock: 50},
		{Name: "Pen", Price: 1.49, Stock: 200},
		{Name: "Eraser", Price: 0.99, Stock: 100},
	}
}
EOF
git add .
git commit -m "add inventory with three items"

# Commit 4
cat > main.go << 'EOF'
package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Welcome to the store, %s!", name)
}

func listItems(items []Item) {
	for _, item := range items {
		fmt.Printf("  %-10s $%.2f (%d in stock)\n", item.Name, item.Price, item.Stock)
	}
}

func main() {
	fmt.Println(greet("customer"))
	fmt.Println("\nToday's inventory:")
	listItems(defaultInventory())
}
EOF
git add .
git commit -m "display inventory in main output"

# Commit 5
cat > inventory.go << 'EOF'
package main

type Item struct {
	Name  string
	Price float64
	Stock int
}

func defaultInventory() []Item {
	return []Item{
		{Name: "Notebook", Price: 4.99, Stock: 50},
		{Name: "Pen", Price: 1.49, Stock: 200},
		{Name: "Eraser", Price: 0.99, Stock: 100},
		{Name: "Ruler", Price: 2.49, Stock: 75},
	}
}
EOF
git add .
git commit -m "add ruler to inventory"

echo ""
echo "Done. The repo is in ./$REPO_DIR"
echo "cd into it and start answering questions.md"

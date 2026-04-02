package main

import (
	"fmt"
	"information-hiding/inventory"
)

func main() {
	store := inventory.NewStore("Auburn Supply Co.")

	sku1, _ := store.AddItem("Notebook", 499, 50, "supplies")
	sku2, _ := store.AddItem("Pen", 149, 200, "supplies")
	_, _ = store.AddItem("Backpack", 3999, 25, "bags")

	fmt.Printf("Store: %s\n", store.Name())
	fmt.Printf("Items: %d\n", store.ItemCount())
	fmt.Printf("Total inventory value: $%.2f\n", float64(store.TotalValue())/100)

	fmt.Println()

	// Sell some items — through the interface, which enforces invariants.
	if err := store.Sell(sku1, 5); err != nil {
		fmt.Println("Error:", err)
	}
	if err := store.Sell(sku2, 10); err != nil {
		fmt.Println("Error:", err)
	}

	// This would have been possible with the old design:
	//   store.Items[0].Stock = -999
	//
	// Now it's impossible. The fields are unexported. The only way
	// to reduce stock is through Sell(), which checks for negatives.

	// Try to oversell — the error is informative.
	if err := store.Sell(sku1, 9999); err != nil {
		fmt.Println("Oversell blocked:", err)
	}

	fmt.Printf("After sales: $%.2f\n", float64(store.TotalValue())/100)
}

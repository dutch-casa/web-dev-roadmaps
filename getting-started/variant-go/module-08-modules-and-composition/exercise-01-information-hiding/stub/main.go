package main

import (
	"fmt"
	"information-hiding/inventory"
)

// This code reaches into the package's internals.
// After you seal the package, update this to use only the public interface.

func main() {
	store := &inventory.Store{
		Name:    "Auburn Supply Co.",
		NextSku: 0,
	}

	// Add some items
	sku1 := inventory.AddItem(store, "Notebook", 499, 50, "supplies")
	sku2 := inventory.AddItem(store, "Pen", 149, 200, "supplies")
	sku3 := inventory.AddItem(store, "Backpack", 3999, 25, "bags")

	fmt.Printf("Store: %s\n", store.Name)
	fmt.Printf("Items: %d\n", len(store.Items))
	fmt.Printf("Total inventory value: $%.2f\n", float64(inventory.TotalValue(store))/100)

	fmt.Println()

	// Sell some items
	if err := inventory.RemoveStock(store, sku1, 5); err != nil {
		fmt.Println("Error:", err)
	}
	if err := inventory.RemoveStock(store, sku2, 10); err != nil {
		fmt.Println("Error:", err)
	}

	// This is the kind of thing callers shouldn't be able to do:
	// directly reaching in and breaking invariants
	store.Items[0].Stock = -999 // whoops

	fmt.Printf("After sales: $%.2f\n", float64(inventory.TotalValue(store))/100)

	_ = sku3
}

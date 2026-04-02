package inventory

import "fmt"

// Information hiding exercise
//
// This package exposes everything. Every field is exported. Every
// helper function is exported. The caller can reach into the guts
// of the inventory and break invariants.
//
// Your job: seal it. Make the interface as small as possible while
// keeping the same functionality. The rules:
//
//   - Only export what callers actually need
//   - Enforce invariants through the interface (e.g., stock can't go negative)
//   - Provide a constructor that ensures valid initial state
//   - Every internal detail should be unexported
//
// After you're done, go to main.go and update the usage to work
// with the sealed interface.

type Item struct {
	Name     string
	PriceCents int
	Stock    int
	Category string
	Sku      string
}

type Store struct {
	Items    []Item
	Name     string
	NextSku  int
}

func GenerateSku(store *Store) string {
	store.NextSku++
	return fmt.Sprintf("SKU-%04d", store.NextSku)
}

func FindItemBySku(store *Store, sku string) *Item {
	for i := range store.Items {
		if store.Items[i].Sku == sku {
			return &store.Items[i]
		}
	}
	return nil
}

func AddItem(store *Store, name string, priceCents int, stock int, category string) string {
	sku := GenerateSku(store)
	store.Items = append(store.Items, Item{
		Name:       name,
		PriceCents: priceCents,
		Stock:      stock,
		Category:   category,
		Sku:        sku,
	})
	return sku
}

func RemoveStock(store *Store, sku string, quantity int) error {
	item := FindItemBySku(store, sku)
	if item == nil {
		return fmt.Errorf("item %s not found", sku)
	}
	item.Stock -= quantity // BUG: no check for negative stock!
	return nil
}

func TotalValue(store *Store) int {
	total := 0
	for _, item := range store.Items {
		total += item.PriceCents * item.Stock
	}
	return total
}

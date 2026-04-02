package inventory

import "fmt"

// The sealed interface. Callers see: NewStore, Store.AddItem,
// Store.Sell, Store.TotalValue, Store.ItemCount, Store.Name.
// Nothing else. They can't touch stock directly, can't generate
// SKUs, can't access the slice of items.

type item struct {
	name       string
	priceCents int
	stock      int
	category   string
	sku        string
}

type Store struct {
	name    string
	items   []item
	nextSku int
}

func NewStore(name string) *Store {
	return &Store{name: name}
}

func (s *Store) Name() string { return s.name }

func (s *Store) ItemCount() int { return len(s.items) }

func (s *Store) generateSku() string {
	s.nextSku++
	return fmt.Sprintf("SKU-%04d", s.nextSku)
}

func (s *Store) findBySku(sku string) *item {
	for i := range s.items {
		if s.items[i].sku == sku {
			return &s.items[i]
		}
	}
	return nil
}

// AddItem creates a new inventory item and returns its SKU.
// Price must be positive. Stock must be non-negative.
func (s *Store) AddItem(name string, priceCents int, stock int, category string) (string, error) {
	if priceCents <= 0 {
		return "", fmt.Errorf("price must be positive, got %d", priceCents)
	}
	if stock < 0 {
		return "", fmt.Errorf("stock must be non-negative, got %d", stock)
	}
	sku := s.generateSku()
	s.items = append(s.items, item{
		name:       name,
		priceCents: priceCents,
		stock:      stock,
		category:   category,
		sku:        sku,
	})
	return sku, nil
}

// Sell reduces stock for an item. Quantity must be positive and
// cannot exceed available stock.
func (s *Store) Sell(sku string, quantity int) error {
	if quantity <= 0 {
		return fmt.Errorf("quantity must be positive, got %d", quantity)
	}
	it := s.findBySku(sku)
	if it == nil {
		return fmt.Errorf("item %s not found", sku)
	}
	if it.stock < quantity {
		return fmt.Errorf("insufficient stock for %s: have %d, need %d",
			it.name, it.stock, quantity)
	}
	it.stock -= quantity
	return nil
}

// TotalValue returns the total inventory value in cents.
func (s *Store) TotalValue() int {
	total := 0
	for _, it := range s.items {
		total += it.priceCents * it.stock
	}
	return total
}

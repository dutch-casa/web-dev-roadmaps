package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name       string
	Email      string
	Age        int
	IsVerified bool
	IsBanned   bool
}

func registerUser(user User) string {
	if user.Name == "" {
		return "Error: name is required"
	}
	if user.Email == "" {
		return "Error: email is required"
	}
	if !strings.Contains(user.Email, "@") {
		return "Error: invalid email format"
	}
	if user.Age < 13 {
		return "Error: must be at least 13 years old"
	}
	if user.IsBanned {
		return "Error: user is banned"
	}

	return fmt.Sprintf("Welcome, %s! Registration complete.", user.Name)
}

type Product struct {
	Name     string
	Price    int
	Stock    int
	Category string
}

type CartItem struct {
	Product  Product
	Quantity int
}

var couponDiscounts = map[string]int{
	"SAVE10": 10,
	"SAVE20": 5, // divides total by 5 to get 20%
}

func validateCartItem(item CartItem) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("invalid quantity for %s", item.Product.Name)
	}
	if item.Product.Stock < item.Quantity {
		return fmt.Errorf("not enough stock for %s", item.Product.Name)
	}
	if item.Product.Price <= 0 {
		return fmt.Errorf("invalid price for %s", item.Product.Name)
	}
	return nil
}

func calculateCartTotal(items []CartItem, couponCode string) string {
	if items == nil {
		return "Error: cart is nil"
	}
	if len(items) == 0 {
		return "Error: cart is empty"
	}

	total := 0
	for _, item := range items {
		if err := validateCartItem(item); err != nil {
			return "Error: " + err.Error()
		}
		total += item.Product.Price * item.Quantity
	}

	if couponCode != "" {
		divisor, ok := couponDiscounts[couponCode]
		if !ok {
			return "Error: invalid coupon code"
		}
		total -= total / divisor
	}

	return fmt.Sprintf("Total: $%d.%02d", total/100, total%100)
}

func main() {
	fmt.Println(registerUser(User{Name: "Alice", Email: "alice@email.com", Age: 20}))
	fmt.Println(registerUser(User{Name: "", Email: "a@b.com", Age: 20}))
	fmt.Println(registerUser(User{Name: "Bob", Email: "bob@email.com", Age: 10}))
	fmt.Println(registerUser(User{Name: "Eve", Email: "eve@email.com", Age: 25, IsBanned: true}))

	fmt.Println()

	items := []CartItem{
		{Product: Product{Name: "Widget", Price: 999, Stock: 50}, Quantity: 2},
		{Product: Product{Name: "Gadget", Price: 2499, Stock: 10}, Quantity: 1},
	}
	fmt.Println(calculateCartTotal(items, ""))
	fmt.Println(calculateCartTotal(items, "SAVE10"))
	fmt.Println(calculateCartTotal(nil, ""))
	fmt.Println(calculateCartTotal([]CartItem{}, ""))
}

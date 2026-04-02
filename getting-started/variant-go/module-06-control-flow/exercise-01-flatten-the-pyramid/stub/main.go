package main

import (
	"fmt"
	"strings"
)

// Flatten the pyramid
//
// Each function below has deeply nested control flow.
// Refactor each one using guard clauses so the maximum
// nesting depth is 2 (one level inside the function body,
// one more for a loop if needed).
//
// Rules:
//   - Don't change the behavior. Output must be identical.
//   - Use guard clauses (early returns) to handle edge cases at the top.
//   - The "happy path" logic should be at the shallowest indentation level.

type User struct {
	Name       string
	Email      string
	Age        int
	IsVerified bool
	IsBanned   bool
}

// Problem 1: Nested validation
func registerUser(user User) string {
	if user.Name != "" {
		if user.Email != "" {
			if strings.Contains(user.Email, "@") {
				if user.Age >= 13 {
					if !user.IsBanned {
						return fmt.Sprintf("Welcome, %s! Registration complete.", user.Name)
					} else {
						return "Error: user is banned"
					}
				} else {
					return "Error: must be at least 13 years old"
				}
			} else {
				return "Error: invalid email format"
			}
		} else {
			return "Error: email is required"
		}
	} else {
		return "Error: name is required"
	}
}

type Product struct {
	Name     string
	Price    int // cents
	Stock    int
	Category string
}

type CartItem struct {
	Product  Product
	Quantity int
}

// Problem 2: Nested loop with conditions
func calculateCartTotal(items []CartItem, couponCode string) string {
	if items != nil {
		if len(items) > 0 {
			total := 0
			for _, item := range items {
				if item.Quantity > 0 {
					if item.Product.Stock >= item.Quantity {
						if item.Product.Price > 0 {
							total += item.Product.Price * item.Quantity
						} else {
							return fmt.Sprintf("Error: invalid price for %s", item.Product.Name)
						}
					} else {
						return fmt.Sprintf("Error: not enough stock for %s", item.Product.Name)
					}
				} else {
					return fmt.Sprintf("Error: invalid quantity for %s", item.Product.Name)
				}
			}
			if couponCode != "" {
				if couponCode == "SAVE10" {
					total = total - total/10
				} else if couponCode == "SAVE20" {
					total = total - total/5
				} else {
					return "Error: invalid coupon code"
				}
			}
			return fmt.Sprintf("Total: $%d.%02d", total/100, total%100)
		} else {
			return "Error: cart is empty"
		}
	} else {
		return "Error: cart is nil"
	}
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

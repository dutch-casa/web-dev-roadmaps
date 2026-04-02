package main

import "fmt"

// --- Domain types ---

type Account struct {
	Owner   string
	Balance int // cents
}

// --- Error types ---
// Each error type carries the data needed to understand what went wrong.
// The caller can type-switch on these to handle specific failures.

type ErrInvalidAmount struct {
	Amount int
}

func (e ErrInvalidAmount) Error() string {
	return fmt.Sprintf("invalid amount: %d cents (must be positive)", e.Amount)
}

type ErrInsufficientFunds struct {
	Requested int
	Available int
}

func (e ErrInsufficientFunds) Error() string {
	return fmt.Sprintf("insufficient funds: requested %d cents, available %d cents",
		e.Requested, e.Available)
}

type ErrSameAccount struct {
	Owner string
}

func (e ErrSameAccount) Error() string {
	return fmt.Sprintf("cannot transfer to the same account (%s)", e.Owner)
}

// --- Operations ---

func (a *Account) Deposit(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount{Amount: amount}
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount{Amount: amount}
	}
	if a.Balance < amount {
		return ErrInsufficientFunds{Requested: amount, Available: a.Balance}
	}
	a.Balance -= amount
	return nil
}

func Transfer(from, to *Account, amount int) error {
	if from == to {
		return ErrSameAccount{Owner: from.Owner}
	}
	if err := from.Withdraw(amount); err != nil {
		return err
	}
	// Deposit can't fail here — amount is already validated as positive
	// by the Withdraw call above. But we handle the error anyway because
	// the signature says it can fail, and ignoring errors is a bad habit.
	return to.Deposit(amount)
}

func main() {
	alice := &Account{Owner: "Alice", Balance: 10000}
	bob := &Account{Owner: "Bob", Balance: 5000}

	fmt.Println("--- Successful operations ---")
	if err := alice.Deposit(2500); err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("Alice deposited $25.00. Balance: $%.2f\n", float64(alice.Balance)/100)
	}

	if err := Transfer(alice, bob, 3000); err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("Transferred $30.00. Alice: $%.2f, Bob: $%.2f\n",
			float64(alice.Balance)/100, float64(bob.Balance)/100)
	}

	fmt.Println("\n--- Error cases ---")

	if err := alice.Deposit(-100); err != nil {
		fmt.Println("Negative deposit:", err)
	}

	if err := bob.Withdraw(99999); err != nil {
		fmt.Println("Overdraw:", err)
	}

	if err := Transfer(alice, alice, 1000); err != nil {
		fmt.Println("Self-transfer:", err)
	}
}

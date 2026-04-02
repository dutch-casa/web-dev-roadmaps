package main

import "fmt"

// Errors as values
//
// You're building a small bank account system. The business rules:
//
//   - An account has an owner name and a balance in cents.
//   - You can deposit any positive amount.
//   - You can withdraw if you have sufficient funds.
//   - You can transfer between two accounts.
//   - Negative amounts, zero amounts, and overdrafts are errors.
//
// Your job: implement the functions below so that errors are
// values in the return type — never panics, never printed-and-ignored.
// Each function returns an error that tells the caller exactly
// what went wrong, using typed errors that callers can inspect.
//
// The three error types to define:
//   - ErrInvalidAmount: deposit or withdrawal amount was <= 0
//   - ErrInsufficientFunds: withdrawal would overdraw the account
//   - ErrSameAccount: transfer source and destination are the same
//
// Implement these as concrete types that implement the error interface.

// TODO: Define the Account type. An account has an Owner (string)
// and Balance (int, in cents).

// TODO: Define error types:
//   ErrInvalidAmount       — should include the rejected amount
//   ErrInsufficientFunds   — should include the requested amount and available balance
//   ErrSameAccount         — should include the account owner name

// TODO: Implement Deposit.
// Rules:
//   - Amount must be > 0, otherwise return ErrInvalidAmount
//   - Add the amount to the balance
//   - Return nil on success

// TODO: Implement Withdraw.
// Rules:
//   - Amount must be > 0, otherwise return ErrInvalidAmount
//   - Balance must be >= amount, otherwise return ErrInsufficientFunds
//   - Subtract the amount from the balance
//   - Return nil on success

// TODO: Implement Transfer.
// Rules:
//   - Source and destination must be different accounts (compare pointers),
//     otherwise return ErrSameAccount
//   - Withdraw from source, then deposit to destination
//   - If the withdrawal fails, return that error (don't deposit)
//   - Amount validation is handled by Withdraw, don't duplicate it

func main() {
	// Uncomment this when your types are defined:

	// alice := &Account{Owner: "Alice", Balance: 10000} // $100.00
	// bob := &Account{Owner: "Bob", Balance: 5000}      // $50.00

	// // Successful operations
	// fmt.Println("--- Successful operations ---")
	// if err := alice.Deposit(2500); err != nil {
	// 	fmt.Println("ERROR:", err)
	// } else {
	// 	fmt.Printf("Alice deposited $25.00. Balance: $%.2f\n", float64(alice.Balance)/100)
	// }

	// if err := Transfer(alice, bob, 3000); err != nil {
	// 	fmt.Println("ERROR:", err)
	// } else {
	// 	fmt.Printf("Transferred $30.00. Alice: $%.2f, Bob: $%.2f\n",
	// 		float64(alice.Balance)/100, float64(bob.Balance)/100)
	// }

	// // Error cases — each should return a specific error type
	// fmt.Println("\n--- Error cases ---")

	// if err := alice.Deposit(-100); err != nil {
	// 	fmt.Println("Negative deposit:", err)
	// }

	// if err := bob.Withdraw(99999); err != nil {
	// 	fmt.Println("Overdraw:", err)
	// }

	// if err := Transfer(alice, alice, 1000); err != nil {
	// 	fmt.Println("Self-transfer:", err)
	// }

	fmt.Println("Define the types and uncomment the code above.")
}

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
// values in the return type — never thrown exceptions, never
// printed-and-ignored. Each function returns a Result that tells
// the caller exactly what went wrong using typed errors the
// caller can inspect.
//
// TypeScript approach: use a Result discriminated union instead
// of try/catch. This makes the error path explicit in the type.

// The Result type — every operation returns one of these.
// TODO: Define the Result type:
//   type Result<T, E> = { ok: true; value: T } | { ok: false; error: E };

// The three error types to define as a discriminated union:
// TODO: Define BankError:
//   - { tag: "invalidAmount"; amount: number }
//   - { tag: "insufficientFunds"; requested: number; available: number }
//   - { tag: "sameAccount"; owner: string }

// TODO: Define the Account class with:
//   - owner: string (readonly)
//   - balance: number (cents)
//   - deposit(amount: number): Result<void, BankError>
//   - withdraw(amount: number): Result<void, BankError>

// TODO: Define the transfer function:
//   transfer(from: Account, to: Account, amount: number): Result<void, BankError>
// Rules:
//   - Source and destination must be different objects,
//     otherwise return sameAccount error
//   - Withdraw from source, then deposit to destination
//   - If the withdrawal fails, return that error (don't deposit)
//   - Amount validation is handled by withdraw, don't duplicate it

// TODO: Define formatError(error: BankError): string
// that returns a human-readable message for each error variant.

// Uncomment this when your types are defined:

// const alice = new Account("Alice", 10000); // $100.00
// const bob = new Account("Bob", 5000);      // $50.00

// // Successful operations
// console.log("--- Successful operations ---");
// const depositResult = alice.deposit(2500);
// if (depositResult.ok) {
//   console.log(`Alice deposited $25.00. Balance: $${(alice.balance / 100).toFixed(2)}`);
// } else {
//   console.log("ERROR:", formatError(depositResult.error));
// }

// const transferResult = transfer(alice, bob, 3000);
// if (transferResult.ok) {
//   console.log(`Transferred $30.00. Alice: $${(alice.balance / 100).toFixed(2)}, Bob: $${(bob.balance / 100).toFixed(2)}`);
// } else {
//   console.log("ERROR:", formatError(transferResult.error));
// }

// // Error cases — each should return a specific error type
// console.log("\n--- Error cases ---");

// const negDeposit = alice.deposit(-100);
// if (!negDeposit.ok) {
//   console.log("Negative deposit:", formatError(negDeposit.error));
// }

// const overdraw = bob.withdraw(99999);
// if (!overdraw.ok) {
//   console.log("Overdraw:", formatError(overdraw.error));
// }

// const selfTransfer = transfer(alice, alice, 1000);
// if (!selfTransfer.ok) {
//   console.log("Self-transfer:", formatError(selfTransfer.error));
// }

console.log("Define the types and uncomment the code above.");

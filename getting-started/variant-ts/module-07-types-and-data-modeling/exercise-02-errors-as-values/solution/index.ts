// --- Result type ---

type Result<T, E> = { ok: true; value: T } | { ok: false; error: E };

// --- Error types ---
// Each variant carries the data needed to understand what went wrong.
// The caller can narrow on the tag to handle specific failures.

type BankError =
  | { tag: "invalidAmount"; amount: number }
  | { tag: "insufficientFunds"; requested: number; available: number }
  | { tag: "sameAccount"; owner: string };

const formatError = (error: BankError): string => {
  switch (error.tag) {
    case "invalidAmount":
      return `invalid amount: ${error.amount} cents (must be positive)`;
    case "insufficientFunds":
      return `insufficient funds: requested ${error.requested} cents, available ${error.available} cents`;
    case "sameAccount":
      return `cannot transfer to the same account (${error.owner})`;
  }
};

// --- Account ---

class Account {
  readonly owner: string;
  balance: number; // cents

  constructor(owner: string, balance: number) {
    this.owner = owner;
    this.balance = balance;
  }

  deposit(amount: number): Result<void, BankError> {
    if (amount <= 0) {
      return { ok: false, error: { tag: "invalidAmount", amount } };
    }
    this.balance += amount;
    return { ok: true, value: undefined };
  }

  withdraw(amount: number): Result<void, BankError> {
    if (amount <= 0) {
      return { ok: false, error: { tag: "invalidAmount", amount } };
    }
    if (this.balance < amount) {
      return { ok: false, error: { tag: "insufficientFunds", requested: amount, available: this.balance } };
    }
    this.balance -= amount;
    return { ok: true, value: undefined };
  }
}

// --- Transfer ---

const transfer = (from: Account, to: Account, amount: number): Result<void, BankError> => {
  if (from === to) {
    return { ok: false, error: { tag: "sameAccount", owner: from.owner } };
  }
  const withdrawResult = from.withdraw(amount);
  if (!withdrawResult.ok) {
    return withdrawResult;
  }
  // Deposit can't fail here — amount is already validated as positive
  // by the withdraw call above. But we return the result anyway because
  // ignoring return values is a bad habit.
  return to.deposit(amount);
};

// --- Main ---

const alice = new Account("Alice", 10000); // $100.00
const bob = new Account("Bob", 5000);      // $50.00

console.log("--- Successful operations ---");
const depositResult = alice.deposit(2500);
if (depositResult.ok) {
  console.log(`Alice deposited $25.00. Balance: $${(alice.balance / 100).toFixed(2)}`);
} else {
  console.log("ERROR:", formatError(depositResult.error));
}

const transferResult = transfer(alice, bob, 3000);
if (transferResult.ok) {
  console.log(`Transferred $30.00. Alice: $${(alice.balance / 100).toFixed(2)}, Bob: $${(bob.balance / 100).toFixed(2)}`);
} else {
  console.log("ERROR:", formatError(transferResult.error));
}

console.log("\n--- Error cases ---");

const negDeposit = alice.deposit(-100);
if (!negDeposit.ok) {
  console.log("Negative deposit:", formatError(negDeposit.error));
}

const overdraw = bob.withdraw(99999);
if (!overdraw.ok) {
  console.log("Overdraw:", formatError(overdraw.error));
}

const selfTransfer = transfer(alice, alice, 1000);
if (!selfTransfer.ok) {
  console.log("Self-transfer:", formatError(selfTransfer.error));
}

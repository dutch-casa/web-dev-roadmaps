// Named conditions
//
// Each function below has complex boolean expressions used inline.
// Extract each one into a named boolean variable that explains
// the intent. The code should read like English when you're done.
//
// Example:
//   Before: if (x > 0 && x < 100 && y !== 0)
//   After:  const inRange = x > 0 && x < 100 && y !== 0;
//           if (inRange) { ... }
//
// Run with: bun run index.js

// Problem 1: Who qualifies for a bonus?
const qualifiesForBonus = (emp) => {
  return emp.isFullTime && emp.yearsEmployed >= 2 && !emp.onProbation && emp.annualSalary < 150000 && (Date.now() - emp.lastReviewDate) < 365 * 24 * 60 * 60 * 1000;
};

// Problem 2: Who can approve a purchase?
const canApprovePurchase = (emp, amount) => {
  if (emp.isManager && emp.isFullTime && !emp.onProbation) {
    if (amount <= 5000 || (amount <= 50000 && emp.yearsEmployed >= 5) || (emp.department === "Finance" && emp.yearsEmployed >= 10)) {
      return true;
    }
  }
  return false;
};

// Problem 3: Should we send a reminder email?
const shouldSendReminder = (emp, lastEmailDays, hasUnreadNotifications, emailOptIn) => {
  return emailOptIn && !emp.onProbation && emp.isFullTime && lastEmailDays > 7 && hasUnreadNotifications && emp.department !== "Legal";
};

const emp = {
  name: "Jordan",
  department: "Engineering",
  yearsEmployed: 6,
  isFullTime: true,
  isManager: true,
  annualSalary: 95000,
  lastReviewDate: Date.now() - (6 * 30 * 24 * 60 * 60 * 1000), // ~6 months ago
  onProbation: false,
};

console.log(`${emp.name} qualifies for bonus: ${qualifiesForBonus(emp)}`);
console.log(`${emp.name} can approve $3000: ${canApprovePurchase(emp, 3000)}`);
console.log(`${emp.name} can approve $25000: ${canApprovePurchase(emp, 25000)}`);
console.log(`Send reminder to ${emp.name}: ${shouldSendReminder(emp, 10, true, true)}`);

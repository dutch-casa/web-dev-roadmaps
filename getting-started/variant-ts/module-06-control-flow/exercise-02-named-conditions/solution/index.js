const qualifiesForBonus = (emp) => {
  const isEligibleStatus = emp.isFullTime && !emp.onProbation;
  const hasEnoughTenure = emp.yearsEmployed >= 2;
  const isBelowSalaryCap = emp.annualSalary < 150000;
  const hasRecentReview = (Date.now() - emp.lastReviewDate) < 365 * 24 * 60 * 60 * 1000;

  return isEligibleStatus && hasEnoughTenure && isBelowSalaryCap && hasRecentReview;
};

const canApprovePurchase = (emp, amount) => {
  const isActiveManager = emp.isManager && emp.isFullTime && !emp.onProbation;
  if (!isActiveManager) {
    return false;
  }

  const isSmallPurchase = amount <= 5000;
  const isMediumWithSeniority = amount <= 50000 && emp.yearsEmployed >= 5;
  const isFinanceVeteran = emp.department === "Finance" && emp.yearsEmployed >= 10;

  return isSmallPurchase || isMediumWithSeniority || isFinanceVeteran;
};

const shouldSendReminder = (emp, lastEmailDays, hasUnreadNotifications, emailOptIn) => {
  const wantsEmail = emailOptIn && emp.department !== "Legal";
  const isActiveEmployee = emp.isFullTime && !emp.onProbation;
  const isDueForReminder = lastEmailDays > 7 && hasUnreadNotifications;

  return wantsEmail && isActiveEmployee && isDueForReminder;
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

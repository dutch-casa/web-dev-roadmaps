// Problem 1: The result is a single expression built from string operations.
// If there's a middle initial, include it. Otherwise, skip it.
const formatName = (first, middle, last) => {
  const base = `${last.toUpperCase()}, ${first}`;
  if (middle !== "") {
    return `${base} ${middle[0]}.`;
  }
  return base;
};

// Problem 2: The boolean expression IS the answer.
// No temporary variables. No if/else. The condition is the return value.
const isEligible = (age, hasPermission) => age >= 18 && hasPermission;

// Problem 3: Early returns turn a nested ladder into a linear scan.
// Each threshold gets one line. The reader sees the mapping instantly.
const letterGrade = (score) => {
  if (score >= 90) return "A";
  if (score >= 80) return "B";
  if (score >= 70) return "C";
  if (score >= 60) return "D";
  return "F";
};

// Problem 4: A helper function expresses the decision.
// The formatting and the decision are separate concerns.
const tempLabel = (f) => {
  if (f > 100) return "scorching";
  if (f > 80) return "hot";
  if (f > 60) return "pleasant";
  if (f > 40) return "cold";
  return "freezing";
};

const describeTemp = (celsius) => {
  const fahrenheit = celsius * 9 / 5 + 32;
  return `${Math.round(fahrenheit)}\u00B0F \u2014 ${tempLabel(fahrenheit)}`;
};

console.log(formatName("Rosa", "Louise", "Parks"));
console.log(formatName("Guido", "", "van Rossum"));
console.log();

console.log("Eligible (21, true):", isEligible(21, true));
console.log("Eligible (16, true):", isEligible(16, true));
console.log();

for (const score of [95, 82, 74, 65, 48]) {
  console.log(`Score ${score} \u2192 ${letterGrade(score)}`);
}
console.log();

for (const temp of [42, 18, -5, 30, 55]) {
  console.log(describeTemp(temp));
}

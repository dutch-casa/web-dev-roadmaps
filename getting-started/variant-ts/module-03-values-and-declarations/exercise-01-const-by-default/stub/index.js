// Every value below is declared with `let`.
// Your job: convert every one that COULD be `const` into `const`.
// Only use `let` when the variable is reassigned later.
// Never use `var`.
//
// For each one you leave as `let`, add a comment explaining
// why it needs to stay mutable.
//
// Run with: bun run index.js

let appName = "Grade Calculator";
let version = "1.0.0";
let maxScore = 100;

let studentName = "Jordan";
let scores = [88, 92, 75, 96, 84];

let total = 0;
for (let i = 0; i < scores.length; i++) {
  total += scores[i];
}

let average = total / scores.length;

let passing = average >= 70;

let grade = "F";
if (average >= 90) {
  grade = "A";
} else if (average >= 80) {
  grade = "B";
} else if (average >= 70) {
  grade = "C";
} else if (average >= 60) {
  grade = "D";
}

console.log(appName, version);
console.log(`Student: ${studentName}`);
console.log(`Scores: ${scores.join(", ")}`);
console.log(`Average: ${average.toFixed(1)} / ${maxScore}`);
console.log(`Passing: ${passing}`);
console.log(`Grade: ${grade}`);

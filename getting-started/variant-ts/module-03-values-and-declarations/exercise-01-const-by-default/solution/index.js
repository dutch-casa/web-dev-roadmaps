const appName = "Grade Calculator";
const version = "1.0.0";
const maxScore = 100;

const studentName = "Jordan";

// scores is an array — `const` prevents reassignment of the binding,
// but we never reassign it. The array contents could be mutated,
// but we don't do that either. `const` is correct here.
const scores = [88, 92, 75, 96, 84];

// total is accumulated in a loop — it changes on every iteration.
let total = 0;
for (let i = 0; i < scores.length; i++) {
  total += scores[i];
}

// average is computed once and never changes after this line.
// Even though it depends on runtime values, `const` works because
// JS `const` just means "this binding won't be reassigned."
const average = total / scores.length;

// passing is derived from average. Computed once, never reassigned.
const passing = average >= 70;

// grade is reassigned inside the if/else chain. Must be `let`.
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

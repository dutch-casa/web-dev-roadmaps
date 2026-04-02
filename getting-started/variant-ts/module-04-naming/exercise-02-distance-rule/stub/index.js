// Distance rule exercise
//
// Each function below has naming problems related to scope distance.
// Some names are too long for their scope. Some are too short.
// Fix each one and add a comment explaining your reasoning.
//
// Run with: bun run index.js

// Problem 1: Names are too verbose for a tight loop.
const sumOfSquares = (inputNumbers) => {
  let totalSumOfAllSquaredValues = 0;
  for (let currentIndex = 0; currentIndex < inputNumbers.length; currentIndex++) {
    const currentSquaredValue = inputNumbers[currentIndex] * inputNumbers[currentIndex];
    totalSumOfAllSquaredValues += currentSquaredValue;
  }
  return totalSumOfAllSquaredValues;
};

// Problem 2: Names are too short for a function this long.
// The variable 'r' is used 15 lines after it's assigned.
const processStudentGrades = (grades) => {
  for (const n of Object.keys(grades)) {
    const g = grades[n];
    let t = 0;
    for (const v of g) {
      t += v;
    }
    const a = t / g.length;

    let r = "fail";
    if (a >= 90) {
      r = "excellent";
    } else if (a >= 80) {
      r = "good";
    } else if (a >= 70) {
      r = "satisfactory";
    } else if (a >= 60) {
      r = "needs improvement";
    }

    console.log(`${n.padEnd(10)} avg=${a.toFixed(1)}  ${r}`);
  }
};

// Problem 3: This function has a name that doesn't describe
// what it actually does. It also has parameter names that don't help.
const doIt = (s, n) => {
  let result = "";
  for (let i = 0; i < n; i++) {
    result += s;
  }
  return result;
};

console.log("Sum of squares:", sumOfSquares([1, 2, 3, 4, 5]));

console.log();
processStudentGrades({
  Alice: [92, 88, 95],
  Bob: [71, 65, 78],
  Clara: [86, 91, 83],
});

console.log();
console.log(doIt("JS! ", 3));

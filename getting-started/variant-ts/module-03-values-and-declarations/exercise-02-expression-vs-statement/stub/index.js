// Expression vs. statement
//
// Expressions produce values. Statements perform actions.
// The practical difference: expressions compose (you can nest them),
// statements are sequential steps.
//
// This exercise is about recognizing when statement-heavy code
// can be simplified by thinking in terms of expressions.
// Each function below does something clunky with temporary variables
// and reassignment. Rewrite each one so the logic flows through
// expressions instead of being stapled together with statements.
//
// Run "bun run index.js" before and after — output must be identical.

// Problem 1: String building through reassignment.
// This builds a formatted name by mutating a variable five times.
// Rewrite it so the result is computed in one or two expressions.
const formatName = (first, middle, last) => {
  let result = "";

  result = last.toUpperCase();

  result = result + ", ";

  result = result + first;

  if (middle !== "") {
    result = result + " " + middle[0] + ".";
  }

  return result;
};

// Problem 2: Boolean logic buried under if/else.
// The if/else statements assign true or false to a variable.
// That's just... the boolean expression itself. Simplify.
const isEligible = (age, hasPermission) => {
  let ageCheck;
  if (age >= 18) {
    ageCheck = true;
  } else {
    ageCheck = false;
  }

  let result;
  if (ageCheck && hasPermission) {
    result = true;
  } else {
    result = false;
  }

  return result;
};

// Problem 3: Computing a letter grade through a chain of reassignment.
// Each branch sets the same variable. There's a simpler way to express
// "map a number to a category."
const letterGrade = (score) => {
  let grade;
  if (score >= 90) {
    grade = "A";
  } else {
    if (score >= 80) {
      grade = "B";
    } else {
      if (score >= 70) {
        grade = "C";
      } else {
        if (score >= 60) {
          grade = "D";
        } else {
          grade = "F";
        }
      }
    }
  }
  return grade;
};

// Problem 4: Describing a temperature by mutating through branches.
// This function reassigns `desc` in every branch, then uses it once.
// Express the decision as a value instead.
const describeTemp = (celsius) => {
  const fahrenheit = celsius * 9 / 5 + 32;

  let desc;
  if (fahrenheit > 100) {
    desc = "scorching";
  } else if (fahrenheit > 80) {
    desc = "hot";
  } else if (fahrenheit > 60) {
    desc = "pleasant";
  } else if (fahrenheit > 40) {
    desc = "cold";
  } else {
    desc = "freezing";
  }

  return `${Math.round(fahrenheit)}\u00B0F \u2014 ${desc}`;
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

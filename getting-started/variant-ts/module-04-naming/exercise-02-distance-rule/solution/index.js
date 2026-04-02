// Fixed: Short names in a tight loop. The scope is 3 lines.
// 'total' and 'n' are fine here — 'n' is used immediately.
const sumOfSquares = (numbers) => {
  let total = 0;
  for (const n of numbers) {
    total += n * n;
  }
  return total;
};

// Fixed: Longer names for a function where variables are used far
// from their declaration. 'rating' is used 10 lines below where
// it's assigned — the reader needs the name to remind them what it is.
const printGradeReport = (gradesByStudent) => {
  for (const name of Object.keys(gradesByStudent)) {
    const grades = gradesByStudent[name];
    let total = 0;
    for (const g of grades) {
      total += g;
    }
    const average = total / grades.length;

    let rating = "fail";
    if (average >= 90) {
      rating = "excellent";
    } else if (average >= 80) {
      rating = "good";
    } else if (average >= 70) {
      rating = "satisfactory";
    } else if (average >= 60) {
      rating = "needs improvement";
    }

    console.log(`${name.padEnd(10)} avg=${average.toFixed(1)}  ${rating}`);
  }
};

// Fixed: The function name describes the transformation.
// Parameters tell the caller what to pass.
const repeatString = (text, count) => text.repeat(count);

console.log("Sum of squares:", sumOfSquares([1, 2, 3, 4, 5]));

console.log();
printGradeReport({
  Alice: [92, 88, 95],
  Bob: [71, 65, 78],
  Clara: [86, 91, 83],
});

console.log();
console.log(repeatString("JS! ", 3));

// File organization exercise
//
// All the code below lives in one file and it's a mess.
// Your job: split it into multiple files based on what each piece does.
//
// Suggested split:
//   index.js     — just the main logic (composition)
//   student.js   — the student-related functions
//   format.js    — formatting and display functions
//   validate.js  — validation logic
//
// Rules:
//   - Use ES modules: export function foo() { } and import { foo } from './foo.js'
//   - "bun run index.js" must produce the same output
//   - Each file should be understandable without reading the others

const formatCurrency = (cents) => {
  const dollars = Math.floor(cents / 100);
  const remainder = cents % 100;
  return `$${dollars}.${String(remainder).padStart(2, "0")}`;
};

const validateEmail = (email) => {
  return email.includes("@") && email.includes(".");
};

const validateStudent = (student) => {
  const errors = [];
  if (student.name === "") {
    errors.push("name is required");
  }
  if (!validateEmail(student.email)) {
    errors.push("invalid email");
  }
  if (student.scores.length === 0) {
    errors.push("at least one score required");
  }
  for (const score of student.scores) {
    if (score < 0 || score > 100) {
      errors.push(`score ${score} out of range`);
    }
  }
  if (student.tuition < 0) {
    errors.push("tuition cannot be negative");
  }
  return errors;
};

const averageScore = (scores) => {
  let total = 0;
  for (const s of scores) {
    total += s;
  }
  return total / scores.length;
};

const formatStudentSummary = (student) => {
  const avg = averageScore(student.scores);
  return `${student.name.padEnd(10)} ${student.email.padEnd(25)} avg=${avg.toFixed(1)} tuition=${formatCurrency(student.tuition)}`;
};

const formatValidationErrors = (name, errors) => {
  return `INVALID ${name.padEnd(10)}: ${errors.join("; ")}`;
};

const students = [
  { name: "Alice", email: "alice@auburn.edu", scores: [92, 88, 95], tuition: 1250000 },
  { name: "Bob", email: "bob@auburn.edu", scores: [71, 65, 78], tuition: 1250000 },
  { name: "", email: "nobody", scores: [], tuition: -100 },
  { name: "Clara", email: "clara@auburn.edu", scores: [86, 91, 83], tuition: 1100000 },
];

console.log("=== Student Report ===");
console.log();

for (const student of students) {
  const errors = validateStudent(student);
  if (errors.length > 0) {
    console.log(formatValidationErrors(student.name, errors));
    continue;
  }
  console.log(formatStudentSummary(student));
}

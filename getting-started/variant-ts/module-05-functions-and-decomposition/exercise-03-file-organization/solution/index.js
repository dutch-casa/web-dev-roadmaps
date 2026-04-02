import { validateStudent } from "./validate.js";
import { formatStudentSummary, formatValidationErrors } from "./format.js";

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

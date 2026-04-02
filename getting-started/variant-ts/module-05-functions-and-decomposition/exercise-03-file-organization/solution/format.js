import { averageScore } from "./student.js";

const formatCurrency = (cents) => {
  const dollars = Math.floor(cents / 100);
  const remainder = cents % 100;
  return `$${dollars}.${String(remainder).padStart(2, "0")}`;
};

export const formatStudentSummary = (student) => {
  const avg = averageScore(student.scores);
  return `${student.name.padEnd(10)} ${student.email.padEnd(25)} avg=${avg.toFixed(1)} tuition=${formatCurrency(student.tuition)}`;
};

export const formatValidationErrors = (name, errors) => {
  return `INVALID ${name.padEnd(10)}: ${errors.join("; ")}`;
};

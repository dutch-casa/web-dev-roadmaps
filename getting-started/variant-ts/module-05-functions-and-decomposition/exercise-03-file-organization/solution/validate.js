const validateEmail = (email) => {
  return email.includes("@") && email.includes(".");
};

export const validateStudent = (student) => {
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

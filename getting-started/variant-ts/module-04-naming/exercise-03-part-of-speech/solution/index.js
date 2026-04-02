// Booleans are predicates — reads as "is eligible"
const isEligible = (applicant) => {
  return applicant.age >= 18 && applicant.isActive;
};

// Functions are verbs — "calculate average score"
const calculateAverageScore = (scores) => {
  if (scores.length === 0) {
    return 0;
  }
  let total = 0;
  for (const s of scores) {
    total += s;
  }
  return total / scores.length;
};

// Function is a verb — "format decision"
const formatDecision = (applicant) => {
  if (!isEligible(applicant)) {
    return "DENIED";
  }
  const average = calculateAverageScore(applicant.scores);

  let standing = "regular";
  if (average >= 90) {
    standing = "honors";
  }

  return `${applicant.firstName} ${applicant.lastName} (approved) - avg: ${average.toFixed(1)} - ${standing}`;
};

// Collection is a plural noun
const applicants = [
  { firstName: "Jordan", lastName: "Lee", age: 20, isActive: true, scores: [92, 88, 95] },
  { firstName: "Sam", lastName: "Park", age: 17, isActive: true, scores: [85, 79, 91] },
  { firstName: "Alex", lastName: "Chen", age: 22, isActive: false, scores: [96, 94, 98] },
  { firstName: "Casey", lastName: "Jones", age: 19, isActive: true, scores: [72, 68, 75] },
];

for (const applicant of applicants) {
  console.log(formatDecision(applicant));
}

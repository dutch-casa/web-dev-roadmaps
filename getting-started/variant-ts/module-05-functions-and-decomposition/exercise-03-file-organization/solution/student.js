export const averageScore = (scores) => {
  let total = 0;
  for (const s of scores) {
    total += s;
  }
  return total / scores.length;
};

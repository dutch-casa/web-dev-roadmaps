// Extract till you drop
//
// This function does everything in one place: parsing, validating,
// computing, formatting, and printing. Your job is to break it into
// small functions where each one does one thing.
//
// Guidelines:
//   - No function should exceed ~25 lines
//   - Each extracted function gets a clear verb name
//   - The main composition function should read like a summary
//   - Keep pure logic separate from printing (side effects)
//   - Run "bun run index.js" before and after — output must be identical

const generateReport = (rawData) => {
  const lines = rawData.split("\n");
  const names = [];
  const scores = [];

  for (const rawLine of lines) {
    const line = rawLine.trim();
    if (line === "") {
      continue;
    }
    const parts = line.split(":");
    if (parts.length !== 2) {
      console.log(`WARNING: skipping malformed line: "${line}"`);
      continue;
    }
    const name = parts[0].trim();
    if (name === "") {
      console.log("WARNING: skipping line with empty name");
      continue;
    }
    const scoreStrs = parts[1].trim().split(",");
    const studentScores = [];
    let valid = true;
    for (const raw of scoreStrs) {
      const s = raw.trim();
      const n = parseInt(s, 10);
      if (isNaN(n)) {
        console.log(`WARNING: invalid score "${s}" for ${name}`);
        valid = false;
        break;
      }
      if (n < 0 || n > 100) {
        console.log(`WARNING: score ${n} out of range for ${name}`);
        valid = false;
        break;
      }
      studentScores.push(n);
    }
    if (!valid) {
      continue;
    }
    names.push(name);
    scores.push(studentScores);
  }

  if (names.length === 0) {
    console.log("No valid student data found.");
    return;
  }

  console.log("=== Grade Report ===");
  console.log();

  let classTotal = 0;
  for (let i = 0; i < names.length; i++) {
    let total = 0;
    for (const s of scores[i]) {
      total += s;
    }
    const avg = total / scores[i].length;
    classTotal += avg;

    let grade = "F";
    if (avg >= 90) {
      grade = "A";
    } else if (avg >= 80) {
      grade = "B";
    } else if (avg >= 70) {
      grade = "C";
    } else if (avg >= 60) {
      grade = "D";
    }

    const bar = "\u2588".repeat(Math.floor(avg / 5));
    console.log(`${names[i].padEnd(12)} avg=${avg.toFixed(1).padStart(5)}  grade=${grade}  ${bar}`);
  }

  const classAvg = classTotal / names.length;
  console.log();
  console.log(`Class average: ${classAvg.toFixed(1)}`);
  console.log(`Students: ${names.length}`);
};

const data = `
    Alice: 92, 88, 95, 91
    Bob: 71, 65, 78, 72
    Clara: 86, 91, 83, 89
    Dave: 55, 62, 48, 51
    Eve: 98, 95, 100, 97
`;
generateReport(data);

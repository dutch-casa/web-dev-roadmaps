// Grade thresholds as a data table instead of branching.
const gradeThresholds = [
  { minAverage: 90, letter: "A" },
  { minAverage: 80, letter: "B" },
  { minAverage: 70, letter: "C" },
  { minAverage: 60, letter: "D" },
];

// --- Parsing (boundary layer: handles raw input) ---

const parseScore = (raw) => {
  const s = raw.trim();
  const n = parseInt(s, 10);
  if (isNaN(n)) {
    return { value: 0, ok: false };
  }
  if (n < 0 || n > 100) {
    return { value: 0, ok: false };
  }
  return { value: n, ok: true };
};

const parseLine = (line) => {
  const parts = line.split(":");
  if (parts.length !== 2) {
    return { record: null, error: `malformed line: "${line}"` };
  }
  const name = parts[0].trim();
  if (name === "") {
    return { record: null, error: `empty name in line: "${line}"` };
  }
  const scoreStrs = parts[1].trim().split(",");
  const scores = [];
  for (const raw of scoreStrs) {
    const { value, ok } = parseScore(raw);
    if (!ok) {
      return { record: null, error: `invalid score "${raw.trim()}" for ${name}` };
    }
    scores.push(value);
  }
  return { record: { name, scores }, error: null };
};

const parseStudentData = (rawData) => {
  const records = [];
  const warnings = [];
  for (const rawLine of rawData.split("\n")) {
    const line = rawLine.trim();
    if (line === "") continue;
    const { record, error } = parseLine(line);
    if (error !== null) {
      warnings.push(error);
      continue;
    }
    records.push(record);
  }
  return { records, warnings };
};

// --- Pure computation (no I/O) ---

const averageScore = (scores) => {
  let total = 0;
  for (const s of scores) {
    total += s;
  }
  return total / scores.length;
};

const letterGrade = (avg) => {
  for (const threshold of gradeThresholds) {
    if (avg >= threshold.minAverage) {
      return threshold.letter;
    }
  }
  return "F";
};

const classAverage = (records) => {
  let total = 0;
  for (const r of records) {
    total += averageScore(r.scores);
  }
  return total / records.length;
};

// --- Formatting (pure: produces strings, no side effects) ---

const formatStudentLine = (record) => {
  const avg = averageScore(record.scores);
  const grade = letterGrade(avg);
  const bar = "\u2588".repeat(Math.floor(avg / 5));
  return `${record.name.padEnd(12)} avg=${avg.toFixed(1).padStart(5)}  grade=${grade}  ${bar}`;
};

// --- Composition (the only place with side effects) ---

const printReport = (records) => {
  if (records.length === 0) {
    console.log("No valid student data found.");
    return;
  }

  console.log("=== Grade Report ===");
  console.log();
  for (const r of records) {
    console.log(formatStudentLine(r));
  }
  console.log();
  console.log(`Class average: ${classAverage(records).toFixed(1)}`);
  console.log(`Students: ${records.length}`);
};

const data = `
    Alice: 92, 88, 95, 91
    Bob: 71, 65, 78, 72
    Clara: 86, 91, 83, 89
    Dave: 55, 62, 48, 51
    Eve: 98, 95, 100, 97
`;
const { records, warnings } = parseStudentData(data);
for (const w of warnings) console.log(`WARNING: ${w}`);
printReport(records);

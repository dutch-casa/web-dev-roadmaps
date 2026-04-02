// Name audit
//
// This program works correctly. The problem is that every name is
// meaningless. Your job: rename every variable, function, and parameter
// so that someone can understand the code without reading the implementation.
//
// Rules:
//   - Don't change the behavior. Only change names.
//   - The code should read almost like English when you're done.
//   - Run "bun run index.js" before and after to verify it still works.

const do1 = (data) => {
  const result = [];
  for (const d of data) {
    if (d.w) {
      result.push(d);
    }
  }
  return result;
};

const do2 = (data) => {
  let temp = 0;
  let n = 0;
  for (const d of data) {
    temp += d.z;
    n++;
  }
  if (n === 0) {
    return 0;
  }
  return temp / n;
};

const do3 = (data) => {
  const m = {};
  for (const d of data) {
    if (!m[d.y]) {
      m[d.y] = [];
    }
    m[d.y].push(d);
  }
  return m;
};

const do4 = (d) => {
  const parts = [];
  parts.push(d.x);
  parts.push(`(year ${d.y})`);
  parts.push(`$${d.z.toFixed(2)}`);
  if (d.w) {
    parts.push("[available]");
  } else {
    parts.push("[sold]");
  }
  return parts.join(" ");
};

const stuff = [
  { x: "JS Programming", y: 2023, z: 49.99, w: true },
  { x: "The Art of SQL", y: 2021, z: 39.95, w: false },
  { x: "Systems Design", y: 2024, z: 54.99, w: true },
  { x: "Network Protocols", y: 2022, z: 44.50, w: true },
  { x: "Data Structures", y: 2023, z: 42.00, w: false },
];

console.log("All:");
for (const d of stuff) {
  console.log(" ", do4(d));
}

const temp = do1(stuff);
console.log(`\nAvailable: ${temp.length}`);
for (const d of temp) {
  console.log(" ", do4(d));
}

console.log(`\nAverage price: $${do2(stuff).toFixed(2)}`);

const m = do3(stuff);
console.log("\nBy year:");
for (const k of Object.keys(m)) {
  console.log(`  ${k}: ${m[k].length} items`);
}

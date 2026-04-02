// Part of speech as design
//
// The naming conventions:
//   Objects    -> nouns          (student, order)
//   Functions  -> verbs          (calculate, send, parse)
//   Booleans   -> predicates     (isActive, hasPermission)
//   Collections-> plural nouns   (users, scores)
//
// Rename everything in this program to follow these conventions.
// The code should read almost like English sentences when you're done.
//
// Run with: bun run index.js

const check = (d) => {
  return d.c >= 18 && d.d;
};

const getStuff = (d) => {
  let t = 0;
  for (const x of d.list) {
    t += x;
  }
  if (d.list.length === 0) {
    return 0;
  }
  return t / d.list.length;
};

const doThing = (d) => {
  if (!check(d)) {
    return "DENIED";
  }
  const avg = getStuff(d);
  let status = "regular";
  if (avg >= 90) {
    status = "honors";
  }
  return `${d.a} ${d.b} (approved) - avg: ${avg.toFixed(1)} - ${status}`;
};

const things = [
  { a: "Jordan", b: "Lee", c: 20, d: true, list: [92, 88, 95] },
  { a: "Sam", b: "Park", c: 17, d: true, list: [85, 79, 91] },
  { a: "Alex", b: "Chen", c: 22, d: false, list: [96, 94, 98] },
  { a: "Casey", b: "Jones", c: 19, d: true, list: [72, 68, 75] },
];

for (const t of things) {
  console.log(doThing(t));
}

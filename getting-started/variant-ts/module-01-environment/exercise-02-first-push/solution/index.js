const name = "Your Name";
const today = new Date().toLocaleDateString("en-US", {
  year: "numeric",
  month: "long",
  day: "numeric",
});

console.log(`Hello, I'm ${name}. Today is ${today}.`);

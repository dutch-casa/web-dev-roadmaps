// main is the composition -- the imperative shell.
// It creates data, orchestrates operations, and prints output.
// The logic lives in the task and display modules.

import { createTask, startTask, completeTask } from "./task/index";
import { formatTask, formatSummary } from "./display/index";

const tasks = [
  createTask("Set up development environment", 1),
  createTask("Learn Git basics", 1),
  createTask("Complete naming exercises", 2),
  createTask("Read about control flow", 2),
  createTask("Build terminal Wordle", 3),
];

startTask(tasks[0]);
completeTask(tasks[0]);
startTask(tasks[1]);

console.log("=== Task Report ===");
console.log();
for (const t of tasks) {
  console.log("  " + formatTask(t));
}
console.log();
console.log(formatSummary(tasks));

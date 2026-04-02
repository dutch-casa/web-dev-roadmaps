// Display module: formatting tasks for terminal output.
// Knows about the task module but nothing about main or I/O.
// Functions return strings -- the caller decides where to print them.

import { type Task, countByStatus } from "../task/index";

export const formatTask = (task: Task): string => {
  const priority = "!".repeat(task.priority);
  return `[${task.status}] ${task.title.padEnd(30)} ${priority}`;
};

export const formatSummary = (tasks: Task[]): string => {
  const counts = countByStatus(tasks);
  return `Total: ${tasks.length} | Todo: ${counts.get("todo") ?? 0} | In Progress: ${counts.get("in-progress") ?? 0} | Done: ${counts.get("done") ?? 0}`;
};

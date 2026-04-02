// Boundary drawing exercise
//
// This program is a tiny task manager. Everything is in one file.
// Your job: split it into modules at the natural boundaries.
//
// Suggested modules:
//   task/index.ts    -- the Task type, status values, and task logic
//   display/index.ts -- formatting and display functions
//   index.ts         -- just the composition (create tasks, print report)
//
// Rules:
//   - Each module should be understandable without reading the others
//   - The task module should know nothing about printing or formatting
//   - The display module should know about tasks but not about main
//   - Only export what callers actually need
//   - "bun run index" must produce the same output
//
// Think about: where are the natural seams? What knowledge belongs
// where? What should be hidden?

type TaskStatus = "todo" | "in-progress" | "done";

type Task = {
  title: string;
  status: TaskStatus;
  priority: number; // 1 = highest
  createdAt: Date;
  doneAt: Date | null;
};

const createTask = (title: string, priority: number): Task => ({
  title,
  status: "todo",
  priority,
  createdAt: new Date(),
  doneAt: null,
});

const startTask = (task: Task): void => {
  task.status = "in-progress";
};

const completeTask = (task: Task): void => {
  task.status = "done";
  task.doneAt = new Date();
};

const isOverdue = (task: Task, deadline: Date): boolean =>
  task.status !== "done" && new Date() > deadline;

const filterByStatus = (tasks: Task[], status: TaskStatus): Task[] =>
  tasks.filter((t) => t.status === status);

const countByStatus = (tasks: Task[]): Map<TaskStatus, number> => {
  const counts = new Map<TaskStatus, number>();
  for (const t of tasks) {
    counts.set(t.status, (counts.get(t.status) ?? 0) + 1);
  }
  return counts;
};

const formatTask = (task: Task): string => {
  const priority = "!".repeat(task.priority);
  return `[${task.status}] ${task.title.padEnd(30)} ${priority}`;
};

const formatSummary = (tasks: Task[]): string => {
  const counts = countByStatus(tasks);
  return `Total: ${tasks.length} | Todo: ${counts.get("todo") ?? 0} | In Progress: ${counts.get("in-progress") ?? 0} | Done: ${counts.get("done") ?? 0}`;
};

const printReport = (tasks: Task[]): void => {
  console.log("=== Task Report ===");
  console.log();
  for (const t of tasks) {
    console.log("  " + formatTask(t));
  }
  console.log();
  console.log(formatSummary(tasks));
};

// --- Main ---

const tasks: Task[] = [
  createTask("Set up development environment", 1),
  createTask("Learn Git basics", 1),
  createTask("Complete naming exercises", 2),
  createTask("Read about control flow", 2),
  createTask("Build terminal Wordle", 3),
];

startTask(tasks[0]);
completeTask(tasks[0]);
startTask(tasks[1]);

printReport(tasks);

// Suppress unused-variable warnings for functions the student
// should move but might not call directly.
void isOverdue;
void filterByStatus;

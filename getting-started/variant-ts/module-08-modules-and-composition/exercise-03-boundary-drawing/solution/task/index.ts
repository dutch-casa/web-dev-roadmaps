// Task module: the Task type, status values, and pure operations on tasks.
// Knows nothing about display or I/O.

export type TaskStatus = "todo" | "in-progress" | "done";

export type Task = {
  title: string;
  status: TaskStatus;
  priority: number;
  createdAt: Date;
  doneAt: Date | null;
};

export const createTask = (title: string, priority: number): Task => ({
  title,
  status: "todo",
  priority,
  createdAt: new Date(),
  doneAt: null,
});

export const startTask = (task: Task): void => {
  task.status = "in-progress";
};

export const completeTask = (task: Task): void => {
  task.status = "done";
  task.doneAt = new Date();
};

export const filterByStatus = (tasks: Task[], status: TaskStatus): Task[] =>
  tasks.filter((t) => t.status === status);

export const countByStatus = (tasks: Task[]): Map<TaskStatus, number> => {
  const counts = new Map<TaskStatus, number>();
  for (const t of tasks) {
    counts.set(t.status, (counts.get(t.status) ?? 0) + 1);
  }
  return counts;
};

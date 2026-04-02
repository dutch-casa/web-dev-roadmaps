// Problem 1: Media player — one variant per state.
// A stopped player has no track. A playing player must have one.
// You can't construct a "stopped with track" — the type doesn't allow it.

type PlayerState =
  | { tag: "stopped" }
  | { tag: "playing"; track: string; position: number }
  | { tag: "paused"; track: string; position: number };

const describePlayer = (state: PlayerState): string => {
  switch (state.tag) {
    case "stopped":
      return "Player is stopped";
    case "playing":
      return `Playing "${state.track}" at ${state.position}s`;
    case "paused":
      return `Paused "${state.track}" at ${state.position}s`;
  }
};

// Problem 2: Traffic light — string literal union.
// Only three values exist. "purple" is not a LightColor.

type LightColor = "red" | "yellow" | "green";

type TrafficLight = {
  color: LightColor;
  duration: number; // seconds
};

// Problem 3: Login form — one variant per stage.
// Each stage carries only the data that's meaningful for that stage.
// A Failed form has an error message. An Authenticated form has a token.
// You can't have both.

type LoginState =
  | { tag: "empty" }
  | { tag: "partial"; username: string }
  | { tag: "complete"; username: string; password: string }
  | { tag: "submitted"; username: string }
  | { tag: "failed"; username: string; error: string }
  | { tag: "authenticated"; username: string; sessionToken: string };

const describeLogin = (state: LoginState): string => {
  switch (state.tag) {
    case "empty":
      return "Login form: empty";
    case "partial":
      return `Login form: username=${state.username}, awaiting password`;
    case "complete":
      return `Login form: ready to submit (${state.username})`;
    case "submitted":
      return `Login form: submitted, waiting... (${state.username})`;
    case "failed":
      return `Login form: failed — ${state.error}`;
    case "authenticated":
      return `Login form: authenticated (token=${state.sessionToken})`;
  }
};

// --- Problem 1: Each state carries exactly its data. No nonsense possible. ---
console.log(describePlayer({ tag: "stopped" }));
console.log(describePlayer({ tag: "playing", track: "Song.mp3", position: 42 }));
console.log(describePlayer({ tag: "paused", track: "Song.mp3", position: 42 }));

console.log();

// --- Problem 2: Only red, yellow, green exist. ---
const lights: TrafficLight[] = [
  { color: "red", duration: 30 },
  { color: "yellow", duration: 5 },
  { color: "green", duration: 25 },
];
for (const l of lights) {
  console.log(`Light: ${l.color} for ${l.duration}s`);
}

console.log();

// --- Problem 3: Each stage is a different variant. Can't mix them up. ---
console.log(describeLogin({ tag: "empty" }));
console.log(describeLogin({ tag: "partial", username: "jordan" }));
console.log(describeLogin({ tag: "complete", username: "jordan", password: "***" }));
console.log(describeLogin({ tag: "submitted", username: "jordan" }));
console.log(describeLogin({ tag: "failed", username: "jordan", error: "wrong password" }));
console.log(describeLogin({ tag: "authenticated", username: "jordan", sessionToken: "tok_abc123" }));

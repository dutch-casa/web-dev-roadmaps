// Illegal states
//
// The types below allow nonsense states. Your job: redesign them
// so that illegal states are impossible to construct.
//
// Problem 1: A media player that can be stopped, playing, or paused.
// When playing, it has a current track and position.
// When paused, it has a current track and position (where it paused).
// When stopped, it has no track and no position.
//
// The current design uses one type with everything optional.

type PlayerState = {
  status: "stopped" | "playing" | "paused";
  track?: string;   // undefined when stopped
  position?: number; // seconds, undefined when stopped
};

// Problem 2: A traffic light that can be red, yellow, or green.
// Each color has a duration in seconds.
//
// The current design uses a string — any string is valid.

type TrafficLight = {
  color: string; // any string works, even "purple"
  duration: number;
};

// Problem 3: A login form that goes through stages:
//   - Empty: no input yet
//   - Partial: username entered but not password
//   - Complete: both username and password entered
//   - Submitted: form has been sent, waiting for response
//   - Failed: submission returned an error message
//   - Authenticated: got back a session token
//
// The current design has boolean flags for each stage.

type LoginForm = {
  username: string;
  password: string;
  isSubmitted: boolean;
  isComplete: boolean;
  hasError: boolean;
  errorMessage: string;  // only meaningful if hasError
  sessionToken: string;  // only meaningful if authenticated
};

// After redesigning the types, write a main() that creates
// one value of each valid state and prints it.
// The key test: try to create an invalid state (like a paused player
// with no track). If the compiler stops you, you succeeded.
//
// Hints:
//   - Use discriminated unions with a "tag" field
//   - Use string literal unions for bounded sets
//   - Each variant carries only the data meaningful for that state

// Current (bad) usage — these nonsense states compile fine:
const bad1: PlayerState = { status: "stopped", track: "Song.mp3", position: 42 };
const bad2: TrafficLight = { color: "purple", duration: 10 };
const bad3: LoginForm = {
  username: "",
  password: "",
  isSubmitted: true,
  isComplete: false,
  hasError: true,
  errorMessage: "oops",
  sessionToken: "abc",
};

console.log("These should be impossible but aren't:");
console.log(`  Stopped player with track: ${JSON.stringify(bad1)}`);
console.log(`  Purple traffic light: ${JSON.stringify(bad2)}`);
console.log(`  Submitted incomplete form with error AND token: ${JSON.stringify(bad3)}`);

console.log();
console.log("TODO: Replace the types above and create valid states here.");

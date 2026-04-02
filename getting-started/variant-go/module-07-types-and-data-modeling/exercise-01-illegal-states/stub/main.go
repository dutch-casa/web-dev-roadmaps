package main

import "fmt"

// Illegal states
//
// The types below allow nonsense states. Your job: redesign them
// so that illegal states are impossible to construct.
//
// Start with Problem 1 (traffic light) — it's the simplest.
// Then Problem 2 (media player). Problem 3 (login form) is the hardest.
//
// ============================================================
// WORKED EXAMPLE: A door can be open, closed, or locked.
// When locked, it has a key code. When open or closed, it doesn't.
//
// Bad version (allows nonsense):
//
//   type Door struct {
//       State   string // "open", "closed", "locked"
//       KeyCode string // meaningless unless locked
//   }
//
// Good version (nonsense is impossible):
//
//   type DoorState interface{ doorState() }
//
//   type Open struct{}
//   type Closed struct{}
//   type Locked struct{ KeyCode string }
//
//   func (Open) doorState()   {}
//   func (Closed) doorState() {}
//   func (Locked) doorState() {}
//
// Now an Open door can't have a KeyCode — the field doesn't exist.
// A Locked door MUST have a KeyCode — the compiler requires it.
// The unexported doorState() method means only this package can
// add new states. It's a sealed set.
// ============================================================

// Problem 1: A traffic light that can be red, yellow, or green.
// Each color has a duration in seconds.
//
// The current design uses a string — any string is valid.
// Hint: use iota to create a closed set of colors.

type TrafficLight struct {
	Color    string // any string works, even "purple"
	Duration int
}

// Problem 2: A media player that can be stopped, playing, or paused.
// When playing, it has a current track and position.
// When paused, it has a current track and position (where it paused).
// When stopped, it has no track and no position.
//
// The current design uses one struct with everything optional.
// Hint: use separate types for each state, like the door example above.

type PlayerState struct {
	Status   string // "stopped", "playing", "paused"
	Track    string // empty when stopped
	Position int    // seconds, 0 when stopped
}

// Problem 3: A login form that goes through stages:
//   - Empty: no input yet
//   - Partial: username entered but not password
//   - Complete: both username and password entered
//   - Submitted: form has been sent, waiting for response
//   - Failed: submission returned an error message
//   - Authenticated: got back a session token
//
// The current design has boolean flags for each stage.
// Hint: one type per stage. Each carries only the data that
// makes sense for that stage.

type LoginForm struct {
	Username     string
	Password     string
	IsSubmitted  bool
	IsComplete   bool
	HasError     bool
	ErrorMessage string // only meaningful if HasError
	SessionToken string // only meaningful if authenticated
}

// After redesigning the types, write a main() that creates
// one value of each valid state and prints it.
// The key test: try to create an invalid state (like a paused player
// with no track). If the compiler stops you, you succeeded.

func main() {
	// Current (bad) usage — these nonsense states compile fine:
	bad1 := TrafficLight{Color: "purple", Duration: 10}
	bad2 := PlayerState{Status: "stopped", Track: "Song.mp3", Position: 42}
	bad3 := LoginForm{IsSubmitted: true, IsComplete: false, HasError: true, SessionToken: "abc"}

	fmt.Println("These should be impossible but aren't:")
	fmt.Printf("  Purple traffic light: %+v\n", bad1)
	fmt.Printf("  Stopped player with track: %+v\n", bad2)
	fmt.Printf("  Submitted incomplete form with error AND token: %+v\n", bad3)

	fmt.Println()
	fmt.Println("TODO: Replace the types above and create valid states here.")
}

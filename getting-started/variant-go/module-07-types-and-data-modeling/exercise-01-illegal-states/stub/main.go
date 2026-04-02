package main

import "fmt"

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
// The current design uses one struct with everything optional.

type PlayerState struct {
	Status   string // "stopped", "playing", "paused"
	Track    string // empty when stopped
	Position int    // seconds, 0 when stopped
}

// Problem 2: A traffic light that can be red, yellow, or green.
// Each color has a duration in seconds.
//
// The current design uses a string — any string is valid.

type TrafficLight struct {
	Color    string // any string works, even "purple"
	Duration int
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

type LoginForm struct {
	Username      string
	Password      string
	IsSubmitted   bool
	IsComplete    bool
	HasError      bool
	ErrorMessage  string // only meaningful if HasError
	SessionToken  string // only meaningful if authenticated
}

// After redesigning the types, write a main() that creates
// one value of each valid state and prints it.
// The key test: try to create an invalid state (like a paused player
// with no track). If the compiler stops you, you succeeded.

func main() {
	// Current (bad) usage — these nonsense states compile fine:
	bad1 := PlayerState{Status: "stopped", Track: "Song.mp3", Position: 42}
	bad2 := TrafficLight{Color: "purple", Duration: 10}
	bad3 := LoginForm{IsSubmitted: true, IsComplete: false, HasError: true, SessionToken: "abc"}

	fmt.Println("These should be impossible but aren't:")
	fmt.Printf("  Stopped player with track: %+v\n", bad1)
	fmt.Printf("  Purple traffic light: %+v\n", bad2)
	fmt.Printf("  Submitted incomplete form with error AND token: %+v\n", bad3)

	fmt.Println()
	fmt.Println("TODO: Replace the types above and create valid states here.")
}

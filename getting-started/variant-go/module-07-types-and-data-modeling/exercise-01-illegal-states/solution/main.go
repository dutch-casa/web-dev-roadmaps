package main

import (
	"fmt"
	"time"
)

// Problem 1: Media player — one type per state.
// A stopped player has no track. A playing player must have one.
// You can't construct a "stopped with track" — the type doesn't allow it.

type Stopped struct{}

type Playing struct {
	Track    string
	Position time.Duration
}

type Paused struct {
	Track    string
	Position time.Duration
}

// PlayerState is a sealed interface — only this package can implement it.
type PlayerState interface {
	playerState()
}

func (Stopped) playerState() {}
func (Playing) playerState() {}
func (Paused) playerState()  {}

func describePlayer(state PlayerState) string {
	switch s := state.(type) {
	case Stopped:
		return "Player is stopped"
	case Playing:
		return fmt.Sprintf("Playing %q at %v", s.Track, s.Position)
	case Paused:
		return fmt.Sprintf("Paused %q at %v", s.Track, s.Position)
	default:
		panic("unreachable: sealed interface")
	}
}

// Problem 2: Traffic light — enum with iota.
// Only three values exist. "purple" is not a LightColor.

type LightColor int

const (
	Red LightColor = iota
	Yellow
	Green
)

func (c LightColor) String() string {
	return [...]string{"Red", "Yellow", "Green"}[c]
}

type TrafficLight struct {
	Color    LightColor
	Duration time.Duration
}

// Problem 3: Login form — one type per stage.
// Each stage carries only the data that's meaningful for that stage.
// A Failed form has an error message. An Authenticated form has a token.
// You can't have both.

type LoginEmpty struct{}

type LoginPartial struct {
	Username string
}

type LoginComplete struct {
	Username string
	Password string
}

type LoginSubmitted struct {
	Username string
}

type LoginFailed struct {
	Username string
	Error    string
}

type LoginAuthenticated struct {
	Username     string
	SessionToken string
}

type LoginState interface {
	loginState()
}

func (LoginEmpty) loginState()         {}
func (LoginPartial) loginState()       {}
func (LoginComplete) loginState()      {}
func (LoginSubmitted) loginState()     {}
func (LoginFailed) loginState()        {}
func (LoginAuthenticated) loginState() {}

func describeLogin(state LoginState) string {
	switch s := state.(type) {
	case LoginEmpty:
		return "Login form: empty"
	case LoginPartial:
		return fmt.Sprintf("Login form: username=%s, awaiting password", s.Username)
	case LoginComplete:
		return fmt.Sprintf("Login form: ready to submit (%s)", s.Username)
	case LoginSubmitted:
		return fmt.Sprintf("Login form: submitted, waiting... (%s)", s.Username)
	case LoginFailed:
		return fmt.Sprintf("Login form: failed — %s", s.Error)
	case LoginAuthenticated:
		return fmt.Sprintf("Login form: authenticated (token=%s)", s.SessionToken)
	default:
		panic("unreachable: sealed interface")
	}
}

func main() {
	// Problem 1: Each state carries exactly its data. No nonsense possible.
	fmt.Println(describePlayer(Stopped{}))
	fmt.Println(describePlayer(Playing{Track: "Song.mp3", Position: 42 * time.Second}))
	fmt.Println(describePlayer(Paused{Track: "Song.mp3", Position: 42 * time.Second}))

	fmt.Println()

	// Problem 2: Only Red, Yellow, Green exist.
	lights := []TrafficLight{
		{Color: Red, Duration: 30 * time.Second},
		{Color: Yellow, Duration: 5 * time.Second},
		{Color: Green, Duration: 25 * time.Second},
	}
	for _, l := range lights {
		fmt.Printf("Light: %s for %v\n", l.Color, l.Duration)
	}

	fmt.Println()

	// Problem 3: Each stage is a different type. Can't mix them up.
	fmt.Println(describeLogin(LoginEmpty{}))
	fmt.Println(describeLogin(LoginPartial{Username: "jordan"}))
	fmt.Println(describeLogin(LoginComplete{Username: "jordan", Password: "***"}))
	fmt.Println(describeLogin(LoginSubmitted{Username: "jordan"}))
	fmt.Println(describeLogin(LoginFailed{Username: "jordan", Error: "wrong password"}))
	fmt.Println(describeLogin(LoginAuthenticated{Username: "jordan", SessionToken: "tok_abc123"}))
}

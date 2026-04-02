package ui

import (
	"wordle/game"
)

// ANSI color codes for terminal output.
const (
	colorReset  = "\033[0m"
	colorGreen  = "\033[42;30m" // green background, black text
	colorYellow = "\033[43;30m" // yellow background, black text
	colorGray   = "\033[100;37m" // gray background, white text
)

// ColorForResult returns the ANSI color code for a letter result.
func ColorForResult(r game.LetterResult) string {
	switch r {
	case game.Correct:
		return colorGreen
	case game.Misplaced:
		return colorYellow
	default:
		return colorGray
	}
}

// ReadGuess reads a guess from stdin.
// It should:
//   - Print a prompt showing the attempt number and remaining guesses
//   - Read a line from stdin
//   - Trim whitespace
//   - Convert to lowercase
//   - Return the cleaned string
func ReadGuess(attemptNum int, remaining int) string {
	// TODO: implement
	return ""
}

// DisplayResult prints a guess with colored letter tiles.
// Each letter gets a background color based on its LetterResult.
func DisplayResult(result game.GuessResult) {
	// TODO: implement
	// For each letter in the guess, print it with the appropriate
	// background color using the ANSI codes above.
}

// DisplayWin prints a congratulatory message.
func DisplayWin(attempts int) {
	// TODO: implement
}

// DisplayLoss prints the target word after a loss.
func DisplayLoss(target string) {
	// TODO: implement
}

// DisplayWelcome prints the game title and instructions.
func DisplayWelcome() {
	// TODO: implement
}

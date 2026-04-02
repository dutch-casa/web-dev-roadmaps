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
//   - Print a prompt showing the attempt number (e.g., "Guess 3/6: ")
//   - Read a line from stdin
//   - Trim whitespace
//   - Convert to lowercase
//   - Return the cleaned string
func ReadGuess(attemptNum int) string {
	// TODO: implement
	return ""
}

// DisplayTurn redraws the full game board (all guesses so far, with
// empty rows for remaining attempts) and the keyboard showing which
// letters have been used and their status.
func DisplayTurn(guesses []game.GuessResult) {
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

// DisplayError shows an inline error message (e.g., "not in word list").
func DisplayError(msg string) {
	// TODO: implement
}

// DisplayWelcome prints the game title and instructions.
func DisplayWelcome() {
	// TODO: implement
}

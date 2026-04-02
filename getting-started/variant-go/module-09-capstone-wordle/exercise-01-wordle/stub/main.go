package main

import (
	"fmt"
	"wordle/game"
	"wordle/ui"
)

// main ties the game logic and the terminal UI together.
// This is the imperative shell — it orchestrates the pure core.
//
// The game loop:
//   1. Pick a random word
//   2. Create a new game
//   3. Display welcome message
//   4. Loop:
//      a. Read a guess from the user
//      b. Submit it to the game
//      c. If error (invalid word, wrong length), show error and retry
//      d. Display the result (colored tiles)
//      e. Check if won or lost, break if so
//   5. Display win or loss message

func main() {
	target := game.RandomWord()
	g := game.New(target)

	ui.DisplayWelcome()

	for g.Status() == game.Playing {
		attempt := len(g.Guesses()) + 1
		guess := ui.ReadGuess(attempt, g.AttemptsRemaining())

		result, err := g.MakeGuess(guess)
		if err != nil {
			fmt.Printf("  %s — try again.\n", err)
			continue
		}

		ui.DisplayResult(result)
	}

	switch g.Status() {
	case game.Won:
		ui.DisplayWin(len(g.Guesses()))
	case game.Lost:
		ui.DisplayLoss(target)
	}
}

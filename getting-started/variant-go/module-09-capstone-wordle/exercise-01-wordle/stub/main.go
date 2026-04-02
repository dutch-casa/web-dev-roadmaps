package main

import (
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
//      d. Display the full board (all guesses + keyboard)
//      e. Check if won or lost, break if so
//   5. Display win or loss message

func main() {
	target := game.RandomWord()
	g := game.New(target)

	ui.DisplayWelcome()

	for g.Status() == game.Playing {
		attempt := len(g.Guesses()) + 1
		guess := ui.ReadGuess(attempt)

		_, err := g.MakeGuess(guess)
		if err != nil {
			ui.DisplayError(err.Error())
			continue
		}

		ui.DisplayTurn(g.Guesses())
	}

	switch g.Status() {
	case game.Won:
		ui.DisplayWin(len(g.Guesses()))
	case game.Lost:
		ui.DisplayLoss(target)
	}
}

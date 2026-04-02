package main

import (
	"wordle/game"
	"wordle/ui"
)

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

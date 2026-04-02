package game

import (
	"fmt"
	"strings"
)

const MaxGuesses = 6
const WordLength = 5

type LetterResult int

const (
	Absent    LetterResult = iota
	Misplaced
	Correct
)

func (r LetterResult) String() string {
	return [...]string{"absent", "misplaced", "correct"}[r]
}

type GuessResult struct {
	Guess   string
	Letters [WordLength]LetterResult
}

type GameStatus int

const (
	Playing GameStatus = iota
	Won
	Lost
)

type Game struct {
	target  string
	guesses []GuessResult
	status  GameStatus
}

func New(target string) *Game {
	if len(target) != WordLength {
		panic(fmt.Sprintf("target must be %d letters, got %d", WordLength, len(target)))
	}
	return &Game{
		target: strings.ToLower(target),
		status: Playing,
	}
}

func (g *Game) MakeGuess(guess string) (GuessResult, error) {
	if g.status != Playing {
		return GuessResult{}, fmt.Errorf("game is already over")
	}
	guess = strings.ToLower(guess)
	if len(guess) != WordLength {
		return GuessResult{}, fmt.Errorf("guess must be %d letters", WordLength)
	}
	if !IsValidWord(guess) {
		return GuessResult{}, fmt.Errorf("%q is not in the word list", guess)
	}

	letters := EvaluateGuess(g.target, guess)
	result := GuessResult{Guess: guess, Letters: letters}
	g.guesses = append(g.guesses, result)

	if guess == g.target {
		g.status = Won
	} else if len(g.guesses) >= MaxGuesses {
		g.status = Lost
	}

	return result, nil
}

// EvaluateGuess is the heart of Wordle. Two-pass algorithm:
//
// Pass 1: Mark exact matches (Correct). For each correct letter,
// "consume" that position in the target so it can't also count
// as Misplaced for a duplicate letter elsewhere.
//
// Pass 2: For each non-Correct letter, check if it exists in the
// target at a position that hasn't been consumed. If yes, Misplaced
// and consume that target position. If no, Absent.
//
// This handles duplicate letters correctly. If the target has one 'e'
// and the guess has two, only one gets marked — Correct takes priority
// over Misplaced.
func EvaluateGuess(target, guess string) [WordLength]LetterResult {
	var result [WordLength]LetterResult
	targetRunes := []rune(target)
	guessRunes := []rune(guess)

	// Track which positions in the target have been "consumed" by a match.
	consumed := [WordLength]bool{}

	// Pass 1: exact matches.
	for i := range WordLength {
		if guessRunes[i] == targetRunes[i] {
			result[i] = Correct
			consumed[i] = true
		}
	}

	// Pass 2: misplaced letters.
	for i := range WordLength {
		if result[i] == Correct {
			continue
		}
		for j := range WordLength {
			if !consumed[j] && guessRunes[i] == targetRunes[j] {
				result[i] = Misplaced
				consumed[j] = true
				break
			}
		}
		// If no match found, result[i] stays Absent (zero value).
	}

	return result
}

func (g *Game) Status() GameStatus       { return g.status }
func (g *Game) Guesses() []GuessResult   { return g.guesses }
func (g *Game) AttemptsRemaining() int   { return MaxGuesses - len(g.guesses) }

package game

const MaxGuesses = 6
const WordLength = 5

// LetterResult describes how a single letter in a guess relates
// to the target word. A guess produces a slice of five of these.
type LetterResult int

const (
	Absent    LetterResult = iota // letter is not in the target word
	Misplaced                     // letter is in the word but wrong position
	Correct                       // letter is in the correct position
)

func (r LetterResult) String() string {
	return [...]string{"absent", "misplaced", "correct"}[r]
}

// GuessResult holds the outcome of evaluating one guess.
type GuessResult struct {
	Guess   string
	Letters [WordLength]LetterResult
}

// GameStatus represents the three possible states of a game.
// A game that's still going has guesses remaining and no win.
// This is a sealed set — only these three values exist.
type GameStatus int

const (
	Playing GameStatus = iota
	Won
	Lost
)

// Game holds the complete state of a Wordle game.
// The target word is hidden — callers interact through MakeGuess.
type Game struct {
	target  string
	guesses []GuessResult
	status  GameStatus
}

// New creates a game with the given target word.
// The target must be exactly WordLength letters.
func New(target string) *Game {
	// TODO: create and return a new game.
	// Validate that target is the right length.
	return nil
}

// MakeGuess processes a guess and returns the result.
// Returns an error if:
//   - the game is already over (Won or Lost)
//   - the guess is not exactly WordLength letters
//   - the guess is not a valid word (use IsValidWord)
func (g *Game) MakeGuess(guess string) (GuessResult, error) {
	// TODO: validate the guess, evaluate it, update game state.
	return GuessResult{}, nil
}

// EvaluateGuess compares a guess against a target and returns
// the letter-by-letter result. This is the core algorithm.
//
// Rules:
//   - A letter in the correct position is Correct.
//   - A letter that exists in the target but in a different position
//     is Misplaced — BUT only if that letter hasn't already been
//     matched by a Correct or a previous Misplaced.
//   - All other letters are Absent.
//
// Example:
//   target: "apple"
//   guess:  "paper"
//   result: [Misplaced, Misplaced, Correct, Absent, Absent]
//
//   The first 'p' is Misplaced (exists in target, wrong position).
//   The 'a' is Misplaced (exists in target, wrong position).
//   The second 'p' is Correct (right position).
//   The 'e' is Absent — even though 'e' is in "apple", it's at
//   position 4 and the guess has it at position 3. But wait, there's
//   only one 'e' in the target word, and... this is the tricky part.
//   Think carefully about duplicate letters.
func EvaluateGuess(target, guess string) [WordLength]LetterResult {
	// TODO: implement the evaluation algorithm.
	// Hint: do two passes. First pass: mark all Correct letters.
	// Second pass: for each non-Correct letter, check if it's
	// Misplaced (exists in target and hasn't been "consumed" yet).
	var result [WordLength]LetterResult
	return result
}

// Status returns the current game status.
func (g *Game) Status() GameStatus { return g.status }

// Guesses returns all guesses made so far.
func (g *Game) Guesses() []GuessResult { return g.guesses }

// AttemptsRemaining returns how many guesses are left.
func (g *Game) AttemptsRemaining() int {
	return MaxGuesses - len(g.guesses)
}

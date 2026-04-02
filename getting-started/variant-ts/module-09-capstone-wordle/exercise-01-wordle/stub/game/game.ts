import { isValidWord } from "./words";

export const MAX_GUESSES = 6;
export const WORD_LENGTH = 5;

// LetterResult describes how a single letter in a guess relates
// to the target word. A guess produces an array of five of these.
export type LetterResult = "correct" | "misplaced" | "absent";

// GuessResult holds the outcome of evaluating one guess.
export type GuessResult = {
  guess: string;
  letters: LetterResult[];
};

// GameStatus represents the three possible states of a game.
// A game that's still going has guesses remaining and no win.
export type GameStatus = "playing" | "won" | "lost";

// Game holds the complete state of a Wordle game.
// The target word is hidden -- callers interact through makeGuess.
export class Game {
  #target: string;
  #guesses: GuessResult[] = [];
  #status: GameStatus = "playing";

  // The target must be exactly WORD_LENGTH letters.
  constructor(target: string) {
    // TODO: validate target length, store lowercase.
    this.#target = target;
  }

  // makeGuess processes a guess and returns the result.
  // Returns an error string if:
  //   - the game is already over (won or lost)
  //   - the guess is not exactly WORD_LENGTH letters
  //   - the guess is not a valid word (use isValidWord)
  makeGuess(guess: string): { result: GuessResult } | { error: string } {
    // TODO: validate the guess, evaluate it, update game state.
    void isValidWord;
    return { error: "not implemented" };
  }

  get status(): GameStatus { return this.#status; }
  get guesses(): GuessResult[] { return this.#guesses; }
  get attemptsRemaining(): number { return MAX_GUESSES - this.#guesses.length; }
}

// evaluateGuess compares a guess against a target and returns
// the letter-by-letter result. This is the core algorithm.
//
// Rules:
//   - A letter in the correct position is "correct".
//   - A letter that exists in the target but in a different position
//     is "misplaced" -- BUT only if that letter hasn't already been
//     matched by a "correct" or a previous "misplaced".
//   - All other letters are "absent".
//
// Examples:
//   target: "apple"
//   guess:  "paper"
//   result: ["misplaced", "misplaced", "correct", "misplaced", "absent"]
//
//   p -> misplaced (exists in "apple" at index 2, but guess has it at 0)
//   a -> misplaced (exists in "apple" at index 0, but guess has it at 1)
//   p -> correct   (index 2 matches)
//   e -> misplaced (exists in "apple" at index 4, guess has it at 3)
//   r -> absent    (not in "apple")
//
//   target: "hello"
//   guess:  "llama"
//   result: ["misplaced", "misplaced", "absent", "absent", "absent"]
//
//   First 'l' -> misplaced (consumes the 'l' at target index 2)
//   Second 'l' -> misplaced (consumes the 'l' at target index 3)
//   'a' -> absent
//   'm' -> absent
//   'a' -> absent
//
//   "hello" has two l's, so both l's get marked.
//   If the target only had one 'l', only the first would be misplaced.
//   Think carefully about duplicate letters.
export const evaluateGuess = (target: string, guess: string): LetterResult[] => {
  // TODO: implement the evaluation algorithm.
  // Hint: do two passes. First pass: mark all correct letters.
  // Second pass: for each non-correct letter, check if it's
  // misplaced (exists in target and hasn't been "consumed" yet).
  return Array.from({ length: WORD_LENGTH }, () => "absent" as LetterResult);
};

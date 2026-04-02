import { isValidWord } from "./words";

export const MAX_GUESSES = 6;
export const WORD_LENGTH = 5;

export type LetterResult = "correct" | "misplaced" | "absent";

export type GuessResult = {
  guess: string;
  letters: LetterResult[];
};

export type GameStatus = "playing" | "won" | "lost";

export type GuessOutcome =
  | { tag: "ok"; result: GuessResult }
  | { tag: "error"; message: string };

export class Game {
  #target: string;
  #guesses: GuessResult[] = [];
  #status: GameStatus = "playing";

  constructor(target: string) {
    if (target.length !== WORD_LENGTH) {
      throw new Error(`target must be ${WORD_LENGTH} letters, got ${target.length}`);
    }
    this.#target = target.toLowerCase();
  }

  makeGuess(guess: string): GuessOutcome {
    if (this.#status !== "playing") {
      return { tag: "error", message: "game is already over" };
    }
    const normalized = guess.toLowerCase();
    if (normalized.length !== WORD_LENGTH) {
      return { tag: "error", message: `guess must be ${WORD_LENGTH} letters` };
    }
    if (!isValidWord(normalized)) {
      return { tag: "error", message: `"${normalized}" is not in the word list` };
    }

    const letters = evaluateGuess(this.#target, normalized);
    const result: GuessResult = { guess: normalized, letters };
    this.#guesses.push(result);

    if (normalized === this.#target) {
      this.#status = "won";
    } else if (this.#guesses.length >= MAX_GUESSES) {
      this.#status = "lost";
    }

    return { tag: "ok", result };
  }

  get status(): GameStatus { return this.#status; }
  get guesses(): GuessResult[] { return [...this.#guesses]; }
  get attemptsRemaining(): number { return MAX_GUESSES - this.#guesses.length; }
}

// evaluateGuess is the heart of Wordle. Two-pass algorithm:
//
// Pass 1: Mark exact matches (correct). For each correct letter,
// "consume" that position in the target so it can't also count
// as misplaced for a duplicate letter elsewhere.
//
// Pass 2: For each non-correct letter, check if it exists in the
// target at a position that hasn't been consumed. If yes, misplaced
// and consume that target position. If no, absent.
//
// This handles duplicate letters correctly. If the target has one 'e'
// and the guess has two, only one gets marked -- correct takes priority
// over misplaced.
export const evaluateGuess = (target: string, guess: string): LetterResult[] => {
  const result: LetterResult[] = Array.from({ length: WORD_LENGTH }, () => "absent");
  const consumed: boolean[] = Array.from({ length: WORD_LENGTH }, () => false);

  // Pass 1: exact matches.
  for (let i = 0; i < WORD_LENGTH; i++) {
    if (guess[i] === target[i]) {
      result[i] = "correct";
      consumed[i] = true;
    }
  }

  // Pass 2: misplaced letters.
  for (let i = 0; i < WORD_LENGTH; i++) {
    if (result[i] === "correct") continue;
    for (let j = 0; j < WORD_LENGTH; j++) {
      if (!consumed[j] && guess[i] === target[j]) {
        result[i] = "misplaced";
        consumed[j] = true;
        break;
      }
    }
    // If no match found, result[i] stays "absent".
  }

  return result;
};

import { type LetterResult, type GuessResult, MAX_GUESSES, WORD_LENGTH } from "../game/game";

// ANSI color codes for terminal output.
const RESET = "\x1b[0m";
const BOLD = "\x1b[1m";
const DIM = "\x1b[2m";
const GREEN = "\x1b[48;5;34;97m";   // green bg, bright white text
const YELLOW = "\x1b[48;5;178;97m"; // amber bg, bright white text
const GRAY = "\x1b[48;5;240;97m";   // gray bg, bright white text
const DIM_BG = "\x1b[48;5;236;37m"; // dark bg, dim text (empty tiles)

// colorForResult returns the ANSI color code for a letter result.
const colorForResult = (r: LetterResult): string => {
  switch (r) {
    case "correct": return GREEN;
    case "misplaced": return YELLOW;
    case "absent": return GRAY;
  }
};

// readGuess reads a guess from stdin.
// It should:
//   - Print a prompt showing the attempt number (e.g., "Guess 3/6: ")
//   - Read a line from stdin (use Bun's prompt())
//   - Trim whitespace
//   - Convert to lowercase
//   - Return the cleaned string
export const readGuess = (attemptNum: number): string => {
  // TODO: implement
  void attemptNum;
  return "";
};

// displayTurn redraws the full game board (all guesses so far, with
// empty rows for remaining attempts) and the keyboard showing which
// letters have been used and their status.
export const displayTurn = (guesses: GuessResult[]): void => {
  // TODO: implement
  // For each letter in the guess, print it with the appropriate
  // background color using the ANSI codes above.
  void guesses;
};

// displayWin prints a congratulatory message.
export const displayWin = (attempts: number): void => {
  // TODO: implement
  void attempts;
};

// displayLoss prints the target word after a loss.
export const displayLoss = (target: string): void => {
  // TODO: implement
  void target;
};

// displayError shows an inline error message (e.g., "not in word list").
export const displayError = (msg: string): void => {
  // TODO: implement
  void msg;
};

// displayWelcome prints the game title and instructions.
export const displayWelcome = (): void => {
  // TODO: implement
};

// Suppress unused import warnings for values the student needs.
void RESET;
void BOLD;
void DIM;
void DIM_BG;
void colorForResult;
void MAX_GUESSES;
void WORD_LENGTH;

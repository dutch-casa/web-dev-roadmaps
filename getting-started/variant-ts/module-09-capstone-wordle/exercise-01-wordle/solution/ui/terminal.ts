import { type LetterResult, type GuessResult, MAX_GUESSES, WORD_LENGTH } from "../game/game";

const RESET = "\x1b[0m";
const BOLD = "\x1b[1m";
const DIM = "\x1b[2m";
const GREEN = "\x1b[48;5;34;97m";   // green bg, bright white
const YELLOW = "\x1b[48;5;178;97m"; // amber bg, bright white
const GRAY = "\x1b[48;5;240;97m";   // gray bg, bright white
const DIM_BG = "\x1b[48;5;236;37m"; // dark bg, dim text (empty tiles)

// Tracks the best-known state per letter across all guesses.
// correct > misplaced > absent. Unseen letters aren't in the map.
const RESULT_PRIORITY = { absent: 0, misplaced: 1, correct: 2 } as const;
const keyboard = new Map<string, LetterResult>();

const colorFor = (r: LetterResult): string => {
  switch (r) {
    case "correct": return GREEN;
    case "misplaced": return YELLOW;
    case "absent": return GRAY;
  }
};

const updateKeyboard = (result: GuessResult): void => {
  for (let i = 0; i < result.guess.length; i++) {
    const ch = result.guess[i];
    const newResult = result.letters[i];
    const prev = keyboard.get(ch);
    if (!prev || RESULT_PRIORITY[newResult] > RESULT_PRIORITY[prev]) {
      keyboard.set(ch, newResult);
    }
  }
};

// --- Tile rendering ---

const tile = (ch: string, r: LetterResult): string =>
  `${colorFor(r)} ${ch.toUpperCase()} ${RESET}`;

const emptyTile = (): string =>
  `${DIM_BG}   ${RESET}`;

const renderGuessRow = (result: GuessResult): string => {
  const tiles = Array.from(result.guess).map((ch, i) => tile(ch, result.letters[i]));
  return "    " + tiles.join(" ");
};

const renderEmptyRow = (): string => {
  const tiles = Array.from({ length: WORD_LENGTH }, () => emptyTile());
  return "    " + tiles.join(" ");
};

// --- Board ---

const printBoard = (guesses: GuessResult[]): void => {
  console.log();
  for (let i = 0; i < MAX_GUESSES; i++) {
    if (i < guesses.length) {
      console.log(renderGuessRow(guesses[i]));
    } else {
      console.log(renderEmptyRow());
    }
  }
};

// --- Keyboard ---

const KEY_ROWS = ["qwertyuiop", "asdfghjkl", "zxcvbnm"] as const;

const printKeyboard = (): void => {
  console.log();
  for (let i = 0; i < KEY_ROWS.length; i++) {
    const pad = " ".repeat((i * 2) + 4);
    const keys = Array.from(KEY_ROWS[i]).map((ch) => {
      const result = keyboard.get(ch);
      const upper = ch.toUpperCase();
      if (!result) return ` ${upper} `;
      return `${colorFor(result)} ${upper} ${RESET}`;
    });
    console.log(pad + keys.join(""));
  }
  console.log();
};

// --- Public API ---

export const displayWelcome = (): void => {
  console.log();
  console.log(`    ${BOLD} W  O  R  D  L  E ${RESET}`);
  console.log();
  console.log("    Guess the 5-letter word in 6 tries.");
  console.log();
  console.log(`    ${GREEN} E ${RESET} correct   ${YELLOW} E ${RESET} wrong spot   ${GRAY} E ${RESET} not in word`);
};

// displayTurn redraws the full board and keyboard after every guess.
export const displayTurn = (guesses: GuessResult[]): void => {
  for (const g of guesses) {
    updateKeyboard(g);
  }
  printBoard(guesses);
  printKeyboard();
};

export const readGuess = (attemptNum: number): string => {
  const raw = prompt(`    Guess ${attemptNum}/${MAX_GUESSES}:`) ?? "";
  return raw.trim().toLowerCase();
};

export const displayError = (msg: string): void => {
  console.log(`    ${DIM}-- ${msg}${RESET}\n`);
};

export const displayWin = (attempts: number): void => {
  const msgs = ["", "Genius!", "Magnificent!", "Impressive!", "Solid!", "Good job!", "Phew!"];
  console.log(`\n    ${BOLD}${msgs[attempts]}${RESET} Got it in ${attempts}.\n`);
};

export const displayLoss = (target: string): void => {
  console.log(`\n    The word was ${BOLD}${target.toUpperCase()}${RESET}. Better luck next time.\n`);
};

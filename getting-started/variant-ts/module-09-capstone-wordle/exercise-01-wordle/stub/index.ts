// main ties the game logic and the terminal UI together.
// This is the imperative shell -- it orchestrates the pure core.
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

import { Game } from "./game/game";
import { randomWord } from "./game/words";
import {
  displayWelcome,
  readGuess,
  displayTurn,
  displayError,
  displayWin,
  displayLoss,
} from "./ui/terminal";

const target = randomWord();
const game = new Game(target);

displayWelcome();

while (game.status === "playing") {
  const attempt = game.guesses.length + 1;
  const guess = readGuess(attempt);

  const outcome = game.makeGuess(guess);
  if ("error" in outcome) {
    displayError(outcome.error);
    continue;
  }

  displayTurn(game.guesses);
}

if (game.status === "won") {
  displayWin(game.guesses.length);
} else {
  displayLoss(target);
}

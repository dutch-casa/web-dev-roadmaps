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
  if (outcome.tag === "error") {
    displayError(outcome.message);
    continue;
  }

  displayTurn(game.guesses);
}

if (game.status === "won") {
  displayWin(game.guesses.length);
} else {
  displayLoss(target);
}

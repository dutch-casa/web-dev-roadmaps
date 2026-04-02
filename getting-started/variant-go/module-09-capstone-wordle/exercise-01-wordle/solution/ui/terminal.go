package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"wordle/game"
)

const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	dim    = "\033[2m"
	green  = "\033[48;5;34;97m"  // green bg, bright white
	yellow = "\033[48;5;178;97m" // amber bg, bright white
	gray   = "\033[48;5;240;97m" // gray bg, bright white
	dimBg  = "\033[48;5;236;37m" // dark bg, dim text (empty tiles)
)

var reader = bufio.NewReader(os.Stdin)

// Tracks the best-known state per letter across all guesses.
// Correct > Misplaced > Absent. Unseen letters aren't in the map.
var keyboard = map[rune]int{}

func colorFor(r game.LetterResult) string {
	switch r {
	case game.Correct:
		return green
	case game.Misplaced:
		return yellow
	default:
		return gray
	}
}

func updateKeyboard(result game.GuessResult) {
	for i, ch := range result.Guess {
		v := int(result.Letters[i])
		if prev, seen := keyboard[ch]; !seen || v > prev {
			keyboard[ch] = v
		}
	}
}

// --- Tile rendering ---

func tile(ch rune, r game.LetterResult) string {
	return fmt.Sprintf("%s %c %s", colorFor(r), unicode.ToUpper(ch), reset)
}

func emptyTile() string {
	return fmt.Sprintf("%s   %s", dimBg, reset)
}

func renderGuessRow(result game.GuessResult) string {
	tiles := make([]string, game.WordLength)
	for i, ch := range result.Guess {
		tiles[i] = tile(ch, result.Letters[i])
	}
	return "    " + strings.Join(tiles, " ")
}

func renderEmptyRow() string {
	tiles := make([]string, game.WordLength)
	for i := range tiles {
		tiles[i] = emptyTile()
	}
	return "    " + strings.Join(tiles, " ")
}

// --- Board ---

func printBoard(guesses []game.GuessResult) {
	fmt.Println()
	for i := range game.MaxGuesses {
		if i < len(guesses) {
			fmt.Println(renderGuessRow(guesses[i]))
		} else {
			fmt.Println(renderEmptyRow())
		}
	}
}

// --- Keyboard ---

var keyRows = [3]string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}

func printKeyboard() {
	fmt.Println()
	for i, row := range keyRows {
		pad := strings.Repeat(" ", (i*2)+4)
		var keys []string
		for _, ch := range row {
			v, seen := keyboard[ch]
			upper := unicode.ToUpper(ch)
			if !seen {
				keys = append(keys, fmt.Sprintf(" %c ", upper))
			} else {
				keys = append(keys, fmt.Sprintf("%s %c %s", colorFor(game.LetterResult(v)), upper, reset))
			}
		}
		fmt.Println(pad + strings.Join(keys, ""))
	}
	fmt.Println()
}

// --- Public API ---

func DisplayWelcome() {
	fmt.Println()
	fmt.Printf("    %s W  O  R  D  L  E %s\n", bold, reset)
	fmt.Println()
	fmt.Println("    Guess the 5-letter word in 6 tries.")
	fmt.Println()
	fmt.Printf("    %s E %s correct   ", green, reset)
	fmt.Printf("%s E %s wrong spot   ", yellow, reset)
	fmt.Printf("%s E %s not in word\n", gray, reset)
}

// DisplayTurn redraws the full board and keyboard after every guess.
func DisplayTurn(guesses []game.GuessResult) {
	for _, g := range guesses {
		updateKeyboard(g)
	}
	printBoard(guesses)
	printKeyboard()
}

func ReadGuess(attemptNum int) string {
	fmt.Printf("    Guess %d/%d: ", attemptNum, game.MaxGuesses)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(line))
}

func DisplayError(msg string) {
	fmt.Printf("    %s-- %s%s\n\n", dim, msg, reset)
}

func DisplayWin(attempts int) {
	msgs := [7]string{"", "Genius!", "Magnificent!", "Impressive!", "Solid!", "Good job!", "Phew!"}
	fmt.Printf("\n    %s%s%s Got it in %d.\n\n", bold, msgs[attempts], reset, attempts)
}

func DisplayLoss(target string) {
	fmt.Printf("\n    The word was %s%s%s. Better luck next time.\n\n",
		bold, strings.ToUpper(target), reset)
}

package game

// A curated list of common five-letter words.
// Kept short on purpose, you can add more if you would like.
var wordList = []string{
	"apple", "baker", "candy", "dance", "eagle",
	"flame", "grape", "house", "index", "joker",
	"knife", "lemon", "mango", "nerve", "ocean",
	"piano", "queen", "river", "snake", "tiger",
	"ultra", "voice", "whale", "youth", "zebra",
	"amber", "blaze", "charm", "draft", "elder",
	"frost", "globe", "happy", "ivory", "jolly",
	"karma", "lucky", "maple", "noble", "olive",
	"pearl", "quiet", "rapid", "sonic", "trace",
	"unity", "vigor", "witty", "xenon", "yacht",
}

// RandomWord returns a random word from the word list.
func RandomWord() string {
	// TODO: pick a random word from wordList.
	return ""
}

// IsValidWord checks whether a guess is in the word list.
func IsValidWord(word string) bool {
	// TODO: check if the word exists in wordList.
	return false
}

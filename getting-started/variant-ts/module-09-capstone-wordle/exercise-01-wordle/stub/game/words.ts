// A curated list of common five-letter words.
// Kept short on purpose -- you can add more if you would like.
const wordList = [
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
] as const;

// randomWord returns a random word from the word list.
export const randomWord = (): string => {
  // TODO: pick a random word from wordList.
  return "";
};

// isValidWord checks whether a guess is in the word list.
export const isValidWord = (word: string): boolean => {
  // TODO: check if the word exists in wordList.
  void word;
  return false;
};

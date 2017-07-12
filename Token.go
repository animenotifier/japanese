package japanese

// Token represents a single token in a sentence.
type Token struct {
	Original string
	Hiragana string
	Katakana string
	Romaji   string
	Furigana bool
}

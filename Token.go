package japanese

import "github.com/gojp/kana"

// Token represents a single token in a sentence.
type Token struct {
	Original string
	Hiragana string
	Katakana string
	Romaji   string
}

// NeedsFurigana tells you whether furigana are needed or not.
func (token *Token) NeedsFurigana() bool {
	return !kana.IsHiragana(token.Original) && !kana.IsKatakana(token.Original) && !kana.IsLatin(token.Original)
}

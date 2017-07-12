package tokenizer

import (
	"github.com/animenotifier/japanese"
	"github.com/gojp/kana"
)

// NeedsFurigana tells you whether furigana are needed or not.
func NeedsFurigana(token *japanese.Token) bool {
	return !kana.IsHiragana(token.Original) && !kana.IsKatakana(token.Original) && !kana.IsLatin(token.Original)
}

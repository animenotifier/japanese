package tokenizer

import (
	"strings"

	"github.com/animenotifier/japanese"
	"github.com/gojp/kana"
	kagome "github.com/ikawaha/kagome/tokenizer"
)

var globalTokenizer = kagome.New()

// Tokenize splits the given sentence into tokens.
func Tokenize(sentence string) []*japanese.Token {
	var tokens []*japanese.Token

	for _, token := range globalTokenizer.Tokenize(sentence) {
		// Ignore start and end of sentence tokens
		if token.Class == kagome.DUMMY {
			continue
		}

		features := token.Features()
		hiragana := ""
		katakana := ""
		romaji := ""

		if len(features) >= 9 {
			katakana = features[8]
			romaji = kana.KanaToRomaji(katakana)
			hiragana = kana.RomajiToHiragana(romaji)
		}

		// Add some custom fixes
		romaji = strings.Replace(romaji, "hu", "fu", -1)

		// Create token
		token := &japanese.Token{
			Original: token.Surface,
			Hiragana: hiragana,
			Katakana: katakana,
			Romaji:   romaji,
		}
		token.Furigana = NeedsFurigana(token)

		// Add to token list
		tokens = append(tokens, token)
	}

	return tokens
}

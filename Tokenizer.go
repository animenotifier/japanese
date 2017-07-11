package japanese

import (
	"strings"

	"github.com/gojp/kana"
	"github.com/ikawaha/kagome/tokenizer"
)

var globalTokenizer = tokenizer.New()

// Tokenize splits the given sentence into tokens.
func Tokenize(japanese string) []*Token {
	var tokens []*Token

	for _, token := range globalTokenizer.Tokenize(japanese) {
		// Ignore start and end of sentence tokens
		if token.Class == tokenizer.DUMMY {
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

		// Add to token list
		tokens = append(tokens, &Token{
			Original: token.Surface,
			Hiragana: hiragana,
			Katakana: katakana,
			Romaji:   romaji,
		})
	}

	return tokens
}

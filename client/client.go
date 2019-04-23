package client

import (
	"github.com/animenotifier/japanese"
	"github.com/parnurzeal/gorequest"
)

// Tokenizer using the HTTP API.
type Tokenizer struct {
	Endpoint string
}

// Tokenize splits the given sentence into tokens by querying the HTTP server.
func (tokenizer *Tokenizer) Tokenize(sentence string) []*japanese.Token {
	var result []*japanese.Token
	_, _, errs := gorequest.New().Get(tokenizer.Endpoint + sentence).EndStruct(&result)

	if len(errs) > 0 {
		return []*japanese.Token{
			{
				Original: sentence,
				Furigana: false,
			},
		}
	}

	return result
}

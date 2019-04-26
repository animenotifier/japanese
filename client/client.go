package client

import (
	"github.com/aerogo/http/client"
	"github.com/animenotifier/japanese"
)

// Tokenizer using the HTTP API.
type Tokenizer struct {
	Endpoint string
}

// Tokenize splits the given sentence into tokens by querying the HTTP server.
func (tokenizer *Tokenizer) Tokenize(sentence string) []*japanese.Token {
	var result []*japanese.Token

	// Send HTTP request and capture response in "result"
	_, err := client.Get(tokenizer.Endpoint + sentence).EndStruct(&result)

	if err != nil {
		return []*japanese.Token{
			{
				Original: sentence,
				Furigana: false,
			},
		}
	}

	return result
}

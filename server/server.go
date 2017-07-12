package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/animenotifier/japanese/tokenizer"
)

func japaneseTokenizer(w http.ResponseWriter, req *http.Request) {
	data := tokenizer.Tokenize(strings.TrimPrefix(req.URL.Path, "/"))
	buffer, err := json.Marshal(data)

	if err != nil {
		io.WriteString(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(buffer)
}

var port = "8080"

func init() {
	flag.StringVar(&port, "port", "", "Port the HTTP server should listen on")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", japaneseTokenizer)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

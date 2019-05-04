package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/animenotifier/japanese/tokenizer"
	jsoniter "github.com/json-iterator/go"
)

var port string

func init() {
	flag.StringVar(&port, "port", "6000", "Port the HTTP server should listen on")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", onRequest)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}

// onRequest will tokenize everything after the slash in the requested path.
func onRequest(response http.ResponseWriter, request *http.Request) {
	data := tokenizer.Tokenize(strings.TrimPrefix(request.URL.Path, "/"))
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	encoder := jsoniter.NewEncoder(response)
	err := encoder.Encode(data)

	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		return
	}
}

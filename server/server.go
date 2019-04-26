package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/animenotifier/japanese/tokenizer"
	jsoniter "github.com/json-iterator/go"
)

var port = "6000"

func init() {
	flag.StringVar(&port, "port", "", "Port the HTTP server should listen on")
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
	buffer, err := jsoniter.Marshal(data)

	if err != nil {
		io.WriteString(response, err.Error())
		io.WriteString(os.Stderr, err.Error())
		return
	}

	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	response.Write(buffer)
}

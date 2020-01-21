// +build wasm

package main

import (
	"log"
	"strings"
	"syscall/js"

	"github.com/martinlindhe/gohash"
)

func main() {
	
}

//go:export encode
func encode(enc string, text string) []byte {
	log.Println("encode run in wasm")
	r := strings.NewReader(text)
	encodings := strings.Split(enc, "+")
	res, err := gohash.RecodeInput(encodings, r, false, false)
	if err != nil {
		log.Fatal("error:", err)
	}
	return res
}

//go:export update
func update() {
	log.Println("update run in wasm")
	document := js.Global().Get("document")
	text := document.Call("getElementById", "text").Get("value").String()
	encodings := document.Call("getElementById", "encoding").Get("value").String()
	result := encode(encodings, text)
	document.Call("getElementById", "result").Set("value", string(result))
}

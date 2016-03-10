package main

import (
	"fmt"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/martinlindhe/gohash"
)

var (
	hash = kingpin.Flag("hash", "Hash to crack, in hex string").Required().String()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	totalTime := time.Now()

	expectedHash := gohash.HexStringToBytes(*hash)

	chk := gohash.FindMatchingOnionURL(expectedHash)
	fmt.Println(chk)

	fmt.Printf("total time: ")
	fmt.Println(time.Since(totalTime))
}

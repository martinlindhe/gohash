package main

import (
	"fmt"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/martinlindhe/gohash"
)

var (
	hash   = kingpin.Flag("hash", "Hash to crack, in hex string").Required().String()
	random = kingpin.Flag("random", "Random mutation mode.").Bool()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	totalTime := time.Now()

	expectedHash := gohash.HexStringToBytes(*hash)

	res := ""
	if *random {
		res = gohash.FindMatchingOnionURLByRandom(expectedHash)
	} else {
		res = gohash.FindMatchingOnionURL(expectedHash)
	}

	fmt.Println("result: ", res)
	fmt.Println("total time: ", time.Since(totalTime))
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/martinlindhe/gohash"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	hash      = kingpin.Flag("hash", "Hash to crack, in hex string").Required().String()
	algo      = kingpin.Flag("algo", "Hash algorithm to use. sha512, sha256 etc").Required().String()
	random    = kingpin.Flag("random", "Random mutation mode.").Bool()
	startTime = time.Now()
	result    = ""
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	expectedHash := gohash.HexStringToBytes(*hash)

	if len(expectedHash) != 64 {
		fmt.Println("hash is wrong size, is ", len(expectedHash), ", should be 64 byte")
		os.Exit(1)
	}

	// XXX todo make use of algo

	// catch ctrl-c interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it

			// XXX * when exited, show number of tries and time ran, and tries/sec
			fmt.Println("total time: ", time.Since(startTime))
			os.Exit(0)
		}
	}()

	if *random {
		result = gohash.FindMatchingOnionURLByRandom(expectedHash)
	} else {
		result = gohash.FindMatchingOnionURL(expectedHash)
	}

	fmt.Println("result: ", result)
}

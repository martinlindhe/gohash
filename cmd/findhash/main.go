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
	hash        = kingpin.Flag("hash", "Hash to crack, in hex string").Required().String()
	algo        = kingpin.Flag("algo", "Hash algorithm to use. sha1, sha512 etc").Required().String()
	allowedKeys = kingpin.Flag("allowed", "Allowed keys to use.").Required().String()
	minLength   = kingpin.Flag("min-length", "Minimum length.").Required().Int()
	suffix      = kingpin.Flag("suffix", "Suffix (optional).").String()
	random      = kingpin.Flag("random", "Random mutation mode.").Bool()
	startTime   = time.Now()
	result      = ""
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	hasher := gohash.NewHasher()

	hasher.Algo(*algo)
	hasher.AllowedKeys(*allowedKeys)
	hasher.Suffix(*suffix)
	hasher.ExpectedHash(*hash)
	hasher.Length(*minLength)

	// catch ctrl-c interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it

			// XXX * when exited, show number of tries and time ran, and tries/sec
			fmt.Println("")
			fmt.Println("total time: ", time.Since(startTime))
			os.Exit(0)
		}
	}()

	var err error
	if *random {
		result, err = hasher.FindRandom()
	} else {
		result, err = hasher.FindSequential()
	}

	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

	fmt.Println("result: ", result)
}

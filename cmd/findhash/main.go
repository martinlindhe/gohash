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
	algo        = kingpin.Flag("algo", "Hash algorithm to use. sha1, sha512 etc").String()
	allowedKeys = kingpin.Flag("allowed", "Allowed keys to use.").String()
	minLength   = kingpin.Flag("min-length", "Minimum length.").Int()
	prefix      = kingpin.Flag("prefix", "Prefix.").String()
	suffix      = kingpin.Flag("suffix", "Suffix.").String()
	random      = kingpin.Flag("random", "Random mutation mode.").Bool()
	reverse     = kingpin.Flag("reverse", "Reverse order (if not random mode)").Bool()
	dictionary  = kingpin.Flag("dictionary", "Dictionary file.").String()
	startTime   = time.Now()
	result      = ""
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	rand.Seed(startTime.UTC().UnixNano())

	// catch ctrl-c interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it

			// XXX when exited, show number of tries and time ran, and tries/sec
			fmt.Println("")
			fmt.Println("total time: ", time.Since(startTime))
			os.Exit(0)
		}
	}()

	if *dictionary != "" {
		if *algo != "" {
			fmt.Println("ERROR dictionary and algo dont mix")
			os.Exit(1)
		}
		if *minLength != 0 {
			fmt.Println("ERROR dictionary and minLength dont mix")
			os.Exit(1)
		}
		if *reverse {
			fmt.Println("ERROR dictionary and reverse dont mix")
			os.Exit(1)
		}

		runDictionary()

	} else {

		if *algo == "" {
			fmt.Println("ERROR algo must be set")
			os.Exit(1)
		}
		if *allowedKeys == "" {
			fmt.Println("ERROR allowed must be set")
			os.Exit(1)
		}
		if *minLength == 0 {
			fmt.Println("ERROR minLength must be set")
			os.Exit(1)
		}

		runHasher()
	}
}

func runDictionary() {

	dict, err := gohash.NewDictionary(*dictionary)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	dict.Prefix(*prefix)
	dict.Suffix(*suffix)
	dict.ExpectedHash(*hash)

	result, algo, err := dict.Find()

	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	if result == "" {
		fmt.Printf("no match")
	} else {
		fmt.Println("result: ", result, ", algo: ", algo)
	}
}

func runHasher() {

	hasher := gohash.NewHasher()
	hasher.Algo(*algo)
	hasher.AllowedKeys(*allowedKeys)
	hasher.Prefix(*prefix)
	hasher.Suffix(*suffix)
	hasher.ExpectedHash(*hash)
	hasher.Length(*minLength)
	hasher.Reverse(*reverse)

	var err error
	if *random {
		result, err = hasher.FindRandom()
	} else {
		result, err = hasher.FindSequential()
	}

	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	fmt.Println("result: ", result)
}

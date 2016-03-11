package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/martinlindhe/gohash"

	"golang.org/x/crypto/ssh/terminal"
)

import "io/ioutil"

var (
	algo = kingpin.Flag("algo", "Hash algorithm to use. sha1, sha512 etc").Required().String()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	var b []byte

	fileName := ""

	if !terminal.IsTerminal(0) {
		b, _ = ioutil.ReadAll(os.Stdin)
		fileName = "-"
	} else {
		fmt.Println("no piped data")

		// TODO read file if filename is provided
		os.Exit(1)
	}

	calc := gohash.NewCalculator(b)

	hash := calc.Sum(*algo)
	if hash == nil {
		fmt.Println("ERROR: unknown algo", *algo)
		os.Exit(1)
	}

	fmt.Println(hex.EncodeToString(*hash), fileName)
}

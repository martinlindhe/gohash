package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/martinlindhe/gohash"

	"golang.org/x/crypto/ssh/terminal"
)

import "io/ioutil"

var (
	algo       = kingpin.Flag("algo", "Hash algorithm to use. sha1, sha512 etc").String()
	fileName   = kingpin.Flag("file", "File to read").Short('i').String()
	listHashes = kingpin.Flag("list-hashes", "List available hash algorithms").Bool()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	if *listHashes {
		fmt.Println("Available hashes")
		fmt.Println("")
		fmt.Println(gohash.AvailableHashes())
		os.Exit(0)
	}

	if *algo == "" {
		fmt.Println("error: required flag --algo not provided, try --help")
		os.Exit(1)
	}

	var b []byte
	var err error

	if !terminal.IsTerminal(0) {
		b, _ = ioutil.ReadAll(os.Stdin)
		*fileName = "-"
	} else {

		if *fileName == "" {
			fmt.Println("error: no piped data and no --file provided")
			os.Exit(1)
		}

		b, err = readBinaryFile(*fileName)
		if err != nil {
			panic(err)
		}
	}

	calc := gohash.NewCalculator(b)

	hash := calc.Sum(*algo)
	if hash == nil {
		fmt.Println("error: --algo", *algo, "is unknown")
		os.Exit(1)
	}

	fmt.Printf("%s  %s", hex.EncodeToString(*hash), *fileName)
	fmt.Println("")
}

func readBinaryFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	size := stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes, err
}

package main

import (
	"fmt"
	"os"

	"github.com/aybabtme/color/brush"
	"github.com/martinlindhe/gohash"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	fileName      = kingpin.Flag("file", "Input file to read.").Short('i').String()
	algo          = kingpin.Arg("algo", "Hash algorithm to use.").String()
	listAlgos     = kingpin.Flag("list-algos", "List available hash algorithms.").Short('A').Bool()
	encoding      = kingpin.Flag("encoding", "Output encoding.").Short('e').Default("hex").String()
	listEncodings = kingpin.Flag("list-encodings", "List available encodings.").Short('E').Bool()
	skipNewline   = kingpin.Flag("skip-newline", "Don't output newline.").Short('n').Bool()
	skipFilename  = kingpin.Flag("skip-filename", "Don't output filename.").Bool()
	reverseBytes  = kingpin.Flag("reverse-bytes", "Reverse byte order of displayed hex value.").Bool()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	if *listAlgos {
		fmt.Println(gohash.AvailableHashes())
		os.Exit(0)
	}

	if *listEncodings {
		fmt.Println(gohash.AvailableEncodings())
		os.Exit(0)
	}

	if *algo == "" {
		fmt.Println("error: required algorithm not provided, try --help")
		os.Exit(1)
	}

	appInputData, err := gohash.ReadPipeOrFile(*fileName)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	calc := gohash.NewCalculator(appInputData.Data)
	hash, err := calc.Sum(*algo)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	if *reverseBytes {
		rev := []byte{}
		for i := len(hash) - 1; i >= 0; i-- {
			rev = append(rev, hash[i])
		}
		hash = rev
	}

	coder := gohash.NewCoder(*encoding)
	encodedHash, err := coder.Encode(hash)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	fmt.Printf("%s", brush.Yellow(encodedHash))
	if !*skipFilename {
		if appInputData.IsPipe {
			*fileName = "-"
		}
		fmt.Printf("  %s", brush.White(*fileName))
	}
	if !*skipNewline {
		fmt.Println()
	}
}

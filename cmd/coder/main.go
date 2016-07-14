package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/martinlindhe/gohash"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	encoding      = kingpin.Arg("encoding", "Output encoding.").String()
	listEncodings = kingpin.Flag("list-encodings", "List available encodings.").Short('E').Bool()
	fileName      = kingpin.Arg("file", "Input file to read.").String()
	encode        = kingpin.Flag("encode", "Encode input (default).").Short('e').Bool()
	decode        = kingpin.Flag("decode", "Decode input.").Short('d').Bool()
	outFileName   = kingpin.Flag("output", "Write output to file.").Short('o').String()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	if *listEncodings {
		fmt.Println(gohash.AvailableEncodings())
		os.Exit(0)
	}

	encodings := strings.Split(*encoding, "+")

	if len(encodings) == 0 {
		fmt.Println("error: required argument 'encoding' not provided, try --help")
		os.Exit(1)
	}

	if *decode && *encode {
		fmt.Println("error: --decode and --encode don't mix")
		os.Exit(1)
	}

	appInputData, err := gohash.ReadPipeOrFile(*fileName)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	res, err := gohash.RecodeInput(encodings, appInputData.Data, *decode)
	if err != nil {
		fmt.Println("error:", err)
		if len(res) > 0 {
			fmt.Println(string(res))
		}
		os.Exit(1)
	}

	if *outFileName != "" {
		f, err := os.Create(*outFileName)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		defer f.Close()
		_, err = f.Write(res)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println(string(res))
	}
}

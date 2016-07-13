package main

import (
	"fmt"
	"os"

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

	if *encoding == "" {
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

	coder := gohash.NewCoder(*encoding)

	res := ""
	if *decode {
		var decoded []byte
		decoded, err = coder.Decode(string(appInputData.Data))
		res = string(decoded)
	} else {
		res, err = coder.Encode(appInputData.Data)
	}
	if err != nil {
		fmt.Println("error:", err)
		if res != "" {
			fmt.Println(res)
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
		_, err = f.WriteString(res)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	} else {
		fmt.Print(res)
	}
}

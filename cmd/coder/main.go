package main

import (
	"fmt"
	"os"

	"github.com/martinlindhe/gohash"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	encoding      = kingpin.Arg("encoding", "Output encoding. hex is default").Required().String()
	listEncodings = kingpin.Flag("list-encodings", "List available encodings").Short('E').Bool()

	fileName = kingpin.Arg("file", "File to read").String()

	decode = kingpin.Flag("decode", "Decode").Short('d').Bool()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	if *listEncodings {
		fmt.Println(gohash.AvailableEncodings())
		os.Exit(0)
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
		fmt.Println(res)
		os.Exit(1)
	}

	fmt.Print(res)
}

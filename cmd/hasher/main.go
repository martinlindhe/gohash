package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/alecthomas/kingpin"
	"github.com/fatih/color"
	"github.com/martinlindhe/gohash"
)

var (
	fileName      = kingpin.Flag("file", "Input file to read (optional).").Short('i').String()
	algo          = kingpin.Arg("algo", "Hash algorithm to use.").String()
	listAlgos     = kingpin.Flag("list-algos", "List available hash algorithms.").Short('A').Bool()
	encoding      = kingpin.Flag("encoding", "Output encoding.").Short('e').Default("hex").String()
	listEncodings = kingpin.Flag("list-encodings", "List available encodings.").Short('E').Bool()
	skipNewline   = kingpin.Flag("skip-newline", "Don't output newline.").Short('n').Bool()
	skipFilename  = kingpin.Flag("skip-filename", "Don't output filename.").Bool()
	noColors      = kingpin.Flag("no-colors", "Don't output colors.").Bool()
	reverseBytes  = kingpin.Flag("reverse-bytes", "Reverse byte order of displayed hex value.").Bool()
	debugAllocs   = kingpin.Flag("debug-allocs", "Debugging: print memory allocations at end of execution.").Bool()
	bsdSyntax     = kingpin.Flag("bsd", "Output result in BSD syntax.").Bool()

	white  = color.New(color.FgWhite).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
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

	r, err := gohash.ReadPipeOrFile(*fileName)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	calc := gohash.NewCalculator(r.Reader)
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
	encodedHash, err := coder.Encode(bytes.NewReader(hash))
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	if r.IsPipe {
		*fileName = "-"
	}
	if *bsdSyntax {
		fmt.Print(strings.ToUpper(*algo))
		if !*skipFilename {
			fmt.Print(" (")
			if !*noColors {
				fmt.Print(white(*fileName))
			} else {
				fmt.Print(*fileName)
			}
			fmt.Print(")")
		}
		fmt.Print(" = ")
		if !*noColors {
			fmt.Print(yellow(encodedHash))
		} else {
			fmt.Print(string(encodedHash))
		}
	} else {
		if !*noColors {
			fmt.Print(yellow(encodedHash))
		} else {
			fmt.Print(string(encodedHash))
		}
		if !*skipFilename {
			if r.IsPipe {
				*fileName = "-"
			}
			fmt.Print("  ")
			if !*noColors {
				fmt.Print(white(*fileName))
			} else {
				fmt.Print(*fileName)
			}
		}
	}

	if !*skipNewline {
		fmt.Println()
	}

	if *debugAllocs {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("\nAlloc = %v\nTotalAlloc = %v\nSys = %v\nNumGC = %v\n\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
	}
}

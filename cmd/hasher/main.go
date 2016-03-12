package main

import (
	"bufio"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/martinlindhe/bubblebabble"
	"github.com/martinlindhe/gohash"
	"github.com/tilinna/z85"

	"golang.org/x/crypto/ssh/terminal"
)

import "io/ioutil"

var (
	algo       = kingpin.Flag("algo", "Hash algorithm to use. sha1, sha512 etc").Short('a').String()
	fileName   = kingpin.Flag("file", "File to read").Short('i').String()
	listHashes = kingpin.Flag("list-hashes", "List available hash algorithms").Bool()
	encoding   = kingpin.Flag("encoding", "Output encoding: hex (default), hexup, base32, base64, bb, bin, oct, dec, z85").Short('e').String()
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

	if *encoding == "" {
		*encoding = "hex"
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

	encodedHash := ""
	switch *encoding {
	case "base16":
		fallthrough
	case "hex":
		encodedHash = hex.EncodeToString(*hash)

	case "base32":
		encodedHash = base32.StdEncoding.EncodeToString(*hash)

	case "base64":
		encodedHash = base64.StdEncoding.EncodeToString(*hash)

	case "bb":
		fallthrough
	case "bubblebabble":
		encodedHash = bubblebabble.EncodeToString(*hash)

	case "bin":
		fallthrough
	case "binary":
		encodedHash = toBinaryString(*hash, " ")

	case "dec":
		fallthrough
	case "decimal":
		encodedHash = toDecimalString(*hash, " ")

	case "hexup":
		encodedHash = hex.EncodeToString(*hash)
		encodedHash = strings.ToUpper(encodedHash)

	case "oct":
		fallthrough
	case "octal":
		encodedHash = toOctalString(*hash, " ")

	case "z85":
		b85 := make([]byte, z85.EncodedLen(len(*hash)))
		z85.Encode(b85, *hash)
		encodedHash = string(b85)

	default:
		fmt.Println("error: unknown --encoding", *encoding)
		os.Exit(1)
	}

	fmt.Printf("%s  %s", encodedHash, *fileName)
	fmt.Println("")
}

func toBinaryString(src []byte, separator string) string {

	res := ""
	for _, b := range src {
		res += fmt.Sprintf("%08b", b) + separator
	}

	return res
}

func toOctalString(src []byte, separator string) string {

	res := ""
	for _, b := range src {
		res += fmt.Sprintf("%#o", b) + separator
	}

	return res
}

func toDecimalString(src []byte, separator string) string {

	res := ""
	for _, b := range src {
		res += fmt.Sprintf("%d", b) + separator
	}

	return res
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

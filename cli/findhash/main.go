package main

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/martinlindhe/gohash"
)

func main() {

	totalTime := time.Now()

	// TODO take a cli parameter specifying possible keys
	possibleKeysStr := "580%(=QWI+qwi*Nn"

	possibleKeys := gohash.StrToDistinctByteSlice(possibleKeysStr)

	fmt.Printf("Using the following keys:\n")
	fmt.Printf("%s", hex.Dump(possibleKeys))

	// TODO take hash as cli param
	hashString := "67ae1a64661ac8b4494666f58c4822408dd0a3e4"
	if !gohash.CanBeSha1(hashString) {
		fmt.Printf("ERROR str is not sha1: %s\n", hashString)
		return
	}

	expectedHash := gohash.HexStringToBytes(hashString)

	// TODO do this concurrently
	for length := 1; length < 10; length++ {
		start := time.Now()

		iterations := hashutil.Pow(len(possibleKeys), length)
		fmt.Printf("length = %2d, iterations = %-15d", length, iterations)

		chk := hashutil.FindMatchingHash(possibleKeys, hashutil.MatchSha1, expectedHash, length)
		fmt.Printf(", took %s\n", time.Since(start))

		if chk != nil {
			fmt.Printf("YAY, matched to: %s\n", string(chk))
			fmt.Printf("%s", hex.Dump(chk))
			break
		}
	}

	fmt.Printf("total time: ")
	fmt.Println(time.Since(totalTime))
}

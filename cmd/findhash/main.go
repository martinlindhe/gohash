package main

import (
	"fmt"
	"time"

	"github.com/martinlindhe/gohash"
)

func main() {

	totalTime := time.Now()

	hashString := "36367763ab73783c7af284446c59466b4cd653239a311cb7116d4618dee09a8425893dc7500b464fdaf1672d7bef5e891c6e2274568926a49fb4f45132c2a8b4"
	expectedHash := gohash.HexStringToBytes(hashString)

	chk := gohash.FindMatchingOnionURL(expectedHash)
	fmt.Println(chk)

	fmt.Printf("total time: ")
	fmt.Println(time.Since(totalTime))
}

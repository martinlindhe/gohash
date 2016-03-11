package gohash

import (
	"encoding/hex"
	"fmt"
)

// dump as hex
func dh(data []byte) {

	fmt.Println(len(data), "bytes:")
	fmt.Println(hex.Dump(data))
}

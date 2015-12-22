package gohash

import (
	"encoding/hex"
	"fmt"
)

// dump as hex
func dh(data []byte) {

	fmt.Printf("%d bytes:\n", len(data))
	fmt.Printf("%s", hex.Dump(data))
}

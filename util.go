package gohash

import (
	"encoding/hex"
	"fmt"
	"sort"
)

// return byte array from hex string
func hexStringToBytes(s string) []byte {

	res, err := hex.DecodeString(s)
	if err != nil {
		fmt.Println("ERROR decoding")
		return nil
	}

	return res
}

// sha1.Sum() returns 20 bytes (160 bit)
func byte20ArrayEquals(a [20]byte, b []byte) bool {

	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// sha512.Sum512() returns 64 bytes (512 bit)
func byte64ArrayEquals(a [64]byte, b []byte) bool {

	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func isByteInSlice(a byte, list []byte) bool {

	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func strToDistinctByteSlice(s string) []byte {

	res := []byte{}

	ptr := 0
	for i := 0; i < len(s); i++ {
		if isByteInSlice(s[i], res) {
			continue
		}
		res = append(res, s[i])
		ptr++
	}

	// XXX TODO sort it too
	sort.Sort(byteSlice(res))

	return res
}

// byteSlice implements sort.Interface for []Person based on
// the Age field.
type byteSlice []byte

func (a byteSlice) Len() int           { return len(a) }
func (a byteSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byteSlice) Less(i, j int) bool { return a[i] < a[j] }

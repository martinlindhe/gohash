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

func byteArrayEquals(a []byte, b []byte) bool {

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

	// sort it too
	sort.Sort(byteSlice(res))

	return res
}

// byteSlice implements sort.Interface to sort a []byte
type byteSlice []byte

func (a byteSlice) Len() int           { return len(a) }
func (a byteSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byteSlice) Less(i, j int) bool { return a[i] < a[j] }

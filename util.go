package gohash

import (
	"encoding/hex"
	"fmt"
)

// return byte array from hex string
func HexStringToBytes(s string) []byte {

	res, err := hex.DecodeString(s)
	if err != nil {
		fmt.Println("ERROR decoding")
		return nil
	}

	return res
}

// XXX sha1.Sum() returns 20 byte arrays
func ByteArrayEquals(a [20]byte, b []byte) bool {

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

func IsByteInSlice(a byte, list []byte) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func StrToDistinctByteSlice(s string) []byte {
	res := []byte{}

	ptr := 0
	for i := 0; i < len(s); i++ {
		if IsByteInSlice(s[i], res) {
			continue
		}
		res = append(res, s[i])
		ptr++
	}

	return res
}

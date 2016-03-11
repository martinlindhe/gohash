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

// md5.Sum() returns 16 bytes (128 bit)
func byte16ArrayEquals(a [16]byte, b []byte) bool {

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

// sha256.Sum224() returns 28 bytes (224 bit)
func byte28ArrayEquals(a [28]byte, b []byte) bool {

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

// sha256.Sum256() returns 32 bytes (256 bit)
func byte32ArrayEquals(a [32]byte, b []byte) bool {

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

// sha512.Sum384() returns 48 bytes (384 bit)
func byte48ArrayEquals(a [48]byte, b []byte) bool {

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

	// sort it too
	sort.Sort(byteSlice(res))

	return res
}

// byteSlice implements sort.Interface to sort a []byte
type byteSlice []byte

func (a byteSlice) Len() int           { return len(a) }
func (a byteSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byteSlice) Less(i, j int) bool { return a[i] < a[j] }

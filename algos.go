package gohash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"

	"github.com/jzelinskie/whirlpool"

	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

var (
	algos = map[string]int{
		"md5":        128,
		"sha1":       160,
		"sha224":     224,
		"sha256":     256,
		"sha384":     384,
		"sha512":     512,
		"sha512-224": 224,
		"sha512-256": 256,
		"sha3-224":   224,
		"sha3-256":   256,
		"sha3-384":   384,
		"sha3-512":   512,
		"shake128":   256,
		"shake256":   512,
		"whirlpool":  512,
		"ripemd160":  160,
	}

	algoEquals = map[string]func(*[]byte, *[]byte) bool{
		"md5":        md5Equals,
		"sha1":       sha1Equals,
		"sha224":     sha224Equals,
		"sha256":     sha256Equals,
		"sha384":     sha384Equals,
		"sha512":     sha512Equals,
		"sha512-224": sha512_224Equals,
		"sha512-256": sha512_256Equals,
		"sha3-224":   sha3_224Equals,
		"sha3-256":   sha3_256Equals,
		"sha3-384":   sha3_384Equals,
		"sha3-512":   sha3_512Equals,
		"shake128":   shake128Equals,
		"shake256":   shake256Equals,
		"whirlpool":  whirlpoolEquals,
		"ripemd160":  ripemd160Equals,
	}
)

func md5Equals(b *[]byte, expected *[]byte) bool {
	return byte16ArrayEquals(md5.Sum(*b), *expected)
}

func sha1Equals(b *[]byte, expected *[]byte) bool {
	return byte20ArrayEquals(sha1.Sum(*b), *expected)
}

func sha224Equals(b *[]byte, expected *[]byte) bool {
	return byte28ArrayEquals(sha256.Sum224(*b), *expected)
}
func sha256Equals(b *[]byte, expected *[]byte) bool {
	return byte32ArrayEquals(sha256.Sum256(*b), *expected)
}

func sha384Equals(b *[]byte, expected *[]byte) bool {
	return byte48ArrayEquals(sha512.Sum384(*b), *expected)
}

func sha512Equals(b *[]byte, expected *[]byte) bool {
	return byte64ArrayEquals(sha512.Sum512(*b), *expected)
}

func sha512_224Equals(b *[]byte, expected *[]byte) bool {
	return byte28ArrayEquals(sha512.Sum512_224(*b), *expected)
}

func sha512_256Equals(b *[]byte, expected *[]byte) bool {
	return byte32ArrayEquals(sha512.Sum512_256(*b), *expected)
}

func sha3_224Equals(b *[]byte, expected *[]byte) bool {
	return byte28ArrayEquals(sha3.Sum224(*b), *expected)
}

func sha3_256Equals(b *[]byte, expected *[]byte) bool {
	return byte32ArrayEquals(sha3.Sum256(*b), *expected)
}

func sha3_384Equals(b *[]byte, expected *[]byte) bool {
	return byte48ArrayEquals(sha3.Sum384(*b), *expected)
}

func sha3_512Equals(b *[]byte, expected *[]byte) bool {
	return byte64ArrayEquals(sha3.Sum512(*b), *expected)
}

func shake128Equals(b *[]byte, expected *[]byte) bool {

	h := make([]byte, 32)
	sha3.ShakeSum128(h, *b)
	return byteArrayEquals(h, *expected)
}

func shake256Equals(b *[]byte, expected *[]byte) bool {

	h := make([]byte, 64)
	sha3.ShakeSum256(h, *b)

	return byteArrayEquals(h, *expected)
}

func whirlpoolEquals(b *[]byte, expected *[]byte) bool {

	w := whirlpool.New()
	w.Write(*b)

	return byteArrayEquals(w.Sum(nil), *expected)
}

func ripemd160Equals(b *[]byte, expected *[]byte) bool {

	w := ripemd160.New()
	w.Write(*b)

	return byteArrayEquals(w.Sum(nil), *expected)
}
